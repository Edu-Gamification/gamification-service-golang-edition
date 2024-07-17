-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_role
(
    user_id BIGINT NOT NULL REFERENCES users(id),
    role_id BIGINT NOT NULL REFERENCES roles(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_role;
-- +goose StatementEnd
