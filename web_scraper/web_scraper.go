package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
 	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatal("Cannot create file", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("internshala.com"),
	)

	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("a"),
			e.ChildText("span"),
		})
	})

	for i := 1; i <= 194; i++ {
		fmt.Printf("Scraping page: %d\n", i)
		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Printf("Scraping complete, check file %q for results\n", fName)
	log.Println(c)

}