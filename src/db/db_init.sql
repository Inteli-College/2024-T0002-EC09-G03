CREATE TABLE Sensors (
    id TEXT PRIMARY KEY,
    name VARCHAR(255),
    coordinate_x FLOAT(53),
    coordinate_y FLOAT(53)
);

CREATE TABLE Sensors_Data (
    id TEXT PRIMARY KEY,
    sensor_id TEXT REFERENCES Sensors(id),
    created_at TIMESTAMP,
    data TEXT
);
