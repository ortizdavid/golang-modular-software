-- tables for 'reference' schema 

DROP TYPE IF EXISTS TYPE_GENDER;
CREATE TYPE TYPE_GENDER AS ENUM('Masculino', 'Feminino');

DROP TABLE IF EXISTS reference.marital_statuses;
CREATE TABLE IF NOT EXISTS reference.marital_statuses (
	status_id SERIAL NOT NULL PRIMARY KEY,
	status_name VARCHAR(100) NOT NULL UNIQUE,
	code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS reference.identification_types;
CREATE TABLE IF NOT EXISTS reference.identification_types (
	type_id SERIAL NOT NULL PRIMARY KEY,
	type_name VARCHAR(100) NOT NULL UNIQUE,
	code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS reference.document_types;
CREATE TABLE IF NOT EXISTS reference.document_types (
	type_id SERIAL NOT NULL PRIMARY KEY,
	type_name VARCHAR(100) NOT NULL UNIQUE,
	code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS reference.contact_types;
CREATE TABLE IF NOT EXISTS reference.contact_types (
	type_id SERIAL NOT NULL PRIMARY KEY,
    type_name VARCHAR(30),
    code VARCHAR(20) UNIQUE
);

DROP TABLE IF EXISTS reference.person;
CREATE TABLE IF NOT EXISTS reference.person (
	person_id SERIAL NOT NULL PRIMARY KEY,
	identification_type_id INT,
	marital_status_id INT,
	first_name VARCHAR(100) NOT NULL,
	last_name VARCHAR(100) NOT NULL,
	birth_date DATE,
	gender TYPE_GENDER,
	identification_number VARCHAR(30),
	unique_id VARCHAR(50) UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_identification_person FOREIGN KEY(identification_type_id) REFERENCES reference.identification_types(type_id),
	CONSTRAINT fk_marital_statuses FOREIGN KEY(marital_status_id) REFERENCES reference.marital_statuses(status_id)
);

DROP TABLE IF EXISTS reference.contacts;
CREATE TABLE IF NOT EXISTS reference.contacts (
	contact_id SERIAL NOT NULL PRIMARY KEY,
    person_id INT NOT NULL,
    contact_type_id INT NOT NULL,
	email VARCHAR(150) UNIQUE,
    phone INT UNIQUE,
    unique_id VARCHAR(50) UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_person FOREIGN KEY(person_id) REFERENCES reference.person(person_id),
    CONSTRAINT fk_contact_type FOREIGN KEY(contact_type_id) REFERENCES reference.contact_types(type_id)
);

DROP TABLE IF EXISTS reference.addresses;
CREATE TABLE IF NOT EXISTS reference.addresses (
	address_id SERIAL NOT NULL PRIMARY KEY,
    person_id INT NOT NULL,
    state VARCHAR(150),
	city VARCHAR(150),
    district VARCHAR(150),
    postal_code VARCHAR(20),
    unique_id VARCHAR(50) UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_person FOREIGN KEY(person_id) REFERENCES reference.person(person_id)
);

DROP TABLE IF EXISTS reference.documents;
CREATE TABLE IF NOT EXISTS reference.documents (
	document_id SERIAL NOT NULL PRIMARY KEY,
    person_id INT NOT NULL,
    document_type_id INT NOT NULL,
	file_name VARCHAR(150),
    unique_id VARCHAR(50) UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_person FOREIGN KEY(person_id) REFERENCES reference.person(person_id),
    CONSTRAINT fk_document_type FOREIGN KEY(document_type_id) REFERENCES reference.document_types(type_id)
);

-- Initial inserts
-- marital_statuses
INSERT INTO reference.marital_statuses (code, status_name) VALUES ('solteiro', 'Solteiro');  
INSERT INTO reference.marital_statuses (code, status_name) VALUES ('casado_com_registo', 'Casado (com registo)');  
INSERT INTO reference.marital_statuses (code, status_name) VALUES ('casado_sem_registo', 'Casado (sem registo)');  
INSERT INTO reference.marital_statuses (code, status_name) VALUES ('divorciado', 'Divorciado');  
INSERT INTO reference.marital_statuses (code, status_name) VALUES ('separado', 'Separado');  
INSERT INTO reference.marital_statuses (code, status_name) VALUES ('viuvo', 'Viúvo');  
INSERT INTO reference.marital_statuses (code, status_name) VALUES ('outro', 'Outro');  
-- identification_types
INSERT INTO reference.identification_types (code, type_name) VALUES ('bi', 'Bilhete de Identidade');
INSERT INTO reference.identification_types (code, type_name) VALUES ('passaporte','Passaporte');
INSERT INTO reference.identification_types (code, type_name) VALUES ('residente', 'Cartão de Residente');
INSERT INTO reference.identification_types (code, type_name) VALUES ('bi-cverde', 'Bilhete de Identidade (Cabo Verde)');
INSERT INTO reference.identification_types (code, type_name) VALUES ('autorizacao', 'Autorização de Residência');
INSERT INTO reference.identification_types (code, type_name) VALUES ('bi-militar', 'Bilhete de Identidade (militar)');
INSERT INTO reference.identification_types (code, type_name) VALUES ('certificado', 'Certificado de Registo de Cidadão UE');
INSERT INTO reference.identification_types (code, type_name) VALUES ('bi-estrangeiro', 'Bilhete de Identidade (estrangeiro)');
INSERT INTO reference.identification_types (code, type_name) VALUES ('outro', 'Outro');
-- marital_document_types
INSERT INTO reference.document_types (type_name, code) VALUES ('Diploma', 'diploma');  
INSERT INTO reference.document_types (type_name, code) VALUES ('Nº de Identificação', 'nif');  
INSERT INTO reference.document_types (type_name, code) VALUES ('Currículo Vitae', 'curriculo');  
INSERT INTO reference.document_types (type_name, code) VALUES ('Bilhete de Identidade', 'bilhete');  
INSERT INTO reference.document_types (type_name, code) VALUES ('Registo Militar', 'registo-militar');  
INSERT INTO reference.document_types (type_name, code) VALUES ('Documento Bancário', 'doc-bancario'); 
INSERT INTO reference.document_types (type_name, code) VALUES ('Registo Criminal', 'registo-criminal');  
INSERT INTO reference.document_types (type_name, code) VALUES ('Recenseamento Militar', 'recenseamento');  
INSERT INTO reference.document_types (type_name, code) VALUES ('Certificado de Habilitações', 'certificado');  
-- contact_types
INSERT INTO reference.contact_types (type_name, code) VALUES ('Casa', 'casa');  
INSERT INTO reference.contact_types (type_name, code) VALUES ('pessoal', 'pessoal');  
INSERT INTO reference.contact_types (type_name, code) VALUES ('Empresa', 'empresa');  
INSERT INTO reference.contact_types (type_name, code) VALUES ('Familiar', 'familiar');  
INSERT INTO reference.contact_types (type_name, code) VALUES ('Outro', 'outro');  


