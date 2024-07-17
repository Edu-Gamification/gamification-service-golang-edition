-- +goose Up
-- +goose StatementBegin
CREATE TABLE events
(
    id          BIGINT PRIMARY KEY,
    end_time    TIMESTAMP WITH TIME ZONE NOT NULL,
    start_time  TIMESTAMP WITH TIME ZONE NOT NULL,
    quote       INTEGER DEFAULT 10       NOT NULL,
    clan_only   BOOLEAN DEFAULT false    NOT NULL,
    description VARCHAR(255)             NOT NULL,
    title       VARCHAR(100)             NOT NULL,
    type        BIGINT REFERENCES event_types (id)
);
CREATE SEQUENCE events_id_seq INCREMENT BY 1 MINVALUE 1;
ALTER TABLE events
    ALTER COLUMN id SET DEFAULT NEXTVAL('events_id_seq');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events
-- +goose StatementEnd
