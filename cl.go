package clparser

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"log"
	"strings"
)

var (
	fp = gofeed.NewParser()
)

/*
CraigsListParser is an object that scrapes CraigsList.
*/
type CraigsListParser struct {
	city     string
	category string
	terms    []string
}

func NewCraigsListParser() *CraigsListParser {
	log.Printf("Created new CraigsListParser\n")
	return &CraigsListParser{}
}

func (c *CraigsListParser) Get() ([]string, error) {
	url := fmt.Sprintf("https://%s.craigslist.org/search/%s?format=rss", c.city, c.category)
	if len(c.terms) > 0 {
		query := ""
		for i := range c.terms {
			query += strings.TrimSpace(c.terms[i])
			if i != len(c.terms)-1 {
				query += "+"
			}
		}
		url = fmt.Sprintf("https://%s.craigslist.org/search/%s?query=%s&format=rss", c.city, c.category, query)
	}
	feed, err := fp.ParseURL(url)

	if err != nil {
		return []string{}, err
	}



	return []string{}, nil
}

// SetCity changes the city we're looking at
func (c *CraigsListParser) SetCity(name string) {
	c.city = name
}

// SetCategory changes the category we're looking at
func (c *CraigsListParser) SetCategory(name string) {
	c.category = name
}

// SetTerms changes the search terms that we're interested in
func (c *CraigsListParser) SetTerms(name []string) {
	c.terms = name
}

// AddTerms adds some strings to search for
func (c *CraigsListParser) AddTerms(name ...string) {
	for i := range name {
		c.terms = append(c.terms, name[i])
	}
}

// DeleteTerms gets rid of some search terms
func (c *CraigsListParser) DeleteTerms(name ...string) {
	fmt.Println(name)
	fmt.Println(c.terms)
	for i := range name {
		for j := range c.terms {
			if c.terms[j] == name[i] {
				fmt.Println("found it")
				fmt.Println(c.terms[:j])
				fmt.Println(c.terms[j])
				fmt.Println(c.terms[j+1:])
				toRet := make([]string, 0)
				for k := 0; k < j; k++ {
					toRet = append(toRet, c.terms[k])
				}
				for k := j + 1; k < len(c.terms); k++ {
					toRet = append(toRet, c.terms[k])
				}
				c.terms = toRet
				fmt.Println("now...")
				fmt.Println(c.terms)
				break
			}
		}
	}
}
