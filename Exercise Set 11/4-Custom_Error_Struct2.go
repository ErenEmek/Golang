package main

import (
	"fmt"
	"log"
	"errors"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		s1 := sqrtError{				//Error occurs, we set the values of struct to produce error code.
			lat:  "31,3948",
			long: "27,234",
			err:  errors.New("math error occured"), //errors.New is used to create custom error.
		}

		return 0, errors.New(s1.Error())		//since func Error() returns a string, this string is called and converted to error type via errors.New
	}
	return 42, nil
}
