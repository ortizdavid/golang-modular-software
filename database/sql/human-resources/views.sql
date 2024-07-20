-- Schema: human_resources

-- view: view_personal_data
DROP VIEW IF EXISTS human_resources.view_personal_data;
CREATE VIEW human_resources.view_personal_data AS 
SELECT pe.person_id, pe.unique_id,
    pe.first_name, pe.last_name,
    pe.birth_date, pe.gender,
	pe.identification_number,
	TO_CHAR(pe.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(pe.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
	it.type_id, it.type_name,
	ms.status_id, ms.status_name
FROM human_resources.person pe
JOIN human_resources.identification_types it ON(it.type_id = pe.identification_type_id)
JOIN human_resources.marital_statuses ms ON(ms.status_id = pe.marital_status_id)
ORDER BY pe.created_at DESC;


-- view: view_contact_data
DROP VIEW IF EXISTS human_resources.view_contact_data;
CREATE VIEW human_resources.view_contact_data AS 
SELECT co.contact_id, co.unique_id,
	co.phone, co.email,
	TO_CHAR(co.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(co.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
	pe.person_id, pe.first_name,
	pe.last_name, pe.identification_number,
	ct.type_id, ct.type_name
FROM human_resources.contacts co 
JOIN human_resources.person pe ON(pe.person_id = co.person_id)
JOIN human_resources.contact_types ct ON(ct.type_id = co.contact_type_id)
ORDER BY co.created_at DESC;


-- view: view_address_data
DROP VIEW IF EXISTS human_resources.view_address_data;
CREATE VIEW human_resources.view_address_data AS 
SELECT ad.address_id, ad.unique_id,
	ad.state, ad.city,
	ad.district, ad.postal_code,
	TO_CHAR(ad.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(ad.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
	pe.person_id, pe.first_name,
	pe.last_name, pe.identification_number
FROM human_resources.addresses ad
JOIN human_resources.person pe ON(pe.person_id = ad.person_id)
ORDER BY ad.created_at DESC;


-- view: view_document_data
DROP VIEW IF EXISTS human_resources.view_document_data;
CREATE VIEW human_resources.view_document_data AS 
SELECT doc.document_id, doc.unique_id,
	doc.file_name,
	TO_CHAR(doc.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(doc.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
	pe.person_id, pe.first_name,
	pe.last_name, pe.identification_number
FROM human_resources.documents doc
JOIN human_resources.person pe ON(pe.person_id = doc.person_id)
JOIN human_resources.document_types dt ON(dt.type_id = doc.document_type_id)
ORDER BY doc.created_at DESC;


