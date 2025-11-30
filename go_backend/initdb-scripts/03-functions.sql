
-- Utility function: automatically update 'updated_at' on any row update.
-- Use in triggers: sets NEW.updated_at = NOW() before every UPDATE.
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- --------------------------------------------------------------------
