package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"go_chatible/model/user"
)

type fetchJob struct {
	usr              *user.User
	retryCount       int
	finishJobChannel chan bool
}

type UserFetcherPool struct {
	jobChannel      chan fetchJob
	urlFormatString string
	maxRetryCount   int
}

func NewUserFetcherPool() *UserFetcherPool {
	url := "https://graph.facebook.com/v5.0/%s?access_token=" + os.Getenv("PAGE_ACCESS_TOKEN")
	fetcher := UserFetcherPool{
		jobChannel:      make(chan fetchJob),
		urlFormatString: url,
		maxRetryCount:   4,
	}
	for w := 0; w < 2; w++ {
		go fetcher.SpawnWorker(w)
	}
	return &fetcher
}

func getUserData(client *http.Client, id string, urlFormatString string) (user.User, error) {
	url := fmt.Sprintf(urlFormatString, id)
	resp, err := client.Get(url)
	newUser := user.User{}
	if err != nil {
		return newUser, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		return newUser, err
	}
	if newUser == (user.User{}) {
		return newUser, fmt.Errorf("unable to fetch user data for user id %s", id)
	}
	newUser.SetFullName()
	return newUser, nil
}

func (fetcher *UserFetcherPool) SpawnWorker(workerID int) {
	client := http.Client{}
	for job := range fetcher.jobChannel {
		usr, err := getUserData(&client, job.usr.ID, fetcher.urlFormatString)
		if err != nil {
			log.Println("User data fetcher, worker", workerID, "error", err)
			if job.retryCount >= fetcher.maxRetryCount {
				log.Println("User data fetcher, worker", workerID, "maximum retrial exceeded")
				job.finishJobChannel <- false
			} else {
				job.retryCount++
				fetcher.jobChannel <- job
			}
		} else {
			*job.usr = usr
			job.finishJobChannel <- true
		}
	}
}

func (fetcher *UserFetcherPool) FetchUserData(usr *user.User) error {
	finishJobChannel := make(chan bool)
	newJob := fetchJob{
		usr:              usr,
		retryCount:       0,
		finishJobChannel: finishJobChannel,
	}
	fetcher.jobChannel <- newJob
	jobFinish := <-finishJobChannel
	if jobFinish {
		return nil
	}
	return fmt.Errorf("cannot fetch user data for user id %s", usr.ID)
}

func (fetcher *UserFetcherPool) Close() {
	close(fetcher.jobChannel)
}
