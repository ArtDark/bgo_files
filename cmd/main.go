package main

import (
	"encoding/csv"
	"github.com/ArtDark/bgo_files/pkg/card"
	"io"
	"log"
	"os"
	"strconv"
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
	} // Инициализация карточки пользователя

	err := ivanPetrov.MakeTransactions(1_0) // создание транзакций
	if err != nil {
		log.Println(err)
	}

	err = exporter(ivanPetrov)
	if err != nil {
		log.Println(err)
	}

}

// Экспорт пользовательских транзакций в .csv
func exporter(user *card.Card) (err error) {

	file, err := os.Create("export.csv")

	if err != nil {
		log.Println(err)
		return err
	}

	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println(cerr)
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range user.Transactions {
		err = writer.Write(transactionToSlice(value))
		if err != nil {
			log.Println(err)
			return err
		}

	}

	return nil
}

func transactionToSlice(transaction card.Transaction) []string {

	var data []string

	data = append(data, transaction.Id)
	data = append(data, strconv.Itoa(int(transaction.Bill)))
	data = append(data, strconv.Itoa(int(transaction.Time)))
	data = append(data, transaction.MCC)
	data = append(data, transaction.Status)

	return data

}
