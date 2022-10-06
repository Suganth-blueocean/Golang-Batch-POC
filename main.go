package main

import (
	"fmt"
	"time"

	"github.com/blu-ocean/bo-common-module/unitofwork"
	"github.com/blu-ocean/bo-customer-service/service"
)

// These workers will receive work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// Two channels - to send them work and to collect their results.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts up 5 workers, initially blocked because there are no jobs yet.
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}
	a := 1
	unitOfWork := unitofwork.NewUnitOfWork()
	// Here we send 10 `jobs` and then `close` that channel to indicate that's all the work we have.
	for true {
		customers, err := service.GetAllCustomerByLimit(unitOfWork.Context, 100)
		if err != nil {
			return
		}
		jobs <- customers[a].Mobile
		a += 1
	}
	close(jobs)

	// collect all the results of the work.
	for a := 1; a <= 5; a++ {
		<-results
	}
}
