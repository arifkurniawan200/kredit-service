package usecase

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/sync/errgroup"
	"kredit-service/internal/consts"
	models "kredit-service/internal/model"
	"kredit-service/internal/repository"
	"kredit-service/internal/utils"
	"time"
)

type UserHandler struct {
	u repository.UserRepository
	t repository.TransactionRepository
}

func (u UserHandler) ListRequestLoan(ctx echo.Context) ([]models.CustomerLoan, error) {
	return u.u.CustomerLoanRequest(consts.LoanRequestStatusRequested)
}

func (u UserHandler) RequestLoan(ctx echo.Context, loan models.LoanRequestParam) error {
	now := time.Now()
	tenor, err := u.t.GetTenorByID(loan.TenorID)
	if err != nil {
		return err
	}

	userInfo, err := u.u.GetUserByEmail(loan.Email)
	if err != nil {
		return err
	}

	err = u.u.RequestLoan(models.CustomerLoan{
		CustomerID: loan.CustomerID,
		Tenor:      loan.TenorID,
		LoanDate:   now,
		LoanAmount: tenor.Value * userInfo.Salary,
		Status:     consts.LoanRequestStatusRequested,
	})
	if err != nil {
		return err
	}
	return err
}

func (u UserHandler) GetUserLimit(ctx echo.Context, userID int) (models.LimitInformation, error) {
	limit, err := u.u.GetUserLimit(userID)
	if err != nil {
		return models.LimitInformation{}, err
	}
	return LoanToLimitInfo(limit), nil
}

func NewUserUsecase(u repository.UserRepository, t repository.TransactionRepository) UserUcase {
	return &UserHandler{u, t}
}

func (u UserHandler) GetUserInfoByEmail(ctx echo.Context, email string) (models.Customer, error) {
	return u.u.GetUserByEmail(email)
}

const (
	secret = "abc&1*~#^2^#s0^=)^^7%b34"
)

func (u UserHandler) RegisterCustomer(ctx echo.Context, c models.CustomerParam) error {
	var (
		err error
		g   errgroup.Group
	)

	g.Go(func() error {
		// hash password
		c.Password, err = utils.HashPassword(c.Password)
		if err != nil {
			log.Errorf("error when hash password ")
			return err
		}
		return err
	})

	g.Go(func() error {
		//encrypt sensitive data
		c.NIK, err = utils.Encrypt(c.NIK, secret)
		if err != nil {
			log.Errorf("error when hash password ")
			return err
		}
		return err
	})

	g.Go(func() error {
		c.FotoKTP, err = utils.Encrypt(c.FotoKTP, secret)
		if err != nil {
			log.Errorf("error when hash password ")
			return err
		}
		return err
	})

	g.Go(func() error {
		c.FotoSelfie, err = utils.Encrypt(c.FotoKTP, secret)
		if err != nil {
			log.Errorf("error when hash password ")
			return err
		}
		return err
	})

	if err = g.Wait(); err != nil {
		return err
	}

	err = u.u.RegisterUser(c)
	if err != nil {
		return err
	}

	return nil
}
