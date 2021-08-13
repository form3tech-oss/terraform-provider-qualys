package gcp

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
)

const apiPath = "/cloudview-api/rest/v1/gcp/connectors"

type ConnectorService struct {
	client *resty.Client
}

type Connector struct {
	ErrorResponse

	ConnectorID  string  `json:"connectorId"`
	Description  string  `json:"description"`
	Groups       []Group `json:"groups"`
	LastSyncedOn string  `json:"lastSyncedOn"`
	Name         string  `json:"name"`
	Project      string  `json:"projectId"`
	Provider     string  `json:"provider"`
	State        string  `json:"state"`
	TotalAssets  int     `json:"totalAssets"`
}

type Group struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type ErrorResponse struct {
	Timestamp    string `json:"timestamp"`
	Status       int    `json:"status"`
	ServiceError string `json:"error"`
	ErrorCode    string `json:"errorCode"`
	Message      string `json:"message"`
	Path         string `json:"path"`
}

type Pageable struct {
	Offset     int  `json:"offset"`
	PageNumber int  `json:"pageNumber"`
	PageSize   int  `json:"pageSize"`
	Paged      bool `json:"paged"`
	Sort       Sort `json:"sort"`
	UnPaged    bool `json:"unpaged"`
}

type Sort struct {
	Sorted   bool `json:"sorted"`
	Unsorted bool `json:"unsorted"`
}

type ConnectorList struct {
	ErrorResponse

	List     []Connector `json:"content"`
	Pageable *Pageable   `json:"pagable"`
	IsFirst  bool        `json:"first"`
	IsLast   bool        `json:"last"`
	Number   int         `json:"number"`
	Total    int         `json:"numberOfElements"`
}

func NewService(baseURL, username, password string) *ConnectorService {
	client := resty.New().
		SetHostURL(baseURL).
		SetBasicAuth(username, password).
		SetRetryCount(3)

	return &ConnectorService{client: client}
}

func (s *ConnectorService) Get(id string) (*Connector, error) {
	path := fmt.Sprintf("%s/%s", apiPath, id)

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

	resp, err := req.Post(apiPath)
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
	path := fmt.Sprintf("%s/%s", apiPath, id)

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
		SetBody(struct {
			ConnectorIds []string `url:"connectorIds,omitempty" json:"connectorIds,omitempty"`
		}{connectorIds}).
		SetHeader("Content-Type", "application/json").
		Delete(apiPath)

	if err != nil {
		return err
	}

	if resp.StatusCode() >= 400 {
		return fmt.Errorf("delete failed with status code %d", resp.StatusCode())
	}
	return nil
}
