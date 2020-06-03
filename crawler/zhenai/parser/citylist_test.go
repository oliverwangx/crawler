package parser

import (
	"crawler/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := fetcher.Fetch(
		"http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	ParseCityList(contents)

	// verify results
}
