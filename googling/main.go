package main

import (
	"hello/googlescraper"
	"os"

	"github.com/labstack/echo"
)

const fileName string = "results.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	searchTerm := c.FormValue("searchTerm")
	countryCode := c.FormValue("countryCode")
	languageCode := c.FormValue("languageCode")
	googlescraper.GoogleScrape(searchTerm, countryCode, languageCode)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("https://geek-or-nerd.github.io/googling/", handleHome)
	e.POST("googling/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":443"))
	// TODO: 演算子の入力
	// var keywords = []string{"site:facebook.com | site:twitter.com | site:youtube.com \"女女し\""}
	// for _, keyword := range keywords {
	// 	res, _ := googlescraper.GoogleScrape(keyword, "jp", "jp")
	// 	fmt.Println("以下のキーワードの検索結果を収集します。")
	// 	fmt.Println(keyword)
	// 	for _, item := range res {
	// 		fmt.Println(item)
	// 	}
	// 	time.Sleep(time.Second * 5)
	// 	fmt.Println("検索結果の収集が終わりました。")
	// }
}
