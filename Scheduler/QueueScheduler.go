package Scheduler

//import (
//	"github.com/payne/awesomeCrawler/Engine"
//)
//
//type QueueScheduler struct {
//	RequestChan chan Engine.Request
//	WorkChan    chan chan Engine.Request
//}
//
//func (q *QueueScheduler) Submit(r Engine.Request) {
//	q.RequestChan <- r
//}
//
//func (q *QueueScheduler) WorkReady(w chan Engine.Request)  {
//	q.WorkChan <- w
//}
//func (q *QueueScheduler) QueueRUn()  {
//	go func() {
//		for {
//			select {
//			case r<-q.RequestChan:
//				requestQueue = append(q.RequestChan, r)
//			case w := <- q.WorkChan:
//				q.WorkChan <- append()
//
//			}
//		}
//	}()
//}
