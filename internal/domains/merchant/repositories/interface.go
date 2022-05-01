package repositories

import (
	"context"

	"github.com/ibidathoillah/majoo-test/internal/domains/merchant/models"
)

type Interface interface {
	GetMerchantsByUserID(ctx context.Context, userID int64) ([]*models.Merchant, error)
}
