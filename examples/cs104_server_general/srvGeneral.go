package main

import (
	"log"
	"time"

	"github.com/thinkgos/go-iecp5/asdu"
	"github.com/thinkgos/go-iecp5/cs104"
)

func main() {
	srv, err := cs104.NewServer(&cs104.Config{}, asdu.ParamsWide, &mysrv{})
	if err != nil {
		panic(err)
	}
	srv.LogMode(true)
	// go func() {
	// 	time.Sleep(time.Second * 20)
	// 	log.Println("try ooooooo", err)
	// 	err := srv.Close()
	// 	log.Println("ooooooo", err)
	// }()
	srv.ListenAndServer(":2404")
}

type mysrv struct{}

func (this *mysrv) InterrogationHandler(c asdu.Connect, asduPack *asdu.ASDU, qoi asdu.QualifierOfInterrogation) error {
	log.Println("qoi", qoi)
	asduPack.SendReplyMirror(c, asdu.Actcon)
	err := asdu.Single(c, false, asdu.CauseOfTransmission{Cause: asdu.Inrogen}, asdu.GlobalCommonAddr,
		asdu.SinglePointInfo{})
	if err != nil {
		// log.Println("falied")
	} else {
		// log.Println("success")
	}
	// go func() {
	// 	for {
	// 		err := asdu.Single(c, false, asdu.CauseOfTransmission{Cause: asdu.Spont}, asdu.GlobalCommonAddr,
	// 			asdu.SinglePointInfo{})
	// 		if err != nil {
	// 			log.Println("falied", err)
	// 		} else {
	// 			log.Println("success", err)
	// 		}

	// 		time.Sleep(time.Second * 1)
	// 	}
	// }()
	asduPack.SendReplyMirror(c, asdu.Actterm)
	return nil
}
func (this *mysrv) CounterInterrogationHandler(asdu.Connect, *asdu.ASDU, asdu.QualifierCountCall) error {
	return nil
}
func (this *mysrv) ReadHandler(asdu.Connect, *asdu.ASDU, asdu.InfoObjAddr) error { return nil }
func (this *mysrv) ClockSyncHandler(asdu.Connect, *asdu.ASDU, time.Time) error   { return nil }
func (this *mysrv) ResetProcessHandler(asdu.Connect, *asdu.ASDU, asdu.QualifierOfResetProcessCmd) error {
	return nil
}
func (this *mysrv) DelayAcquisitionHandler(asdu.Connect, *asdu.ASDU, uint16) error { return nil }
func (this *mysrv) ASDUHandler(asdu.Connect, *asdu.ASDU) error                     { return nil }