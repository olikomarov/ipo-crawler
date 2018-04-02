package main

import (
	"github.com/visheratin/ico-crawler/crawler/icorating"
	"github.com/visheratin/ico-crawler/misc"

	//"crawler-olga/crawler/icorating"
	//"crawler-olga/misc"

	"fmt"
)

func main() {
	fmt.Println("Let's go")
	misc.InitLog()
	config := misc.ReadConfig("config.json")
	manager := crawler.ICORatingCrawler{}
	err := manager.Init(config)
	if err != nil {
		misc.LogError(err)
	}


}
