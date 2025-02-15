package wishlists

import (
	businesses "aprian1337/thukul-service/business"
	"aprian1337/thukul-service/business/users"
	"context"
	"time"
)

type WishlistUsecase struct {
	Repo        Repository
	UserUsecase users.Usecase
	Timeout     time.Duration
}

func NewWishlistUsecase(repo Repository, userUsecase users.Usecase, timeout time.Duration) *WishlistUsecase {
	return &WishlistUsecase{
		Repo:        repo,
		UserUsecase: userUsecase,
		Timeout:     timeout,
	}
}

func (uc *WishlistUsecase) GetList(ctx context.Context, userId int) ([]Domain, error) {
	data, err := uc.Repo.GetList(ctx, userId)
	if err != nil {
		return []Domain{}, err
	}
	return data, nil
}

func (uc *WishlistUsecase) GetById(ctx context.Context, userId int, wishlistId int) (Domain, error) {
	data, err := uc.Repo.GetById(ctx, userId, wishlistId)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}

func (uc *WishlistUsecase) Create(ctx context.Context, domain Domain, userId int) (Domain, error) {
	_, err := uc.UserUsecase.GetById(ctx, userId)
	if err != nil {
		return Domain{}, businesses.ErrUserIdNotFound
	}
	data, err := uc.Repo.Create(ctx, domain, userId)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}

func (uc *WishlistUsecase) Update(ctx context.Context, domain Domain, userId int, wishlistId int) (Domain, error) {
	data, err := uc.Repo.Update(ctx, domain, userId, wishlistId)
	if err != nil {
		return Domain{}, err
	}
	return data, nil
}

func (uc *WishlistUsecase) Delete(ctx context.Context, userId int, wishlistId int) error {
	rowsAffected, err := uc.Repo.Delete(ctx, userId, wishlistId)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return businesses.ErrNothingDestroy
	}
	return nil
}
