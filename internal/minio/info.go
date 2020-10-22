package minio

import (
	"context"
	"github.com/minio/minio/pkg/madmin"
	"github.com/venezia/minio-grpc-admin/internal/minio/models"
)

func GetServerInfo(endpoint models.Endpoint) (*models.InfoMessage, error) {

	client, err := madmin.New(endpoint.Endpoint, endpoint.AccessKeyID, endpoint.SecretAccessKey, endpoint.SSL)
	if err != nil {
		return nil, err
	}

	serverInformation, err := client.ServerInfo(context.Background())
	if err != nil {
		return nil, err
	}

	return &models.InfoMessage{
		Mode:         serverInformation.Mode,
		Domain:       serverInformation.Domain,
		Region:       serverInformation.Region,
		SQSARN:       serverInformation.SQSARN,
		DeploymentID: serverInformation.DeploymentID,
		Buckets:      models.Buckets{Count: serverInformation.Buckets.Count},
		Objects:      models.Objects{Count: serverInformation.Objects.Count},
		Usage:        models.Usage{Size: serverInformation.Usage.Size},
		Services:     convertServices(serverInformation.Services),
		Backend:      convertBackend(serverInformation.Backend),
		Servers:      convertServerProperties(serverInformation.Servers),
	}, nil
}

func convertServices(input madmin.Services) models.Services {
	return models.Services{
		Vault:         models.Vault{Status: input.Vault.Status, Encrypt: input.Vault.Encrypt, Decrypt: input.Vault.Decrypt},
		LDAP:          models.LDAP{Status: input.LDAP.Status},
		Logger:        convertLogger(input.Logger),
		Audit:         convertAudit(input.Audit),
		Notifications: convertServiceNotifcations(input.Notifications),
	}
}

func convertLogger(input []madmin.Logger) []models.Logger {
	var output []models.Logger

	for _, item := range input {
		tempMap := make(models.Logger)
		for k, v := range item {
			tempMap[k] = models.Status(v)
		}
		output = append(output, tempMap)
	}

	return output
}

func convertAudit(input []madmin.Audit) []models.Audit {
	var output []models.Audit

	for _, item := range input {
		tempMap := make(models.Audit)
		for k, v := range item {
			tempMap[k] = models.Status(v)
		}
		output = append(output, tempMap)
	}

	return output
}

func convertServiceNotifcations(input []map[string][]madmin.TargetIDStatus) []map[string][]models.TargetIDStatus {
	// This literally is []map[string][]map[string]Status where status is a struct
	var output []map[string][]models.TargetIDStatus

	// Because this is largely meaningless, let's go with one, two, three, four for the four layers deep...
	for _, one := range input {
		// So now we're working with one == map[string][]madmin.TargetIDStatus
		oneMap := make(map[string][]models.TargetIDStatus)
		for i, two := range one {
			// So now we're working with two == []models.TargetIDStatus, with key value i
			var twoArray []models.TargetIDStatus
			for _, three := range two {
				// So now we're working with three = models.TargetIDStatus
				// which is also known as map[string]madmin.Status
				threeMap := make(map[string]models.Status)
				for j, four := range three {
					// Finally unwound, now just clone the item into the map
					threeMap[j] = models.Status{Status: four.Status}
				}
				twoArray = append(twoArray, threeMap)
			}
			oneMap[i] = twoArray
		}
		output = append(output, oneMap)
	}

	return output
}

func convertBackend(input interface{}) interface{} {
	// Right now we only know of two backends - FSBackend and ErasureType

	// Unsure why I can't do a type assertion here!
	backendMap, ok := input.(map[string]interface{})
	if ok {
		if backend, ok := backendMap["backendType"].(string); ok {
			if backend == string(models.FsType) {
				return models.FSBackend{Type: models.FsType}
			} else if backend == string(models.ErasureType) {
				output := models.ErasureBackend{Type: models.ErasureType}
				if offlineDisks, ok := backendMap["offlineDisks"].(float64); ok {
					output.OfflineDisks = int(offlineDisks)
				}
				if onlineDisks, ok := backendMap["onlineDisks"].(float64); ok {
					output.OnlineDisks = int(onlineDisks)
				}
				if rrSCData, ok := backendMap["rrSCData"].(float64); ok {
					output.RRSCData = int(rrSCData)
				}
				if rrSCParity, ok := backendMap["rrSCParity"].(float64); ok {
					output.RRSCParity = int(rrSCParity)
				}
				if standardSCData, ok := backendMap["standardSCData"].(float64); ok {
					output.StandardSCData = int(standardSCData)
				}
				if standardSCParity, ok := backendMap["standardSCParity"].(float64); ok {
					output.StandardSCParity = int(standardSCParity)
				}
				return output
			}
		}
	}

	/*	logger.Infof("%v", input)
		logger.Infof("%T", input)
		// FSBackend is pretty generic, don't really need it, just need to see if it is it
		_, ok := input.(madmin.FSBackend)
		if ok {
			return models.FSBackend{Type: models.FsType}
		}

		// Ok, checking for erasureBackend instead...
		erasureBackend, ok := input.(madmin.ErasureBackend)
		if ok {
			logger.Infof("Hi Mom3")
			return models.ErasureBackend{
				Type: models.ErasureType,
				OnlineDisks: erasureBackend.OnlineDisks,
				OfflineDisks: erasureBackend.OfflineDisks,
				StandardSCData: erasureBackend.StandardSCData,
				StandardSCParity: erasureBackend.StandardSCParity,
				RRSCData: erasureBackend.RRSCData,
				RRSCParity: erasureBackend.RRSCParity,
			}
		}
	*/
	return nil
}

func convertServerProperties(input []madmin.ServerProperties) []models.ServerProperties {
	var output []models.ServerProperties

	for _, entry := range input {
		output = append(output, models.ServerProperties{
			State:    entry.State,
			Endpoint: entry.Endpoint,
			Uptime:   entry.Uptime,
			Version:  entry.Version,
			CommitID: entry.CommitID,
			Network:  convertServerPropertiesNetwork(entry.Network),
			Disks:    convertServerPropertiesDisk(entry.Disks),
		})
	}

	return output
}

func convertServerPropertiesNetwork(input map[string]string) map[string]string {
	output := make(map[string]string)

	for key, value := range input {
		output[key] = value
	}

	return output
}

func convertServerPropertiesDisk(input []madmin.Disk) []models.Disk {
	var output []models.Disk

	for _, entry := range input {
		output = append(output, models.Disk{
			Endpoint:        entry.Endpoint,
			RootDisk:        entry.RootDisk,
			DrivePath:       entry.DrivePath,
			Healing:         entry.Healing,
			State:           entry.State,
			UUID:            entry.UUID,
			Model:           entry.Model,
			TotalSpace:      entry.TotalSpace,
			UsedSpace:       entry.UsedSpace,
			AvailableSpace:  entry.AvailableSpace,
			ReadThroughput:  entry.ReadThroughput,
			WriteThroughPut: entry.WriteThroughPut,
			ReadLatency:     entry.ReadLatency,
			WriteLatency:    entry.WriteLatency,
			Utilization:     entry.Utilization,
		})
	}

	return output
}
