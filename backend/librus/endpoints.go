package librus

import "encoding/json"

func (L *Librus) LoadEndpoints(endpoints ...string) error {
	jobs := make(chan *mRequestParams, 5)
	results := make(chan []byte, 5)
	errs := make(chan error, 5)
	defer close(jobs)
	defer close(results)
	defer close(errs)

	for i := 0; i < 3 || i < len(endpoints); i += 1 {
		go requestWorker(jobs, results, errs)
	}

	for _, endpoint := range endpoints {
		jobs <- &mRequestParams{
			path:      API_ROOT + "2.0/" + endpoint,
			authToken: "Bearer " + L.AccessToken,
			method:    "GET",
		}
	}

	for i := 0; i < len(endpoints); i += 1 {
		select {
		case response := <-results:
			json.Unmarshal(response, L)
		case err := <-errs:
			return err
		}
	}

	return nil
}
