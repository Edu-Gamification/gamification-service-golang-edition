package repository

import (
	"GamificationEducation/internal/domain"
	errors2 "GamificationEducation/internal/errors"
	"database/sql"
	"errors"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (eventRepository *EventRepository) FindAllEventTypes() ([]domain.EventType, error) {
	var types []domain.EventType

	rows, err := eventRepository.db.Query("select * from event_types")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var eType domain.EventType
		err := rows.Scan(&eType.Id, &eType.Name, &eType.CoinsAmount, &eType.ClanPointsAmount)
		if err != nil {
			return nil, err
		}
		types = append(types, eType)
	}
	return types, nil
}

func (eventRepository *EventRepository) FindEventTypeByName(name string) (domain.EventType, error) {
	var eType domain.EventType
	if err := eventRepository.db.QueryRow("select * from event_types where name = $1", name).Scan(&eType.Id,
		&eType.Name, &eType.CoinsAmount, &eType.ClanPointsAmount); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.EventType{}, errors2.ErrNotFound
		}
		return domain.EventType{}, err
	}
	return eType, nil
}

func (eventRepository *EventRepository) FindEventsByType(eType domain.EventType) ([]domain.Event, error) {
	var events []domain.Event

	rows, err := eventRepository.db.Query("select * from events where type = $1", eType.Id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var event domain.Event
		err := rows.Scan(&event.Id, &event.StartTime, &event.EndTime, &event.Quote, &event.ClanOnly, &event.Description,
			&event.Title, &event.Type)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
