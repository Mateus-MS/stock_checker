SELECT 
	c.sku,
	c.piece,
	c.status
FROM spreadsheet s
LEFT JOIN spreadsheet_row c ON c.spreadsheet_key = s.id
WHERE s.id = '21bb5c0e-66c8-4ec5-bb09-e659e379a105'
LIMIT 2;

SELECT * FROM spreadsheet;

TRUNCATE spreadsheet;
TRUNCATE spreadsheet_row;

DROP TABLE spreadsheet;
DROP TABLE spreadsheet_row;