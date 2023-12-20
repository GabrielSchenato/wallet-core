package gateway

import "github.com.br/GabrielSchenato/wallet-core/internal/entity"

type AccountGateway interface {
	FindById(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
