package main

import (
	"github.com/mars1385/exercism/robotname"
)

func main() {

	for i := 0; i < 26*26*10*10*10; i++ {
		d := robotname.Robot{}
		d.Name()
	}

}
