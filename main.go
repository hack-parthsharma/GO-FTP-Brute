package main

import (
	"github.com/cyberkhalid/gosec/goftpbrute/goftpbrute"
	"flag"
	"sync"
)

func main()  {
	//cmd flags
	//cmd flags
	hostname := flag.String("host","","Hostname or ip of the target.")
	port := flag.String("port", "21", "Port number of target.")
	user := flag.String("user","","username of target.")
	//password := flag.String("p","","password of the target")
	thread := flag.Int("t",5,"number of thread to use")
	passwordfile := flag.String("P","","password file")
	flag.Parse()
	

	var goFtpBrute goftpbrute.GoFtpBrute
	goFtpBrute.Init(*hostname,*port,*user,"",*passwordfile, *thread) 
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	
	goFtpBrute.Bruteforce(&wg,cond)
	wg.Wait()
}