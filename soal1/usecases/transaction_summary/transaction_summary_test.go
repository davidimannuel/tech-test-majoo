package transaction_summary_test

import (
	"majoo/soal1/usecases"
	"testing"
	"time"
)

func TestLastDate(t *testing.T) {
	year := 2021
	month := 5
	firstDayOfThisMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	t.Log(firstDayOfThisMonth)
	t.Log(firstDayOfThisMonth.AddDate(0, 0, -1))
	t.Log(firstDayOfThisMonth.AddDate(0, 0, 5))
	endOfThisMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.Local)
	t.Log(endOfThisMonth)
	t.Log(endOfThisMonth.Sub(firstDayOfThisMonth).Hours() / 24)
}

func TestPaginationDate(t *testing.T) {
	year := 2021
	month := 5

	firstDayOfThisMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	t.Log(firstDayOfThisMonth)
	endOfThisMonth := time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.Local)
	t.Log(endOfThisMonth)

	p := usecases.Pagination{
		CurrentPage: 2,
		PerPage:     5,
		TotalData:   endOfThisMonth.Day(),
	}

	startDay := firstDayOfThisMonth.AddDate(0, 0, p.GetOffset())
	endDay := startDay.AddDate(0, 0, p.PerPage-1)

	t.Log("=", startDay, endDay)
}

func TestPagination(t *testing.T) {
	total := 14
	limit := 5
	totalPage := total / limit
	t.Log("totalPage", totalPage)
}
