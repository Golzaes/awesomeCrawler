package Engine

import (
	"github.com/payne/awesomeCrawler/Fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	configureWorkChan(chan Request)
}

type SimpleScheduler struct {
	workerCh chan Request
}

func (s SimpleScheduler) Submit(r Request) {
	s.workerCh <- r
}

func (s *SimpleScheduler) configureWorkChan(c chan Request) {
	s.workerCh = c

}

// ConcurrentRun Concurrent execute
func (c *ConcurrentEngine) ConcurrentRun(seed ...Request) {
	taskCh := make(chan Request)
	resultCh := make(chan ParseResult)
	c.Scheduler.configureWorkChan(taskCh)
	for i := 0; i <= c.WorkCount; i++ {
		CreateWorker(taskCh, resultCh)
	}
	for _, request := range seed {
		c.Scheduler.Submit(request)
	}
	for {
		resultCh := <-resultCh
		for _, item := range resultCh.Item {
			log.Printf(`Got item: %s`, item)
		}
		for _, request := range resultCh.Request {
			c.Scheduler.Submit(request)
		}
	}
}

// CreateWorker
func CreateWorker(taskCh <-chan Request, resultCh chan<- ParseResult) {
	go func() {
		for {
			Getter := <-taskCh
			result, err := worker(Getter)
			if err != nil {
				log.Printf(`Woker Error: %s`, err)
				continue
			}
			resultCh <- result
		}
	}()
}

func worker(getter Request) (ParseResult, error) {
	log.Printf(`Fetching URL:%s`, getter.URL)
	body, err := Fetcher.Fetch(getter.Method, getter.URL)
	if err != nil {
		log.Printf(`Worker Fetch Error %s`, getter.URL)
		return ParseResult{}, err
	}
	return getter.ParseFunc(body), nil
}

// ==============================================
//type ConcurrentEngine struct {
//	Scheduler Scheduler
//	WorkCount int
//}
//type Scheduler interface {
//	Submit(Request)
//	configureWorkChan(chan Request)
//}
//type SimpleScheduler struct {
//	workerChan chan Request
//}
//
//func (s *SimpleScheduler) Submit(r Request) {
//	s.workerChan <- r
//}
//
//func (s *SimpleScheduler) configureWorkChan(c chan Request) {
//	s.workerChan = c
//}
//
//func (c *ConcurrentEngine) Run(seeds ...Request) {
//	in := make(chan Request)
//	out := make(chan ParseResult)
//	for i := 0; i <= c.WorkCount; i++ {
//		CreateWork(in, out)
//	}
//	for _, seed := range seeds {
//		c.Scheduler.Submit(seed)
//	}
//}
//
//func CreateWork(in chan Request, out chan ParseResult) {
//	go func() {
//		for {
//			request := <-in
//			result, err := worker(request)
//			if err != nil {
//				continue
//			}
//			out <- result
//		}
//	}()
//}
//
//func worker(r Request) (ParseResult, error) {
//	fmt.Printf(`Fech URL %s`, r.URL)
//	body, err := Fetcher.Fetch(r.Method, r.URL)
//	if err != nil {
//		fmt.Printf(`Worker Err, %s`, err)
//		return ParseResult{}, err
//	}
//	return r.ParseFunc(body), err
//
//}
