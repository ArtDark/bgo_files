package main

import (
	"encoding/xml"
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
		Issuer:   "Master Card",
		Balance:  48234_63,
		Currency: "RUB",
		Number:   "5106212365738734",
		Icon:     "https://www.mastercard.ru/content/dam/public/enterprise/resources/images/icons/favicon.ico",
		Transactions: card.Transactions{
			XMLName: xml.Name{
				Space: "",
				Local: "",
			},
			Transactions: []card.Transaction{},
		},
	} // Инициализация карточки пользователя Иван Петров

	petrIvanov := &card.Card{
		Id: 2,
		Owner: card.Owner{
			FirstName: "Petr",
			LastName:  "Ivanov",
		},
		Issuer:   "Master Card",
		Balance:  50_000_00,
		Currency: "RUB",
		Number:   "5106213743272347",
		Icon:     "https://www.mastercard.ru/content/dam/public/enterprise/resources/images/icons/favicon.ico",
		Transactions: card.Transactions{
			XMLName: xml.Name{
				Space: "",
				Local: "",
			},
			Transactions: []card.Transaction{},
		},
	}

	err := ivanPetrov.MakeTransactions(1_0) // создание транзакций
	if err != nil {
		log.Println(err)
	}

	err = card.ExporterToCsv(ivanPetrov)
	if err != nil {
		log.Println(err)
	}

	log.Println(petrIvanov)

	err = card.ExporterToJson(ivanPetrov, "export.json")
	if err != nil {
		log.Println(err)
	}

	err = card.ExporterToXml(ivanPetrov, "export.xml")
	if err != nil {
		log.Println(err)
	}

	err = card.ExporterToXml(ivanPetrov, "export.xml")
	log.Println(petrIvanov.Transactions.Transactions)

	if err != nil {
		log.Println(err)
	}

	err = card.ImporterFromXml(petrIvanov, "export.xml")
	if err != nil {
		log.Println(err)
	}

	log.Println(petrIvanov.Transactions.Transactions)

}
