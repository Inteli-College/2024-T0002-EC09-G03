CREATE TABLE Sensors (
    id TEXT PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE Sensors_Data (
    id TEXT PRIMARY KEY,
    sensor_id TEXT REFERENCES Sensors(id),
    created_at TIMESTAMP,
    data TEXT,
    coordinate_x FLOAT(53),
    coordinate_y FLOAT(53)
);

CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY, -- 'SERIAL' para PostgreSQL
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL, -- Armazena o hash da senha, não a senha em texto claro
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Alerts (
    alert_id SERIAL PRIMARY KEY,
    sensor_id TEXT REFERENCES Sensors(id),
    alert_description TEXT,
    alert_level VARCHAR(50),
    alert_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
