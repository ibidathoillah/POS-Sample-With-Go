package container

import (
	"github.com/go-kit/kit/log"
	"github.com/ibidathoillah/majoo-test/internal/domains/auth"
	"github.com/ibidathoillah/majoo-test/internal/domains/merchant"
	merchantRepository "github.com/ibidathoillah/majoo-test/internal/domains/merchant/repositories"
	"github.com/ibidathoillah/majoo-test/internal/domains/outlet"
	outletRepository "github.com/ibidathoillah/majoo-test/internal/domains/outlet/repositories"
	"github.com/ibidathoillah/majoo-test/internal/domains/transaction"
	transactionRepository "github.com/ibidathoillah/majoo-test/internal/domains/transaction/repositories"
	"github.com/ibidathoillah/majoo-test/internal/domains/user"
	userRepository "github.com/ibidathoillah/majoo-test/internal/domains/user/repositories"
)

type Container struct {
	AuthService        auth.UseCase
	UserService        user.UseCase
	TransactionService transaction.UseCase
}

func New(
	logger log.Logger,
) Container {

	userService := user.NewUserService(logger, userRepository.NewUserRepository())
	authService := auth.NewAuthService(logger, userService)
	outletService := outlet.NewOutletService(logger, outletRepository.NewOutletRepository())
	merchantService := merchant.NewMerchantService(logger, merchantRepository.NewMerchantRepository())

	transactionService := transaction.NewTransactionService(
		logger,
		transactionRepository.NewTransactionRepository(),
		outletService,
		merchantService,
	)

	return Container{
		UserService:        userService,
		AuthService:        authService,
		TransactionService: transactionService,
	}
}
