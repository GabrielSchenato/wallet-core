package createAccount

import (
	"github.com.br/GabrielSchenato/wallet-core/internal/entity"
	"github.com.br/GabrielSchenato/wallet-core/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountIOutputtDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(AccountGateway gateway.AccountGateway, ClientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: AccountGateway,
		ClientGateway:  ClientGateway,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountIOutputtDTO, error) {
	client, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}
	account := entity.NewAccount(client)
	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}
	return &CreateAccountIOutputtDTO{
		ID: account.ID,
	}, nil
}
