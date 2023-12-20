package gateway

import "github.com.br/GabrielSchenato/wallet-core/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
