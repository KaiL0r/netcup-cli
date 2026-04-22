package api

import "time"

// ======================
// ROOT
// ======================

type Server struct {
	ID                       int                  `json:"id"`
	Name                     string               `json:"name"`
	Hostname                 *string              `json:"hostname"`
	Nickname                 *string              `json:"nickname"`
	Disabled                 bool                 `json:"disabled"`
	Template                 *Template            `json:"template"`
	ServerLiveInfo           *ServerLiveInfo      `json:"serverLiveInfo"`
	IPv4Addresses            []IPv4AddressMinimal `json:"ipv4Addresses"`
	IPv6Addresses            []IPv6AddressMinimal `json:"ipv6Addresses"`
	Site                     Site                 `json:"site"`
	SnapshotCount            int                  `json:"snapshotCount"`
	MaxCPUCount              int                  `json:"maxCpuCount"`
	DisksAvailableSpaceInMiB int64                `json:"disksAvailableSpaceInMiB"`
	RescueSystemActive       bool                 `json:"rescueSystemActive"`
	SnapshotAllowed          bool                 `json:"snapshotAllowed"`
	Architecture             Architecture         `json:"architecture"`
	GPUDriverAvailable       bool                 `json:"gpuDriverAvailable"`
}

type Template struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ======================
// LIVE INFO
// ======================

type ServerLiveInfo struct {
	State                       ServerState         `json:"state"`
	Autostart                   bool                `json:"autostart"`
	UEFI                        bool                `json:"uefi"`
	Interfaces                  []ServerInterface   `json:"interfaces"`
	Disks                       []ServerDisk        `json:"disks"`
	Bootorder                   []BootOrder         `json:"bootorder"`
	RequiredStorageOptimization StorageOptimization `json:"requiredStorageOptimization"`
	Template                    *string             `json:"template,omitempty"`
	UptimeInSeconds             int                 `json:"uptimeInSeconds"`
	CurrentServerMemoryInMiB    int64               `json:"currentServerMemoryInMiB"`
	MaxServerMemoryInMiB        int64               `json:"maxServerMemoryInMiB"`
	CPUCount                    int                 `json:"cpuCount"`
	CPUMaxCount                 int                 `json:"cpuMaxCount"`
	Sockets                     int                 `json:"sockets"`
	CoresPerSocket              int                 `json:"coresPerSocket"`
	LatestQemu                  bool                `json:"latestQemu"`
	ConfigChanged               bool                `json:"configChanged"`
	OSOptimization              OsOptimization      `json:"osOptimization"`
	NestedGuest                 bool                `json:"nestedGuest"`
	MachineType                 string              `json:"machineType"`
	KeyboardLayout              string              `json:"keyboardLayout"`
	CloudinitAttached           bool                `json:"cloudinitAttached"`
}

// ======================
// INTERFACES
// ======================

type ServerInterface struct {
	MAC                    string   `json:"mac"`
	Driver                 string   `json:"driver"`
	MTU                    int      `json:"mtu"`
	SpeedInMBits           int      `json:"speedInMBits"`
	RXMonthlyInMiB         int      `json:"rxMonthlyInMiB"`
	TXMonthlyInMiB         int      `json:"txMonthlyInMiB"`
	IPv4Addresses          []string `json:"ipv4Addresses"`
	IPv6LinkLocalAddresses []string `json:"ipv6LinkLocalAddresses"`
	IPv6NetworkPrefixes    []string `json:"ipv6NetworkPrefixes"`
	TrafficThrottled       bool     `json:"trafficThrottled"`
	VLANInterface          bool     `json:"vlanInterface"`
	VLANID                 int      `json:"vlanId"`
}

// ======================
// DISKS
// ======================

type ServerDisk struct {
	Dev             string `json:"dev"`
	Driver          string `json:"driver"`
	CapacityInMiB   int64  `json:"capacityInMiB"`
	AllocationInMiB int64  `json:"allocationInMiB"`
}

// ======================
// IP ADDRESSES
// ======================

