package communication

import (
	"log"
	"strings"
)


func Deciphor(original string)(string){

	msg_prefix := "[[::"
	msg_suffix := "::]]"

	last := strings.Replace(original, msg_prefix, "", -1)
	a := strings.Replace(last, msg_suffix, "", -1)
	return a
}
