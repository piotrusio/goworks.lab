package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const maxFileSize = 235 * 1024 * 1024

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <csv-file-path>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	if err := splitCSV(inputFile); err != nil {
		log.Fatal(err)
	}
}

func splitCSV(inputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	header, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	baseDir := filepath.Dir(inputPath)
	baseName := filepath.Base(inputPath)
	ext := filepath.Ext(baseName)
	nameWithoutExt := baseName[:len(baseName)-len(ext)]

	var currentFile *os.File
	var currentWriter *csv.Writer
	var currentSize int64
	chunkNum := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read record: %w", err)
		}

		recordSize := calculateRecordSize(record)

		if currentFile == nil || currentSize+recordSize > maxFileSize {
			if currentWriter != nil {
				currentWriter.Flush()
				currentFile.Close()
			}

			outputPath := filepath.Join(baseDir, fmt.Sprintf("%s_chunk_%d%s", nameWithoutExt, chunkNum, ext))
			currentFile, err = os.Create(outputPath)
			if err != nil {
				return fmt.Errorf("failed to create output file: %w", err)
			}

			currentWriter = csv.NewWriter(currentFile)
			if err := currentWriter.Write(header); err != nil {
				return fmt.Errorf("failed to write header: %w", err)
			}

			currentSize = calculateRecordSize(header)
			chunkNum++
			fmt.Printf("Created chunk: %s\n", outputPath)
		}

		if err := currentWriter.Write(record); err != nil {
			return fmt.Errorf("failed to write record: %w", err)
		}

		currentSize += recordSize
	}

	if currentWriter != nil {
		currentWriter.Flush()
		currentFile.Close()
	}

	fmt.Printf("Successfully split CSV into %d chunks\n", chunkNum-1)
	return nil
}

func calculateRecordSize(record []string) int64 {
	size := int64(0)
	for i, field := range record {
		fieldSize := int64(len(field))
		if containsSpecialChars(field) {
			fieldSize += 2
		}
		size += fieldSize
		if i < len(record)-1 {
			size += 1
		}
	}
	size += 1
	return size
}

func containsSpecialChars(field string) bool {
	for _, char := range field {
		if char == '"' || char == ',' || char == '\n' || char == '\r' {
			return true
		}
	}
	return false
}
