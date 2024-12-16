package robotname

import (
	"errors"
	"math/rand"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var robotNames = map[string]bool{}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numberRunes = []rune("0123456789")

var total = len(letterRunes) * len(letterRunes) * len(numberRunes) * len(numberRunes) * len(numberRunes)

func randString() string {

	res := make([]rune, 5)
	for i := range res {

		if i <= 1 {
			res[i] = letterRunes[rand.Intn(len(letterRunes))]
		} else {
			res[i] = numberRunes[rand.Intn(len(numberRunes))]
		}
	}
	return string(res)
}

func generateName() (string, error) {

	for i := 0; i < total; i++ {
		name := randString()

		_, has := robotNames[name]

		if !has {
			robotNames[name] = true
			return name, nil
		}
	}

	return "", errors.New("all names are allocated")
}

func (r *Robot) Name() (string, error) {

	if len(r.name) > 0 {
		return r.name, nil
	}

	name, err := generateName()

	if err == nil {
		r.name = name
	}

	return name, err

}

func (r *Robot) Reset() {

	if len(r.name) > 0 {
		robotNames[r.name] = false
	}
	r.name = ""

}
