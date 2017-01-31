package main

import (
	"flag"
	"os"

	"github.com/go-kit/kit/log"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kiubernetes/vendor/github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/client/unversioned"
	kubectl_util "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

var (
	appName    = "podfilter"
	appVersion = "DEV"
	appTag     = "HEAD"
)

var (
	stdErr = os.Stderr
	logger *log.Context
)

var (
	flags = pflag.NewFlagSet("", pflag.ExitOnError)
)

func main() {
	flags.AddGoFlagSet(flag.CommandLine)
	flags.Parse(os.Args)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(stdErr))
	logger = log.NewContext(logger).With(
		"t", log.DefaultTimestamp,
		"appName", appName,
		"appVersion", appVersion,
		"appTag", appTag,
		"caller", log.DefaultCaller,
	)

	var clientConfig clientcmd.ClientConfig
	var kubeClient unversioned.Client
	{
		var err error
		clientConfig = kubectl_util.DefaultClientConfig(flags)
		kubeClient, err = unversioned.NewInCluster()
		if err != nil {
			config, err := clientConfig.ClientConfig()
			if err != nil {
				logger.Log("level", "error",
					"msg", "error configuring the kube client",
					"err", err,
				)
				os.Exit(1)
			}
			kubeClient, err = unversioned.New(config)
			if err != nil {
				logger.Log("level", "error",
					"msg", "error creating kube client",
					"err", err,
				)
				os.Exit(1)
			}
		}
	}
}
