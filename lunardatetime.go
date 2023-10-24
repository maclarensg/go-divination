package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	calendar "github.com/Lofanmi/chinese-calendar-golang/calendar"
)

// DateTime struct
type DateTime struct {
	Year, Month, Day, Hour, Minute, Second int
}

// countNumber(n int) , where n is the number input and will return an int looping from 1 to 6, e.g. 1=1, 6=6, 11=5, 12=12.
func countNumber(n int) int {
	if n%6 == 0 {
		return 6
	}
	return n % 6
}

// return the string representation of the DateTime struct
func (dt *DateTime) String() string {
	return fmt.Sprintf("Lunar %d/%d/%d, %02d:%02d:%02d",
		dt.Year, dt.Month, dt.Day,
		dt.Hour, dt.Minute, dt.Second)
}

// printDivineResult returns the divine message for the current time
func (dt *DateTime) PrintDivineResult() {
	fmt.Println("Lunar DateTime:", dt.String())

	fmt.Println(divineTable[countNumber(dt.Month)])
	fmt.Println(divineTable[countNumber(dt.Day)])
	fmt.Println(divineTable[countNumber(TimeTable(dt.Hour))])
}

// getLunarDateTime retrieves the current Lunar Date and Time
func getLunarDateTime() (*DateTime, error) {
	t := time.Now()

	c := calendar.BySolar(int64(t.Year()), int64(t.Month()), int64(t.Day()), int64(t.Hour()), int64(t.Minute()), int64(t.Second()))
	bytes, _ := c.ToJSON()

	// Create a map to store the JSON data
	var resultMap map[string]interface{}
	if err := json.Unmarshal(bytes, &resultMap); err != nil {
		return nil, err
	}

	// Accessing the values from the map
	lunarData, ok := resultMap["lunar"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error extracting lunar data")
	}

	solarData, ok := resultMap["solar"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error extracting solar data")
	}

	day, dayOk := lunarData["day"].(float64)
	month, monthOk := lunarData["month"].(float64)
	year, yearOk := lunarData["year"].(float64)
	hour, hourOk := solarData["hour"].(float64)
	minute, minuteOk := solarData["minute"].(float64)
	second, secondOk := solarData["second"].(float64)

	if dayOk && monthOk && yearOk && hourOk && minuteOk && secondOk {
		return &DateTime{
			Year:   int(year),
			Month:  int(month),
			Day:    int(day),
			Hour:   int(hour),
			Minute: int(minute),
			Second: int(second),
		}, nil
	}

	return nil, fmt.Errorf("error accessing date and time values from lunar data")
}

// parseLunarDateTime parses the Lunar Date and Time from a string
func parseLunarDateTime(s string) (*DateTime, error) {

	//split string into "YYYY/MM/DD HH:MM" to year, month, day, hour, minute
	var syear, smonth, sday, shour, sminute int64
	_, err := fmt.Sscanf(s, "%d/%d/%d %d:%d", &syear, &smonth, &sday, &shour, &sminute)

	if err != nil {
		log.Fatalln("Error:", err)
	}

	c := calendar.BySolar(syear, smonth, sday, shour, sminute, 0)

	bytes, _ := c.ToJSON()

	// Create a map to store the JSON data
	var resultMap map[string]interface{}
	if err := json.Unmarshal(bytes, &resultMap); err != nil {
		return nil, err
	}

	// Accessing the values from the map
	lunarData, ok := resultMap["lunar"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error extracting lunar data")
	}

	solarData, ok := resultMap["solar"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("error extracting solar data")
	}

	day, dayOk := lunarData["day"].(float64)
	month, monthOk := lunarData["month"].(float64)
	year, yearOk := lunarData["year"].(float64)
	hour, hourOk := solarData["hour"].(float64)
	minute, minuteOk := solarData["minute"].(float64)
	second, secondOk := solarData["second"].(float64)

	if dayOk && monthOk && yearOk && hourOk && minuteOk && secondOk {
		return &DateTime{
			Year:   int(year),
			Month:  int(month),
			Day:    int(day),
			Hour:   int(hour),
			Minute: int(minute),
			Second: int(second),
		}, nil
	}

	return nil, fmt.Errorf("error accessing date and time values from lunar data")
}
