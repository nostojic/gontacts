-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_name VARCHAR(255) UNIQUE NOT NULL,
  user_email VARCHAR(255) UNIQUE NOT NULL,
  password TEXT NOT NULL,
  date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE groups (
  group_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  group_name VARCHAR(255) UNIQUE NOT NULL,
  user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
  date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_group_name ON groups(group_name);

CREATE TABLE contacts (
  contact_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  contact_first_name VARCHAR(255) NOT NULL,
  contact_last_name VARCHAR(255),
  contact_screen_name VARCHAR(255),
  user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
  date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_contact_first_name ON contacts(contact_first_name);

CREATE TABLE channels (
  channel_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  channel_name VARCHAR(255) NOT NULL,
  channel_icon_path TEXT,
  channel_description TEXT,
  user_id UUID REFERENCES users(user_id) ON DELETE CASCADE,
  date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (user_id, channel_name)
);

CREATE INDEX idx_channel_name ON channels(channel_name);

CREATE TABLE contact_channels (
  contact_channel_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  channel_id UUID NOT NULL REFERENCES channels(channel_id) ON DELETE CASCADE,
  contact_channel_value VARCHAR(255) NOT NULL,
  date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  date_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_contact_channels_channel_id ON contact_channels(channel_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
DROP TABLE groups
DROP TABLE contacts
DROP TABLE channels
DROP TABLE contact_channels;
-- +goose StatementEnd
