-- +goose Up
-- +goose StatementBegin
CREATE FUNCTION date_updated ()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
AS $$
BEGIN
  NEW.date_updated = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS date_updated;
-- +goose StatementEnd
