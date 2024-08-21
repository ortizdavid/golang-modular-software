-- Table: task_statuses
DROP TABLE IF EXISTS reference.task_statuses;
CREATE TABLE reference.task_statuses (
    id SERIAL PRIMARY KEY,
    status_name VARCHAR(100) NOT NULL,
    code VARCHAR(30) NOT NULL UNIQUE,
    lbl_color VARCHAR(20),
    bg_color VARCHAR(20),
    description TEXT NOT NULL,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO reference.task_statuses (status_name, code, lbl_color, bg_color, description) VALUES
('Pending', 'PENDING', 'yellow', 'ui yellow', 'The task has been created but not yet started.'),
('In Progress', 'IN_PROGRESS', 'blue', 'ui blue', 'The task is currently being worked on.'),
('Completed', 'COMPLETED', 'green', 'ui positive', 'The task has been completed.'),
('On Hold', 'ON_HOLD', 'orange', 'ui orange', 'The task is temporarily paused.'),
('Cancelled', 'CANCELLED', 'red', 'ui negative', 'The task has been cancelled.'),
('Review', 'REVIEW', 'teal', 'ui teal', 'The task is completed and under review.'),
('Failed', 'FAILED', 'red', 'ui negative', 'The task was attempted but did not succeed.');


-- Table: approval_statuses 
DROP TABLE IF EXISTS reference.approval_statuses;
CREATE TABLE reference.approval_statuses (
    id SERIAL PRIMARY KEY,
    status_name VARCHAR(100) NOT NULL,
    code VARCHAR(30) NOT NULL UNIQUE,
    lbl_color VARCHAR(20),
    bg_color VARCHAR(20),
    description TEXT NOT NULL,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO reference.approval_statuses (status_name, code, lbl_color, bg_color, description) VALUES
('Pending Approval', 'PENDING_APPROVAL', 'yellow', 'ui yellow', 'Waiting for approval from the designated authority.'),
('Approved', 'APPROVED', 'green', 'ui positive', 'The item has been approved and can proceed.'),
('Rejected', 'REJECTED', 'red', 'ui negative', 'The item has been rejected and requires changes.'),
('Under Review', 'UNDER_REVIEW', 'teal', 'ui teal', 'The item is under review and awaiting a decision.'),
('Escalated', 'ESCALATED', 'orange', 'ui orange', 'The approval request has been escalated for further review.'),
('Deferred', 'DEFERRED', 'grey', 'ui grey', 'The approval decision has been postponed.'),
('Awaiting Input', 'AWAITING_INPUT', 'yellow', 'ui yellow', 'The item is awaiting additional information or input before proceeding with the approval.'),
('Final Review', 'FINAL_REVIEW', 'blue', 'ui blue', 'The item is in the final review stage before the final approval or rejection decision.'),
('Resubmitted', 'RESUBMITTED', 'purple', 'ui purple', 'The item has been resubmitted after a previous rejection or request for changes.'),
('Reviewed', 'REVIEWED', 'teal', 'ui teal', 'The item has been reviewed but not yet approved or rejected.'),
('Pending Documentation', 'PENDING_DOCUMENTATION', 'orange', 'ui orange', 'The approval is pending the submission of required documentation.');


-- Table: workflow_statuses
DROP TABLE IF EXISTS reference.workflow_statuses;
CREATE TABLE reference.workflow_statuses (
    id SERIAL PRIMARY KEY,
    status_name VARCHAR(100) NOT NULL UNIQUE,
    code VARCHAR(30) NOT NULL UNIQUE,
    lbl_color VARCHAR(20),
    bg_color VARCHAR(20),
    description TEXT,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO reference.workflow_statuses (status_name, code, lbl_color, bg_color, description) VALUES
('Not Started', 'NOT_STARTED', 'grey', 'ui grey', 'The task or process has not yet begun.'),
('In Review', 'IN_REVIEW', 'teal', 'ui teal', 'The task or process is under review.'),
('Approved', 'APPROVED', 'green', 'ui positive', 'The task or process has been approved to move forward.'),
('Rejected', 'REJECTED', 'red', 'ui negative', 'The task or process has been rejected and may require correction.'),
('Completed', 'COMPLETED', 'green', 'ui positive', 'The task or process has been fully completed.'),
('Archived', 'ARCHIVED', 'grey', 'ui grey', 'The task or process has been archived.');


-- Table: document_statuses
DROP TABLE IF EXISTS reference.document_statuses;
CREATE TABLE reference.document_statuses (
    id SERIAL PRIMARY KEY,
    status_name VARCHAR(100) NOT NULL UNIQUE,
    code VARCHAR(30) UNIQUE NOT NULL,
    lbl_color VARCHAR(20),
    bg_color VARCHAR(20),
    description TEXT,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Update document_statuses with Semantic UI bg_color classes
INSERT INTO reference.document_statuses (status_name, code, lbl_color, bg_color, description) VALUES
('Draft', 'DRAFT', 'blue', 'ui blue', 'The document is in the draft stage and is not yet finalized.'),
('Submitted', 'SUBMITTED', 'yellow', 'ui yellow', 'The document has been submitted for review or approval.'),
('Under Review', 'UNDER_REVIEW', 'teal', 'ui teal', 'The document is currently being reviewed by the designated authority.'),
('Approved', 'APPROVED', 'green', 'ui positive', 'The document has been approved and is finalized.'),
('Rejected', 'REJECTED', 'red', 'ui negative', 'The document has been rejected and may require changes.'),
('Archived', 'ARCHIVED', 'grey', 'ui grey', 'The document has been archived and is no longer actively used.'),
('Expired', 'EXPIRED', 'orange', 'ui orange', 'The document is no longer valid due to expiration.'),
('Canceled', 'CANCELED', 'red', 'ui negative', 'The document process has been canceled and will not proceed.');


-- Table: user_statuses
DROP TABLE IF EXISTS reference.user_statuses;
CREATE TABLE reference.user_statuses (
    id SERIAL PRIMARY KEY,
    status_name VARCHAR(100) NOT NULL UNIQUE,
    code VARCHAR(30) NOT NULL UNIQUE,
    lbl_color VARCHAR(20),
    bg_color VARCHAR(20),
    description TEXT,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO reference.user_statuses (status_name, code, lbl_color, bg_color, description) VALUES
('Active', 'ACTIVE', 'green', 'ui positive', 'The user is currently active and using the system.'),
('Inactive', 'INACTIVE', 'grey', 'ui grey', 'The user is not currently active but retains an account.'),
('Locked', 'LOCKED', 'red', 'ui negative', 'The user’s account is locked due to security reasons.'),
('Suspended', 'SUSPENDED', 'orange', 'ui orange', 'The user’s account is temporarily suspended.'),
('Deleted', 'DELETED', 'red', 'ui negative', 'The user’s account has been deleted from the system.');


-- Table: marital_statuses
DROP TABLE IF EXISTS reference.marital_statuses;
CREATE TABLE IF NOT EXISTS reference.marital_statuses (
    status_id SERIAL PRIMARY KEY,
    status_name VARCHAR(100) NOT NULL UNIQUE,
    code VARCHAR(30) UNIQUE NOT NULL,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO reference.marital_statuses (code, status_name) VALUES 
('SINGLE', 'Single'),
('MARRIED_WITH_REGISTRATION', 'Married (with registration)'),  
('MARRIED_WITHOUT_REGISTRATION', 'Married (without registration)'),  
('DIVORCED', 'Divorced'),  
('SEPARATED', 'Separated'),  
('WIDOWED', 'Widowed'),  
('OTHER', 'Other');


-- Table: identification_types
DROP TABLE IF EXISTS reference.identification_types;
CREATE TABLE IF NOT EXISTS reference.identification_types (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(100) NOT NULL UNIQUE,
    code VARCHAR(30) UNIQUE NOT NULL,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO reference.identification_types (code, type_name) VALUES 
('ID_CARD', 'Identity Card'),
('PASSPORT', 'Passport'),
('RESIDENCE_CARD', 'Residence Card'),
('ID_CARD_CAPE_VERDE', 'Identity Card (Cape Verde)'),
('RESIDENCE_PERMIT', 'Residence Permit'),
('MILITARY_ID_CARD', 'Military Identity Card'),
('EU_REGISTRATION_CERTIFICATE', 'EU Citizen Registration Certificate'),
('FOREIGN_ID_CARD', 'Foreign Identity Card'),
('OTHER', 'Other');


-- Table: contact_types
DROP TABLE IF EXISTS reference.contact_types;
CREATE TABLE IF NOT EXISTS reference.contact_types (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(100) NOT NULL UNIQUE,
    code VARCHAR(30) UNIQUE NOT NULL,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO reference.contact_types (code, type_name) VALUES 
('HOME', 'Home'),
('PERSONAL', 'Personal'),
('BUSINESS', 'Business'),
('FAMILY', 'Family'),
('OTHER', 'Other');

-- -------------------------------------------------------------------------------------------
-- Table countries
DROP TABLE IF EXISTS reference.countries;
CREATE TABLE reference.countries (
    country_id SERIAL PRIMARY KEY,
    country_name VARCHAR(150) UNIQUE NOT NULL,
    iso_code CHAR(2) UNIQUE,
    dialing_code VARCHAR(7),
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO reference.countries (country_name, iso_code, dialing_code) VALUES
('Afghanistan', 'AF', '+93'),
('Albania', 'AL', '+355'),
('Algeria', 'DZ', '+213'),
('Andorra', 'AD', '+376'),
('Angola', 'AO', '+244'),
('Antigua and Barbuda', 'AG', '+1-268'),
('Argentina', 'AR', '+54'),
('Armenia', 'AM', '+374'),
('Australia', 'AU', '+61'),
('Austria', 'AT', '+43'),
('Azerbaijan', 'AZ', '+994'),
('Bahamas', 'BS', '+1-242'),
('Bahrain', 'BH', '+973'),
('Bangladesh', 'BD', '+880'),
('Barbados', 'BB', '+1-246'),
('Belarus', 'BY', '+375'),
('Belgium', 'BE', '+32'),
('Belize', 'BZ', '+501'),
('Benin', 'BJ', '+229'),
('Bhutan', 'BT', '+975'),
('Bolivia', 'BO', '+591'),
('Bosnia and Herzegovina', 'BA', '+387'),
('Botswana', 'BW', '+267'),
('Brazil', 'BR', '+55'),
('Brunei Darussalam', 'BN', '+673'),
('Bulgaria', 'BG', '+359'),
('Burkina Faso', 'BF', '+226'),
('Burundi', 'BI', '+257'),
('Cabo Verde', 'CV', '+238'),
('Cambodia', 'KH', '+855'),
('Cameroon', 'CM', '+237'),
('Canada', 'CA', '+1'),
('Central African Republic', 'CF', '+236'),
('Chad', 'TD', '+235'),
('Chile', 'CL', '+56'),
('China', 'CN', '+86'),
('Colombia', 'CO', '+57'),
('Comoros', 'KM', '+269'),
('Congo', 'CG', '+242'),
('Congo, Democratic Republic of the', 'CD', '+243'),
('Costa Rica', 'CR', '+506'),
('Croatia', 'HR', '+385'),
('Cuba', 'CU', '+53'),
('Cyprus', 'CY', '+357'),
('Czech Republic', 'CZ', '+420'),
('Denmark', 'DK', '+45'),
('Djibouti', 'DJ', '+253'),
('Dominica', 'DM', '+1-767'),
('Dominican Republic', 'DO', '+1-809'),
('Ecuador', 'EC', '+593'),
('Egypt', 'EG', '+20'),
('El Salvador', 'SV', '+503'),
('Equatorial Guinea', 'GQ', '+240'),
('Eritrea', 'ER', '+291'),
('Estonia', 'EE', '+372'),
('Eswatini', 'SZ', '+268'),
('Ethiopia', 'ET', '+251'),
('Fiji', 'FJ', '+679'),
('Finland', 'FI', '+358'),
('France', 'FR', '+33'),
('Gabon', 'GA', '+241'),
('Gambia', 'GM', '+220'),
('Georgia', 'GE', '+995'),
('Germany', 'DE', '+49'),
('Ghana', 'GH', '+233'),
('Greece', 'GR', '+30'),
('Grenada', 'GD', '+1-473'),
('Guatemala', 'GT', '+502'),
('Guinea', 'GN', '+224'),
('Guinea-Bissau', 'GW', '+245'),
('Guyana', 'GY', '+592'),
('Haiti', 'HT', '+509'),
('Honduras', 'HN', '+504'),
('Hungary', 'HU', '+36'),
('Iceland', 'IS', '+354'),
('India', 'IN', '+91'),
('Indonesia', 'ID', '+62'),
('Iran', 'IR', '+98'),
('Iraq', 'IQ', '+964'),
('Ireland', 'IE', '+353'),
('Israel', 'IL', '+972'),
('Italy', 'IT', '+39'),
('Jamaica', 'JM', '+1-876'),
('Japan', 'JP', '+81'),
('Jordan', 'JO', '+962'),
('Kazakhstan', 'KZ', '+7'),
('Kenya', 'KE', '+254'),
('Kiribati', 'KI', '+686'),
('Korea, North', 'KP', '+850'),
('Korea, South', 'KR', '+82'),
('Kosovo', 'XK', '+383'),
('Kuwait', 'KW', '+965'),
('Kyrgyzstan', 'KG', '+996'),
('Laos', 'LA', '+856'),
('Latvia', 'LV', '+371'),
('Lebanon', 'LB', '+961'),
('Lesotho', 'LS', '+266'),
('Liberia', 'LR', '+231'),
('Libya', 'LY', '+218'),
('Liechtenstein', 'LI', '+423'),
('Lithuania', 'LT', '+370'),
('Luxembourg', 'LU', '+352'),
('Madagascar', 'MG', '+261'),
('Malawi', 'MW', '+265'),
('Malaysia', 'MY', '+60'),
('Maldives', 'MV', '+960'),
('Mali', 'ML', '+223'),
('Malta', 'MT', '+356'),
('Marshall Islands', 'MH', '+692'),
('Mauritania', 'MR', '+222'),
('Mauritius', 'MU', '+230'),
('Mexico', 'MX', '+52'),
('Micronesia', 'FM', '+691'),
('Moldova', 'MD', '+373'),
('Monaco', 'MC', '+377'),
('Mongolia', 'MN', '+976'),
('Montenegro', 'ME', '+382'),
('Morocco', 'MA', '+212'),
('Mozambique', 'MZ', '+258'),
('Myanmar', 'MM', '+95'),
('Namibia', 'NA', '+264'),
('Nauru', 'NR', '+674'),
('Nepal', 'NP', '+977'),
('Netherlands', 'NL', '+31'),
('New Zealand', 'NZ', '+64'),
('Nicaragua', 'NI', '+505'),
('Niger', 'NE', '+227'),
('Nigeria', 'NG', '+234'),
('North Macedonia', 'MK', '+389'),
('Norway', 'NO', '+47'),
('Oman', 'OM', '+968'),
('Pakistan', 'PK', '+92'),
('Palau', 'PW', '+680'),
('Panama', 'PA', '+507'),
('Papua New Guinea', 'PG', '+675'),
('Paraguay', 'PY', '+595'),
('Peru', 'PE', '+51'),
('Philippines', 'PH', '+63'),
('Poland', 'PL', '+48'),
('Portugal', 'PT', '+351'),
('Qatar', 'QA', '+974'),
('Romania', 'RO', '+40'),
('Russia', 'RU', '+7'),
('Rwanda', 'RW', '+250'),
('Saint Kitts and Nevis', 'KN', '+1-869'),
('Saint Lucia', 'LC', '+1-758'),
('Saint Vincent and the Grenadines', 'VC', '+1-784'),
('Samoa', 'WS', '+685'),
('San Marino', 'SM', '+378'),
('Sao Tome and Principe', 'ST', '+239'),
('Saudi Arabia', 'SA', '+966'),
('Senegal', 'SN', '+221'),
('Serbia', 'RS', '+381'),
('Seychelles', 'SC', '+248'),
('Sierra Leone', 'SL', '+232'),
('Singapore', 'SG', '+65'),
('Slovakia', 'SK', '+421'),
('Slovenia', 'SI', '+386'),
('Solomon Islands', 'SB', '+677'),
('Somalia', 'SO', '+252'),
('South Africa', 'ZA', '+27'),
('South Sudan', 'SS', '+211'),
('Spain', 'ES', '+34'),
('Sri Lanka', 'LK', '+94'),
('Sudan', 'SD', '+249'),
('Suriname', 'SR', '+597'),
('Sweden', 'SE', '+46'),
('Switzerland', 'CH', '+41'),
('Syria', 'SY', '+963'),
('Taiwan', 'TW', '+886'),
('Tajikistan', 'TJ', '+992'),
('Tanzania', 'TZ', '+255'),
('Thailand', 'TH', '+66'),
('Timor-Leste', 'TL', '+670'),
('Togo', 'TG', '+228'),
('Tonga', 'TO', '+676'),
('Trinidad and Tobago', 'TT', '+1-868'),
('Tunisia', 'TN', '+216'),
('Turkey', 'TR', '+90'),
('Turkmenistan', 'TM', '+993'),
('Tuvalu', 'TV', '+688'),
('Uganda', 'UG', '+256'),
('Ukraine', 'UA', '+380'),
('United Arab Emirates', 'AE', '+971'),
('United Kingdom', 'GB', '+44'),
('United States', 'US', '+1'),
('Uruguay', 'UY', '+598'),
('Uzbekistan', 'UZ', '+998'),
('Vanuatu', 'VU', '+678'),
('Vatican City', 'VA', '+379'),
('Venezuela', 'VE', '+58'),
('Vietnam', 'VN', '+84'),
('Yemen', 'YE', '+967'),
('Zambia', 'ZM', '+260'),
('Zimbabwe', 'ZW', '+263');


-- Table: currencies
DROP TABLE IF EXISTS reference.currencies;
CREATE TABLE reference.currencies (
    currency_id SERIAL PRIMARY KEY,
    currency_name VARCHAR(100) NOT NULL UNIQUE,
    code CHAR(3) NOT NULL,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
DROP INDEX IF EXISTS idx_currency_code;
CREATE INDEX idx_currency_code ON reference.currencies(code);

-- Insert currencies
INSERT INTO reference.currencies (currency_name, code) VALUES
('Afghan Afghani', 'AFN'),
('Euro', 'EUR'),
('United States Dollar', 'USD'),
('Albanian Lek', 'ALL'),
('Algerian Dinar', 'DZD'),
('Angolan Kwanza', 'AOA'),
('East Caribbean Dollar', 'XCD'),
('Argentine Peso', 'ARS'),
('Australian Dollar', 'AUD'),
('Bahamian Dollar', 'BSD'),
('Bahraini Dinar', 'BHD'),
('Bangladeshi Taka', 'BDT'),
('Barbadian Dollar', 'BBD'),
('Belarusian Ruble', 'BYN'),
('Bermudian Dollar', 'BMD'),
('Bhutanese Ngultrum', 'BTN'),
('Bolivian Boliviano', 'BOB'),
('Bosnia and Herzegovina Convertible Mark', 'BAM'),
('Botswana Pula', 'BWP'),
('Brazilian Real', 'BRL'),
('British Pound Sterling', 'GBP'),
('Brunei Dollar', 'BND'),
('Bulgarian Lev', 'BGN'),
('Burundian Franc', 'BIF'),
('Cabo Verdean Escudo', 'CVE'),
('Cambodian Riel', 'KHR'),
('Canadian Dollar', 'CAD'),
('Central African CFA Franc', 'XAF'),
('Chilean Peso', 'CLP'),
('Chinese Yuan', 'CNY'),
('Colombian Peso', 'COP'),
('Comorian Franc', 'KMF'),
('Congolese Franc', 'CDF'),
('Costa Rican Colón', 'CRC'),
('Croatian Kuna', 'HRK'),
('Cuban Peso', 'CUP'),
('Czech Koruna', 'CZK'),
('Danish Krone', 'DKK'),
('Djiboutian Franc', 'DJF'),
('Dominican Peso', 'DOP'),
('Dram', 'AMD'),
('Ecuadorian Dollar', 'USD'),
('Egyptian Pound', 'EGP'),
('El Salvadoran Colón', 'SVC'),
('Equatorial Guinean Franco', 'XAF'),
('Eritrean Nakfa', 'ERN'),
('Estonian Kroon', 'EEK'),
('Ethiopian Birr', 'ETB'),
('Fijian Dollar', 'FJD'),
('Gabonese Franc', 'XAF'),
('Gambian Dalasi', 'GMD'),
('Georgian Lari', 'GEL'),
('Ghanaian Cedi', 'GHS'),
('Gibraltar Pound', 'GIP'),
('Guatemalan Quetzal', 'GTQ'),
('Guinean Franc', 'GNF'),
('Guyanaese Dollar', 'GYD'),
('Haitian Gourde', 'HTG'),
('Honduran Lempira', 'HNL'),
('Hong Kong Dollar', 'HKD'),
('Hungarian Forint', 'HUF'),
('Icelandic Króna', 'ISK'),
('Indian Rupee', 'INR'),
('Indonesian Rupiah', 'IDR'),
('Iranian Rial', 'IRR'),
('Iraqi Dinar', 'IQD'),
('Israeli New Shekel', 'ILS'),
('Italian Lira', 'ITL'),
('Jamaican Dollar', 'JMD'),
('Japanese Yen', 'JPY'),
('Jordanian Dinar', 'JOD'),
('Kazakhstani Tenge', 'KZT'),
('Kenyan Shilling', 'KES'),
('Kiribati Dollar', 'KID'),
('Kuwaiti Dinar', 'KWD'),
('Kyrgystani Som', 'KGS'),
('Lao Kip', 'LAK'),
('Latvian Lats', 'LVL'),
('Lebanese Pound', 'LBP'),
('Lesotho Loti', 'LSL'),
('Liberian Dollar', 'LRD'),
('Libyan Dinar', 'LYD'),
('Lithuanian Litas', 'LTL'),
('Macanese Pataca', 'MOP'),
('Malagasy Ariary', 'MGA'),
('Malawian Kwacha', 'MWK'),
('Malaysian Ringgit', 'MYR'),
('Maldivian Rufiyaa', 'MVR'),
('Mauritanian Ouguiya', 'MRU'),
('Mauritian Rupee', 'MUR'),
('Mexican Peso', 'MXN'),
('Moldovan Leu', 'MDL'),
('Mongolian Tugrik', 'MNT'),
('Moroccan Dirham', 'MAD'),
('Mozambican Metical', 'MZN'),
('Myanmar Kyat', 'MMK'),
('Namibian Dollar', 'NAD'),
('Nauruan Dollar', 'NAD'),
('Nepalese Rupee', 'NPR'),
('Netherlands Antillean Guilder', 'ANG'),
('New Zealand Dollar', 'NZD'),
('Nicaraguan Córdoba', 'NIO'),
('Nigerian Naira', 'NGN'),
('North Korean Won', 'KPW'),
('Norwegian Krone', 'NOK'),
('Omani Rial', 'OMR'),
('Pakistani Rupee', 'PKR'),
('Panamanian Balboa', 'PAB'),
('Papua New Guinean Kina', 'PGK'),
('Paraguayan Guarani', 'PYG'),
('Peruvian Nuevo Sol', 'PEN'),
('Philippine Peso', 'PHP'),
('Polish Zloty', 'PLN'),
('Qatari Rial', 'QAR'),
('Romanian Leu', 'RON'),
('Russian Ruble', 'RUB'),
('Rwandan Franc', 'RWF'),
('Saint Helena Pound', 'SHP'),
('Samoan Tala', 'WST'),
('San Marinese Lira', 'SML'),
('Sao Tomean Dobra', 'STN'),
('Saudi Riyal', 'SAR'),
('Senegalese Franc', 'XOF'),
('Serbian Dinar', 'RSD'),
('Seychellois Rupee', 'SCR'),
('Sierra Leonean Leone', 'SLL'),
('Singapore Dollar', 'SGD'),
('Slovak Koruna', 'SKK'),
('Slovenian Tolar', 'SIT'),
('Solomon Islands Dollar', 'SBD'),
('Somali Shilling', 'SOS'),
('South African Rand', 'ZAR'),
('South Korean Won', 'KRW'),
('South Sudanese Pound', 'SSP'),
('Spanish Peseta', 'ESP'),
('Sri Lankan Rupee', 'LKR'),
('Sudanese Pound', 'SDG'),
('Surinamese Dollar', 'SRD'),
('Swedish Krona', 'SEK'),
('Swiss Franc', 'CHF'),
('Syrian Pound', 'SYP'),
('Tajikistani Somoni', 'TJS'),
('Tanzanian Shilling', 'TZS'),
('Thai Baht', 'THB'),
('Tongan Paʻanga', 'TOP'),
('Trinidad and Tobago Dollar', 'TTD'),
('Tunisian Dinar', 'TND'),
('Turkish Lira', 'TRY'),
('Turkmenistani Manat', 'TMT'),
('Ugandan Shilling', 'UGX'),
('Ukrainian Hryvnia', 'UAH'),
('United Arab Emirates Dirham', 'AED'),
('Uruguayan Peso', 'UYU'),
('Uzbekistani Som', 'UZS'),
('Vanuatu Vatu', 'VUV'),
('Venezuelan Bolívar', 'VES'),
('Vietnamese Dong', 'VND'),
('Yemeni Rial', 'YER'),
('Zambian Kwacha', 'ZMW'),
('Zimbabwean Dollar', 'ZWL');
