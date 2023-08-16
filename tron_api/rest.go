package tron_api

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func (cli *TronClient) get(path string, body []byte) (*http.Response, error) {
	preparedRequest, err := cli.prepareRequest("GET", path, body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare request")
	}

	resp, err := cli.client.Do(preparedRequest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to process request")
	}

	return resp, nil
}

func (cli *TronClient) post(path string, body []byte) (*http.Response, error) {
	preparedRequest, err := cli.prepareRequest("POST", path, body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare request")
	}

	resp, err := cli.client.Do(preparedRequest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to process request")
	}

	return resp, nil
}

func (cli *TronClient) prepareRequest(method, path string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, cli.rpcUrl+path, strings.NewReader(string(body)))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new request")
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("TRON-PRO-API-KEY", cli.apiKey)

	return req, nil
}

func (cli *TronClient) processResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal json", logan.F{
			"response_body": string(body),
		})
	}
	return nil
}
