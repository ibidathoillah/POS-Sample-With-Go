package repositories

import (
	"context"

	"github.com/ibidathoillah/majoo-test/internal/domains/merchant/models"

	"github.com/ibidathoillah/majoo-test/internal/globals"
	"github.com/jmoiron/sqlx"
)

type mysql struct{}

func (p *mysql) GetMerchantsByUserID(ctx context.Context, userID int64) ([]*models.Merchant, error) {
	var result []*models.Merchant
	var err error
	var rows *sqlx.Rows
	var arg map[string]interface{} = make(map[string]interface{})

	query := `
	SELECT id, user_id, merchant_name
	FROM Merchants
	WHERE user_id IN (:user_id)`

	arg["user_id"] = userID

	rows, err = globals.GetQuery(ctx).NamedQueryxContext(ctx, query, arg)

	if nil != err {
		return nil, err
	}

	for rows.Next() {
		var model models.Merchant
		err = rows.StructScan(&model)
		if nil != err {
			return nil, err
		}
		result = append(result, &model)
	}

	_ = rows.Close()

	return result, nil

}

func NewMerchantRepository() Interface {
	return &mysql{}
}
