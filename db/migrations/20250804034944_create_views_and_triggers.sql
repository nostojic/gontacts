-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS v1;

CREATE OR REPLACE VIEW v1.users AS SELECT * FROM users;
CREATE OR REPLACE VIEW v1.groups AS SELECT * FROM groups;
CREATE OR REPLACE VIEW v1.contacts AS SELECT * FROM contacts;
CREATE OR REPLACE VIEW v1.channels AS SELECT * FROM channels;
CREATE OR REPLACE VIEW v1.contact_channels AS SELECT * FROM contact_channels;

CREATE TRIGGER trigger_update_users BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION date_updated();
CREATE TRIGGER trigger_update_groups BEFORE UPDATE ON groups FOR EACH ROW EXECUTE FUNCTION date_updated();
CREATE TRIGGER trigger_update_contacts BEFORE UPDATE ON contacts FOR EACH ROW EXECUTE FUNCTION date_updated();
CREATE TRIGGER trigger_update_channels BEFORE UPDATE ON channels FOR EACH ROW EXECUTE FUNCTION date_updated();
CREATE TRIGGER trigger_update_contact_channels BEFORE UPDATE ON contact_channels FOR EACH ROW EXECUTE FUNCTION date_updated();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW v1.users;
DROP VIEW v1.groups;
DROP VIEW v1.contacts;
DROP VIEW v1.channels;
DROP VIEW v1.contact_channels;

DROP TRIGGER trigger_update_users ON users;
DROP TRIGGER trigger_update_groups ON groups;
DROP TRIGGER trigger_update_contacts ON contacts;
DROP TRIGGER trigger_update_channels ON channels;
DROP TRIGGER trigger_update_contact_channels ON contact_channels;
-- +goose StatementEnd
