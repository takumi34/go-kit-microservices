package main


import (
    "context"
    "flag"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
		"syscall"
)

func main (){

	var (
        httpAddr = flag.String("http", ":3000", "http listen address")
		)
		
		// TODO make endpoint
}