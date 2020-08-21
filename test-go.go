package main
import (
	"fmt"
	"time"
	"github.com/robfig/cron"
	// cron "github.com/robfig/cron/v3"
)
func main() {
	// sh, _ := time.LoadLocation("Asia/Shanghai")
	// fmt.Println("sh = ", sh)
	// cron.WithLocation(sh)
	c := cron.New()
	c.Start()
	startTime := time.Now().Add(10*time.Second)
	fmt.Println("startTime = ", startTime)
	fmt.Println("next = ", startTime.Add(3*time.Second))
	fmt.Println("next = ", startTime.Add(6*time.Second))
	go func() {
		fmt.Println("now:", time.Now())
		fmt.Println("start time = ", startTime)
		c.AddFunc("@every 3s", func() {
			if time.Now().After(startTime) {
				fmt.Println("time = ", time.Now())
			}
		})
	}()
	time.Sleep(1000 * time.Second)
}