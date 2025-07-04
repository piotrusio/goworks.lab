CREATE TABLE IF NOT EXISTS bikes (
  id INT GENERATED ALWAYS AS IDENTITY,
  brand TEXT,
  model TEXT
);

INSERT INTO bikes (brand, model) VALUES
('Trek', 'Madonne'),
('Specialized', 'Tarmac'),
('Cannondale', 'SuperSix'),
('Giant', 'Propel'),
('Colnago', 'V5RS');