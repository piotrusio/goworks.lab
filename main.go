package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"
	"reflect"

    "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
    tUUID         = reflect.TypeOf(uuid.UUID{})
    uuidSubtype   = byte(0x04)
    mongoRegistry = bson.NewRegistry()
)

type config struct {
    port    int
    env     string
    uri     string
}

type MongoDB struct {
    Client *mongo.Client
}

func main() {
    var cfg config

    flag.IntVar(&cfg.port, "port", 4000, "API server port")
    flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
    flag.StringVar(&cfg.uri, "db-uri", os.Getenv("MONGODB_URI"), "MongoDB URI")
    flag.Parse()

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    // main application context
    appCtx, stop := signal.NotifyContext(context.Background(),
        syscall.SIGINT, syscall.SIGTERM)
    defer stop()

    // database startup context
    startupCtx, startupCancel := context.WithTimeout(appCtx, 30*time.Second)
    defer startupCancel()

    // TODO: retry
    mongoDB, err := NewMongoDB(startupCtx, cfg.uri)
    if err != nil {
        logger.Error("Failed to connect to MongoDB", "error", err)
        os.Exit(1)
    }

    <-appCtx.Done()
    logger.Info("Shutdown signal received, attempting graceful shutdown.")

    // graceful shutdown context
    shutdownCtx, shutdownCancel := context.WithTimeout(appCtx, 10*time.Second)
    defer shutdownCancel()

    shutdownErr := false

    if err := mongoDB.Close(shutdownCtx); err != nil {
        logger.Error("Error during MongoDB disconnect", "error", err)
        shutdownErr = true
    } else {
        logger.Info("MongoDB gracefully disconnected")
    }

    logger.Info("Service exiting.")
    if shutdownErr {
        os.Exit(1)
    } else {
        os.Exit(0)
    }
}

func NewMongoDB(ctx context.Context, uri string) (*MongoDB, error) {
    clientOptions := options.Client().ApplyURI(uri).SetRegistry(mongoRegistry)
    client, err := mongo.Connect(clientOptions)
    if err != nil {
        return nil, err
    }

    if err := client.Ping(ctx, nil); err != nil {
        _ = client.Disconnect(ctx)
        return nil, err
    }

    return &MongoDB{Client: client}, nil
}

func (m *MongoDB) Close(ctx context.Context) error {
    return m.Client.Disconnect(ctx)
}

func init() {
    mongoRegistry.RegisterTypeEncoder(tUUID, bson.ValueEncoderFunc(uuidEncodeValue))
    mongoRegistry.RegisterTypeDecoder(tUUID, bson.ValueDecoderFunc(uuidDecodeValue))
}

func uuidEncodeValue(ec bson.EncodeContext, vw bson.ValueWriter, val reflect.Value) error {
    if !val.IsValid() || val.Type() != tUUID {
        return bson.ValueEncoderError{Name: "uuidEncodeValue", Types: []reflect.Type{tUUID}, Received: val}
    }
    b := val.Interface().(uuid.UUID)
    return vw.WriteBinaryWithSubtype(b[:], uuidSubtype)
}

func uuidDecodeValue(dc bson.DecodeContext, vr bson.ValueReader, val reflect.Value) error {
    if !val.CanSet() || val.Type() != tUUID {
        return bson.ValueDecoderError{Name: "uuidDecodeValue", Types: []reflect.Type{tUUID}, Received: val}
    }

    var data []byte
    var subtype byte
    var err error
    switch vrType := vr.Type(); vrType {
    case bson.TypeBinary:
        data, subtype, err = vr.ReadBinary()
        if subtype != uuidSubtype {
            return fmt.Errorf("unsupported binary subtype %v for UUID", subtype)
        }
    case bson.TypeNull:
        err = vr.ReadNull()
    case bson.TypeUndefined:
        err = vr.ReadUndefined()
    default:
        return fmt.Errorf("cannot decode %v into a UUID", vrType)
    }

    if err != nil {
        return err
    }
    uuid2, err := uuid.FromBytes(data)
    if err != nil {
        return err
    }
    val.Set(reflect.ValueOf(uuid2))
    return nil
}