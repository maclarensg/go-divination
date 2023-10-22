package main

import (
	"encoding/json"
	"fmt"
	"time"

	calendar "github.com/Lofanmi/chinese-calendar-golang/calendar"
)

// DateTime struct
type DateTime struct {
	Year, Month, Day, Hour, Minute, Second int
}

// return the string representation of the DateTime struct
func (dt *DateTime) String() string {
	return fmt.Sprintf("Lunar %d/%d/%d, %02d:%02d:%02d",
		dt.Year, dt.Month, dt.Day,
		dt.Hour, dt.Minute, dt.Second)
}

// printDivineResult returns the divine message for the current time
func (dt *DateTime) PrintDivineResult() {
	fmt.Println(divineTable[dt.Month%6])
	fmt.Println(divineTable[dt.Day%6])
	fmt.Println(divineTable[TimeTable(dt.Hour)%6])
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
