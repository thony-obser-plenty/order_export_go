package main

import (
	"order_export_go/pkg/exporter"
)

func main() {
	var exporterService exporter.ExporterServiceContract = exporter.ExporterService{}
	result := exporterService.InitExport()
	println(result)
}
