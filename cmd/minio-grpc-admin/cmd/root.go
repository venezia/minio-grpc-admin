package cmd

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/juju/loggo"
	"github.com/soheilhy/cmux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/venezia/minio-grpc-admin/internal/apiserver"
	"github.com/venezia/minio-grpc-admin/internal/flags"
	"github.com/venezia/minio-grpc-admin/pkg/util/log"
)

const (
	loggerModuleName = "cmd.minio-grpc-admin.cmd"
)

var (
	rootCmd = &cobra.Command{
		Use:   "minio-grpc-admin",
		Short: "Minio gRPC Admin API Server",
		Long:  `The minio gRPC Admin API Service`,
		Run: func(cmd *cobra.Command, args []string) {
			runWebServer()
		},
	}

	logger loggo.Logger
)

func init() {
	viper.SetEnvPrefix(flags.EnvironmentPrefix)
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)

	// using standard library "flag" package
	rootCmd.Flags().Bool(flags.GRPCAuthentication, flags.GRPCAuthenticationDefault, flags.GRPCAuthenticationDescription)
	rootCmd.Flags().Bool(flags.GRPCAuthorization, flags.GRPCAuthorizationDefault, flags.GRPCAuthorizationDescription)
	rootCmd.Flags().Bool(flags.GRPCUI, flags.GRPCUIDefault, flags.GRPCUIDescription)
	rootCmd.Flags().Int(flags.GRPCUIPort, flags.GRPCUIPortDefault, flags.GRPCUIPortDescription)
	rootCmd.Flags().Bool(flags.GRPCRecovery, flags.GRPCRecoveryDefault, flags.GRPCRecoveryDescription)
	rootCmd.Flags().Bool(flags.GRPCReflection, flags.GRPCReflectionDefault, flags.GRPCReflectionDescription)
	rootCmd.Flags().String(flags.LogLevel, flags.LogLevelDefault, flags.LogLevelDescription)
	rootCmd.Flags().String(flags.OIDCAudience, flags.OIDCAudienceDefault, flags.OIDCAudienceDescription)
	rootCmd.Flags().String(flags.OIDCEndpoint, flags.OIDCEndpointDefault, flags.OIDCEndpointDescription)
	rootCmd.Flags().Int(flags.Port, flags.PortDefault, flags.PortDescription)

	handleViperError(flags.GRPCAuthentication, viper.BindPFlag(flags.GRPCAuthentication, rootCmd.Flags().Lookup(flags.GRPCAuthentication)))
	handleViperError(flags.GRPCAuthorization, viper.BindPFlag(flags.GRPCAuthorization, rootCmd.Flags().Lookup(flags.GRPCAuthorization)))
	handleViperError(flags.GRPCUI, viper.BindPFlag(flags.GRPCUI, rootCmd.Flags().Lookup(flags.GRPCUI)))
	handleViperError(flags.GRPCUIPort, viper.BindPFlag(flags.GRPCUIPort, rootCmd.Flags().Lookup(flags.GRPCUIPort)))
	handleViperError(flags.GRPCRecovery, viper.BindPFlag(flags.GRPCRecovery, rootCmd.Flags().Lookup(flags.GRPCRecovery)))
	handleViperError(flags.GRPCReflection, viper.BindPFlag(flags.GRPCReflection, rootCmd.Flags().Lookup(flags.GRPCReflection)))
	handleViperError(flags.LogLevel, viper.BindPFlag(flags.LogLevel, rootCmd.Flags().Lookup(flags.LogLevel)))
	handleViperError(flags.OIDCAudience, viper.BindPFlag(flags.OIDCAudience, rootCmd.Flags().Lookup(flags.OIDCAudience)))
	handleViperError(flags.OIDCEndpoint, viper.BindPFlag(flags.OIDCEndpoint, rootCmd.Flags().Lookup(flags.OIDCEndpoint)))
	handleViperError(flags.Port, viper.BindPFlag(flags.Port, rootCmd.Flags().Lookup(flags.Port)))

	viper.AutomaticEnv()
	rootCmd.Flags().AddGoFlagSet(flag.CommandLine)

	sanitizeLogLevel()
	logger = log.GetModuleLogger(loggerModuleName)
}

func handleViperError(flag string, err error) {
	if err != nil {
		logger.Warningf("viper - could not bind %s, error was: %s", flag, err.Error())
	}
}

func sanitizeLogLevel() {
	if logLevel, ok := loggo.ParseLevel(viper.GetString(flags.LogLevel)); !ok {
		loggo.GetLogger("").SetLogLevel(loggo.INFO)
		//logger.SetLogLevel(loggo.INFO)
	} else {
		loggo.GetLogger("").SetLogLevel(logLevel)
		//logger.SetLogLevel(logLevel)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runWebServer() {
	// get flags
	portNumber := viper.GetInt(flags.Port)
	grpcuiPortNumber := viper.GetInt(flags.GRPCUIPort)

	// Debug for now
	logger.Criticalf("\nParsed Variables: \n"+
		"\tGRPC Authentication: \t%t\n"+
		"\tGRPC Authorization: \t%t\n"+
		"\tgRPC-UI Enabled: \t%t\n"+
		"\tgRPC-UI Port: \t\t%d\n"+
		"\tGRPC Recovery: \t\t%t\n"+
		"\tGRPC Reflection: \t%t\n"+
		"\tDefault Log Level: \t%s\n"+
		"\tOIDC Endpoint: \t\t%s\n"+
		"\tOIDC Audience: \t\t%s\n"+
		"\tPort: \t\t\t%d\n",
		viper.GetBool(flags.GRPCAuthentication),
		viper.GetBool(flags.GRPCAuthorization),
		viper.GetBool(flags.GRPCUI),
		grpcuiPortNumber,
		viper.GetBool(flags.GRPCRecovery),
		viper.GetBool(flags.GRPCReflection),
		loggo.GetLogger("").LogLevel(),
		viper.GetString(flags.OIDCEndpoint),
		viper.GetString(flags.OIDCAudience),
		portNumber,
	)

	var wg sync.WaitGroup
	stop := make(chan struct{})

	logger.Infof("Creating Web Server")
	tcpMux := createWebServer(&apiserver.ServerOptions{PortNumber: portNumber})
	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Infof("Starting to serve requests on port %d", portNumber)
		err := tcpMux.Serve()
		if err != nil {
			logger.Criticalf("could not serve requests on port %d, error was: ", portNumber, err.Error())
		}
	}()

	if viper.GetBool(flags.GRPCUI) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(2 * time.Second)
			logger.Infof("Starting to grpcui requests on port %d", grpcuiPortNumber)
			output, err := exec.Command("grpcui", "-plaintext", "-port", fmt.Sprintf("%d", grpcuiPortNumber), fmt.Sprintf("localhost:%d", portNumber)).Output()
			logger.Errorf("error was: %v, %v", output, err)
		}()
	}

	<-stop
	logger.Infof("Wating for controllers to shut down gracefully")
	wg.Wait()
}

func createWebServer(options *apiserver.ServerOptions) cmux.CMux {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", options.PortNumber))
	if err != nil {
		panic(err)
	}
	tcpMux := cmux.New(conn)

	apiserver.AddServersToMux(tcpMux, options)

	return tcpMux
}
