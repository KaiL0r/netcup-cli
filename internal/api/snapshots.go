package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// ======================
// TYPES
// ======================

type Snapshot struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"createdAt"`
	Disk        string    `json:"disk"`
	SizeInMiB   int64     `json:"sizeInMiB"`
	Description *string   `json:"description,omitempty"`
}

type SnapshotCreate struct {
	Name           string  `json:"name"`
	Diskname       *string `json:"disk,omitempty"`
	Description    *string `json:"description,omitempty"`
	OnlineSnapshot *bool   `json:"onlineSnapshot,omitempty"`
}

// ======================
// SNAPSHOTS
// ======================

func (c *Client) ListServerSnapshots(serverID int) ([]Snapshot, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/snapshots", serverID))

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var out []Snapshot
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) GetServerSnapshot(serverID int, snapshotName string) (Snapshot, error) {
	path := fmt.Sprintf("/servers/%d/snapshots/%s", serverID, snapshotName)

	resp, err := c.Do("GET", path, nil)
	if err != nil {
		return Snapshot{}, err
	}

	var out Snapshot
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) CreateServerSnapshot(serverID int, create SnapshotCreate) (Task, error) {
	path := fmt.Sprintf("/servers/%d/snapshots", serverID)

	if create.OnlineSnapshot == nil {
		// default to true if not set
		create.OnlineSnapshot = new(bool)
		*create.OnlineSnapshot = true
	}

	if create.Diskname == nil && *create.OnlineSnapshot == false {
		return Task{}, fmt.Errorf("diskname must be set if onlineSnapshot is false")
	}

	resp, err := c.Do("POST", path, create)
	if err != nil {
		return Task{}, err
	}

	var out Task
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) DeleteServerSnapshot(serverID int, snapshotname string) (Task, error) {
	path := fmt.Sprintf("/servers/%d/snapshots/%s", serverID, snapshotname)

	resp, err := c.Do("DELETE", path, nil)
	if err != nil {
		return Task{}, err
	}

	var out Task
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) ExportServerSnapshot(serverID int, snapshotname string) (Task, error) {
	path := fmt.Sprintf("/servers/%d/snapshots/%s/export", serverID, snapshotname)

	resp, err := c.Do("POST", path, nil)
	if err != nil {
		return Task{}, err
	}

	var out Task
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) RevertServerSnapshot(serverID int, snapshotname string) (Task, error) {
	path := fmt.Sprintf("/servers/%d/snapshots/%s/revert", serverID, snapshotname)

	resp, err := c.Do("POST", path, nil)
	if err != nil {
		return Task{}, err
	}

	var out Task
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}
