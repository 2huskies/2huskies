CREATE TABLE IF NOT EXISTS abiturient (
    id           SERIAL PRIMARY KEY,
    first_name   TEXT   NOT NULL,
    last_last    TEXT   NOT NULL,
    birth_date   DATE   NOT NULL,
    birth_place  TEXT   NOT NULL,  
    address      TEXT   NOT NULL,   
    phone_number TEXT   NOT NULL     
);

CREATE TABLE IF NOT EXISTS subject (
    id   SERIAL PRIMARY KEY,
    name TEXT   NOT NULL
);

CREATE TABLE IF NOT EXISTS specialty (
    id   SERIAL  PRIMARY KEY,
	code INTEGER NOT NULL,
    name TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS university (
    id      SERIAL  PRIMARY KEY,
    name    TEXT    NOT NULL,
	active  BOOLEAN DEFAULT TRUE,
	url     TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS score (
    abiturient_id INTEGER NOT NULL,
    subject_id    INTEGER NOT NULL,
    score         INTEGER CHECK (score >= 0 AND score <= 100),
	
    PRIMARY KEY (abiturient_id, subject_id),
    CONSTRAINT abiturient_fkey FOREIGN KEY (abiturient_id) REFERENCES  abiturient(id) MATCH SIMPLE ON DELETE NO ACTION,
    CONSTRAINT subject_fkey    FOREIGN KEY (subject_id)    REFERENCES  subject(id)    MATCH SIMPLE ON DELETE NO ACTION
);

