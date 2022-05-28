package main

import (
	"github.com/BreCkver/st-codeChallenge/handlers"
)

func main() {
	/*
		inputDate := "20220728"
		fmt.Println("\n\nTest 2 date : ", inputDate)
		layOut := "20060102"
		timeStamp, err := time.Parse(layOut, inputDate)
		if err != nil {
			fmt.Println(err)
		} else {
			year, month, day := timeStamp.Date()
			fmt.Printf("Date : [%d]year : [%d]month : [%d]day \n", year, month, day)
			// dd-mm-yyyy
			fmt.Println("dd-mm-yyyy date format : ", timeStamp.Format("02-01-2006"))
		}*/

	handlers.Handler()
}
