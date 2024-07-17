-- +goose Up
-- +goose StatementBegin
CREATE TABLE event_types
(
    id                  BIGINT PRIMARY KEY,
    coins_amount        INTEGER DEFAULT 0,
    name                VARCHAR(50),
    tribe_points_amount INTEGER DEFAULT 0
);
CREATE SEQUENCE event_types_id_seq INCREMENT BY 1 MINVALUE 1;
ALTER TABLE event_types
    ALTER COLUMN id SET DEFAULT NEXTVAL('event_types_id_seq');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE event_types
-- +goose StatementEnd
