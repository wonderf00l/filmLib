package actor

import (
	"context"

	entity "github.com/wonderf00l/filmLib/internal/entity/actor"
)

var _ Service = (*actorService)(nil)

func (s *actorService) AddActor(ctx context.Context, actor entity.Actor) error {
	if err := validateActorData(actor); err != nil {
		return err
	}

	if err := s.repo.AddActorData(ctx, actor); err != nil {
		return err
	}

	return nil
}

func (s *actorService) GetActor(ctx context.Context, actorID int) (*entity.Actor, error) {
	actor, err := s.repo.GetActorData(ctx, actorID)
	if err != nil {
		return nil, err
	}

	return actor, nil
}

// name, gender can't be empty or "" - in delivery
func (s *actorService) UpdateActorData(ctx context.Context, actorID int, updFields map[string]any) error {
	_, err := s.repo.GetActorData(ctx, actorID)
	if err != nil {
		return err
	}
	// if null --> ok no change --> set old
	// проверка явного "" в delivery
	// тогда тут "" это null --> перезапись --> неявно, полагаемся на валидацию delivery
	// if updActor.DateOfBirth.IsZero() {
	// 	updActor.DateOfBirth = actor.DateOfBirth
	// }

	if err := s.repo.UpdateActorData(ctx, actorID, updFields); err != nil {
		return err
	}

	return nil
}

func (s *actorService) DeleteActorData(ctx context.Context, actorID int) error {
	_, err := s.repo.GetActorData(ctx, actorID)
	if err != nil {
		return err
	}

	if err := s.repo.DeleteActorData(ctx, actorID); err != nil {
		return err
	}

	return nil
}
