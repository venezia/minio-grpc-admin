/*
 * MinIO Cloud Storage, (C) 2017 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Cloned effectively from https://github.com/minio/minio/blob/6e138f955e6bbd033a6d51753a72ccbebfb1e559/pkg/madmin/info-commands.go

package models

import "time"

// BackendType - represents different backend types.
type BackendType int

// Enum for different backend types.
const (
	Unknown BackendType = iota
	// Filesystem backend.
	FS
	// Multi disk Erasure (single, distributed) backend.
	Erasure

	// Add your own backend.
)

// StorageInfo - represents total capacity of underlying storage.
type StorageInfo struct {
	Disks []Disk

	// Backend type.
	Backend struct {
		// Represents various backend types, currently on FS and Erasure.
		Type BackendType

		// Following fields are only meaningful if BackendType is Erasure.
		OnlineDisks      BackendDisks // Online disks during server startup.
		OfflineDisks     BackendDisks // Offline disks during server startup.
		StandardSCData   int          // Data disks for currently configured Standard storage class.
		StandardSCParity int          // Parity disks for currently configured Standard storage class.
		RRSCData         int          // Data disks for currently configured Reduced Redundancy storage class.
		RRSCParity       int          // Parity disks for currently configured Reduced Redundancy storage class.
	}
}

// BackendDisks - represents the map of endpoint-disks.
type BackendDisks map[string]int

// DataUsageInfo represents data usage of an Object API
type DataUsageInfo struct {
	// LastUpdate is the timestamp of when the data usage info was last updated.
	// This does not indicate a full scan.
	LastUpdate       time.Time `json:"lastUpdate"`
	ObjectsCount     uint64    `json:"objectsCount"`
	ObjectsTotalSize uint64    `json:"objectsTotalSize"`

	// ObjectsSizesHistogram contains information on objects across all buckets.
	// See ObjectsHistogramIntervals.
	ObjectsSizesHistogram map[string]uint64 `json:"objectsSizesHistogram"`

	BucketsCount uint64 `json:"bucketsCount"`

	// BucketsSizes is "bucket name" -> size.
	BucketsSizes map[string]uint64 `json:"bucketsSizes"`
}

// InfoMessage container to hold server admin related information.
type InfoMessage struct {
	Mode         string             `json:"mode,omitempty"`
	Domain       []string           `json:"domain,omitempty"`
	Region       string             `json:"region,omitempty"`
	SQSARN       []string           `json:"sqsARN,omitempty"`
	DeploymentID string             `json:"deploymentID,omitempty"`
	Buckets      Buckets            `json:"buckets,omitempty"`
	Objects      Objects            `json:"objects,omitempty"`
	Usage        Usage              `json:"usage,omitempty"`
	Services     Services           `json:"services,omitempty"`
	Backend      interface{}        `json:"backend,omitempty"`
	Servers      []ServerProperties `json:"servers,omitempty"`
}

// Services contains different services information
type Services struct {
	Vault         Vault                         `json:"vault,omitempty"`
	LDAP          LDAP                          `json:"ldap,omitempty"`
	Logger        []Logger                      `json:"logger,omitempty"`
	Audit         []Audit                       `json:"audit,omitempty"`
	Notifications []map[string][]TargetIDStatus `json:"notifications,omitempty"`
}

// Buckets contains the number of buckets
type Buckets struct {
	Count uint64 `json:"count,omitempty"`
}

// Objects contains the number of objects
type Objects struct {
	Count uint64 `json:"count,omitempty"`
}

// Usage contains the tottal size used
type Usage struct {
	Size uint64 `json:"size,omitempty"`
}

// Vault - Fetches the Vault status
type Vault struct {
	Status  string `json:"status,omitempty"`
	Encrypt string `json:"encrypt,omitempty"`
	Decrypt string `json:"decrypt,omitempty"`
}

// LDAP contains ldap status
type LDAP struct {
	Status string `json:"status,omitempty"`
}

// Status of endpoint
type Status struct {
	Status string `json:"status,omitempty"`
}

// Audit contains audit logger status
type Audit map[string]Status

// Logger contains logger status
type Logger map[string]Status

// TargetIDStatus containsid and status
type TargetIDStatus map[string]Status

// backendType - indicates the type of backend storage
type backendType string

const (
	// FsType - Backend is FS Type
	FsType = backendType("FS")
	// ErasureType - Backend is Erasure type
	ErasureType = backendType("Erasure")
)

// FSBackend contains specific FS storage information
type FSBackend struct {
	Type backendType `json:"backendType,omitempty"`
}

// ErasureBackend contains specific erasure storage information
type ErasureBackend struct {
	Type         backendType `json:"backendType,omitempty"`
	OnlineDisks  int         `json:"onlineDisks,omitempty"`
	OfflineDisks int         `json:"offlineDisks,omitempty"`
	// Data disks for currently configured Standard storage class.
	StandardSCData int `json:"standardSCData,omitempty"`
	// Parity disks for currently configured Standard storage class.
	StandardSCParity int `json:"standardSCParity,omitempty"`
	// Data disks for currently configured Reduced Redundancy storage class.
	RRSCData int `json:"rrSCData,omitempty"`
	// Parity disks for currently configured Reduced Redundancy storage class.
	RRSCParity int `json:"rrSCParity,omitempty"`
}

// ServerProperties holds server information
type ServerProperties struct {
	State    string            `json:"state,omitempty"`
	Endpoint string            `json:"endpoint,omitempty"`
	Uptime   int64             `json:"uptime,omitempty"`
	Version  string            `json:"version,omitempty"`
	CommitID string            `json:"commitID,omitempty"`
	Network  map[string]string `json:"network,omitempty"`
	Disks    []Disk            `json:"drives,omitempty"`
}

// Disk holds Disk information
type Disk struct {
	Endpoint        string  `json:"endpoint,omitempty"`
	RootDisk        bool    `json:"rootDisk,omitempty"`
	DrivePath       string  `json:"path,omitempty"`
	Healing         bool    `json:"healing,omitempty"`
	State           string  `json:"state,omitempty"`
	UUID            string  `json:"uuid,omitempty"`
	Model           string  `json:"model,omitempty"`
	TotalSpace      uint64  `json:"totalspace,omitempty"`
	UsedSpace       uint64  `json:"usedspace,omitempty"`
	AvailableSpace  uint64  `json:"availspace,omitempty"`
	ReadThroughput  float64 `json:"readthroughput,omitempty"`
	WriteThroughPut float64 `json:"writethroughput,omitempty"`
	ReadLatency     float64 `json:"readlatency,omitempty"`
	WriteLatency    float64 `json:"writelatency,omitempty"`
	Utilization     float64 `json:"utilization,omitempty"`
}
