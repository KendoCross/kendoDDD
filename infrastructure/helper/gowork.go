package helper

import "sync"

type GoWorker interface {
	Task()
}

type workerPool struct {
	work chan GoWorker
	wg   sync.WaitGroup
}

func NewWorkerPool(maxGoRoutines int) *workerPool {
	p := workerPool{
		work: make(chan GoWorker),
	}
	p.wg.Add(maxGoRoutines)
	for i := 0; i < maxGoRoutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

//会等待之前的任务完成之后，才继续。
func (p *workerPool) Run(w GoWorker) {
	p.work <- w
}

//只是等待最后 maxGoRoutines 个GoRoutine的完成
func (p *workerPool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

type workerSyncPool struct {
	maxGoRoutines int
	workerPool
}

//runtime.NumCPU()
func NewSyncPool(maxGoRoutines int) *workerSyncPool {
	p := workerSyncPool{maxGoRoutines: maxGoRoutines}
	return &p
}

//提交多个任务,并等待执行完成
func (p *workerSyncPool) Run(ws []GoWorker) {
	p.work = make(chan GoWorker, len(ws))
	p.wg.Add(p.maxGoRoutines)
	for i := 0; i < p.maxGoRoutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	for i := range ws {
		p.work <- ws[i]
	}
	p.Shutdown()
}
