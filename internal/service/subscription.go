package service

import (
	"context"
	"fmt"

	"rutube-task/internal/repository"
)

type SubscriptionService struct {
	rep *repository.Repository
}

func NewSubscriptionService(rep *repository.Repository) *SubscriptionService {
	return &SubscriptionService{rep: rep}
}

func (s *SubscriptionService) Subscribe(ctx context.Context, userID int, name string) error {
	err := s.rep.SubscriptionRepositoryInterface.SubscribeDB(ctx, userID, name)
	if err != nil {
		return fmt.Errorf("subscribe: %w ", err)
	}

	return nil
}

func (s *SubscriptionService) Unsubscribe(ctx context.Context, userID int, name string) error {
	err := s.rep.SubscriptionRepositoryInterface.UnsubscribeDB(ctx, userID, name)
	if err != nil {
		return fmt.Errorf("Unsubscribe: %w ", err)
	}
	return nil
}
