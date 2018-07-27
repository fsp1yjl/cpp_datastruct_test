package main

import (
	"fmt"
	"net/http"
)

var (
	MaxWorkers = 100
	MaxQueue = 1000
)
var JobQueue chan string = make(chan string, MaxQueue)



type Worker struct {
	WorkerPool chan chan string
	JobChannel chan string
	quit chan bool
}

func NewWorker(workerpool chan chan string) Worker {
	return Worker{
		WorkerPool: workerpool,
		JobChannel: make(chan string),
		quit: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select{
			case job := <- w.JobChannel:
					fmt.Println("worker process::", job)
			case <-w.quit:
				return
			}
			// fmt.Println("worker start ...") 
		}
		
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}


type Dispatcher struct {
	WorkerPool chan chan string
}


func NewDispatcher(maxWorkers int) *Dispatcher{
	pool := make(chan chan string , maxWorkers)

	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	for i :=0; i< MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	 go d.dispatch()
}



func (d *Dispatcher) dispatch() {

	for {
		select {
		case job := <-JobQueue:
			go func(job string) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}

func main() {

	d := NewDispatcher(MaxWorkers)
	d.Run()
	
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		JobQueue <- "333"
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.ListenAndServe(":80", nil)

	select {}

}
