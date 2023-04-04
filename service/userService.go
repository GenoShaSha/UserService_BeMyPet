package userService

import (
	"context"
	"encoding/json"
	"net/http"
	"userMicroService/user"
)

type Service interface {
	GetUser(context.Context) (*user.User, error)
}

type UserService struct {
	url string
}

func NewUserService(url string) Service {
	return &UserService{
		url: url,
	}
}

func (s *UserService) GetUser(ctx context.Context) (*user.User, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &user.User{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