type IPv4AddressMinimal struct {
	ID        int     `json:"id"`
	IP        string  `json:"ip"`
	Netmask   string  `json:"netmask"`
	Gateway   *string `json:"gateway,omitempty"`
	Broadcast *string `json:"broadcast,omitempty"`
}

type IPv6AddressMinimal struct {
	ID                  int     `json:"id"`
	NetworkPrefix       string  `json:"networkPrefix"`
	NetworkPrefixLength int     `json:"networkPrefixLength"`
	Gateway             *string `json:"gateway,omitempty"`
}

// ======================
// SITE
// ======================

type Site struct {
	ID   int    `json:"id"`
	City string `json:"city"`
}

// ======================
// ENUMS
// ======================

type ServerState string

const (
	ServerStateNoState      ServerState = "NOSTATE"
	ServerStateRunning      ServerState = "RUNNING"
	ServerStateBlocked      ServerState = "BLOCKED"
	ServerStatePaused       ServerState = "PAUSED"
	ServerStateShutdown     ServerState = "SHUTDOWN"
	ServerStateShutoff      ServerState = "SHUTOFF"
	ServerStateCrashed      ServerState = "CRASHED"
	ServerStatePMSuspended  ServerState = "PMSUSPENDED"
	ServerStateDiskSnapshot ServerState = "DISK_SNAPSHOT"
)

type BootOrder string

const (
	BootOrderHDD     BootOrder = "HDD"
	BootOrderCDROM   BootOrder = "CDROM"
	BootOrderNetwork BootOrder = "NETWORK"
)

type StorageOptimization string

const (
	StorageInconsistent StorageOptimization = "INCONSISTENT"
	StorageCompat       StorageOptimization = "COMPAT"
	StorageSlow         StorageOptimization = "SLOW"
	StorageFast         StorageOptimization = "FAST"
	StorageNone         StorageOptimization = "NO"
)

type OsOptimization string

const (
	OSLinux       OsOptimization = "LINUX"
	OSWindows     OsOptimization = "WINDOWS"
	OSBSD         OsOptimization = "BSD"
	OSLinuxLegacy OsOptimization = "LINUX_LEGACY"
	OSUnknown     OsOptimization = "UNKNOWN"
)

type Architecture string

const (
	ArchAMD64 Architecture = "AMD64"
	ArchARM64 Architecture = "ARM64"
)

type SetServerState string
type SetServerStateOption string

const (
	SetStateOn        SetServerState = "ON"
	SetStateOff       SetServerState = "OFF"
	SetStateSuspended SetServerState = "SUSPENDED"
)

const (
	SetStateOptionPowercycle SetServerStateOption = "POWERCYCLE"
	SetStateOptionReset      SetServerStateOption = "RESET"
	SetStateOptionPoweroff   SetServerStateOption = "POWEROFF"
)

// ======================
// OTHER DEFINITIONS
// ======================

type ServerListMinimal struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Hostname *string   `json:"hostname"`
	Nickname *string   `json:"nickname"`
	Disabled bool      `json:"disabled"`
	Template *Template `json:"template"`
}

type S3DownloadInfos struct {
	Filename                            string              `json:"filename"`
	PresignedURL                        string              `json:"presignedUrl"`
	PresignedURLValidityDurationInHours int                 `json:"presignedUrlValidityDurationInHours"`
	Headers                             map[string][]string `json:"headers"`
}

type GuestAgentData struct {
	GuestAgentAvailable bool           `json:"guestAgentAvailable"`
	GuestAgentData      map[string]any `json:"guestAgentData"`
}

// ======================
// Logs
// ======================

type Log struct {
	Type          LogType     `json:"type"`
	ExecutingUser UserMinimal `json:"executingUser"`
	Date          time.Time   `json:"date"`
	LogKey        string      `json:"logKey"`
	Message       string      `json:"message"`
}

type UserMinimal struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Company   string `json:"company"`
}

type LogType string

const (
	LogTypeError   LogType = "ERROR"
	LogTypeWarning LogType = "WARNING"
	LogTypeInfo    LogType = "INFO"
)

// ======================
// Rescue System
// ======================

type RescueSystemStatus struct {
	Active   bool    `json:"active"`
	Password *string `json:"password"` // nullable
}
