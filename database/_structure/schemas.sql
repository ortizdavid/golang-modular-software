-- Create database
CREATE DATABASE golang_modular_software;

\c golang_modular_software;

-- SCHEMAS 

-- for all configurations
CREATE SCHEMA configurations;

-- user, roles, permission, authentication management
CREATE SCHEMA authentication;

-- human resources management
CREATE SCHEMA human_resources;

-- customers management
CREATE SCHEMA customers;
