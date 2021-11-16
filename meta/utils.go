package meta

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetFirstIntReturn(i int, err error) int {
	return i
}

func CurlData(url string) (string, error) {
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

func DownloadFile(filepath string, url string) error {

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

func Filter(arr []int, cond func(int) bool) []int {
	result := []int{}
	for i := range arr {
		if cond(arr[i]) {
			result = append(result, arr[i])
		}
	}
	return result
}
