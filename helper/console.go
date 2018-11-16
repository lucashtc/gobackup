package helper

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint function vai printar resultado in consoole beautiful
func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
