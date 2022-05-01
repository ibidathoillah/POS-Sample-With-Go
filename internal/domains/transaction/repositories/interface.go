package repositories

import (
	"context"

	"github.com/ibidathoillah/majoo-test/internal/domains/transaction/models"
)

type Interface interface {
	FindAllOmzetReport(ctx context.Context, params models.FindAllOmzetReport) ([]*models.TransactionOmzet, error)
}
