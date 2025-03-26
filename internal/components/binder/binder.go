package binder

import (
	"go.uber.org/dig"
	"sync"
)

type ContainerImplement interface {
	Provider() *Provider
	Invoker() *Invoker
}

func NewContainer(container *dig.Container, impl ContainerImplement) *Container {
	return (&Container{
		Container: container,
		impl:      impl,
	}).construction()
}

type Container struct {
	once sync.Once

	impl      ContainerImplement
	Container *dig.Container
}

func (b *Container) construction() *Container {
	b.once.Do(func() {
		if b.impl.Provider() != nil {
			for _, binding := range b.impl.Provider().getBindings() {
				if err := b.Container.Provide(binding.Constructor, binding.Options...); err != nil {
					panic(err)
				}
			}
		}

		if b.impl.Invoker() != nil {
			for _, invocation := range b.impl.Invoker().getInvocations() {
				if err := b.Container.Invoke(invocation.Constructor, invocation.Options...); err != nil {
					panic(err)
				}
			}
		}
	})

	return b
}

// binding represents a dependency with its constructor and options
type binding struct {
	Constructor interface{}
	Options     []dig.ProvideOption
}

func NewProvider() *Provider {
	return &Provider{
		bindings: make([]binding, 0),
	}
}

// Provider is a concrete implementation of the Provider interface
type Provider struct {
	bindings []binding
}

func (p *Provider) Provide(constructor interface{}, opts ...dig.ProvideOption) *Provider {
	p.bindings = append(p.bindings, binding{
		Constructor: constructor,
		Options:     opts,
	})
	return p
}

func (p *Provider) getBindings() []binding {
	return p.bindings
}

// Invocation represents a function to be invoked with options
type invocation struct {
	Constructor interface{}
	Options     []dig.InvokeOption
}

func NewInvoker() *Invoker {
	return &Invoker{
		invocations: make([]invocation, 0),
	}
}

// Invoker is a concrete implementation of the Invoker interface
type Invoker struct {
	invocations []invocation
}

func (i *Invoker) Invoke(constructor interface{}, opts ...dig.InvokeOption) *Invoker {
	i.invocations = append(i.invocations, invocation{
		Constructor: constructor,
		Options:     opts,
	})
	return i
}

func (i *Invoker) getInvocations() []invocation {
	return i.invocations
}
