package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go_chatible/api/worker_pool"
	"go_chatible/model/message"
)

type MessageSenderPool struct {
	workerPool *worker_pool.WorkerPool
}

type MsgWorkerMaker struct{}

type msgWorker struct {
	client *http.Client
	pool   *worker_pool.WorkerPool
}

func (wrk msgWorker) Work(JobData interface{}) error {
	msg, ok := JobData.(message.Message)
	if !ok {
		panic("wtf")
	}
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	bodyReader := bytes.NewReader(body)
	url, ok := wrk.pool.PoolData.(string)
	if !ok {
		panic("wtf")
	}
	resp, err := wrk.client.Post(url, "application/json", bodyReader)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("request error when sending message %s", string(body))
	}
	return nil
}

func (maker MsgWorkerMaker) MakeWorker(pool *worker_pool.WorkerPool) worker_pool.Worker {
	worker := msgWorker{
		client: &http.Client{},
		pool:   pool,
	}
	return worker
}

func NewMessageSenderPool() *MessageSenderPool {
	url := "https://graph.facebook.com/v5.0/me/messages?access_token=" + os.Getenv("PAGE_ACCESS_TOKEN")
	workerMaker := MsgWorkerMaker{}
	res := MessageSenderPool{
		workerPool: worker_pool.NewPool(url, workerMaker, 4, 4),
	}
	return &res
}

func (pool *MessageSenderPool) SendMessage(msg message.Message) error {
	err := pool.workerPool.DoJob(msg)
	return err
}

func (pool *MessageSenderPool) Close() {
	pool.workerPool.Close()
}
