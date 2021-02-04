package stack

import (
	"github.com/stack-labs/stack/registry"
	"github.com/stack-labs/stack/registry/service"
)

type serviceRegistryPlugin struct{}

func (s *serviceRegistryPlugin) Name() string {
	return "service"
}

func (s *serviceRegistryPlugin) Options() []registry.Option {
	return nil
}

func (s *serviceRegistryPlugin) New(opts ...registry.Option) registry.Registry {
	return service.NewRegistry(opts...)
}
