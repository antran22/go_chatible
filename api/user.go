package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"go_chatible/api/worker_pool"
	"go_chatible/model/user"
)

type UserFetcherPool struct {
	workerPool *worker_pool.WorkerPool
}

type UsrFetcherWorkerMaker struct{}

type usrWorker struct {
	client *http.Client
	pool   *worker_pool.WorkerPool
}

func (wrk usrWorker) Work(JobData interface{}) error {
	usr, ok1 := JobData.(*user.User)
	urlFormatString, ok2 := wrk.pool.PoolData.(string)
	if !ok1 || !ok2 {
		panic("wtf")
	}
	url := fmt.Sprintf(urlFormatString, usr.ID)
	resp, err := wrk.client.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("unable to fetch user data for user id %s", usr.ID)
	}
	body, err := ioutil.ReadAll(resp.Body)
	newUser := user.User{}
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		return err
	}
	if newUser == (user.User{}) {
		return fmt.Errorf("unable to read user data for user id %s", usr.ID)
	}
	newUser.SetFullName()
	*usr = newUser
	return nil
}

func (maker UsrFetcherWorkerMaker) MakeWorker(pool *worker_pool.WorkerPool) worker_pool.Worker {
	worker := usrWorker{
		client: &http.Client{},
		pool:   pool,
	}
	return worker
}

func NewUserFetcherPool() *UserFetcherPool {
	urlFormatString := "https://graph.facebook.com/v5.0/%s?access_token=" + os.Getenv("PAGE_ACCESS_TOKEN")
	workerMaker := UsrFetcherWorkerMaker{}
	res := UserFetcherPool{
		workerPool: worker_pool.NewPool(urlFormatString, workerMaker, 4, 4),
	}
	return &res
}

func (pool *UserFetcherPool) FetchUser(usr *user.User) error {
	err := pool.workerPool.DoJob(usr)
	return err
}

func (pool *UserFetcherPool) Close() {
	pool.workerPool.Close()
}
