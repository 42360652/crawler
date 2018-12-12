package main

import (
	"moscow/crawler/engine"
	"moscow/crawler/scheduler"
	"moscow/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:  &scheduler.QueuedScheduler{},
		WorkeCount: 10,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
