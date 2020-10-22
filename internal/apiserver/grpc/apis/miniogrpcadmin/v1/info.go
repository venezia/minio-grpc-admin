package v1

import (
	"golang.org/x/net/context"

	"github.com/venezia/minio-grpc-admin/internal/minio"
	"github.com/venezia/minio-grpc-admin/internal/minio/models"
	pb "github.com/venezia/minio-grpc-admin/pkg/generated/api/minioadmin/v1"
)

func (s *Server) GetServerInformation(ctx context.Context, in *pb.GetServerInformationMsg) (*pb.GetServerInformationReply, error) {
	info, err := minio.GetServerInfo(models.Endpoint{
		Endpoint:        in.Endpoint.Endpoint,
		AccessKeyID:     in.Endpoint.AccessKeyId,
		SecretAccessKey: in.Endpoint.SecretAccessKey,
		SSL:             in.Endpoint.Ssl,
	})
	if err != nil {
		return nil, err
	}

	output := &pb.GetServerInformationReply{
		Mode:         info.Mode,
		Domain:       info.Domain,
		Region:       info.Region,
		SqsArn:       info.SQSARN,
		DeploymentId: info.DeploymentID,
		Buckets:      &pb.GetServerInformationReply_Buckets{Count: info.Buckets.Count},
		Objects:      &pb.GetServerInformationReply_Objects{Count: info.Objects.Count},
		Usage:        &pb.GetServerInformationReply_Usage{Size: info.Usage.Size},
		Services:     convertInfoServices(info.Services),
		Backend:      nil,
		Servers:      convertInfoServers(info.Servers),
	}

	fsBackend, ok := info.Backend.(models.FSBackend)
	if ok {
		output.Backend = &pb.GetServerInformationReply_FsBackend{
			FsBackend: &pb.GetServerInformationReply_FSBackendType{ BackendType: string(fsBackend.Type) },
		}
	} else if erasureBackend, ok := info.Backend.(models.ErasureBackend); ok {
		output.Backend = &pb.GetServerInformationReply_ErasureBackend{
			ErasureBackend: &pb.GetServerInformationReply_ErasureBackendType{
				BackendType:      string(erasureBackend.Type),
				OnlineDisks:      int32(erasureBackend.OnlineDisks),
				OfflineDisks:     int32(erasureBackend.OfflineDisks),
				StandardScData:   int32(erasureBackend.StandardSCData),
				StandardScParity: int32(erasureBackend.StandardSCParity),
				RrscData:         int32(erasureBackend.RRSCData),
				RrscParity:       int32(erasureBackend.RRSCParity),
			},
		}
	}


	return output, nil
}

func convertInfoServices(input models.Services) *pb.GetServerInformationReply_Services {
	return &pb.GetServerInformationReply_Services{
		Vault:         &pb.GetServerInformationReply_Services_Vault{
			Status:  input.Vault.Status,
			Encrypt: input.Vault.Encrypt,
			Decrypt: input.Vault.Decrypt,
		},
		Ldap:          &pb.GetServerInformationReply_Services_LDAP{Status: input.LDAP.Status},
		Loggers:       convertInfoLoggers(input.Logger),
		Audits:        convertInfoAudits(input.Audit),
		Notifications: convertInfoNotifications(input.Notifications),
	}
}

func convertInfoLoggers(input []models.Logger) []*pb.GetServerInformationReply_Services_Logger {
	var loggers []*pb.GetServerInformationReply_Services_Logger

	for _, loggerEntry := range input {
		var loggerMap []*pb.GetServerInformationReply_Services_Logger_LoggerMap
		for logEntryName, logEntryStatus := range loggerEntry {
			loggerMap = append(loggerMap, &pb.GetServerInformationReply_Services_Logger_LoggerMap{
				Key:    logEntryName,
				Status: logEntryStatus.Status,
			})
		}
		loggers = append(loggers, &pb.GetServerInformationReply_Services_Logger{LoggerMap: loggerMap})
	}

	return loggers
}

