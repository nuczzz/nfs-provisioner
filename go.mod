module nfs-provisioner

go 1.13

require (
	github.com/miekg/dns v1.1.49 // indirect
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.10.0 // indirect
	golang.org/x/time v0.0.0-20220411224347-583f2d630306 // indirect
	k8s.io/api v0.19.1
	k8s.io/apimachinery v0.19.1
	k8s.io/client-go v0.19.1
	sigs.k8s.io/sig-storage-lib-external-provisioner/v8 v8.0.0
)
