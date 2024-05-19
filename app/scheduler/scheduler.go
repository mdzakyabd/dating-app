package scheduler

import (
	"time"

	"github.com/mdzakyabd/dating-app/app/repository"
)

type Scheduler struct {
	userRepo repository.UserRepository
}

func NewScheduler(userRepo repository.UserRepository) *Scheduler {
	return &Scheduler{userRepo}
}

func (s *Scheduler) Start() {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.checkExpiredSubscriptions()
			}
		}
	}()
}

func (s *Scheduler) checkExpiredSubscriptions() {
	users, err := s.userRepo.FindAllPremiumUsers()
	if err != nil {
		// Handle error
		return
	}

	now := time.Now()
	for _, user := range users {
		if user.PremiumExpiryTime.Before(now) {
			s.userRepo.UpdatePremiumStatus(user.ID, false, time.Time{})
		}
	}
}
