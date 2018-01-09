package communication

import (
	"log"
	"os"
	"os/exec"
	"io/ioutil"
	"strings"
)

// globals
var msg_prefix string = "[[::"
var msg_suffix string = "::]]"

// util functions
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func checkfile(f string)(bool){
	_,err := os.Stat(f)
	if(err != nil){
		return false
	}
	return true
}

func Deciphor(original string)(string){
	if !strings.Contains(original, msg_prefix) {
                // No msg_prefix, no where to start
		return ""
	}
	last := strings.Replace(original, msg_prefix, "", -1)
	if !strings.Contains(last, msg_suffix) {
                // After taking out the msg_prefix, no msg_suffix left
		return ""
	}
	a := strings.Replace(last, msg_suffix, "", -1)
	return a
}

func Perform(data string)(string) {

	s := strings.Split(Deciphor(data),":")
	command := s[0]
	var filename, content string
	if len(s) >= 2 {
		filename = s[1]
	}
	if len(s) >= 3 {
		content = s[2]
	}

	if command == "" {
		log.Printf("No command received. Received %s", data)
		return "fail"
	}
	if filename == "" {
		log.Printf("No filename received. Received %s", data)
		return "fail"
	}

	switch command {
	case "deletefile":
		// delete a file
		if !checkfile(filename) {
			log.Printf("File does not exist. Cannot delete. Received %s", filename)
			return "fail"
		}
		cmd := exec.Command("rm", filename)
		cmd.Run()
	case "writefile":
		// write to a file, if not exist, will create
		err := ioutil.WriteFile(filename, []byte(content), 0644)
		check(err)
	case "appendfile":
		// append to a file, if not exist, wil create
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		check(err)
		f.WriteString(string(content))
		f.Close()
	case "exec":
		// execute a file as shell file
		cmd := exec.Command("sh", filename)
		err := cmd.Run()
		check(err)
	case "readfile":
		dat, err := ioutil.ReadFile(filename)
		check(err)
		// 106 is the messageTYPE
		return string(dat)
	default:
		log.Printf("Unkonwn command received:%s", command)
	}

	return "success"
}
