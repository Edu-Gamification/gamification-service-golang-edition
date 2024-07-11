-- +goose Up
-- +goose StatementBegin
CREATE TABLE clans
(
    id            BIGINT PRIMARY KEY,
    name          VARCHAR(50) NOT NULL,
    points_amount INTEGER DEFAULT 0
);
CREATE SEQUENCE clan_id_seq START 1 INCREMENT BY 1;
ALTER TABLE clans
    ALTER COLUMN id SET DEFAULT NEXTVAL('clan_id_seq')
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clans
-- +goose StatementEnd
