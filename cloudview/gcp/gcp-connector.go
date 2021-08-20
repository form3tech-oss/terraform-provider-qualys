package gcp

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

const ApiPath = "/cloudview-api/rest/v1/gcp/connectors"

type ConnectorService struct {
	client *resty.Client
}

func NewService(baseURL, username, password string) *ConnectorService {
	client := resty.New().
		SetHostURL(baseURL).
		SetBasicAuth(username, password).
		SetRetryCount(3)

	return &ConnectorService{client: client}
}

func (s *ConnectorService) Get(id string) (*Connector, error) {
	path := fmt.Sprintf("%s/%s", ApiPath, id)

	connector := &Connector{}

	resp, err := s.client.R().
		SetResult(connector).
		Get(path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 400 {
		return nil, errors.New(resp.String())
	}
	if len(connector.ServiceError) > 0 {
		return nil, errors.New(resp.String())
	}
	if len(connector.ErrorCode) > 0 {
		return nil, errors.New(resp.String())
	}

	return connector, err
}

func (s *ConnectorService) Create(opt *ConnectorConfig) (*Connector, error) {
	connector := new(Connector)
	v, err := query.Values(opt)
	if err != nil {
		return nil, err
	}

	req := s.client.R().
		SetResult(&connector).
		SetFormDataFromValues(v)

	if opt.ConfigFile == "" {
		return nil, errors.New("missing config file")
	} else {
		req.SetFileReader("configFile", "configFile", strings.NewReader(opt.ConfigFile))
	}

	resp, err := req.Post(ApiPath)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 400 {
		return nil, errors.New(string(resp.Body()))
	}
	if len(connector.ServiceError) > 0 {
		return nil, errors.New(string(resp.Body()))
	}
	if len(connector.ErrorCode) > 0 {
		return nil, errors.New(string(resp.Body()))
	}

	return connector, err
}

func (s *ConnectorService) Update(id string, opt *ConnectorConfig) error {
	path := fmt.Sprintf("%s/%s", ApiPath, id)

	connector := new(Connector)
	v, err := query.Values(opt)
	if err != nil {
		return err
	}
	req := s.client.R().
		SetResult(&connector).
		SetFormDataFromValues(v)

	if opt.ConfigFile == "" {
		// needed for work around a bug that prevent update string-with-space values.
		req.SetFileReader("dummy", "dummy", strings.NewReader("dummy"))
	} else {
		req.SetFileReader("configFile", "configFile", strings.NewReader(opt.ConfigFile))
	}

	resp, err := req.Put(path)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 400 {
		return errors.New(string(resp.Body()))
	}
	if len(connector.ServiceError) > 0 {
		return errors.New(string(resp.Body()))
	}
	if len(connector.ErrorCode) > 0 {
		return errors.New(string(resp.Body()))
	}
	return nil
}

func (s *ConnectorService) Delete(connectorIds []string) error {
	resp, err := s.client.R().
		SetBody(ConnectorIds{connectorIds}).
		SetHeader("Content-Type", "application/json").
		Delete(ApiPath)

	if err != nil {
		return err
	}

	if resp.StatusCode() >= 400 {
		return fmt.Errorf("delete failed with status code %d", resp.StatusCode())
	}
	return nil
}
