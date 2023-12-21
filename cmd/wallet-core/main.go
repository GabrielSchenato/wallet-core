package walletcore

import (
	"database/sql"
	"fmt"

	"github.com.br/GabrielSchenato/wallet-core/internal/database"
	"github.com.br/GabrielSchenato/wallet-core/internal/event"
	createaccount "github.com.br/GabrielSchenato/wallet-core/internal/usecase/create_account"
	createclient "github.com.br/GabrielSchenato/wallet-core/internal/usecase/create_client"
	createtransaction "github.com.br/GabrielSchenato/wallet-core/internal/usecase/create_transaction"
	"github.com.br/GabrielSchenato/wallet-core/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))

	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDB := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)
	transactionDB := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDB)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDB, clientDB)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(transactionDB, accountDB, eventDispatcher, transactionCreatedEvent)

}
