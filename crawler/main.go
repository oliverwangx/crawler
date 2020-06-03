package main

import (
	"crawler/crawler/engine"
	"crawler/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: parser.ParseCityList,
	})
}
