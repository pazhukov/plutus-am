package plutus

import (
	"encoding/xml"
)

type CBRCurrencyRate struct {
	Text     string `xml:",chardata"`
	ID       string `xml:"ID,attr"`
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal  string `xml:"Nominal"`
	Name     string `xml:"Name"`
	Value    string `xml:"Value"`
}

type CBRRoot struct {
	XMLName xml.Name          `xml:"ValCurs"`
	Text    string            `xml:",chardata"`
	Date    string            `xml:"Date,attr"`
	Name    string            `xml:"name,attr"`
	Valute  []CBRCurrencyRate `xml:"Valute"`
}

func GetCurrencyRatesCBR(onDate string) []CurrencyRate {

	// resp, err := http.Get("https://www.cbr.ru/scripts/XML_daily.asp")
	// if err != nil {
	// 	log.Fatalln(err) // log.Fatal always exits the program, need to check err != nil first
	// }
	// defer resp.Body.Close()

	// xmlData := &CBRRoot{}

	// body, _ := ioutil.ReadAll(resp.Body)
	// err = xml.Unmarshal(body, xmlData)
	// log.Println(xmlData.Date)

	// d := xml.NewDecoder(resp.Body)
	// d.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
	// 	switch charset {
	// 	case "windows-1251":
	// 		content, _ := ioutil.ReadAll(resp.Body)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		//return charmap.Windows1251.NewDecoder().Reader(input), nil
	// 		return bytes.NewReader(content), nil
	// 	default:
	// 		return nil, fmt.Errorf("unknown charset: %s", charset)
	// 	}
	// }

	// err_dec := d.Decode(&xmlData)
	// if err_dec != nil {
	// 	log.Println(err_dec)
	// 	return []CurrencyRate{}
	// }

	//if err = xml.NewDecoder(resp.Body).Decode(&xmlData); err != nil {
	//	log.Fatalln(err)
	//}

	return []CurrencyRate{}
}
