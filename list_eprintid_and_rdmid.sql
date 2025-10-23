--
-- List the rdmid and eprintid for records that are Technical Reports
--
SELECT
    json->>'id' AS id,
    elem->>'identifier' AS eprintid
FROM
    rdm_records_metadata,
    jsonb_array_elements(json->'metadata'->'identifiers') AS elem
WHERE
    json->'metadata'->'resource_type'->>'id' = 'publication-technicalnote'
    AND elem->>'scheme' = 'eprintid';
