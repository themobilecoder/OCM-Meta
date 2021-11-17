package meta

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getFirstIntReturn(i int, err error) int {
	return i
}

func curlData(url string) (string, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return "", err
		}
		return string(bodyBytes), nil
	} else {
		return "", errors.New("FETCH_ERROR")
	}
}

func downloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func filter(arr []string, cond func(string) bool) []string {
	result := []string{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
