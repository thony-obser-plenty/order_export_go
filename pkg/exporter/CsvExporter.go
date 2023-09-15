package exporter

import (
	"encoding/csv"
	"fmt"
	"order_export_go/pkg/providers"
	"order_export_go/pkg/queue"
	"os"
	"path/filepath"
	"strconv"
)

type Exporter struct{}

func (exporter Exporter) Export(queue queue.Queue, queueName string) bool {
	filePath := getFilePath()
	file := createCsvFile(filePath, queueName)
	writeFile(file, queue)

	return true
}

func writeFile(file *os.File, queue queue.Queue) {
	writer := csv.NewWriter(file)
	defer writer.Flush()
	defer file.Close()

	header := []string{"Id", "Date", "StatusId", "Address", "OrderItems"}

	if err := writer.Write(header); err != nil {
		panic(fmt.Sprintf("Error writing header: %v", err))
	}

	for queue.Peek() != nil {
		var order providers.Order
		order = *queue.Dequeue()
		orderItems := order.OrderItems
		orderItemsString := ""

		for _, orderItem := range orderItems {
			orderItemsString += addIntField(orderItem.Id) + ":"
			orderItemsString += addStringField(orderItem.CreatedAt) + ":"
			orderItemsString += addStringField(orderItem.UpdatedAt) + ":"
			orderItemsString += addStringField(orderItem.DeletedAt) + ":"
			orderItemsString += addStringField(orderItem.ProductName) + ":"
			orderItemsString += addIntField(orderItem.OrderId)
		}

		record := []string{
			addIntField(order.Id),
			addStringField(order.Date),
			addIntField(order.StatusId),
			addStringField(order.Address),
			orderItemsString,
		}

		if err := writer.Write(record); err != nil {
			panic(fmt.Sprintf("Error writing record: %v", err))
		}
	}
}

func addStringField(field *string) string {
	if field != nil {
		return *field
	}

	return ""
}

func addIntField(field *int) string {
	if field != nil {
		return strconv.Itoa(*field)
	}

	return ""
}

func createCsvFile(filePath string, queueName string) *os.File {
	file, err := os.Create(filePath + "/" + queueName + ".csv")

	if err != nil {
		panic("Error creating file")
	}

	return file
}

func getFilePath() string {
	dir, err := os.Getwd()

	if err != nil {
		panic("Error getting current directory")
	}

	filePath := filepath.Join(dir, "output/export")

	return filePath
}
