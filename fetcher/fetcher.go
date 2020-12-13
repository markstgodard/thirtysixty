package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Fetcher struct {
	Client *http.Client
}

func (f *Fetcher) Fetch(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http response: %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
