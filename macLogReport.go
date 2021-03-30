package main

import (
	"fmt"
	"os"

	ftp "github.com/martinr92/goftp"
)

var uploadLog = func() string {

	//connect
	ftpClient, err := ftp.NewFtp("192.168.88.247:21")
	if err != nil {
		panic(err)
	}
	defer ftpClient.Close()

	//send credentials
	if err = ftpClient.Login("YOUR USER HERE", "YOUR PW HERE"); err != nil {
		panic(err)
	}

	//opens the system log directory
	if err = ftpClient.OpenDirectory("/ftp/upload/macbookLogs/"); err != nil {
		panic(err)
	}

	//uploads the system log file. You can upload other files by opening a different directory and then uploading from that directory
	if err = ftpClient.Upload("/var/log/daily.out", "daily.out"); err != nil {
		panic(err)
	}
	println("log uploaded :)")
	return "task completed."
}

func main() {

	if len(os.Args) < 3 {
		println("please provide a host address, username, and password")
	}

	//define our variables from command line. Can optionally
	fmt.Println("starting connection. . .")
}
