package main

import (
	"database/sql"
	"fmt"

	"github.com.br/GabrielSchenato/wallet-core/internal/database"
	"github.com.br/GabrielSchenato/wallet-core/internal/event"
	"github.com.br/GabrielSchenato/wallet-core/internal/usecase/create_account"
	"github.com.br/GabrielSchenato/wallet-core/internal/usecase/create_client"
	"github.com.br/GabrielSchenato/wallet-core/internal/usecase/create_transaction"
	"github.com.br/GabrielSchenato/wallet-core/internal/web"
	"github.com.br/GabrielSchenato/wallet-core/internal/web/webserver"
	"github.com.br/GabrielSchenato/wallet-core/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))

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

	createClientUseCase := create_client.NewCreateClientUseCase(clientDB)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDB, clientDB)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDB, accountDB, eventDispatcher, transactionCreatedEvent)

	webserver := webserver.NewWebServer(":3000")
	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accounttHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accounttHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webserver.Start()
}
