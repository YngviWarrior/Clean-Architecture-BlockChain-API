package apis

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func findTransaction() {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://rest-sandbox.coinapi.io/", nil)

	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-CoinAPI-Key", "540A13CE-0504-41A7-9D95-0494DC440E37")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println(bodyBytes)
}
