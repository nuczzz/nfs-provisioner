package main

import (
	"context"
	"flag"
	"log"

	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/sig-storage-lib-external-provisioner/v8/controller"
)

var (
	server         = flag.String("server", "", "nfs server ip")
	serverPath     = flag.String("serverPath", "", "nfs server root mount path")
	mountPath      = flag.String("mountPath", "/mount", "local mount path")
	provisoner     = flag.String("provisioner", "nfs", "provisioner name")
	kubeconfig     = flag.String("kubeconfig", "", "kubernetes config file")
	leaderElection = flag.Bool("leaderElection", false, "leader election switch")
)

func main() {
	flag.Parse()
	log.Printf("nfs server: %s, rootMountPath: %s", *server, *serverPath)

	if *server == "" || *serverPath == "" || *mountPath == "" {
		log.Fatalf("flag server/serverPath/mountPath not set")
	}

	k8sClient, err := initK8sClient(*kubeconfig)
	if err != nil {
		log.Fatalf("initK8sClient error: %s", err.Error())
	}

	nfs := &nfsProvisioner{
		server:        *server,
		rootMountPath: *serverPath,
		mountPath:     *mountPath,
	}

	controller.NewProvisionController(
		k8sClient,
		*provisoner,
		nfs,
		controller.LeaderElection(*leaderElection),
	).Run(context.TODO())
}

func initK8sClient(config string) (*kubernetes.Clientset, error) {
	cfg, err := clientcmd.BuildConfigFromFlags("", config)
	if err != nil {
		return nil, errors.Wrap(err, "BuildConfigFromFlags error")
	}

	return kubernetes.NewForConfig(cfg)
}
