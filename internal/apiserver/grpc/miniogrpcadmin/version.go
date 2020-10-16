package sosapi

import (
	"gitlab.com/mvenezia/version-info/pkg/version"
	"golang.org/x/net/context"

	pb "github.com/venezia/minio-grpc-admin/pkg/generated/api/minioadmin/v1"
)

func (s *Server) GetVersionInformation(ctx context.Context, in *pb.GetVersionMsg) (*pb.GetVersionReply, error) {
	versionInformation := version.Get()
	reply := &pb.GetVersionReply{
		Ok: true,
		VersionInformation: &pb.GetVersionReply_VersionInformation{
			GitVersion:   versionInformation.GitVersion,
			GitCommit:    versionInformation.GitCommit,
			GitTreeState: versionInformation.GitTreeState,
			BuildDate:    versionInformation.BuildDate,
			GoVersion:    versionInformation.GoVersion,
			Compiler:     versionInformation.Compiler,
			Platform:     versionInformation.Platform,
		},
	}
	return reply, nil
}
