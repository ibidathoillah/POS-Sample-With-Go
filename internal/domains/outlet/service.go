package outlet

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/ibidathoillah/majoo-test/internal/domains/outlet/models"
	"github.com/ibidathoillah/majoo-test/internal/domains/outlet/repositories"
)

type service struct {
	actor      string
	logger     log.Logger
	repository repositories.Interface
}

func (s *service) GetOutletByMerchantIDs(ctx context.Context, merchantIDs []int64) ([]*models.Outlet, error) {
	return s.repository.GetOutletByMerchantIDs(ctx, merchantIDs)
}
func NewOutletService(
	logger log.Logger,
	repository repositories.Interface,
) UseCase {
	service := service{
		actor:      "OUTLET",
		logger:     nil,
		repository: repository,
	}

	service.logger = log.With(logger, "ACTOR", service.actor)

	return &service
}
