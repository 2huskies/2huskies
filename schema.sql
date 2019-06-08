CREATE TABLE IF NOT EXISTS abiturient (
    id           SERIAL PRIMARY KEY,
    first_name   TEXT   NOT NULL,
    middle_name  TEXT,
    last_last    TEXT   NOT NULL,
    birth_date   DATE   NOT NULL,
    birth_place  TEXT   NOT NULL,  
    address      TEXT   NOT NULL,   
    phone_number TEXT   NOT NULL
);

CREATE TABLE IF NOT EXISTS login (
    login         VARCHAR(20) PRIMARY KEY NOT NULL,
    in_role       VARCHAR(20) NOT NULL,
    password      VARCHAR(30) NOT NULL,
    abiturient_id INTEGER NULL
);

-- Справочник предметов ЕГЭ
CREATE TABLE IF NOT EXISTS subject (
    id TEXT PRIMARY KEY,
    name TEXT   NOT NULL
);

-- Справочник специальностей в вузах
CREATE TABLE IF NOT EXISTS specialty (
    code TEXT  PRIMARY KEY,
    name TEXT    NOT NULL
);

CREATE TABLE IF NOT EXISTS university (
    id      SERIAL  PRIMARY KEY,
    name    TEXT    NOT NULL,
	active  BOOLEAN DEFAULT TRUE,
	url     TEXT    NOT NULL,
	short_name TEXT,
	rate integer DEFAULT NULL
);

--Результат ЕГЭ абитуриента
CREATE TABLE IF NOT EXISTS score (
    abiturient_id INTEGER NOT NULL,
    subject_id    TEXT NOT NULL,
    score         INTEGER CHECK (score >= 0 AND score <= 100),
	
    PRIMARY KEY (abiturient_id, subject_id),
    CONSTRAINT abiturient_fkey FOREIGN KEY (abiturient_id) REFERENCES  abiturient(id) MATCH SIMPLE ON DELETE NO ACTION,
    CONSTRAINT subject_fkey    FOREIGN KEY (subject_id)    REFERENCES  subject(id)    MATCH SIMPLE ON DELETE NO ACTION
);

CREATE TABLE IF NOT EXISTS doc_type(
    id   SERIAL PRIMARY KEY,
	name TEXT   NOT NULL,
	code text   NOT NULL
);

CREATE TABLE IF NOT EXISTS document (
    id   SERIAL PRIMARY KEY,
	name TEXT   NOT NULL,
    data BYTEA  NOT NULL,
	type_id INTEGER NOT NULL,
	abiturient_id INTEGER NOT NULL,
	
	CONSTRAINT subject_fkey    FOREIGN KEY (type_id)    REFERENCES  doc_type(id)    MATCH SIMPLE ON DELETE NO ACTION,
	CONSTRAINT abiturient_fkey FOREIGN KEY (abiturient_id) REFERENCES  abiturient(id) MATCH SIMPLE ON DELETE NO ACTION
);

