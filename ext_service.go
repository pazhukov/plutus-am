package plutus

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Text    string   `xml:",chardata"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []struct {
		Text     string `xml:",chardata"`
		ID       string `xml:"ID,attr"`
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Nominal  string `xml:"Nominal"`
		Name     string `xml:"Name"`
		Value    string `xml:"Value"`
	} `xml:"Valute"`
}

func GetCurrencyRatesCBR(onDate string) []CurrencyRate {

	url := "https://www.cbr.ru/scripts/XML_daily.asp"
	if onDate != "01/01/1900" {
		url = url + "?date_req=" + onDate
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return []CurrencyRate{}
	}
	defer resp.Body.Close()

	xmlData := ValCurs{}

	d := xml.NewDecoder(resp.Body)
	d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}
	err_dec := d.Decode(&xmlData)
	if err_dec != nil {
		log.Println("Parse error")
		log.Println(err_dec)
		return []CurrencyRate{}
	}

	var out []CurrencyRate

	var arrDate = strings.Split(xmlData.Date, ".")
	var newDate = arrDate[2] + "-" + arrDate[1] + "-" + arrDate[0]

	for _, element := range xmlData.Valute {

		currcencyID, err1 := strconv.Atoi(element.NumCode)
		if err1 != nil {
			log.Println(err1)
			continue
		}
		rate, err2 := strconv.ParseFloat(strings.Replace(element.Value, ",", ".", -1), 64)
		if err2 != nil {
			log.Println(err2)
			continue
		}
		nominal, err3 := strconv.ParseFloat(element.Nominal, 64)
		if err3 != nil {
			log.Println(err3)
			continue
		}

		var cr CurrencyRate
		cr.Currency = currcencyID
		cr.Period = newDate
		cr.Rate = math.Floor(rate/nominal*10000) / 10000

		out = append(out, cr)

	}

	return out
}
