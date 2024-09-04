-- Insert IDs into core_entity_flag for all modules based on the code
INSERT INTO configurations.core_entity_flag (entity_id, module_id, status)
SELECT ce.entity_id,
       CASE
           WHEN ce.code LIKE 'authentication.%' THEN 1
           WHEN ce.code LIKE 'configurations.%' THEN 2
           WHEN ce.code LIKE 'references.%' THEN 3
           WHEN ce.code LIKE 'company.%' THEN 4
           WHEN ce.code LIKE 'employees.%' THEN 5
           WHEN ce.code LIKE 'reports.%' THEN 6
           ELSE NULL
       END AS module_id,
       'Disabled' AS status
FROM configurations.core_entities ce
WHERE ce.code LIKE 'authentication.%'
   OR ce.code LIKE 'configurations.%'
   OR ce.code LIKE 'references.%'
   OR ce.code LIKE 'company.%'
   OR ce.code LIKE 'employees.%'
   OR ce.code LIKE 'reports.%';


-- Update flags for configurations, set 'Enabled'
UPDATE configurations.core_entity_flag SET status = 'Enabled' WHERE module_id = 2;

