-- +goose Up
-- +goose StatementBegin
CREATE TABLE events_participants
(
    event_id       BIGINT,
    participant_id BIGINT,
    PRIMARY KEY (event_id, participant_id)
);
ALTER TABLE events_participants
    ADD CONSTRAINT fk_events_participants_event_id FOREIGN KEY (event_id) REFERENCES events (id),
    ADD CONSTRAINT fk_events_participants_participant_id FOREIGN KEY (participant_id) REFERENCES users (id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events_participants
-- +goose StatementEnd
