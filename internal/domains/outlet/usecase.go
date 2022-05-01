package outlet

import (
	"context"

	"github.com/ibidathoillah/majoo-test/internal/domains/outlet/models"
)

type UseCase interface {
	GetOutletByMerchantIDs(ctx context.Context, merchantIDs []int64) ([]*models.Outlet, error)
}
