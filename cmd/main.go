package main

import (
	"github.com/ArtDark/bgo_files/pkg/card"
	"log"
)

func main() {

	ivanPetrov := &card.Card{
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
	} // Инициализация карточки пользователя Иван Петров

	petrIvanov := &card.Card{
		Id: 2,
		Owner: card.Owner{
			FirstName: "Petr",
			LastName:  "Ivanov",
		},
		Issuer:       "Master Card",
		Balance:      50_000_00,
		Currency:     "RUB",
		Number:       "5106213743272347",
		Icon:         "https://www.mastercard.ru/content/dam/public/enterprise/resources/images/icons/favicon.ico",
		Transactions: []card.Transaction{},
	}

	err := ivanPetrov.MakeTransactions(1_0) // создание транзакций
	if err != nil {
		log.Println(err)
	}

	err = card.Exporter(ivanPetrov)
	if err != nil {
		log.Println(err)
	}

	err = card.Importer(petrIvanov, "export.csv")
	if err != nil {
		log.Println(err)
	}

	log.Println(ivanPetrov)
	log.Println("=================================================================================================")
	log.Println(petrIvanov)

}

