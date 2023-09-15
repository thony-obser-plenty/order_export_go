package providers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Orders struct {
	Orders []Order `json:"data"`
}

type Order struct {
	Id         *int        `json:"Id"`
	Date       *string     `json:"Date"`
	StatusId   *int        `json:"StatusId"`
	Address    *string     `json:"Address"`
	OrderItems []OrderItem `json:"OrderItems"`
}

type OrderItem struct {
	Id          *int    `json:"Id"`
	CreatedAt   *string `json:"CreatedAt"`
	UpdatedAt   *string `json:"UpdatedAt"`
	DeletedAt   *string `json:"DeletedAt"`
	ProductName *string `json:"ProductName"`
	OrderId     *int    `json:"OrderId"`
}

type PageCount struct {
	PageCount int `json:"pageCount"`
}

func FetchOrders(url, token string, pageId int) (*Orders, error) {
	url = strings.ReplaceAll(url, "{page}", strconv.Itoa(pageId))
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var Orders Orders
	err = json.Unmarshal(body, &Orders)
	if err != nil {
		return nil, err
	}

	return &Orders, nil
}

func FetchPageCount(url, token string) (*PageCount, error) {
	url = strings.ReplaceAll(url, "{page}", strconv.Itoa(1))
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("token", token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var PageCount PageCount
	err = json.Unmarshal(body, &PageCount)
	if err != nil {
		return nil, err
	}

	return &PageCount, nil
}
