package main

import (
	"github.com/ArtDark/bgo_files/pkg/card"
	"io"
	"log"
	"os"
)

func main() {
	user := &card.Card{
		Id: 1,
		Owner: card.Owner{
			FirstName: "Ivan",
			LastName:  "Petrov",
		},
		Issuer:       "Master Card",
		Balance:      48234_63,
		Currency:     "RUB",
		Number:       "5106212365738734",
		Icon:         "https://www.mastercard.ru/content/dam/public/enterprise/resources/images/icons/favicon.ico",
		Transactions: []card.Transaction{},
	} // Инициализация карточки пользователя

	err := user.MakeTransactions(1_0) // создание транзакций
	if err != nil {
		log.Println(err)
	}

	exporter(user)

}

// Экспорт пользовательских транзакций в .csv
func exporter(user *card.Card) (err error) {

	transactions := &user.Transactions

	if transactions == nil {
		log.Println("No transactions")
		return nil
	}

	file, err := os.Create("export.csv")

	if err != nil {
		log.Println(err)
		return err
	}

	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
			if err == nil {
				err = cerr
			}
		}
	}(file)

	// TODO: дописать функцию

	return nil

}

func importer() {

}
