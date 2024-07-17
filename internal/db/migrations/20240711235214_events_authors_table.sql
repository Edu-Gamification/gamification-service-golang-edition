-- +goose Up
-- +goose StatementBegin
CREATE TABLE events_authors
(
    event_id  BIGINT,
    author_id BIGINT,
    PRIMARY KEY (event_id, author_id)
);

ALTER TABLE events_authors
    ADD CONSTRAINT fk_event_authors_event_id FOREIGN KEY (event_id) REFERENCES events (id),
    ADD CONSTRAINT fk_event_authors_author_id FOREIGN KEY (author_id) REFERENCES users (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events_authors
-- +goose StatementEnd
