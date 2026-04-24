package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// ======================
// API Method
// ======================
func (c *Client) ListServers(limit int, offset int, name string, ip string, query string) ([]ServerListMinimal, error) {
	u, _ := url.Parse("/servers")
	q := u.Query()

	if limit != 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if offset != 0 {
		q.Set("offset", strconv.Itoa(offset))
	}
	if name != "" {
		q.Set("name", name)
	}
	if ip != "" {
		q.Set("ip", ip)
	}
	if query != "" {
		q.Set("q", query)
	}

	u.RawQuery = q.Encode()

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var out []ServerListMinimal
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) UpdateServerState(serverID int, state SetServerState, stateOption SetServerStateOption) (Task, error) {
	switch state {
	case SetStateOn:
		if stateOption != "" &&
			stateOption != SetStateOptionPowercycle &&
			stateOption != SetStateOptionReset {
			return Task{}, fmt.Errorf("invalid state-option %q for state %s", stateOption, state)
		}

	case SetStateOff:
		if stateOption != "" && stateOption != SetStateOptionPoweroff {
			return Task{}, fmt.Errorf("invalid state-option %q for state %s", stateOption, state)
		}

	case SetStateSuspended:
		if stateOption != "" {
			return Task{}, fmt.Errorf("invalid state-option %q for state %s", stateOption, state)
		}

	default:
		return Task{}, fmt.Errorf("invalid state %q", state)
	}

	return c.patchServer(serverID, string(stateOption), map[string]any{"state": string(state)})
}

func (c *Client) UpdateServerAutostart(serverID int, autostart bool) (Task, error) {
	return c.patchServer(serverID, "", map[string]any{"autostart": autostart})
}

func (c *Client) UpdateServerBootorder(serverID int, bootorder []BootOrder) (Task, error) {
	for _, v := range bootorder {
		switch v {
		case BootOrderCDROM, BootOrderNetwork, BootOrderHDD:
			continue
		default:
			return Task{}, fmt.Errorf("invalid bootorder %q", v)
		}
	}

	return c.patchServer(serverID, "", map[string]any{"bootorder": bootorder})
}

func (c *Client) UpdateServerOsoptimization(serverID int, osOptimization OsOptimization) (Task, error) {
	switch osOptimization {
	case
		OSLinux,
		OSWindows,
		OSBSD,
		OSLinuxLegacy,
		OSUnknown:
	default:
		return Task{}, fmt.Errorf("invalid OSOptimization %q", osOptimization)
	}

	return c.patchServer(serverID, "", map[string]any{"os_optimization": osOptimization})
}

func (c *Client) UpdateServerCpuTopology(serverID int, sockets int, cores int) (Task, error) {
	return c.patchServer(serverID, "", map[string]any{
		"cpuTopology": map[string]int{
			"socketCount":         sockets,
			"coresPerSocketCount": cores,
		},
	})
}

func (c *Client) UpdateServerUefi(serverID int, uefi bool) (Task, error) {
	return c.patchServer(serverID, "", map[string]any{"uefi": uefi})
}

func (c *Client) UpdateServerHostname(serverID int, hostname string) (Task, error) {
	return c.patchServer(serverID, "", map[string]any{"hostname": hostname})
}

func (c *Client) UpdateServerNickname(serverID int, nickname string) (Task, error) {
	return c.patchServer(serverID, "", map[string]any{"nickname": nickname})
}

func (c *Client) UpdateServerKeyboardLayout(serverID int, layout string) (Task, error) {
	return c.patchServer(serverID, "", map[string]any{"keyboardLayout": layout})
}

func (c *Client) UpdateServerRootPassword(serverID int, password string) (Task, error) {
	return c.patchServer(serverID, "", map[string]any{"rootPassword": password})
}

func (c *Client) patchServer(serverID int, stateOption string, body map[string]any) (Task, error) {
	u, _ := url.Parse("/servers/" + strconv.Itoa(serverID))
	q := u.Query()
	if stateOption != "" {
		q.Set("stateOption", stateOption)
	}
	u.RawQuery = q.Encode()

	resp, err := c.Do("PATCH", u.String(), body)
	if err != nil {
		return Task{}, err
	}

	if len(resp) != 0 {
		var out Task
		return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
	}

	return Task{}, nil
}

func (c *Client) GetServer(serverID int) (Server, error) {
	u, _ := url.Parse("/servers/" + strconv.Itoa(serverID))

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return Server{}, err
	}

	var out Server
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) GetServerGpuDriver(serverID int) (S3DownloadInfos, error) {
	u, _ := url.Parse("/servers/" + strconv.Itoa(serverID) + "/gpu-driver")

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return S3DownloadInfos{}, err
	}

	var out S3DownloadInfos
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) GetServerGuestAgent(serverID int) (GuestAgentData, error) {
	path := fmt.Sprintf("/servers/%d/guest-agent", serverID)

	resp, err := c.Do("GET", path, nil)
	if err != nil {
		return GuestAgentData{}, err
	}

	var out GuestAgentData
	return out, json.NewDecoder(strings.NewReader(resp)).Decode(&out)
}

func (c *Client) ListServerLogs(serverID int, limit, offset int) ([]Log, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/logs", serverID))
	q := u.Query()

	if limit != 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if offset != 0 {
		q.Set("offset", strconv.Itoa(offset))
	}
	u.RawQuery = q.Encode()

	resp, err := c.Do("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	var logs []Log
	return logs, json.Unmarshal([]byte(resp), &logs)
}

func (c *Client) GetServerRescueSystem(serverID int) (RescueSystemStatus, error) {
	path := fmt.Sprintf("/servers/%d/rescuesystem", serverID)

	resp, err := c.Do("GET", path, nil)
	if err != nil {
		return RescueSystemStatus{}, err
	}

	var out RescueSystemStatus
	return out, json.Unmarshal([]byte(resp), &out)
}

func (c *Client) ActivateServerRescueSystem(serverID int) (Task, error) {
	path := fmt.Sprintf("/servers/%d/rescuesystem", serverID)

	resp, err := c.Do("POST", path, nil)
	if err != nil {
		return Task{}, err
	}

	var out Task
	return out, json.Unmarshal([]byte(resp), &out)
}

func (c *Client) DeactivateServerRescueSystem(serverID int) (Task, error) {
	path := fmt.Sprintf("/servers/%d/rescuesystem", serverID)

	resp, err := c.Do("DELETE", path, nil)
	if err != nil {
		return Task{}, err
	}

	var out Task
	return out, json.Unmarshal([]byte(resp), &out)
}

func (c *Client) OptimizeServerStorage(serverID int, disks []string, startAfterOptimization bool) (Task, error) {
	u, _ := url.Parse(fmt.Sprintf("/servers/%d/storageoptimization", serverID))
	q := u.Query()

	if len(disks) == 0 {
		return Task{}, fmt.Errorf("at least one disk must be specified for optimization")
	}

	for _, d := range disks {
		q.Add("disks", d)
	}

	q.Set("startAfterOptimization", strconv.FormatBool(startAfterOptimization))
	u.RawQuery = q.Encode()

	resp, err := c.Do("POST", u.String(), nil)
	if err != nil {
		return Task{}, err
	}

	var out Task
	return out, json.Unmarshal([]byte(resp), &out)
}
