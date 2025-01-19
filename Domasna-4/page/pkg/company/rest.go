package company

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetManyCompanies() ([]Company, error) {
	endpoint := fmt.Sprintf("http://%s%s/%s/%s/companies/",
		os.Getenv("MSEMK_DOMAIN"),
		os.Getenv("MSEMK_PORT"),
		"api",
		os.Getenv("MSEMK_VERSION"))
	res, err := http.Get(endpoint)
	if err != nil {
		errMsg := fmt.Errorf("error getting companies: %v", err)
		return nil, errMsg
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		errMsg := fmt.Errorf("error reading response body: %v", err)
		return nil, errMsg
	}
	var companies []Company
	if err := json.Unmarshal(body, &companies); err != nil {
		errMsg := fmt.Errorf("error unmarshalling response body: %v", err)
		return nil, errMsg
	}
	return companies, nil
}

func GetCompanyByTicker(tkr string) (CompanyDetailedResponse, error) {
	ticker := tkr
	endpoint := fmt.Sprintf("http://%s%s/%s/%s/companies/%s",
		os.Getenv("MSEMK_DOMAIN"),
		os.Getenv("MSEMK_PORT"),
		"api",
		os.Getenv("MSEMK_VERSION"),
		ticker)
	res, err := http.Get(endpoint)
	if err != nil {
		errMsg := fmt.Errorf("error getting companies: %v", err)
		return CompanyDetailedResponse{}, errMsg
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		errMsg := fmt.Errorf("error reading response body: %v", err)
		return CompanyDetailedResponse{}, errMsg
	}
	var company CompanyDetailedResponse
	if err := json.Unmarshal(body, &company); err != nil {
		errMsg := fmt.Errorf("error unmarshalling response body: %v", err)
		return CompanyDetailedResponse{}, errMsg
	}
	return company, nil
}
