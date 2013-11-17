package gorinju

import (
	"time"
)

func ExamplePrintURL() {
	rep := Report{url: "foobar"}
	rep.PrintURL()
	//output:
	//Target URL :: foobar
}

func ExamplePrintTime() {
	rep := Report{time: time.Duration(10)}
	rep.PrintTime()
	//output:
	//Time       :: 10ns
}
