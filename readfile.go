package goftpbrute

import (
	"io/ioutil"
	"fmt"
	"log"
)

//read password file
func readFile(passwordfile string)(passwords []byte)  {
	//check password file exist
	if passwordfile == ""{
		fmt.Println("[-] no file")
		return nil
	}
	//read from file
	var fileErr error
	passwords, fileErr = ioutil.ReadFile(passwordfile)
	if fileErr != nil {
        log.Fatal("[-] failed reading data from file: ", fileErr)
    }
	return
}