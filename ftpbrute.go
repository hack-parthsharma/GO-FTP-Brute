package goftpbrute

type GoFtpBrute struct {
	hostname string
	port string
	thread int
	username string
	password string
	passwordfile string
	passwords []byte
}