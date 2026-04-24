package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// ======================
// Types + Params
// ======================

type ListTasksParams struct {
	Limit    int
	Offset   int
	Q        string
	ServerID int
	State    string
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Task struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	State      string `json:"state"`
	StartedAt  string `json:"startedAt"`
	FinishedAt string `json:"finishedAt"`
	Message    string `json:"message"`
	OnRollback bool   `json:"onRollback"`

	ExecutingUser interface{}    `json:"executingUser"`
	TaskProgress  interface{}    `json:"taskProgress"`
	Steps         interface{}    `json:"steps"`
	Result        interface{}    `json:"result"`
	ResponseError *ResponseError `json:"responseError,omitempty"`
}

// ======================
// API Method
// ======================

func (c *Client) ListTasks(p ListTasksParams) ([]Task, error) {
	u, _ := url.Parse("/tasks")
	q := u.Query()

	if p.Limit != 0 {
		q.Set("limit", strconv.Itoa(p.Limit))
	}
	if p.Offset != 0 {
		q.Set("offset", strconv.Itoa(p.Offset))
	}
	if p.Q != "" {
		q.Set("q", p.Q)
	}
	if p.ServerID != 0 {
		q.Set("serverId", strconv.Itoa(p.ServerID))
	}
	if p.State != "" {
		q.Set("state", p.State)
	}

	u.RawQuery = q.Encode()

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var out []Task
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) GetTask(uuid string) (Task, error) {
	u, _ := url.Parse("/tasks/" + uuid)

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return Task{}, err
	}

	var task Task
	if err := json.NewDecoder(strings.NewReader(resp)).Decode(&task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (c *Client) CancelTask(uuid string) error {
	u, _ := url.Parse(fmt.Sprintf("/tasks/%s:cancel", uuid))

	if _, err := c.Do("PUT", u.String(), nil); err != nil {
		return err
	}

	return nil
}
