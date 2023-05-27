package service

import (
	"github.com/Waelson/internal/model"
)

type ClientService interface {
	Save(client *model.Client) (error, *model.Client)
}

type clientService struct {
}

func (a *clientService) Save(client *model.Client) (error, *model.Client) {
	return nil, nil
}

func NewClientService() ClientService {
	return &clientService{}
}
