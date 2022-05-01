package merchant

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/ibidathoillah/majoo-test/internal/domains/merchant/models"
	"github.com/ibidathoillah/majoo-test/internal/domains/merchant/repositories"
)

type service struct {
	actor      string
	logger     log.Logger
	repository repositories.Interface
}

func (s *service) GetMerchantsByUserID(ctx context.Context, userID int64) ([]*models.Merchant, error) {
	return s.repository.GetMerchantsByUserID(ctx, userID)
}
func NewMerchantService(
	logger log.Logger,
	repository repositories.Interface,
) UseCase {
	service := service{
		actor:      "MERCHANT",
		logger:     nil,
		repository: repository,
	}

	service.logger = log.With(logger, "ACTOR", service.actor)

	return &	
}
