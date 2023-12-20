package gateway

import "github.com.br/GabrielSchenato/wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
