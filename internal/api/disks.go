package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type Disk struct {
	Name            string        `json:"name"`
	AllocationInMiB int           `json:"allocationInMiB"`
	CapacityInMiB   int           `json:"capacityInMiB"`
	StorageDriver   StorageDriver `json:"storageDriver"`
}

func (c *Client) ListDisks(serverId int) ([]Disk, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/disks", serverId))

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var out []Disk
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) GetDisk(serverId int, diskName string) (Disk, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/disks/%s", serverId, diskName))

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return Disk{}, err
	}

	var disk Disk
	if err := json.NewDecoder(strings.NewReader(resp)).Decode(&disk); err != nil {
		return Disk{}, err
	}

	return disk, nil
}

type StorageDriver string

const (
	StorageDriverVirtIo     StorageDriver = "VIRTIO"
	StorageDriverVirtIoScsi StorageDriver = "VIRTIO_SCSI"
	StorageDriverIde        StorageDriver = "IDE"
	StorageDriverSata       StorageDriver = "SATA"
)

func (c *Client) ListDiskSupportedDrivers(serverId int) ([]StorageDriver, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/disks/supported-drivers", serverId))

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var out []StorageDriver
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) UpdateDiskDriver(serverId int, storageDriver StorageDriver) (Task, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/disks", serverId))

	storageDriver = StorageDriver(strings.ToUpper(string(storageDriver)))
	switch storageDriver {
	case
		StorageDriverVirtIo,
		StorageDriverVirtIoScsi,
		StorageDriverIde,
		StorageDriverSata:
	default:
		return Task{}, fmt.Errorf("invalid StorageDriver %q", storageDriver)
	}

	resp, err := c.Do("PATCH", u.String(), map[string]any{"driver": storageDriver})
	if err != nil {
		return Task{}, err
	}

	var task Task
	if err := json.NewDecoder(strings.NewReader(resp)).Decode(&task); err != nil {
		return Task{}, err
	}

	return task, nil
}

func (c *Client) FormatDisk(serverId int, diskName string) (Task, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/disks/%s:format", serverId, diskName))

	resp, err := c.Do("POST", u.String(), nil)
	if err != nil {
		return Task{}, err
	}

	var task Task
	if err := json.NewDecoder(strings.NewReader(resp)).Decode(&task); err != nil {
		return Task{}, err
	}

	return task, nil
}
