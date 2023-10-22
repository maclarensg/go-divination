package main

import (
	"log"
)

func main() {
	lunarDateTime, err := getLunarDateTime()
	if err != nil {
		log.Fatalln("Error:", err)
	}

	lunarDateTime.PrintDivineResult()
}
