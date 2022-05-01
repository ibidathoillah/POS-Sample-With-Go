package repositories

import (
	"context"

	"github.com/ibidathoillah/majoo-test/internal/domains/outlet/models"
)

type Interface interface {
	GetOutletByMerchantIDs(ctx context.Context, merchantIDs []int64) ([]*models.Outlet, error)
}
