package users

import (
	"aprian1337/thukul-service/app/middlewares"
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/wallets"
	"aprian1337/thukul-service/helpers"
	"aprian1337/thukul-service/utilities"
	"aprian1337/thukul-service/utilities/constants"
	"context"
	"log"
	"time"
)

type UserUsecase struct {
	Repo          Repository
	WalletUsecase wallets.Usecase
	Timeout       time.Duration
	jwtAuth       *middlewares.ConfigJWT
}

func NewUserUsecase(repo Repository, walletUsecase wallets.Usecase, timeout time.Duration, jwtAuth *middlewares.ConfigJWT) *UserUsecase {
	return &UserUsecase{
		Repo:          repo,
		WalletUsecase: walletUsecase,
		Timeout:       timeout,
		jwtAuth:       jwtAuth,
	}
}

func (uc *UserUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) GetById(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetById(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	if user.ID == 0 {
		return Domain{}, businesses.ErrIDNotFound
	}
	return user, nil
}

func (uc *UserUsecase) Create(ctx context.Context, domain *Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, businesses.ErrEmailRequired
	}

	if !helpers.IsEmailValid(domain.Email) {
		return Domain{}, businesses.ErrEmailNotValid
	}

	data, err := uc.Repo.GetByEmail(ctx, domain.Email)
	if data.ID > 0 {
		return Domain{}, businesses.ErrEmailHasBeenRegister
	}

	if domain.Password == "" {
		return Domain{}, businesses.ErrPasswordRequired
	}

	if !helpers.IsDate(domain.Birthday) {
		return Domain{}, businesses.ErrInvalidDate
	}

	domain.Password, _ = utilities.HashPassword(domain.Password)

	user, err := uc.Repo.Create(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	err = uc.WalletUsecase.Create(ctx, user.ToWalletDomain())
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (Domain, string, error) {
	if email == "" || password == "" {
		return Domain{}, "", businesses.ErrUsernamePasswordNotFound
	}

	user, err := uc.Repo.GetByEmail(ctx, email)

	if err != nil {
		return Domain{}, "", err
	}

	if !utilities.CheckPassword(password, user.Password) {
		return Domain{}, "", businesses.ErrInvalidAuthentication
	}

	//JWT
	token, errToken := uc.jwtAuth.GenerateTokenJWT(user.ID)
	if errToken != nil {
		log.Println(errToken)
	}
	if token == "" {
		return Domain{}, "", businesses.ErrInvalidAuthentication
	}

	return user, token, nil
}

func (uc *UserUsecase) Update(ctx context.Context, domain *Domain, id uint) (Domain, error) {
	if domain.Birthday != "" {
		_, err := time.Parse(constants.BirthdayFormat, domain.Birthday)
		if err != nil {
			return Domain{}, businesses.ErrInvalidDate
		}
	}
	if domain.Email == "" {
		return Domain{}, businesses.ErrEmailRequired
	}
	if !helpers.IsEmailValid(domain.Email) {
		return Domain{}, businesses.ErrEmailNotValid
	}

	data, err := uc.Repo.GetByEmail(ctx, domain.Email)
	if data.ID > 0 && data.ID != id {
		return Domain{}, businesses.ErrEmailHasBeenRegister
	}

	domain.Password, _ = utilities.HashPassword(domain.Password)
	domain.ID = id
	user, err := uc.Repo.Update(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUsecase) Delete(ctx context.Context, id uint) error {
	err := uc.Repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
