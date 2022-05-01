package repositories

import (
	"context"

	"github.com/ibidathoillah/majoo-test/lib/utils"

	"github.com/ibidathoillah/majoo-test/internal/domains/outlet/models"

	"github.com/ibidathoillah/majoo-test/internal/globals"
	"github.com/jmoiron/sqlx"
)

type mysql struct{}

func (p *mysql) GetOutletByMerchantIDs(ctx context.Context, merchantIDs []int64) ([]*models.Outlet, error) {
	var result []*models.Outlet
	var err error
	var rows *sqlx.Rows
	var arg map[string]interface{} = make(map[string]interface{})

	query := `
	SELECT merchant_id, id, outlet_name
	FROM Outlets
	WHERE merchant_id IN (` + utils.ArrayToString(merchantIDs, ",") + `)`

	rows, err = globals.GetQuery(ctx).NamedQueryxContext(ctx, query, arg)

	if nil != err {
		return nil, err
	}

	for rows.Next() {
		var model models.Outlet
		err = rows.StructScan(&model)
		if nil != err {
			return nil, err
		}
		result = append(result, &model)
	}

	_ = rows.Close()

	return result, nil

}

func NewOutletRepository() Interface {
	return &mysql{}
}
