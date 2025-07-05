CREATE TABLE spreadsheet(
	ID UUID PRIMARY KEY NOT NULL,
	name TEXT DEFAULT 'Planilha sem nome',

    -- Trackers
	unchecked_count INTEGER DEFAULT 0,
	checked_count INTEGER DEFAULT 0,
	absent_count INTEGER DEFAULT 0
);

CREATE TYPE spreadsheet_row_status_enum AS ENUM ('unchecked', 'checked', 'absent');
CREATE TABLE spreadsheet_row (
	-- Identifier for the specifique product (i,e: 244448001)
    SKU TEXT NOT NULL,

	-- Identifier for the product group      (i,e: 765474)
	piece TEXT NOT NULL,

	-- Product status on the list            (i,e: absent)
    status spreadsheet_row_status_enum NOT NULL DEFAULT 'unchecked',
    spreadsheet_key UUID NOT NULL,

    PRIMARY KEY (spreadsheet_key, SKU)
);
CREATE TRIGGER spreadsheet_row_status_trigger
AFTER INSERT OR UPDATE OR DELETE ON spreadsheet_row
FOR EACH ROW EXECUTE FUNCTION update_spreadsheet_status_counts();