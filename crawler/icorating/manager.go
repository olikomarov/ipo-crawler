package crawler

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/visheratin/ico-crawler/misc"

	"fmt"
)

var mainLink = "https://www.iposcoop.com/last-12-months/"

type ICORatingCrawler struct {
	workers []*ICORatingWorker
}

func (manager *ICORatingCrawler) Init(config misc.Configuration) error {
	links, err := manager.GetEntitiesLinks(mainLink)
	fmt.Println("main manager=" + mainLink)

	if err != nil {
		return err
	}
	workersLinks, err := misc.SplitLinks(links, config.WorkersNumber)
	if err != nil {
		return err
	}

	for i := 0; i < config.WorkersNumber; i++ {
		worker := &ICORatingWorker{
			id:    i,
			links: workersLinks[i],
		}
		fmt.Println("workersLinks[i]=" + links[i])
		manager.workers = append(manager.workers, worker)
		go func() {
			worker.Start()
		}()
	}
	timeout, err := time.ParseDuration(config.UpdateTimeout)
	if err != nil {
		timeout, _ = time.ParseDuration("5m")
	}
	for {
		time.Sleep(timeout)
		workersFinished := true
		for _, worker := range manager.workers {
			if !worker.finished {
				workersFinished = false
				break
			}
		}
		if workersFinished {
			break
		}
	}
	return nil
}

func (crawler *ICORatingCrawler) GetEntitiesLinks(mainPageLink string) ([]string, error) {
	doc, err := goquery.NewDocument(mainPageLink)
	fmt.Println("main crawler=" + mainPageLink)

	if err != nil {
		return nil, err
	}

	fmt.Println("hhhhhhhhhhhhhhhhhhh")
	result := []string{}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {

		title := s.Text()
		linkTag := s.Find("a")
		href, found := linkTag.Attr("href")
		fmt.Println(",  title="+title+",   href="+href)


		//href, found := s.Attr("href")

		if found {
			result = append(result, href)
			fmt.Println("resultkkkkkkkkkkkkkkkkkkkkk")
		}

	})
	return result, nil
}
