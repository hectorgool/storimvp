package controller

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"storimvp/config"
	"storimvp/schema"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	FILE           = "./txns.cvs"
	FILE_DELIMITER = ','
	YEAR           = 2024
)

func SendMail(c *gin.Context) {
	readCVSFile()
	c.JSON(200, gin.H{
		"api": "Send Mail",
	})
}

func readCVSFile() {

	file, err := os.Open(FILE)
	printError(err)
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = FILE_DELIMITER

	rows, err := reader.ReadAll()
	printError(err)

	d := new(schema.Document)

	for n, col := range rows {
		d.Id = col[0]
		d.Date = col[1]
		d.Transaction = col[2]
		n++
		if (strings.ToUpper(d.Id) != strings.ToUpper("Id")) || (strings.ToUpper(d.Date) != strings.ToUpper("Date")) || (strings.ToUpper(d.Transaction) != strings.ToUpper("Transaction")) {
			dateCVS := strings.Split(d.Date, "/")
			month := dateCVS[0]
			day := dateCVS[1]
			fmt.Printf("IdTransaction:%v, Moth:%v, Day:%v, Transactio:%v;\n", d.Id, month, day, d.Transaction)

			t := schema.DBDocument{
				IdTransaction: stringToUint(d.Id),
				Date:          fmt.Sprintf("%v-%v-%v", YEAR, month, day),
				Transaction:   stringToUint64(d.Transaction),
			}
			addTransaction(t)

		}
	}
	fmt.Println("Total balance: ", totalBalance())
	fmt.Println("Average debit amount: ", averageDebitAmount())
	fmt.Println("Average credit amount: ", averageCreditAmount())
	numberTransactionsInMonth()
}

func addTransaction(t schema.DBDocument) {
	if err := config.GetDB().Create(&t).Error; err != nil {
		log.Fatalln("Create row fail!")
	}
}

func totalBalance() float64 {
	var total float64
	if err := config.GetDB().Model(&schema.DBDocument{}).Select("SUM(transaction)").Scan(&total).Error; err != nil {
		log.Fatalln("failed to get total transaction")
	}
	return total
}

func averageDebitAmount() float64 {
	var avg float64
	if err := config.GetDB().Model(&schema.DBDocument{}).Where("transaction < ?", 0).Select("AVG(transaction)").Scan(&avg).Error; err != nil {
		log.Fatalln("failed to get total transaction")
	}
	return avg
}

func averageCreditAmount() float64 {
	var avg float64
	if err := config.GetDB().Model(&schema.DBDocument{}).Where("transaction > ?", 0).Select("AVG(transaction)").Scan(&avg).Error; err != nil {
		log.Fatalln("failed to get total transaction")
	}
	return avg
}

func printError(err error) {
	if err != nil {
		fmt.Printf("\nError: %v \n ", err.Error())
		os.Exit(1)
	}
}

func stringToUint(s string) uint {
	// Convertir el string a uint64
	num, _ := strconv.ParseUint(s, 10, 0)
	// Convertir el uint64 a uint
	return uint(num)
}

func stringToUint64(s string) float64 {
	floatValue, _ := strconv.ParseFloat(s, 64)
	// Convertir el float64 a int64
	return floatValue
}

func countTransactionsByMonth(monthNumber int) int64 {
	var count int64
	if err := config.GetDB().Model(&schema.DBDocument{}).Where("MONTH(date) = ?", monthNumber).Count(&count).Error; err != nil {
		log.Fatalln("failed to get total transaction")
	}
	return count
}

func numberTransactionsInMonth() {
	months := map[int]string{
		1:  "Junuary",
		2:  "February",
		3:  "March",
		4:  "Abril",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	for monthNumber, monthName := range months {
		n := countTransactionsByMonth(monthNumber)
		if n != 0 {
			fmt.Printf("Number of transactions in %v: %v\n", monthName, countTransactionsByMonth(monthNumber))
		}
	}
}
