package controller

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"storimvp/config"
	"storimvp/schema"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	FILE           = "./txns.cvs"
	FILE_DELIMITER = ','
	YEAR           = 2024
)

func SendMail(c *gin.Context) {

	var emailData schema.EmailData

	userEmail := c.Param("userEmail")

	readCVSFile()

	emailData.EmailTo = userEmail
	emailData.TotalBalance = totalBalance()
	emailData.AverageDebitAmount = averageDebitAmount()
	emailData.AverageCreditAmount = averageCreditAmount()
	emailData.Transactions = numberTransactionsInMonth()

	fmt.Println(emailData)

	err := SendEmail(emailData)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"api": "Send Mail",
	})
}

func Reset(c *gin.Context) {

	result := config.GetDB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&schema.DBDocument{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"api": "Reset",
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

	for _, col := range rows {
		d.Id = col[0]
		d.Date = col[1]
		d.Transaction = col[2]
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

func numberTransactionsInMonth() []schema.TransactionsByMonth {

	transactions := []schema.TransactionsByMonth{}

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
			newTransaction := schema.TransactionsByMonth{
				Total: n,
				Month: monthName,
			}
			// Append the new transaction to the array
			transactions = append(transactions, newTransaction)
		}
	}
	return transactions

}
func SendEmail(data schema.EmailData) error {
	// Configuración del servidor SMTP
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	senderEmail := os.Getenv("SMTP_SENDER")
	senderPassword := os.Getenv("SMTP_PASSWD")

	// Autenticación con el servidor SMTP
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)

	// Plantilla HTML externa
	templateFile := "template.html"

	// Parseamos la plantilla HTML
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		return err
	}

	// Creamos un buffer para almacenar la salida de la plantilla
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return err
	}

	// Mensaje del correo electrónico
	htmlMessage := tpl.String()

	// Destinatario
	to := []string{os.Getenv("SMTP_SENDER")}
	to = append(to, data.EmailTo)
	subject := "MVP Stori, Héctor González Olmos"

	body := []byte("To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		htmlMessage)

	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, senderEmail, to, body)
	if err != nil {
		return err
	}
	return nil
}
