package models

const (
	InstanceStatus_UNKNOWN      int32 = 0
	InstanceStatus_PENDING      int32 = 1
	InstanceStatus_PROVISIONING int32 = 2
	InstanceStatus_READY        int32 = 3
	InstanceStatus_DEGRADED     int32 = 4
	InstanceStatus_OFFLINE      int32 = 5
	InstanceStatus_DELETING     int32 = 6
	InstanceStatus_DELETED      int32 = 7
)

var (
	InstanceStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "PENDING",
		2: "PROVISIONING",
		3: "READY",
		4: "DEGRADED",
		5: "OFFLINE",
		6: "DELETING",
		7: "DELETED",
	}
	InstanceStatus_value = map[string]int32{
		"UNKNOWN":      0,
		"PENDING":      1,
		"PROVISIONING": 2,
		"READY":        3,
		"DEGRADED":     4,
		"OFFLINE":      5,
		"DELETING":     6,
		"DELETED":      7,
	}
)

type Instance struct {
	ConnectionInformation InstanceConnectionInformation
	Details               InstanceDetails
	Name                  string
	Status                int32
}

type InstanceDetails struct {
	Capacity int64
	Usage    int64
}

type InstanceConnectionInformation struct {
	Endpoint string
	Bucket   string
	Username string
	Password string
}
