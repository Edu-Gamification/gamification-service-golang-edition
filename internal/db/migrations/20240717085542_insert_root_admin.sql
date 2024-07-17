-- +goose Up
-- +goose StatementBegin
INSERT INTO users (name, surname, password, email, active)
VALUES ('admin', 'admin', 'admin', 'admin@edu.com', true);

INSERT INTO roles (name) VALUES ('ROLE_USER');
INSERT INTO roles (name) VALUES ('ROLE_ADMIN');

INSERT INTO user_role (user_id, role_id) VALUES (currval('users_id_seq'), 1);
INSERT INTO user_role (user_id, role_id) VALUES (currval('users_id_seq'), 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users
-- +goose StatementEnd
