package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetUser(context.Context) (*User, error)
}

type UserService struct {
	url string
}

func NewUserService(url string) Service {
	return &UserService{
		url: url,
	}
}

func (s *UserService) GetUser(ctx context.Context) (*User, error) {
	resp, err := http.Get(s.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := &User{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