func convertInfoAudits(input []models.Audit) []*pb.GetServerInformationReply_Services_Audit {
	var audits []*pb.GetServerInformationReply_Services_Audit

	for _, auditEntry := range input {
		var auditMap []*pb.GetServerInformationReply_Services_Audit_AuditMap
		for auditEntryName, auditEntryStatus := range auditEntry {
			auditMap = append(auditMap, &pb.GetServerInformationReply_Services_Audit_AuditMap{
				Key:    auditEntryName,
				Status: auditEntryStatus.Status,
			})
		}
		audits = append(audits, &pb.GetServerInformationReply_Services_Audit{AuditMap: auditMap})
	}

	return audits
}

func convertInfoNotifications(input []map[string][]models.TargetIDStatus) []*pb.GetServerInformationReply_Services_Notifications {
	var notifications []*pb.GetServerInformationReply_Services_Notifications

	for _, notificationEntry := range input {
		// So now we're in the context of map[string][]models.TargetIDStatus
		var notificationMap []*pb.GetServerInformationReply_Services_Notifications_NotificationMap
		for notificationEntryName, notificationEntryValue := range notificationEntry {
			// So now we're in the context of []models.TargetIDStatus, with notificationEntryName being the key
			var targetStatii []*pb.GetServerInformationReply_Services_Notifications_NotificationMap_TargetIDStatus
			for _, targetStatus := range notificationEntryValue {
				// So now we're in the context of models.TargetIDStatus, which is really map[string]models.Status
				var targetMap []*pb.GetServerInformationReply_Services_Notifications_NotificationMap_TargetIDStatus_TargetIDStatusMap
				for targetEntryName, targetEntryValue := range targetStatus {
					// Now we're in the context of models.Status
					targetMap = append(targetMap, &pb.GetServerInformationReply_Services_Notifications_NotificationMap_TargetIDStatus_TargetIDStatusMap{
						Key:    targetEntryName,
						Status: targetEntryValue.Status,
					})
				}
				targetStatii = append(targetStatii, &pb.GetServerInformationReply_Services_Notifications_NotificationMap_TargetIDStatus{
					TargetIdStatusMap: targetMap,
				})
			}
			notificationMap = append(notificationMap, &pb.GetServerInformationReply_Services_Notifications_NotificationMap{
				Key:            notificationEntryName,
				TargetIdStatus: targetStatii,
			})
		}
		notifications = append(notifications, &pb.GetServerInformationReply_Services_Notifications{NotificationMap: notificationMap})
	}

	return notifications
}

func convertInfoServers(input []models.ServerProperties) []*pb.GetServerInformationReply_ServerProperties {
	var output []*pb.GetServerInformationReply_ServerProperties

	for _, server := range input {
		output = append(output, &pb.GetServerInformationReply_ServerProperties{
			State:    server.State,
			Endpoint: server.Endpoint,
			Uptime:   server.Uptime,
			Version:  server.Version,
			CommitId: server.CommitID,
			Network:  convertInfoNetwork(server.Network),
			Disks:    convertInfoDisks(server.Disks),
		})
	}

	return output
}

func convertInfoNetwork(input map[string]string) []*pb.GetServerInformationReply_ServerProperties_Network {
	var output []*pb.GetServerInformationReply_ServerProperties_Network

	for name, value := range input {
		output = append(output, &pb.GetServerInformationReply_ServerProperties_Network{
			Name:  name,
			Value: value,
		})
	}

	return output
}

func convertInfoDisks(input []models.Disk) []*pb.GetServerInformationReply_ServerProperties_Disk {
	var output []*pb.GetServerInformationReply_ServerProperties_Disk

	for _, diskEntry := range input {
		output = append(output, &pb.GetServerInformationReply_ServerProperties_Disk{
			Endpoint:        diskEntry.Endpoint,
			RootDisk:        diskEntry.RootDisk,
			DrivePath:       diskEntry.DrivePath,
			Healing:         diskEntry.Healing,
			State:           diskEntry.State,
			Uuid:            diskEntry.UUID,
			Model:           diskEntry.Model,
			TotalSpace:      diskEntry.TotalSpace,
			UsedSpace:       diskEntry.UsedSpace,
			AvailableSpace:  diskEntry.AvailableSpace,
			ReadThroughput:  float32(diskEntry.ReadThroughput),
			WriteThroughput: float32(diskEntry.WriteThroughPut),
			ReadLatency:     float32(diskEntry.ReadLatency),
			WriteLatency:    float32(diskEntry.WriteLatency),
			Utilization:     float32(diskEntry.Utilization),
		})
	}

	return output
}