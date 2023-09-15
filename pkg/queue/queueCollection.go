package queue

type QueueCollection struct {
	queues map[string]*Queue
}

func NewQueueCollection() *QueueCollection {
	return &QueueCollection{
		queues: make(map[string]*Queue),
	}
}

func (qc *QueueCollection) AddQueue(name string) {
	qc.queues[name] = NewQueue()
}

func (qc *QueueCollection) GetQueue(name string) *Queue {
	return qc.queues[name]
}

func (qc *QueueCollection) GetQueues() map[string]*Queue {
	return qc.queues
}
