package artifactory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

type clientConfig struct {
	user     string
	pass     string
	url      string
	clientMu sync.Mutex // clientMu protects the client during multi-threaded calls]
	client   *http.Client
}

// Client is used to call Artifactory REST APIs
type Client interface {
	Ping() error
	GetRepository(key string, v interface{}) error
	CreateRepository(key string, v interface{}) error
	UpdateRepository(key string, v interface{}) error
	DeleteRepository(key string) error
	GetUser(name string) (*User, error)
	CreateUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(name string) error
	GetGroup(name string) (*Group, error)
	CreateGroup(g *Group) error
	UpdateGroup(g *Group) error
	DeleteGroup(name string) error
	ExpireUserPassword(name string) error
}

var _ Client = clientConfig{}

// NewClient constructs a new artifactory client
func NewClient(username, pass, url string, client *http.Client) *clientConfig {
	return &clientConfig{
		user:   username,
		pass:   pass,
		url:    strings.TrimRight(url, "/"),
		client: client,
	}
}

// Lock the client until release
func (c *clientConfig) Lock() {
	c.clientMu.Lock()
}

// Unlock the client after a lock action
func (c *clientConfig) Unlock() {
	c.clientMu.Unlock()
}

// Ping calls the system to verify connectivity
func (c clientConfig) Ping() error {
	resp, err := c.execute("GET", "system/ping", nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Ping failed. Status: %s", resp.Status)
	}

	return resp.Body.Close()
}

func (c clientConfig) execute(method string, endpoint string, payload interface{}) (resp *http.Response, err error) {
	var req *http.Request
	c.Lock()
	defer c.Unlock()

	url := fmt.Sprintf("%s/api/%s", c.url, endpoint)

	var jsonpayload *bytes.Buffer
	if payload == nil {
		jsonpayload = &bytes.Buffer{}
	} else {
		var jsonbuffer []byte
		jsonpayload = bytes.NewBuffer(jsonbuffer)
		enc := json.NewEncoder(jsonpayload)
		err = enc.Encode(payload)
		if err != nil {
			log.Printf("[ERROR] Error Encoding Payload: %s", err)
			return nil, err
		}
	}

	req, err = http.NewRequest(method, url, jsonpayload)
	if err != nil {
		log.Printf("[ERROR] Error creating new request: %s", err)
		return nil, err
	}
	req.SetBasicAuth(c.user, c.pass)
	req.Header.Add("content-type", "application/json")

	resp, err = c.client.Do(req)
	if err == io.EOF {
		err = nil // ignore EOF errors caused by empty response body
	}

	return resp, err
}

func (c clientConfig) validateResponse(expected int, actual int, action string) (err error) {
	if expected != actual {
		err = fmt.Errorf("Expected %d for '%s', got '%d'", expected, action, actual)
	}
	return
}
