package regclient

import (
	"bytes"
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
	serviceInfo       []byte
	registryURL       string
	mutex             sync.Mutex
	isActive          bool
	heartbeatInterval time.Duration
}

func NewClient(info ServiceInfo, regURL string, interval time.Duration) *Client {
	sInfo, err := info.MarshalJSON()
	if err != nil {
		return nil
	}

	return &Client{
		serviceInfo:       sInfo,
		registryURL:       regURL,
		heartbeatInterval: interval,
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
		for {
			c.mutex.Lock()
			if c.isActive {
				c.performRequest(http.MethodPut)
			}
			c.mutex.Unlock()
		}
		time.Sleep(c.heartbeatInterval)
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

	req, err := http.NewRequest(method, c.registryURL, bytes.NewReader(c.serviceInfo))

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
