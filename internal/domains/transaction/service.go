package transaction

import (
	"context"

	"github.com/ibidathoillah/majoo-test/internal/domains/merchant"
	"github.com/ibidathoillah/majoo-test/internal/domains/outlet"

	"github.com/go-kit/kit/log"
	merchantModels "github.com/ibidathoillah/majoo-test/internal/domains/merchant/models"
	outletModels "github.com/ibidathoillah/majoo-test/internal/domains/outlet/models"
	"github.com/ibidathoillah/majoo-test/internal/domains/transaction/models"
	"github.com/ibidathoillah/majoo-test/internal/domains/transaction/repositories"
)

type service struct {
	actor           string
	logger          log.Logger
	repository      repositories.Interface
	outletService   outlet.UseCase
	merchantService merchant.UseCase
}

func (s *service) FindAllOmzetReport(ctx context.Context, params models.FindAllOmzetReport) ([]*models.TransactionOmzet, error) {

	merchantIds, merchantHash, err := s.getMerchantByUserId(ctx, *params.UserID)
	if err != nil {
		return nil, err
	}
	params.MerchantIDs = merchantIds

	var outletHash map[int64]*outletModels.Outlet
	if models.OmzetGrouBy(params.GroupBy) == models.OMZET_GROUP_BY_OUTLET {
		var outletIds []int64
		outletIds, outletHash, err = s.getOutletByMerchantIDs(ctx, merchantIds)
		if err != nil {
			return nil, err
		}
		params.OutletIDs = outletIds
		params.OutletHash = &outletHash
	}

	omzetList, err := s.repository.FindAllOmzetReport(ctx, params)
	if err != nil {
		return nil, err
	}

	for _, omzet := range omzetList {
		omzet.MerchantName = merchantHash[omzet.MerchantID].MerchantName
		if omzet.OutletID != nil {
			omzet.OutletName = &outletHash[*omzet.OutletID].OutletName
		}

	}
	return omzetList, nil
}

func (s *service) getMerchantByUserId(ctx context.Context, userID int64) ([]int64, map[int64]*merchantModels.Merchant, error) {
	merchantList, err := s.merchantService.GetMerchantsByUserID(ctx, userID)
	if err != nil {
		return nil, nil, err
	}

	merchantIds := []int64{}
	merchantHash := make(map[int64]*merchantModels.Merchant)
	for _, merchant := range merchantList {
		merchantIds = append(merchantIds, merchant.ID)
		merchantHash[merchant.ID] = merchant
	}

	return merchantIds, merchantHash, nil
}

func (s *service) getOutletByMerchantIDs(ctx context.Context, merchantIDs []int64) ([]int64, map[int64]*outletModels.Outlet, error) {
	outletList, err := s.outletService.GetOutletByMerchantIDs(ctx, merchantIDs)
	if err != nil {
		return nil, nil, err
	}

	outletIds := []int64{}
	outletHash := make(map[int64]*outletModels.Outlet)
	for _, outlet := range outletList {
		outletIds = append(outletIds, outlet.ID)
		outletHash[outlet.ID] = outlet
	}

	return outletIds, outletHash, nil
}

func NewTransactionService(
	logger log.Logger,
	repository repositories.Interface,
	outletService outlet.UseCase,
	merchantService merchant.UseCase,
) UseCase {
	service := service{
		actor:           "TRANSACTION",
		logger:          nil,
		repository:      repository,
		outletService:   outletService,
		merchantService: merchantService,
	}

	service.logger = log.With(logger, "ACTOR", service.actor)

	return &service
}
