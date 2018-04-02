package crawler

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	model "github.com/visheratin/ico-crawler/model/icorating"
	"github.com/visheratin/ico-crawler/writer"
	"fmt"
)

type ICORatingWorker struct {
	id       int
	finished bool
	pageType string
	links    []string
}

func (worker *ICORatingWorker) Start() error {
	for _, link := range worker.links {
		entity, _ := worker.GetDetails(link)
		outputPath := "./data/icorating/"
		outFilename := entity.Title + ".json"
		writer.WriteToFS(outputPath, outFilename, entity)
		fmt.Println(outFilename)
	}
	return nil
}

func (worker *ICORatingWorker) GetDetails(detailsLink string) (model.ICORatingCompany, error) {
	fmt.Println("detailsLink="+detailsLink)
	doc, err := goquery.NewDocument(detailsLink)
	if err != nil {
		return model.ICORatingCompany{}, err
	}
	result := model.ICORatingCompany{}

	//titleNode := doc.Find("h1")
	titleNode := doc.Find("h3")

	if len(titleNode.Nodes) > 0 {
		result.Title = titleNode.Text()
	}
	tableCells := doc.Find("td")
	for i := range tableCells.Nodes {
		cell := tableCells.Eq(i)
		text := cell.Text()
		fmt.Println("text="+text)

		if text == "Industry:" {
			result.Industry = clearText(cell.Siblings().Text())
		}
		if text == "Employees:" {
			result.Employees = clearText(cell.Siblings().Text())
		}
		if text == "Address:" {
			result.Address = clearText(cell.Siblings().Text())
		}
		if text == "Phone Number" {
			result.Phone = clearText(cell.Siblings().Text())
		}
		if text == "Web Address" {
			result.Web_address = clearText(cell.Siblings().Text())
		}
		if text == "Market Cap" {
			result.Market_cup = clearText(cell.Siblings().Text())
		}
		if text == "Revenues" {
			result.Revenues = clearText(cell.Siblings().Text())
		}
		if text == "Net Income" {
			result.Net_income = clearText(cell.Siblings().Text())
		}
		if text == "Symbol" {
			result.Symbol = clearText(cell.Siblings().Text())
		}
		if text == "Exchange" {
			result.Exchange = clearText(cell.Siblings().Text())
		}
		if text == "Shares (millions):" {
			result.Shares = clearText(cell.Siblings().Text())
		}
		if text == "Price range" {
			result.Price_range = clearText(cell.Siblings().Text())
		}
		if text == "Est. $ Volume" {
			result.Est_volume = clearText(cell.Siblings().Text())
		}
		if text == "Manager / Joint Managers" {
			result.Manager = clearText(cell.Siblings().Text())
		}
		if text == "CO-Managers" {
			result.CO_managers = clearText(cell.Siblings().Text())
		}
		if text == "Expected To Trade:" {
			result.Exp_to_trade = clearText(cell.Siblings().Text())
		}
		if text == "Status: " {
			result.Status = clearText(cell.Siblings().Text())
		}
		if text == "Quiet Period Expiration Date:" {
			result.Quiet_period = clearText(cell.Siblings().Text())
		}
		if text == "Lock-Up Period Expiration Date:" {
			result.Lock_up_period = clearText(cell.Siblings().Text())
		}
		if text == "SCOOP Rating" {
			result.Scoop_rate = clearText(cell.Siblings().Text())
		}
		if text == "Rating Change" {
			result.Rating_change = clearText(cell.Siblings().Text())
		}
		if text == "Business:" {
			result.Business = clearText(cell.Siblings().Text())
		}

	}
	return result, nil
}



func clearText(input string) string {
	output := strings.Replace(input, "\n", "", -1)
	output = strings.TrimSpace(output)
	return output
}

