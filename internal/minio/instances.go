package minio

import (
	"github.com/venezia/minio-grpc-admin/internal/minio/errors"
	"github.com/venezia/minio-grpc-admin/internal/minio/models"
)

func GetInstances(username string, group string) ([]*models.Instance, error) {
	var results []*models.Instance
	if username == "" {
		return nil, errors.New(errors.ErrUnauthorized, "wrong username provided", nil)
	}

	results = append(results, &models.Instance{
		Name: "dummy",
		Details: models.InstanceDetails{
			Capacity: 536870912000,
			Usage:    322122547200,
		},
		Status: models.InstanceStatus_READY,
		ConnectionInformation: models.InstanceConnectionInformation{
			Endpoint: "https://cnct-harbor-sos.cloud.sdsamerica.net:8501",
			Bucket:   "harbor",
			Username: "dummy",
			Password: "not-real",
		},
	})

	return results, nil
}
