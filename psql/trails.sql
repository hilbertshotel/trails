CREATE DATABASE trails;

\c trails

CREATE TABLE workouts (
  id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  date TIMESTAMPTZ NOT NULL,
  duration VARCHAR(20) NOT NULL,
  distance REAL NOT NULL,
  pace_avg REAL NOT NULL,
  pace_best REAL NOT NULL,
  hr_avg INT NOT NULL,
  hr_max INT NOT NULL,
  elev_gain INT NOT NULL,
  elev_loss INT NOT NULL,
  location VARCHAR(100) NOT NULL,
  terrain VARCHAR(10) NOT NULL,
  footwear VARCHAR(10) NOT NULL 
);
