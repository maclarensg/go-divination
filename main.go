package main

import (
	"flag"
	"log"
)

func main() {
	parseDateStr := flag.String("date", "", "Date to parse")
	flag.Parse()

	// *parseDateStr if empty

	if *parseDateStr == "" {

		lunarDateTime, err := getLunarDateTime()
		if err != nil {
			log.Fatalln("Error:", err)
		}

		lunarDateTime.PrintDivineResult()
	} else {
		lunarDateTime, err := parseLunarDateTime(*parseDateStr)
		if err != nil {
			log.Fatalln("Error:", err)
		}

		lunarDateTime.PrintDivineResult()
	}
}
