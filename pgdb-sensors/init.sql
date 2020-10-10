CREATE TABLE sensors (
  id integer NOT NULL,
  sensorName character varying NOT NULL,
  seonsorUuid character varying NOT NULL,
  measuringUnit character varying NOT NULL,
  minContructiveValue double precision NOT NULL,
  maxContructiveValue double precision NOT NULL,
  minNominalValue double precision NULL,
  maxNominalValue double precision NULL,
  geoLocation character varying NULL
);

CREATE SEQUENCE sensors_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE sensors_id_seq OWNED BY sensors.id;

ALTER TABLE ONLY sensors ALTER COLUMN id SET DEFAULT nextval('sensors_id_seq'::regclass);

/** Insert the 3 sensor names in sensors table **/
INSERT INTO sensors 
( sensorName,
  seonsorUuid,
  measuringUnit,
  minContructiveValue,
  maxContructiveValue,
  minNominalValue,
  maxNominalValue,
  geoLocation)
VALUES
('sensor-A','00001','Celsius', -50, 80, -30, 50, 'someGoogleMapsLocationA'),
('sensor-B','00002','Celsius', -40, 90, -20, 60, 'someGoogleMapsLocationB'),
('sensor-C','00003','Celsius', -30, 100, -10, 70, 'someGoogleMapsLocationC')


