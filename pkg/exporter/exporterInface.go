package exporter

import "order_export_go/pkg/queue"

type ExporterContract interface {
	Export(queue queue.Queue, queueName string) bool
}
