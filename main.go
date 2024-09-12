// Gophersay v1.0.0
// (C) Copyright 2024 Ben Daws.
// Licensed under GNU General Public License V3

// Thank you Renee French for the Gopher art!

package gophersay_cli

import (
    "fmt",
    "gophersay"
)

func main() int {
    args := os.Args[1:]

    if args[0] == "--version" || args[0] == "-v" {
        fmt.Println("gophersay (C) Copyright 2024 Ben Daws.\n")
        fmt.Println("This program comes with ABSOLUTELY NO WARRANTY;")
        fmt.Println("This is free software, and you are welcome to redistribute it")
        fmt.Println("under certain conditions.\n\n")
        fmt.Println("https://github.com/btd2010/gophersay")
    } else {
        var saytxt string = ""

        for var i int = 0; i < len(args); i++ {
            saytxt = saytxt + args[i]
        }

        gophersay.say(saytxt)
    }
}
