package regclient

import (
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	DefaultHeartbeatInterval = 10 * time.Second
)

type Client struct {
	name              string
	registryAddress   string
	mutex             sync.Mutex
	isActive          bool
	heartbeatInterval time.Duration
}

func NewClient(name string, registryAddr string, interval time.Duration) *Client {
	return &Client{
		name: name,
		registryAddress:   registryAddr,
		heartbeatInterval: interval * time.Second,
	}
}

func (c *Client) Register() {
	c.mutex.Lock()
	if !c.isActive {
		c.isActive = true
		c.performRequest(http.MethodPost)
	}
	c.mutex.Unlock()
}

func (c Client) Start() {
	go func() {
		time.Sleep(c.heartbeatInterval)

		c.mutex.Lock()
		if c.isActive {
			c.performRequest(http.MethodPut)
		}
		c.mutex.Unlock()
	}()
}

func (c *Client) Unregister() {
	c.mutex.Lock()
	if c.isActive {
		c.isActive = false
		c.performRequest(http.MethodDelete)
	}
	c.mutex.Unlock()
}

func (c *Client) performRequest(method string) {
	client := &http.Client{}

	req, err := http.NewRequest(method, c.registryAddress+ "/" + c.name, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("registry service not found")
	} else if resp.StatusCode != http.StatusOK {
		rb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		log.Printf("registry service response: %v, %v", resp.StatusCode, string(rb))
	}
}
