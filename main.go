package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/lajunta/myauthd/grpcd"
	"github.com/lajunta/myauthd/utils"
)

var (
	host     string
	port     string
	restPort string
	rest     bool
)

func init() {
	flagParse()

	//config grpcd variables
	grpcd.GrpcdAddress = host + ":" + port
	grpcd.Rest = rest
	grpcd.RestPort = restPort
}

func flagParse() {
	flag.StringVar(&port, "port", "5050", "server running port")
	flag.StringVar(&host, "host", "0.0.0.0", "server host ip ")
	flag.StringVar(&restPort, "P", "8081", "restful http port")
	flag.BoolVar(&rest, "r", false, "weather open rest auth")
	flag.Parse()
}

// handle program exit event
func handleExit() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigc
		fmt.Println(s.String())
		err := utils.DB.Close()
		if err != nil {
			log.Fatalln(err.Error())
		} else {
			log.Println("Program Exit Ok and DB Close Wonderful")
			os.Exit(0)
		}
	}()
}

func main() {
	handleExit()
	go utils.CheckDB()
	grpcd.Serve(host, port)
}
