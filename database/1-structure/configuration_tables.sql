
DROP TABLE IF EXISTS configurations.email_configuration;
CREATE TABLE configurations.email_configuration (
    configuration_id INT PRIMARY KEY,
    smtp_server VARCHAR(50),
    smtp_port INT,
    sender_email VARCHAR(100),
    sender_password VARCHAR(150)
);

DROP TABLE IF EXISTS configurations.email_configuration;
CREATE TABLE configurations.basic_configuration (
    configuration_id INT PRIMARY KEY NOT NULL,
    company_name VARCHAR(100) NOT NULL,
    company_acronym VARCHAR(50),
    company_main_color VARCHAR(10),
    company_image VARCHAR(100),
    company_email VARCHAR(100),
    company_phone VARCHAR(20),
    num_of_records_per_page INT
);
