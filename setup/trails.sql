CREATE DATABASE trails;

\c trails

CREATE TABLE workouts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    date VARCHAR(10) NOT NULL,      
	distance NUMERIC(2) NOT NULL,
    duration INT NOT NULL, -- seconds
    elevation INT NOT NULL,
    avg_pace NUMERIC(2) NOT NULL,
    avg_hr INT NOT NULL
);