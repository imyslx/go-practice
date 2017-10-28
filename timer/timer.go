package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

// MonTick : Structure for ticker and parameters.
type MonTick struct {
	Tick     *time.Ticker
	sec      int
	text     string
	exitChan chan bool
}

// Create Ticker.
func createMonTick(sec int, str string) (mt MonTick) {

	tick := time.NewTicker(time.Second * time.Duration(sec))
	mt = MonTick{tick, sec, str, make(chan bool)}

	return
}

func createWorkers() (mtArray []MonTick) {
	log.Print("Creating Monitors... ")

	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 1)
		rand.Seed(time.Now().UnixNano())
		sec := rand.Intn(60) + 1
		ssec := "job-" + strconv.Itoa(sec)
		log.Print("    Created: ", ssec)
		mtArray = append(mtArray, createMonTick(sec, ssec))
	}
	log.Print("Created Monitors... [ OK ]")

	return
}

func runAllTickers(mtArray []MonTick) {
	log.Print("Running All Monitors... ")
	for _, v := range mtArray {
		v.work()
	}
	log.Print("[ OK ]")
}

func exitAllTickers(mtArray []MonTick) {
	log.Print("Exiting All Monitors... ")
	for _, v := range mtArray {
		v.exitChan <- true
	}
}

// Definition the worker to monitoring by ticker.
func (mt MonTick) work() {

	go func() {
	forLoop:
		for {
			select {
			case <-mt.Tick.C:
				log.Print(mt.text + ": I'm working !")
			case <-mt.exitChan:
				break forLoop
			}
		}
		log.Print("Exited: " + mt.text)
	}()
}

//func readConfig() {
//	file, err := os.Open(`/etc/ismon/ismon.conf`)
//}

// Main function.
func main() {

	mtArray := createWorkers()
	runAllTickers(mtArray)

	for {
		time.Sleep(time.Second * 10)
		break
	}
	exitAllTickers(mtArray)
	time.Sleep(time.Second * 5)
	log.Print("Exit the application. ")

}
