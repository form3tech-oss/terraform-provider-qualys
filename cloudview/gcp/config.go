package gcp

type ConnectorConfig struct {
	Name        string `url:"name,omitempty" json:"name,omitempty"`
	Description string `url:"description,omitempty" json:"description,omitempty"`
	ProjectId   string `url:"projectId,omitempty" json:"projectId,omitempty"`
	ConfigFile  string `url:"-" json:"-"`
}

func NewConnectorConfig() *ConnectorConfig {
	return &ConnectorConfig{}
}

func (c *ConnectorConfig) WithName(name string) *ConnectorConfig {
	c.Name = name
	return c
}

func (c *ConnectorConfig) WithDescription(desc string) *ConnectorConfig {
	c.Description = desc
	return c
}

func (c *ConnectorConfig) WithConfigFile(file string) *ConnectorConfig {
	c.ConfigFile = file
	return c
}

func (c *ConnectorConfig) WithProjectId(projectId string) *ConnectorConfig {
	c.ProjectId = projectId
	return c
}
