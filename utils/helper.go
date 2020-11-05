package utils

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	excelize "github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GetDateTimeT(data string) (time.Time, error) {
	exDate, _ := strconv.ParseFloat(data, 64)

	dateT, err := excelize.ExcelDateToTime(exDate, false)
	return dateT, err
}

func GetDateTime(data string) (string, error) {
	exDate, _ := strconv.ParseFloat(data, 64)

	dateT, err := excelize.ExcelDateToTime(exDate, false)
	return dateT.Format("02/01/2006"), err
}

func GetDateFromStr(dateStr string) string {
	dateT, err := time.Parse("01-02-06", dateStr)

	log.Println(err)

	return dateT.Format("02/01/2006")
}

func GetIntFromStirng(data string) int64 {
	val, err := strconv.ParseInt(data, 0, 64)
	if err != nil {
		val = 0
	}

	return val
}

func GetFloatFromStirng(data string) float64 {
	val, err := strconv.ParseFloat(data, 64)
	if err != nil {
		val = 0
	}

	return math.Round(val*100) / 100
}

func GetDateF2(strDate string) (time.Time, error) {
	clDateT, err := time.Parse("1/2/06 15:04", strDate)

	if err != nil {
		clDateT, err = GetDateTimeT(strDate)
		if err != nil {
			fmt.Println(err)
		}
	}

	return clDateT, err

}

func GetFloatRound(val float64) float64 {
	return math.Round(val*100) / 100
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
