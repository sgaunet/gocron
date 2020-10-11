package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"sgaunet/gocron/cronline"

	"github.com/robfig/cron"
)

func execIt(command []string) error {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	var crontabFile string
	var err error
	sigs := make(chan os.Signal, 1)

	flag.StringVar(&crontabFile, "f", "", "Crontab file")
	flag.Parse()

	file, err := os.Open(crontabFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	c := cron.New()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ligne := scanner.Text()
		fmt.Println("COMMAND=", cronline.GetCommand(ligne))
		fmt.Println("CRON=", cronline.GetCron(ligne))
		fmt.Println("")

		c.AddFunc(cronline.GetCron(ligne), func() { execIt(cronline.SafeSplit(cronline.GetCommand(ligne))) })
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	c.Start()
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	_ = <-sigs
	c.Stop()
}
