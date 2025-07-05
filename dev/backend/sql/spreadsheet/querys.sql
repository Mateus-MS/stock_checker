-- =========== Select rows query =========== --

SELECT 
    c.sku,
    c.piece,
    c.status
FROM spreadsheet s
LEFT JOIN spreadsheet_column c ON c.spreadsheet_key = s.id
WHERE s.id = 'c21ea6f0-5b9a-46fc-9ec9-f8b168af1c49'
ORDER BY c.id;

-- =========== Select spreadsheet query =========== --

SELECT 
    id AS spreadsheet_id,
    name AS spreadsheet_name
FROM spreadsheet
WHERE id = 'c21ea6f0-5b9a-46fc-9ec9-f8b168af1c49';

-- =========== Update the row =========== --

UPDATE spreadsheet_row c
SET status = 'checked'
FROM spreadsheet s
WHERE c.spreadsheet_key = s.id
  AND s.id = 'da81b20e-95b8-46fe-91d9-58ba9105ee83'
  AND c.sku = '242455704';