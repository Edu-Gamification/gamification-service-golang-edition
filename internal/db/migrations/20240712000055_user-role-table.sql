-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles
(
    id   BIGINT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);
CREATE SEQUENCE roles_id_seq START 1 INCREMENT BY 1;
ALTER TABLE roles
    ALTER COLUMN id SET DEFAULT NEXTVAL('roles_id_seq');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_role;
-- +goose StatementEnd
