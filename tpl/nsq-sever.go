package tpl

// NSQSERVER NSQ 服务端代码
const NSQSERVER = `package main

import (
	"os"

	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
)

var workChan chan *nsq.Message
var producer *nsq.Producer

const (
	modulename = "nsq-server"
)

// DataAgent 实现NSQ的HandleMessage函数
type DataAgent struct{}

// HandleMessage 处理NSQ发来的消息
func (da *DataAgent) HandleMessage(m *nsq.Message) error {
	m.DisableAutoResponse()
	workChan <- m
	return nil
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	workChan = make(chan *nsq.Message)

	cfg := nsq.NewConfig()
	cfg.MaxInFlight = 1000

	r, err := nsq.NewConsumer("TOPIC-NAME", modulename, cfg)
	if err != nil {
		logrus.WithFields(logrus.Fields{"Connect Nsq Build Topic Error": err}).Error(modulename)
		os.Exit(-1)
	}

	go func() {
		logrus.WithFields(logrus.Fields{"WorkChan Listen Status": "Listen..."}).Debug(modulename)
		for m := range workChan {
			logrus.WithFields(logrus.Fields{"Recevie NSQ Request": string(m.Body)}).Debug(modulename)

			// err = json.Unmarshal(m.Body, &req)
			// if err != nil {
			// 	logrus.WithFields(logrus.Fields{"Unmarshal Request Error": err}).Error(ModelName)
			// 	continue
			// }

			m.Finish()
		}
	}()
	r.AddConcurrentHandlers(&DataAgent{}, 20)
	err = r.ConnectToNSQD("NSQ-Endpoint")
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	<-r.StopChan
}
`
