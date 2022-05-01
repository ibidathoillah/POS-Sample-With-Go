package transaction

import (
	"context"

	"github.com/ibidathoillah/majoo-test/internal/domains/transaction/models"
)

type UseCase interface {
	FindAllOmzetReport(ctx context.Context, params models.FindAllOmzetReport) ([]*models.TransactionOmzet, error)
}
