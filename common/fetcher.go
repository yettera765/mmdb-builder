package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const lineSep = '\n'

func Fetch(url string) []string {
	buf, err := fetch(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	return Split(string(buf), lineSep)
}

func fetch(url string) ([]byte, error) {
	log.Println("fetch", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("url content %s is empty", url)
	}
	return data, nil
}
