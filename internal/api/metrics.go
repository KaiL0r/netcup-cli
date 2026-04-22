package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

func (c *Client) MetricsCpu(serverId int, hours int) (any, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/metrics/cpu", serverId))

	q := u.Query()
	if hours != 0 {
		q.Set("hours", fmt.Sprintf("%d", hours))
	}
	u.RawQuery = q.Encode()

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	var out any
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) MetricsDisk(serverId int, hours int) (any, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/metrics/disk", serverId))

	q := u.Query()
	if hours != 0 {
		q.Set("hours", fmt.Sprintf("%d", hours))
	}
	u.RawQuery = q.Encode()

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	var out any
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) MetricsNetwork(serverId int, hours int) (any, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/metrics/network", serverId))

	q := u.Query()
	if hours != 0 {
		q.Set("hours", fmt.Sprintf("%d", hours))
	}
	u.RawQuery = q.Encode()

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	var out any
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) MetricsNetworkPackets(serverId int, hours int) (any, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/metrics/network/packet", serverId))

	q := u.Query()
	if hours != 0 {
		q.Set("hours", fmt.Sprintf("%d", hours))
	}
	u.RawQuery = q.Encode()

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp)

	var out any
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}
