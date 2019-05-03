package cronitor

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"github.com/json-iterator/go/extra"
)

const apiEndpoint = "https://cronitor.io/v3"

func init() {
	extra.RegisterFuzzyDecoders()
}

type Client struct {
	secretKey string
}

func (c Client) create(m Monitor) (code *string, err error) {
	b, err := jsoniter.Marshal(m)

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(b)

	req, _ := http.NewRequest("POST", apiEndpoint+"/monitors", buf)

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.secretKey, "")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("failed to create monitor, status code: %d, body: %s", resp.StatusCode, body)
	}

	createdMonitor := Monitor{}

	if err := jsoniter.Unmarshal(body, &createdMonitor); err != nil {
		return nil, err
	}

	return &createdMonitor.Code, nil
}

func (c Client) update(code string, m Monitor) error {
	b, err := jsoniter.Marshal(m)

	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(b)

	req, _ := http.NewRequest("PUT", apiEndpoint+"/monitors/"+code, buf)

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.secretKey, "")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to update monitor, status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}

func (c Client) read(code string) (*Monitor, error) {
	req, _ := http.NewRequest("GET", apiEndpoint+"/monitors/"+code, nil)

	req.SetBasicAuth(c.secretKey, "")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to read monitor, status code: %d, body: %s", resp.StatusCode, body)
	}

	monitor := Monitor{}

	if err := jsoniter.Unmarshal(body, &monitor); err != nil {
		return nil, err
	}

	return &monitor, nil
}

func (c Client) delete(code string) error {
	req, _ := http.NewRequest("DELETE", apiEndpoint+"/monitors/"+code, nil)

	req.SetBasicAuth(c.secretKey, "")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete monitor, status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}
