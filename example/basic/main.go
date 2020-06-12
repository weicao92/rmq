package main

import (
	"encoding/json"
	"github.com/adjust/rmq/v2"
	"log"
	"time"
)

type Task struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
type TaskConsumer struct {

}

func (consumer *TaskConsumer) Consume(delivery rmq.Delivery) {
	var task Task
	var err error
	if err = json.Unmarshal([]byte(delivery.Payload()), &task); err != nil {
		// handle error
		delivery.Reject()
		return
	}

	// perform task
	log.Printf("performing task %v", task)
	delivery.Ack()
}

func main() {
	connection := rmq.OpenConnection("basic service", "tcp", "weivu:6379", 2)
	taskQueue := connection.OpenQueue("tasks")
	task := Task{
		"Yang Zhen",
		39,
	}
	taskBytes, err := json.Marshal(task)
	if err != nil {
		log.Println("err when marshaling task")
		return
	}

	taskQueue.PublishBytes(taskBytes)
	taskQueue.StartConsuming(10, time.Second)

	taskConsumer := &TaskConsumer{}
	taskQueue.AddConsumer("task consumer", taskConsumer)

	log.Println("......")
	log.Println("...end...")

	finishedChan := taskQueue.StopConsuming()
	<- finishedChan
	log.Println("...exit...")
}
