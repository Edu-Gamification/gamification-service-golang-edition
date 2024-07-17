package service

import "GamificationEducation/internal/domain"

type EventRepository interface {
	FindAllEventTypes() ([]domain.EventType, error)
	FindEventTypeByName(name string) (domain.EventType, error)
	FindEventsByType(eType domain.EventType) ([]domain.Event, error)
	FindEventById(id int64) (domain.Event, error)
	FindEventAuthorsID(eventId int64) ([]int64, error)
	FindEventParticipantsID(eventId int64) ([]int64, error)
}

type EventService struct {
	eventRepository EventRepository
	userService     UserService
}

func NewEventService(repository EventRepository, service UserService) *EventService {
	return &EventService{eventRepository: repository, userService: service}
}

func (eventService *EventService) GetEventAuthors(eventId int64) ([]domain.User, error) {
	var authors []domain.User

	authorsID, err := eventService.eventRepository.FindEventAuthorsID(eventId)
	if err != nil {
		return nil, err
	}
	for _, id := range authorsID {
		author, err := eventService.userService.GetById(id)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (eventService *EventService) GetEventParticipants(eventId int64) ([]domain.User, error) {
	var participants []domain.User

	participantsID, err := eventService.eventRepository.FindEventParticipantsID(eventId)
	if err != nil {
		return nil, err
	}
	for _, id := range participantsID {
		participant, err := eventService.userService.GetById(id)
		if err != nil {
			return nil, err
		}
		participants = append(participants, participant)
	}
	return participants, nil
}

func (eventService *EventService) GetById(id int64) (domain.Event, error) {
	event, err := eventService.eventRepository.FindEventById(id)
	if err != nil {
		return domain.Event{}, err
	}
	authors, err := eventService.GetEventAuthors(id)
	if err != nil {
		return domain.Event{}, err
	}
	participants, err := eventService.GetEventParticipants(id)
	if err != nil {
		return domain.Event{}, err
	}
	event.Authors = authors
	event.Participants = participants
	return event, nil
}
