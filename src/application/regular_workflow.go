package application

import (
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

// When it is deployed as a workload it periodically reports status by calling
// an API and submitting battery, cpu usage, cpu temp, disk space.
func (a App) RegularWorkflow() error {
	log.WithFields(log.Fields{"tag": a.Tag}).Info("regular workflow started")
	t := time.NewTicker(time.Second)
	defer t.Stop()
	waitChan := make(chan struct{}, 10)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	for {
		select {
		case <-t.C:
			waitChan <- struct{}{}
			go func() {
				a.Report()
				<-waitChan
				wg.Done()
			}()
			wg.Add(1)
		case <-signalChan:
			log.WithFields(log.Fields{"tag": a.Tag}).Info("interrupted by user")
			wg.Wait()
			return nil
		}
	}
	<-signalChan // second signal, hard exit
	os.Exit(2)
	return nil
}

func (a App) Report() {

	request, err := http.NewRequest("GET", a.Url+"serverinfo", nil)
	if err != nil {
		log.WithFields(log.Fields{
			"tag": a.Tag,
			"err": err,
		}).Error("get server info error")
	}	
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	var bearer = "Bearer " + a.ServerToken
	request.Header.Add("Authorization", bearer)

	// client := &http.Client{}
	// response, err := client.Do(request)
	// if err != nil {
	// 	log.WithFields(log.Fields{
	// 		"tag": a.Tag,
	// 		"err": err,
	// 	}).Error("get server info error")
	// }
	// defer response.Body.Close()

	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.WithFields(log.Fields{ "tag": a.Tag, "err": err,
	// 	}).Errorf("response read error")
	// }

	// data := struct{
	// 	Battery string `json:"battery"`
	// 	CpuUsage string `json:"cpu_usage"`
	// 	CpuTemp string `json:"cpu_temp"`
	// 	DiskSpace string `json:"disk_space"`
	// }{}
	// if err = json.Unmarshal(bodyBytes, &data); err != nil {
	// 	log.WithFields(log.Fields{ "tag": a.Tag, "err": err,
	// 	}).Error("response unmarshal error")

	log.WithFields(log.Fields{
		"tag":        a.Tag,
		"battery":    "test",
		"cpu usage":  "test",
		"cpu temp":   "test",
		"disk space": "test",
	}).Info("client info")
}
