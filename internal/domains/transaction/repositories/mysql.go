package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ibidathoillah/majoo-test/lib/utils"

	"github.com/ibidathoillah/majoo-test/internal/domains/transaction/models"
	"github.com/ibidathoillah/majoo-test/internal/globals"
)

type mysql struct{}

func (u *mysql) FindAllOmzetReport(ctx context.Context, params models.FindAllOmzetReport) ([]*models.TransactionOmzet, error) {
	if len(params.MerchantIDs) == 0 {
		return []*models.TransactionOmzet{}, nil
	}

	firstDate := time.Date(params.Year, time.Month(params.Month), 1, 0, 0, 0, 0, time.UTC)
	endDate := int64(utils.EndOfMonth(firstDate).Day())

	queryDay := []string{}

	if params.OutletHash == nil {
		for _, merchantID := range params.MerchantIDs {
			for i := int64(1); i <= endDate; i++ {
				queryDay = append(queryDay, fmt.Sprintf("SELECT %d as day, %d as qd_merchant_id", i, merchantID))
			}
		}
	} else {
		for _, outlet := range *params.OutletHash {
			for i := int64(1); i <= endDate; i++ {
				queryDay = append(queryDay, fmt.Sprintf("SELECT %d as day, %d as qd_merchant_id, %d as qd_outlet_id", i, outlet.MerchantID, outlet.ID))
			}
		}
	}

	result := []*models.TransactionOmzet{}
	var arg map[string]interface{} = make(map[string]interface{})
	var groupDate = `qd.day, qd.qd_merchant_id`
	var selectQ string = `SELECT`
	var onQ string = `ON DAY(t.created_at) = qd.day AND qd.qd_merchant_id = t.merchant_id AND MONTH(created_at) = :month AND YEAR(created_at) = :year`
	var groupByQ string = `GROUP BY ` + groupDate

	limit := params.Limit
	offset := limit * (params.Page - int64(1))
	limitQ := fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)

	arg["month"] = params.Month
	arg["year"] = params.Year

	dateformat := fmt.Sprintf(`STR_TO_DATE(CONCAT(qd.day,"-%d-%d")`, params.Month, params.Year) + `, "%d-%m-%Y")`

	if len(params.MerchantIDs) > 0 {
		selectQ += ` qd.qd_merchant_id merchant_id, IFNULL(SUM(bill_total),0) omzet, ` + dateformat + ` date, qd.day `
		onQ += ` AND merchant_id IN (` + utils.ArrayToString(params.MerchantIDs, ",") + `)`

		if len(params.OutletIDs) > 0 {
			selectQ += `, qd.qd_outlet_id outlet_id `
			groupByQ += `, qd.qd_outlet_id`
			onQ += ` AND outlet_id IN (` + utils.ArrayToString(params.OutletIDs, ",") + `)`
			onQ += ` AND qd.qd_outlet_id = t.outlet_id`
		}
	}

	query := selectQ + ` FROM (` + strings.Join(queryDay, " UNION ALL ") + `) as qd LEFT JOIN Transactions t ` +
		onQ + ` ` + groupByQ + ` ORDER BY ` + groupDate + ` ASC ` + limitQ
	rows, err := globals.GetQuery(ctx).NamedQueryxContext(ctx, query, arg)

	fmt.Println(rows)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var model models.TransactionOmzet
		err = rows.StructScan(&model)
		if nil != err {
			return nil, err
		}
		result = append(result, &model)
	}

	if err != nil {
		return nil, err
	}

	return result, err
}

func NewTransactionRepository() Interface {
	return &mysql{}
}
