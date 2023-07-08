package main

import (
	"encoding/csv"
	"fmt"
	"os"
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
	header := []string{"georgian_date", "jalali_date", "jalali_persian_date_chars_full", "jalali_day_name", "jalali_month_name", "season"}
	err = writer.Write(header)
	if err != nil {
		panic(err)
	}

	// Seasons mapping
	fasl_map := make(map[string][3]string)
	fasl_map["بهار"] = [3]string{"فروردین", "اردیبهشت", "خرداد"}
	fasl_map["تابستان"] = [3]string{"تیر", "مرداد", "شهریور"}
	fasl_map["پاییز"] = [3]string{"مهر", "آبان", "آذر"}
	fasl_map["زمستان"] = [3]string{"دی", "بهمن", "اسفند"}

	// Going Back and Front in time
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
		jalali_persian_date_chars_full := fa_then.Format("yyyy MMM dd")
		season := ""
		for k, months := range fasl_map {
			for _, v := range months {
				if jalali_dateـmonth_name == v {
					season = k
					fmt.Println(georgian_date, jalali_date, jalali_persian_day, jalali_dateـmonth_name, k)
				}
			}
		}
		// Write the row to the CSV file
		row := []string{georgian_date, jalali_date, jalali_persian_date_chars_full, jalali_persian_day, jalali_dateـmonth_name, season}
		err = writer.Write(row)
		if err != nil {
			panic(err)
		}

	}

	// Flush the CSV writer to write any remaining data to the file
	writer.Flush()

	fmt.Println("Data written to data.csv")
}
