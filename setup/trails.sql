CREATE DATABASE trails;

CREATE TABLE trails.sessions (
  id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  date TIMESTAMTZ NOT NULL,
  duration TIME NOT NULL,
  distance
  avg_hr
  max_hr
  fastest_km
  avg_pace
  elev_gain
  elev_loss
  

)