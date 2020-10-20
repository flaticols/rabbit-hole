package rabbithole

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type UserDefinition struct {
	Name             string `json:"name" yaml:"name"`
	PasswordHash     string `json:"password_hash" yaml:"password_hash"`
	HashingAlgorithm string `json:"hashing_algorithm" yaml:"hashing_algorithm"`
	Tags             string `json:"tags" yaml:"tags"`
}

type VHostDefinitions struct {
	Name string `json:"name" yaml:"name"`
}

type PermissionsDefinitions struct {
	User      string `json:"user" yaml:"user"`
	Vhost     string `json:"vhost" yaml:"vhost"`
	Configure string `json:"configure" yaml:"configure"`
	Write     string `json:"write" yaml:"write"`
	Read      string `json:"read" yaml:"read"`
}

type TopicPermissionsDefinitions struct {
	User     string `json:"user" yaml:"user"`
	Vhost    string `json:"vhost" yaml:"vhost"`
	Exchange string `json:"exchange"`
	Write    string `json:"write" yaml:"write"`
	Read     string `json:"read" yaml:"read"`
}

type GlobalParametersDefinitions struct {
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
}

type QueueDefinitions struct {
	Name       string                 `json:"name" yaml:"name"`
	Vhost      string                 `json:"vhost" yaml:"vhost"`
	Durable    bool                   `json:"durable" yaml:"durable"`
	AutoDelete bool                   `json:"auto_delete" yaml:"auto_delete"`
	Arguments  map[string]interface{} `json:"arguments" yaml:"arguments"`
}

type ExchangeDefinitions struct {
	Name       string                 `json:"name" yaml:"name"`
	Vhost      string                 `json:"vhost" yaml:"vhost"`
	Type       string                 `json:"type" yaml:"type"`
	Durable    bool                   `json:"durable" yaml:"durable"`
	AutoDelete bool                   `json:"auto_delete" yaml:"auto_delete"`
	Internal   bool                   `json:"internal" yaml:"internal"`
	Arguments  map[string]interface{} `json:"arguments" yaml:"arguments"`
}

type BindingDefinitions struct {
	Source          string                 `json:"source" yaml:"source"`
	Vhost           string                 `json:"vhost" yaml:"vhost"`
	Destination     string                 `json:"destination" yaml:"destination"`
	DestinationType string                 `json:"destination_type" yaml:"destination_type"`
	RoutingKey      string                 `json:"routing_key" yaml:"routing_key"`
	Arguments       map[string]interface{} `json:"arguments" yaml:"arguments"`
}

type Definitions struct {
	RabbitVersion string `json:"rabbit_version"`

	Users            UserDefinition                `json:"users" yaml:"users"`
	VirtualHosts     []VHostDefinitions            `json:"virtual_hosts" yaml:"virtual_hosts"`
	Permissions      []PermissionsDefinitions      `json:"permissions" yaml:"permissions"`
	TopicPermissions []TopicPermissionsDefinitions `json:"topic_permissions" yaml:"topic_permissions"`
	Parameters       []map[string]interface{}      `json:"parameters" yaml:"parameters"`
	GlobalParameters []GlobalParametersDefinitions `json:"global_parameters" yaml:"global_parameters"`
	Policies         []map[string]interface{}      `json:"policies" yaml:"policies"`
	Queues           []QueueDefinitions            `json:"queues" yaml:"queues"`
	Exchanges        []ExchangeDefinitions         `json:"exchanges" yaml:"exchanges"`
	Bindings         []BindingDefinitions          `json:"bindings" yaml:"bindings"`
}

// BackupDefinitions back all definitions.
func (c *Client) BackupDefinitions() (def *Definitions, err error) {
	req, err := newGETRequest(c, "/definitions")
	if err != nil {
		return &Definitions{}, err
	}

	if err = executeAndParseRequest(c, req, &def); err != nil {
		return &Definitions{}, err
	}

	return def, nil
}

// BackupVhostDefinitions back vhost definitions.
func (c *Client) BackupVhostDefinitions(vhost string) (def *Definitions, err error) {
	req, err := newGETRequest(c, url.PathEscape(fmt.Sprintf("/definitions/%s", vhost)))
	if err != nil {
		return &Definitions{}, err
	}

	if err = executeAndParseRequest(c, req, &def); err != nil {
		return &Definitions{}, err
	}

	return def, nil
}

// RestoreDefinitions restore all definitions.
func (c *Client) RestoreDefinitions(def *Definitions) (res *http.Response, err error) {
	body, err := json.Marshal(def)
	if err != nil {
		return nil, err
	}

	req, err := newRequestWithBody(c, "POST", "/definitions", body)
	if err != nil {
		return nil, err
	}

	res, err = executeRequest(c, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// RestoreVhostDefinitions restore vhost definitions.
func (c *Client) RestoreVhostDefinitions(def *Definitions, vhost string) (res *http.Response, err error) {
	body, err := json.Marshal(def)
	if err != nil {
		return nil, err
	}

	req, err := newRequestWithBody(c, "POST", url.PathEscape(fmt.Sprintf("/definitions/%s", vhost)), body)
	if err != nil {
		return nil, err
	}

	res, err = executeRequest(c, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
