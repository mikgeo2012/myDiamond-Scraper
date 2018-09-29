package scraper

import (
    "myDiamond-scraper/pkg/models"
    "fmt"
    "strings"


    "net/http"

    "os"
    "io/ioutil"
)

const (
    // James Allen's base URL
    JamesAllenBase = "https://www.jamesallen.com/JSite/Core/jx.ashx?PageUrl=loose-diamonds/"
)



type WholeSaler interface {
    Scrape(page int) (resp string)
    Parse(resp string) []models.Diamond
}

type JamesAllen struct {
    baseUrl string
    path string
}

func (j JamesAllen) Scrape(page int) (resp string) {
    response,err := http.Get(fmt.Sprintf("%s/%s", "https://www.jamesallen.com/JSite/Core/jx.ashx?PageUrl=loose-diamonds", j.path))
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))

    }
    return ""
}

func (j JamesAllen) Parse(resp string) []models.Diamond {
    return nil
}

type WebScraper struct {
    sites []WholeSaler

}

func (w *WebScraper) Initialize(o models.Options) {
    jaPath := fmt.Sprintf("round-cut/page-%s/?Color=%s&Cut=%s&Clarity=%s&PriceFrom=%f&PriceTo=%f&CaratFrom=%f&CaratTo=%f&Lab=%s", "2",
        strings.Join(o.Color, ","), strings.Join(o.Cut, ","), strings.Join(o.Clarity, ","), o.Price[0],
        o.Price[1], o.Carat[0], o.Carat[1], strings.Join(o.Lab, ","))
    w.sites = []WholeSaler{JamesAllen{JamesAllenBase, jaPath}}
    for _,s := range w.sites {
        _ = s.Scrape(3)

    }
}


