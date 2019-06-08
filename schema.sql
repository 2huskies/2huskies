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
    role       VARCHAR(20) NOT NULL,
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
	city    TEXT    NOT NULL,
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

CREATE TABLE IF NOT EXISTS faculty (
    id   SERIAL PRIMARY KEY,
    university_id INTEGER NOT NULL,
	name    TEXT NOT NULL,
	
    CONSTRAINT university_fkey FOREIGN KEY (university_id) REFERENCES  university(id) MATCH SIMPLE ON DELETE NO ACTION
);

--Кол-во баллов ЕГЭ минимально необходимых для подачи документов
CREATE TABLE IF NOT EXISTS pass_score (
    university_id INTEGER NOT NULL,
    subject_id    TEXT NOT NULL,
	specialty_id  TEXT NOT NULL,
	faculty_id INTEGER NOT NULL,
    pass_score    INTEGER CHECK (pass_score >= 0 AND pass_score <= 100),
	
    PRIMARY KEY (university_id, subject_id, specialty_id),
    CONSTRAINT university_fkey FOREIGN KEY (university_id) REFERENCES  university(id) MATCH SIMPLE ON DELETE NO ACTION,
    CONSTRAINT subject_fkey    FOREIGN KEY (subject_id)    REFERENCES  subject(id)    MATCH SIMPLE ON DELETE NO ACTION,
	CONSTRAINT specialty_fkey    FOREIGN KEY (specialty_id)    REFERENCES  specialty(code)    MATCH SIMPLE ON DELETE NO ACTION,
	CONSTRAINT faculty_fkey    FOREIGN KEY (faculty_id)    REFERENCES  faculty(id)    MATCH SIMPLE ON DELETE NO ACTION
);