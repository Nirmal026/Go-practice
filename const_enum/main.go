package main

import "fmt"

const (
	Ldate         = 4 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

func main() {
	fmt.Println("Ldate         :", Ldate)
	fmt.Println("Ltime         :", Ltime)
	fmt.Println("Lmicroseconds :", Lmicroseconds)
	fmt.Println("Llongfile     :", Llongfile)
	fmt.Println("Lshortfile    :", Lshortfile)
	fmt.Println("LUTC          :", LUTC)
	fmt.Println("Lmsgprefix    :", Lmsgprefix)
	fmt.Println("LstdFlags     :", LstdFlags)
}
