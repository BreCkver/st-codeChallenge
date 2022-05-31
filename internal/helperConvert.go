package internal

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/BreCkver/st-codeChallenge/models"
)

const (
	debit      = "debit"
	credit     = "credit"
	layoutDate = "20060102"
)

/*ConvertTransaction conver string to Transaction*/
func ConvertTransaction(line string) (*models.Transaction, error) {

	Tx := new(models.Transaction)

	lineSplit := strings.Split(line, ",")

	if len(lineSplit) != 3 {
		return nil, errors.New(fmt.Sprintf("File's line with structure wrong. %v", line))
	}

	identifier, err := strconv.Atoi(strings.TrimSpace(lineSplit[0]))
	if err != nil {
		return nil, err
	}

	dateSplit := strings.Split(strings.TrimSpace(lineSplit[1]), "/")

	if len(dateSplit) != 2 {
		return nil, errors.New("Field date with structure wrong, date: " + lineSplit[1])
	}

	monthTemp, err := strconv.Atoi(dateSplit[0])
	if err != nil {
		return nil, err
	}

	if monthTemp > 12 || monthTemp < 0 {
		return nil, errors.New("Field month needs between 1 and 12, line: " + lineSplit[1])
	}

	dayTemp, err := strconv.Atoi(dateSplit[1])
	if err != nil {
		return nil, err
	}

	daysMonth := daysIn(time.Now().Year(), monthTemp)

	if dayTemp > daysMonth || dayTemp < 0 {
		return nil, errors.New("Field day needs between 1 and " + strconv.Itoa(daysMonth) + ". line: " + lineSplit[1])
	}

	layOut := layoutDate
	day := fmt.Sprintf("%02s", dateSplit[1])
	month := fmt.Sprintf("%02s", dateSplit[0])
	date, err := time.Parse(layOut, strconv.Itoa(time.Now().Year())+month+day)
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseFloat(strings.TrimSpace(lineSplit[2]), 64)
	if err != nil {
		return nil, err
	}

	Tx.Identifier = int32(identifier)
	Tx.Amount = math.Round(amount*100) / 100
	Tx.Date = date
	Tx.Type = getTransactionType(float32(amount))
	return Tx, nil
}

func getTransactionType(amount float32) string {
	if amount < 0 {
		return debit
	} else {
		return credit
	}
}

func daysIn(month int, year int) int {
	return time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
