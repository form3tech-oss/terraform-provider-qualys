package gcp

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

type ConnectorIds struct {
	ConnectorIds []string `url:"connectorIds,omitempty" json:"connectorIds,omitempty"`
}
