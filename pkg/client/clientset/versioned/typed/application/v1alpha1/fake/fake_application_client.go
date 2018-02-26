package fake

import (
	v1alpha1 "github.com/argoproj/argo-cd/pkg/client/clientset/versioned/typed/application/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeArgoprojV1alpha1 struct {
	*testing.Fake
}

func (c *FakeArgoprojV1alpha1) Applications(namespace string) v1alpha1.ApplicationInterface {
	return &FakeApplications{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeArgoprojV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
