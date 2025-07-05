-- This function auto-update the trackers of each spreadsheet whenever any of rows are updated
CREATE FUNCTION update_spreadsheet_status_counts()
RETURNS TRIGGER AS $$
BEGIN
  -- On DELETE
  IF (TG_OP = 'DELETE') THEN
    UPDATE spreadsheet
    SET
      checked_count = checked_count - (CASE WHEN OLD.status = 'checked' THEN 1 ELSE 0 END),
      unchecked_count = unchecked_count - (CASE WHEN OLD.status = 'unchecked' THEN 1 ELSE 0 END),
      absent_count = absent_count - (CASE WHEN OLD.status = 'absent' THEN 1 ELSE 0 END)
    WHERE id = OLD.spreadsheet_key;

    RETURN OLD;
  END IF;

  -- On INSERT
  IF (TG_OP = 'INSERT') THEN
    UPDATE spreadsheet
    SET
      checked_count = checked_count + (CASE WHEN NEW.status = 'checked' THEN 1 ELSE 0 END),
      unchecked_count = unchecked_count + (CASE WHEN NEW.status = 'unchecked' THEN 1 ELSE 0 END),
      absent_count = absent_count + (CASE WHEN NEW.status = 'absent' THEN 1 ELSE 0 END)
    WHERE id = NEW.spreadsheet_key;

    RETURN NEW;
  END IF;

  -- On UPDATE
  IF (TG_OP = 'UPDATE') THEN
    IF (NEW.status != OLD.status) THEN
      UPDATE spreadsheet
      SET
        checked_count = checked_count
                          + (CASE WHEN NEW.status = 'checked' THEN 1 ELSE 0 END)
                          - (CASE WHEN OLD.status = 'checked' THEN 1 ELSE 0 END),
        unchecked_count = unchecked_count
                          + (CASE WHEN NEW.status = 'unchecked' THEN 1 ELSE 0 END)
                          - (CASE WHEN OLD.status = 'unchecked' THEN 1 ELSE 0 END),
        absent_count = absent_count
                          + (CASE WHEN NEW.status = 'absent' THEN 1 ELSE 0 END)
                          - (CASE WHEN OLD.status = 'absent' THEN 1 ELSE 0 END)
      WHERE id = NEW.spreadsheet_key;
    END IF;

    RETURN NEW;
  END IF;

  RETURN NULL;
END;
$$ LANGUAGE plpgsql;