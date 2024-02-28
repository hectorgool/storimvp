package controller

import (
	"encoding/csv"
	"fmt"
	"os"
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
			fmt.Printf("%v,%v,%v,%v;\n", d.Id, month, day, d.Transaction)

			/*
				t := schema.DBDocument{
						Id:          stringToUint(d.Id),
						Date:        fmt.Sprintf("%v-%v-%v", YEAR, month, day),
						Transaction: d.Transaction,
				}
				AddTransaction(t)
			*/
		}
	}

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
