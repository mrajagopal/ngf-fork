package kubernetes

import "sigs.k8s.io/controller-runtime/pkg/client"

//go:generate go tool counterfeiter -generate

//counterfeiter:generate . Reader

// Reader allows getting and listing resources from a cache.
// This interface is introduced for testing to mock the methods from
// sigs.k8s.io/controller-runtime/pkg/client.Reader.
type Reader interface {
	client.Reader
}
