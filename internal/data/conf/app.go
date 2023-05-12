package conf

import (
	"fmt"
)

// MySQL conf
type MySQL struct {
	Username   string
	Password   string
	Host       string
	Port       string
	Database   string
	Parameters []struct {
		Key   string
		Value string
	}
}

// DSN for mysql connection
func (c *MySQL) DSN() string {
	host := "127.0.0.1"
	if c.Host != "" {
		host = c.Host
	}
	port := "3306"
	if c.Port != "" {
		port = c.Port
	}

	params := ""
	if len(c.Parameters) > 0 {
		params = "?"
		for _, param := range c.Parameters {
			if len(params) > 1 {
				params += "&"
			}
			params += fmt.Sprintf("%s=%s", param.Key, param.Value)
		}
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", c.Username, c.Password, host, port, c.Database, params)
}

// String for hiding password
func (c MySQL) String() string {
	return fmt.Sprintf(`{Username: %s, Password: *****, Host: %s, Port: %s, Database: %s, Parameters: %v}`,
		c.Username, c.Host, c.Port, c.Database, c.Parameters)
}

// AppConf .
type AppConf struct {
	MySQL MySQL `mapstructure:"mysql"`
}
