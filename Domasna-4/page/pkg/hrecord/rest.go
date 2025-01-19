package hrecord

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetTopTen() ([]HRecord, error) {
	endpoint := fmt.Sprintf("http://%s%s/%s/%s/records/%s",
		os.Getenv("MSEMK_DOMAIN"),
		os.Getenv("MSEMK_PORT"),
		"api",
		os.Getenv("MSEMK_VERSION"),
		"top-ten")
	res, err := http.Get(endpoint)
	if err != nil {
		errMsg := fmt.Errorf("error getting top ten records: %v", err)
		return nil, errMsg
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		errMsg := fmt.Errorf("error reading response body: %v", err)
		return nil, errMsg
	}
	var records []HRecord
	if err := json.Unmarshal(body, &records); err != nil {
		errMsg := fmt.Errorf("error unmarshalling response body: %v", err)
		return nil, errMsg
	}
	return records, nil
}
