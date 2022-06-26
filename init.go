package goftpbrute


import (
	"log"
)

func (s *GoFtpBrute)Init(hostname,port,username,password,passwordfile string,thread int){
	//Check error from parsed args 
	if (hostname == ""){
		log.Fatal("[-] missing hostname")
	}else if (username == ""){
		log.Fatal("[-] missing username")
	}
	if ((password == "") && (passwordfile == "")){
		log.Fatal("[-] missing credentials: provide password")
	}
	//initialize
	s.hostname = hostname
	s.port = port
	s.username = username
	s.password = password
	s.passwordfile = passwordfile
	s.thread = thread
	//read password file
	passwords := readFile(passwordfile) 
	if passwords != nil{
		s.passwords = passwords
	}
}