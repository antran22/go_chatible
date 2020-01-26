package worker_pool

type job struct {
	JobData      interface{}
	retryCount   int
	errorChannel chan error
}

type Worker interface {
	Work(JobData interface{}) error
}

type WorkerMaker interface {
	MakeWorker(pool *WorkerPool) Worker
}

type WorkerPool struct {
	jobChannel    chan job
	PoolData      interface{}
	maxRetryCount int
	maker         WorkerMaker
}

func NewPool(poolData interface{}, maker WorkerMaker, workerCount int, maxRetryCount int) *WorkerPool {
	pool := WorkerPool{
		jobChannel:    make(chan job, workerCount),
		maxRetryCount: maxRetryCount,
		PoolData:      poolData,
		maker:         maker,
	}
	for w := 0; w < workerCount; w++ {
		go pool.spawnWorker()
	}
	return &pool
}

func (pool *WorkerPool) spawnWorker() {
	worker := pool.maker.MakeWorker(pool)
	for job := range pool.jobChannel {
		err := worker.Work(job.JobData)
		if err != nil {
			if job.retryCount >= pool.maxRetryCount {
				job.errorChannel <- err
			} else {
				job.retryCount++
				pool.jobChannel <- job
			}
		} else {
			job.errorChannel <- nil
		}
	}
}

func (pool *WorkerPool) DoJob(JobData interface{}) error {
	errorChannel := make(chan error)
	newJob := job{
		JobData:      JobData,
		retryCount:   0,
		errorChannel: errorChannel,
	}
	pool.jobChannel <- newJob
	err := <-errorChannel
	return err
}

func (pool *WorkerPool) Close() {
	close(pool.jobChannel)
}
