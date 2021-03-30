package main

import (
	"fmt"
	"os"

	ftp "github.com/martinr92/goftp"
)

var uploadLog = func() string {
	ip_port := os.Args[1]
	user := os.Args[2]
	pw := os.Args[3]
	filePath := os.Args[4]
	out := os.Args[5]
	remoteDir := os.Args[6]

	//connect
	ftpClient, err := ftp.NewFtp(ip_port)
	if err != nil {
		panic(err)
	}
	defer ftpClient.Close()

	//send credentials
	if err = ftpClient.Login(user, pw); err != nil {
		panic(err)
	}

	//opens the system log directory
	if err = ftpClient.OpenDirectory(remoteDir); err != nil {
		panic(err)
	}

	//uploads the system log file. You can upload other files by opening a different directory and then uploading from that directory
	if err = ftpClient.Upload(filePath, out); err != nil {
		panic(err)
	}
	println("log uploaded :)")
	return "task completed."
}

func main() {

	if len(os.Args) < 5 {
		println("example usage: ./macLogReport_CLI.go IP:PORT USERNAME PASSWORD FILPATH FILEOUT REMOTE_DIRECTORY")
	}

	//define our variables from command line. Can optionally
	fmt.Println("starting connection. . .")
	uploadLog()

}
