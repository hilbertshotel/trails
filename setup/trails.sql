CREATE DATABASE trails;

\c trails

CREATE TABLE workouts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    date VARCHAR(30) NOT NULL,      
	distance FLOAT NOT NULL,
    duration VARCHAR(30) NOT NULL,
    elevation INT NOT NULL,
    avg_pace FLOAT NOT NULL,
    avg_hr INT NOT NULL
);