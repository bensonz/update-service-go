package communication_test

import (
	"os"
	"strconv"
	"testing"
	"../communication"
)

func singleTestDeciphor(s string, expected string, t *testing.T) {
	received := communication.Deciphor(s)
	if  received != expected {
		t.Error("Test",s,"failed")
		t.Error("Received:",received)
		t.Error("Expected:",expected)
	}
}

func TestDeciphor(t *testing.T) {
	s := ""
	singleTestDeciphor(s, "", t)
	s = "[[::a::]]"
	singleTestDeciphor(s, "a", t)
	s = "adf"
	singleTestDeciphor(s, "", t)
	s = "[[::]]"
	singleTestDeciphor(s, "", t)
	s = "[[::::]]"
	singleTestDeciphor(s, "", t)
	s = "[[::a:b:c::]]"
	singleTestDeciphor(s, "a:b:c", t)
	s = "[[::a-b/`12`[[]]1234c::]]"
	singleTestDeciphor(s, "a-b/`12`[[]]1234c", t)
        // not sure if this case is possible
	// s = "[[::a::]][[::b"
	// singleTestDeciphor(s, "a", t)

}


func buildMsg(s string)(string){
	return "[[::" + s + "::]]"
}

func checkfile(f string)(bool){
	_,err := os.Stat(f)
	if(err != nil){
		return false
	}
	return true
}

func TestWrtieFile(t *testing.T) {
	data := buildMsg("writefile:testWrite:ok")
	ret := communication.Perform(data)
	if ret != "success" {
		t.Error("Test write failed.")
	}
	if !checkfile("testWrite"){
		t.Error("Test write failed. File not created.")
	}
}

func TestAppend(t *testing.T){
	if !checkfile("testWrite") {
		os.Create("testWrite")
	}
	data := buildMsg("appendfile:testWrite:123")
	ret := communication.Perform(data)
	if ret != "success" {
		t.Error("Test append failed.")
	}
}

func TestRead(t *testing.T){
        // testWrite must exist
	data := buildMsg("readfile:testWrite")
	ret := communication.Perform(data)
	if ret != "ok123" {
		t.Error("Test read failed.")
	}
}

func TestDelete(t *testing.T){
	if !checkfile("testWrite") {
		os.Create("testWrite")
	}
	data := buildMsg("deletefile:testWrite")
	ret := communication.Perform(data)
	if ret != "success" {
		t.Error("Test delete failed.")
	}
	if checkfile("testWrite") {
		t.Error("Test delete failed. File not deleted.")
	}
}

func TestExec(t *testing.T){
	data := buildMsg("writefile:testExec.sh:#! bin/bash\n")
	ret := communication.Perform(data)
	if ret != "success" {
		t.Error("write in Test exec failed.")
	}
	if !checkfile("testExec.sh") {
		t.Error("write in Test exec failed. File not created.")
	}

	content := "for x in {1..3};do echo $x > $x;done"
	data = buildMsg("appendfile:testExec.sh:"+content)
	ret = communication.Perform(data)
	if ret != "success" {
		t.Error("append in Test exec failed.")
	}

	data = buildMsg("exec:testExec.sh")
	communication.Perform(data)
	for i := 1; i <= 3; i ++ {
		intString := strconv.Itoa(i)
		if !checkfile(intString) {
			t.Error("Test exec failed with file:", intString)
		}else{
			data = buildMsg("deletefile:"+intString)
			communication.Perform(data)
		}
	}

	data = buildMsg("deletefile:testExec.sh")
	communication.Perform(data)

}
