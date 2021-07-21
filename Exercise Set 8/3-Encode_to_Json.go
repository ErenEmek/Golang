package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

/* func NewEncoder(w io.Writer) *Encoder
   func (enc *Encoder) Encode(v interface{}) error => to use this function, we need to add and encoder.
						      to add an encoder. We need to give it a Writer.
						      Write method is attached to type File.
						      Writer interface has the same method type File has.
						      os.Stdout is a NewFile, gives pointer to a file.
*/						      

	enc := json.NewEncoder(os.Stdout)	//new encoder defined
	enc.Encode(users)			//Encode method is called with identifier users
						//We encode data, straight to os.Stdout
}
