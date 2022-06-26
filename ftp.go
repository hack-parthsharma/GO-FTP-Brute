package goftpbrute

import (
	"github.com/jlaffaye/ftp"
	"time"
	"log"
	"fmt"
	"strings"
	"os"
)

func ftpConnect(port, hostname,username,password string){
	socket := fmt.Sprintf("%s:%s", hostname, port)
	c, err := ftp.Dial(socket, ftp.DialWithTimeout(20*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(username, password)
	if err != nil {
		if strings.Contains(err.Error(),"User cannot log in") == true{ //check invalid creds
			fmt.Println("[+] Failed -->", username,":",password)
			if err := c.Quit(); err != nil {
				log.Fatal(err)
			}
			return
		}
		if err := c.Quit(); err != nil {
			log.Fatal(err)
		}
		log.Fatal(err) //other errors
	}

	fmt.Println("[+] Success -->",username,":",password) //success
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}