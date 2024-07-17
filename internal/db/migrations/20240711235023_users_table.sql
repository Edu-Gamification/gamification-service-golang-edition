-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id          BIGINT PRIMARY KEY,
    name        VARCHAR(50)  NOT NULL,
    surname     VARCHAR(100) NOT NULL,
    patronymic  VARCHAR(100),
    password    VARCHAR(255) NOT NULL,
    email       VARCHAR(100) NOT NULL UNIQUE,
    active      BOOLEAN      NOT NULL,
    clan_points INTEGER      NOT NULL DEFAULT 0,
    coins       INTEGER      NOT NULL DEFAULT 0,
    clan        BIGINT REFERENCES clans (id)
);
CREATE SEQUENCE users_id_seq INCREMENT BY 1 MINVALUE 1;
ALTER TABLE users
    ALTER COLUMN id SET DEFAULT NEXTVAL('users_id_seq');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd
