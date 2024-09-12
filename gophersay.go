// Gophersay v1.0.0
// (C) Copyright 2024 Ben Daws.
// Licensed under GNU General Public License V3

// Thank you Renee French for the Gopher art!

package gophersay

import "fmt"

func gophersay_chardupe(character string, strlen int) string {
	// previously spaces
	var spaces string = ""

	for var i int = 0; i < strlen; i++ {
		spaces = spaces + character
	}

	return spaces
}

func gophersay_spaces(strlen int) string {
	var spaces string = ""

	for var i int = 0; i < strlen; i++ {
		spaces = spaces + character
	}

	spaces = spaces + "  "

	return spaces
}

func say(gophertext string) {
	var strlen int = len(gophertext)

	fmt.Println("    `.-::::::-.`    " + gophersay_spaces(strlen))
	fmt.Println(".:-::::::::::::::-:." + " " + gophersay_chardupe("=", strlen) + " ")
	fmt.Println("`_:::    ::    :::_`" + "= " + gophertext + " =")
	fmt.Println(" .:( ^   :: ^   ):. " + " " + gophersay_chardupe("=", strlen) + " ")
	fmt.Println(" `:::   (..)   :::. " + gophersay_spaces(strlen))
	fmt.Println(" `:::::::UU:::::::` " + gophersay_spaces(strlen))
	fmt.Println(" .::::::::::::::::. " + gophersay_spaces(strlen))
	fmt.Println(" O::::::::::::::::O " + gophersay_spaces(strlen))
	fmt.Println(" -::::::::::::::::- " + gophersay_spaces(strlen))
	fmt.Println(" `::::::::::::::::` " + gophersay_spaces(strlen))
	fmt.Println("  .::::::::::::::.  " + gophersay_spaces(strlen))
	fmt.Println("    oO:::::::Oo     " + gophersay_spaces(strlen))
}

// usage: gophersay.say(txt)
