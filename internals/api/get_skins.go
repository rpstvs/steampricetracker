package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rpstvs/steamprice-api/internals/utils"
)

func (cfg *Client) GetSkins(start int) utils.SearchResult {
	url := GetUrl(start)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("error occurred: %s", err)
		return utils.SearchResult{}
	}

	resp, err := cfg.httpClient.Do(req)

	if err != nil {
		fmt.Printf("error occyrred: %s", err)
		cfg.GetSkins(start)
		return utils.SearchResult{}
	}

	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("error: %s", err)
		return utils.SearchResult{}
	}

	searchResult := &utils.SearchResult{}

	err = json.Unmarshal(dat, &searchResult)

	if err != nil {
		fmt.Printf("error occurred: %s ", err)
		return utils.SearchResult{}
	}
	//utils.WriteToFile(*searchResult)
	//utils.ParseResults(*searchResult)

	return *searchResult
}
