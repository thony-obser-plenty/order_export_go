package exporter

import (
	"log"
	"order_export_go/pkg/providers"
	"order_export_go/pkg/queue"
	"strings"
	"sync"
	"time"
)

type ExporterService struct{}

func (e ExporterService) InitExport() bool {
	routes := getRoutes()
	parameters := getParameters()

	queueCollection := setupQueueCollection()
	token := getToken(parameters, routes)
	pageCount := getPageCount(parameters, routes, token)
	setupOrders(pageCount, parameters, routes, token, queueCollection)

	waitGroup := sync.WaitGroup{}
	for queueName, monthQueue := range queueCollection.GetQueues() {
		waitGroup.Add(1)
		go func(monthQueue queue.Queue, queueName string) {
			defer waitGroup.Done()

			var exporter ExporterContract = Exporter{}
			exporter.Export(monthQueue, queueName)
		}(*monthQueue, queueName)
	}
	waitGroup.Wait()
	return true
}

func getPageCount(parameters *providers.Parameters, routes *providers.Routes, token *providers.Token) *providers.PageCount {
	pageCount, error := providers.FetchPageCount(parameters.ApiBaseScheme+"://"+parameters.ApiBaseUri+routes.ApiOrderRoute.Path, token.Token)

	if error != nil {
		panic("Page count could not be fetched")
	}

	return pageCount
}

func getToken(parameters *providers.Parameters, routes *providers.Routes) *providers.Token {
	token, error := providers.FetchToken(parameters.ApiBaseScheme+"://"+parameters.ApiBaseUri+routes.ApiAuthenticatorRoute.Path, parameters.ApiKey)

	if error != nil {
		panic("Token could not be fetched")
	}

	return token
}

func getParameters() *providers.Parameters {
	parameters, error := providers.GetParameters("/Users/thonyobser/Documents/dev/excercise_2/order_export_go/configs/parameters.yml")

	if error != nil {
		panic("Parameters could not be loaded")
	}

	return parameters
}

func getRoutes() *providers.Routes {
	routes, error := providers.GetRoutes("/Users/thonyobser/Documents/dev/excercise_2/order_export_go/configs/routes.yml")

	if error != nil {
		panic("Routes could not be loaded")
	}

	return routes
}

func setupOrders(pageCount *providers.PageCount, parameters *providers.Parameters, routes *providers.Routes, token *providers.Token, queueCollection *queue.QueueCollection) {
	for i := 1; i <= pageCount.PageCount; i++ {
		orders, err := providers.FetchOrders(parameters.ApiBaseScheme+"://"+parameters.ApiBaseUri+routes.ApiOrderRoute.Path, token.Token, i)

		if err != nil {
			log.Fatalf("Request failed: %v", err)
		}

		for _, order := range orders.Orders {
			timestamp, err := time.Parse(time.RFC3339, *order.Date)

			if err != nil {
				log.Fatalf("Date parsing failed: %v", err)
			}
			month := strings.ToLower(timestamp.Month().String())

			queueCollection.GetQueue(month).Enqueue(&order)
		}
	}
}

func setupQueueCollection() *queue.QueueCollection {
	queueCollection := queue.NewQueueCollection()

	queueCollection.AddQueue("january")
	queueCollection.AddQueue("february")
	queueCollection.AddQueue("march")
	queueCollection.AddQueue("april")
	queueCollection.AddQueue("may")
	queueCollection.AddQueue("june")
	queueCollection.AddQueue("july")
	queueCollection.AddQueue("august")
	queueCollection.AddQueue("september")
	queueCollection.AddQueue("october")
	queueCollection.AddQueue("november")
	queueCollection.AddQueue("december")

	return queueCollection
}
