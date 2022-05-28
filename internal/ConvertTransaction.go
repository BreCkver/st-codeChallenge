package internal

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/BreCkver/st-codeChallenge/models"
)

/*ConvertTransaction */
func ConvertTransaction(line string) (*models.Transaction, error) {

	Tx := new(models.Transaction)

	lineSplit := strings.Split(line, ",")

	identifier, err := strconv.Atoi(lineSplit[0])
	if err != nil {
		return nil, err
	}

	dateSplit := strings.Split(lineSplit[1], "/")
	layOut := "20060102"

	day := fmt.Sprintf("%02s", dateSplit[1])
	month := fmt.Sprintf("%02s", dateSplit[0])
	date, err := time.Parse(layOut, strconv.Itoa(time.Now().Year())+month+day)
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseFloat(lineSplit[2], 32)
	if err != nil {
		return nil, err
	}

	Tx.Identifier = int32(identifier)
	Tx.Amount = float32(amount)
	Tx.Date = date
	Tx.Type = GetTransactionType(float32(amount))
	Tx.Period = GetTransactionPeriod(date)

	return Tx, nil
}

func GetTransactionType(amount float32) string {
	if amount < 0 {
		return "debit"
	} else {
		return "credit"
	}
}

func GetTransactionPeriod(date time.Time) string {
	var yyyyMM = "200601"
	return date.Format(yyyyMM)
}
