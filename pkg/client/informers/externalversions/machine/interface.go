// This file was automatically generated by informer-gen

package machine

import (
	internalinterfaces "github.com/gardener/node-controller-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gardener/node-controller-manager/pkg/client/informers/externalversions/machine/v1alpha1"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1alpha1 provides access to shared informers for resources in V1alpha1.
	V1alpha1() v1alpha1.Interface
}

type group struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &group{f}
}

// V1alpha1 returns a new v1alpha1.Interface.
func (g *group) V1alpha1() v1alpha1.Interface {
	return v1alpha1.New(g.SharedInformerFactory)
}