package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	ptime "github.com/yaa110/go-persian-calendar"
)

func main() {
	// File Manipulations
	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	// Write the CSV header row
	header := []string{"georgian_date", "jalali_date", "jalali_year_id", "jalali_month_id", "jalali_day_id", "day_of_week_id", "day_of_week", "month", "season", "season_id"}

	err = writer.Write(header)
	if err != nil {
		panic(err)
	}

	// Seasons mapping
	daysOfWeekMap := make(map[string]int)
	daysOfWeekMap["شنبه"] = 1
	daysOfWeekMap["یک‌شنبه"] = 2
	daysOfWeekMap["دوشنبه"] = 3
	daysOfWeekMap["سه‌شنبه"] = 4
	daysOfWeekMap["چهارشنبه"] = 5
	daysOfWeekMap["پنج‌شنبه"] = 6
	daysOfWeekMap["جمعه"] = 7

	fasl_map := make(map[string][3]string)
	fasl_map["بهار"] = [3]string{"فروردین", "اردیبهشت", "خرداد"}
	fasl_map["تابستان"] = [3]string{"تیر", "مرداد", "شهریور"}
	fasl_map["پاییز"] = [3]string{"مهر", "آبان", "آذر"}
	fasl_map["زمستان"] = [3]string{"دی", "بهمن", "اسفند"}
	faslMapNum := make(map[string]int)
	faslMapNum["بهار"] = 1
	faslMapNum["تابستان"] = 2
	faslMapNum["پاییز"] = 3
	faslMapNum["زمستان"] = 4
	then := time.Date(1980, 01, 02, 01, 01, 01, 651387237, ptime.Iran())
	//then_before := time.Date(1980, 01, 01, 01, 01, 01, 651387237, time.Now().Local().UTC().Location())
	//now := time.Time()
	for i := 0; i <= 365*70; i++ {
		then = then.AddDate(0, 0, 1)
		fa_then := ptime.New(then)
		georgian_date := then.Format("2006-01-02")
		jalali_date := fa_then.Format("yyyy-MM-dd")
		jalali_dateـmonth_name := fa_then.Format("MMM")
		jalali_persian_day := fa_then.Format("E")
		jalali_persian_date_year := fa_then.Format("yyyy")
		jalali_persian_day_id := daysOfWeekMap[jalali_persian_day]
		jalali_normal_day_id := fa_then.Format("dd")
		jalali_persian_season_id := 0
		jalali_persian_month_id := fa_then.Format("M")
		season := ""
		for k, months := range fasl_map {
			for _, v := range months {
				if jalali_dateـmonth_name == v {
					season = k
					jalali_persian_season_id = faslMapNum[season]

				}
			}
		}
		// Write the row to the CSV file
		row := []string{
			georgian_date,
			jalali_date,
			jalali_persian_date_year,
			jalali_persian_month_id,
			jalali_normal_day_id,
			strconv.Itoa(jalali_persian_day_id),
			jalali_persian_day,
			jalali_dateـmonth_name,
			season,
			strconv.Itoa(jalali_persian_season_id),
		}
		err = writer.Write(row)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	fmt.Println("Data written to data.csv")
}
