package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func Job() {
	fmt.Println(time.Now())
}

func main() {
	c := cron.New()

	c.AddFunc(
		"* * * * *",
		Job,
	)

	c.Start()

	select {}
}
