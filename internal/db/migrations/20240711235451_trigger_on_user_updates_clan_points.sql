-- +goose Up
-- +goose StatementBegin
CREATE
    OR REPLACE FUNCTION update_clan_points()
    RETURNS TRIGGER AS
$body$
BEGIN
    UPDATE clans
    SET points_amount = points_amount + (NEW.clan_points - OLD.clan_points)
    WHERE id = NEW.clan;
    RETURN NEW;
END;

$body$
    LANGUAGE plpgsql;

CREATE TRIGGER trg_update_clan_points
    AFTER UPDATE OF clan_points
    ON users
    FOR EACH ROW
EXECUTE PROCEDURE update_clan_points();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER trg_update_clan_points on users
-- +goose StatementEnd
