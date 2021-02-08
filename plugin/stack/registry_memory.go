package stack

import (
	"github.com/stack-labs/stack/registry"
	"github.com/stack-labs/stack/registry/memory"
)

type memoryRegistryPlugin struct{}

func (m *memoryRegistryPlugin) Name() string {
	return "memory"
}

func (m *memoryRegistryPlugin) Options() []registry.Option {
	return nil
}

func (m *memoryRegistryPlugin) New(opts ...registry.Option) registry.Registry {
	return memory.NewRegistry(opts...)
}
