package librus

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type AppError struct {
	Status int `json:"-"`
	Reason string
	Stage  string
}

type ResponseWithErrors struct {
	Message string
	Error   string `json:"error"`
}

type mRequestParams struct {
	path      string
	authToken string
	method    string
	body      string
}

func (e *AppError) Error() string {
	reason, _ := json.Marshal(e)
	return string(reason)
}

func request(params *mRequestParams) ([]byte, error) {
	stage := "fetch"
	req, err := http.NewRequest(params.method, params.path, strings.NewReader(params.body))
	if err != nil {
		return nil, &AppError{500, err.Error(), stage}
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(params.body)))
	req.Header.Add("Authorization", params.authToken)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "LibrusMobileApp")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, &AppError{500, err.Error(), stage}
	}

	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, &AppError{500, err.Error(), stage}
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		responseErrors := ResponseWithErrors{}
		json.Unmarshal(response, &responseErrors)
		return nil, &AppError{
			res.StatusCode,
			responseErrors.Message + responseErrors.Error,
			stage,
		}
	}

	return response, nil
}

func requestWorker(jobs <-chan *mRequestParams, results chan<- []byte, errs chan<- error) {
	for job := range jobs {
		response, err := request(job)
		if err != nil {
			errs <- err
			return
		}
		results <- response
	}
}
