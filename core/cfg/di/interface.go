package di

// ContainerProvider provides and interface to work with DI container.
type ContainerProvider interface {
	// Build builds application dependencies.
	Build() error
	// Get returns the dependency by its name.
	Get(name string) interface{}
	// RegisterDependency registers a dependency in DI container.
	RegisterDependency(
		depName string,
		registrar dependencyRegistrar,
		disposer dependencyDisposer,
	) error
	// Shutdown makes graceful shutdown of whole DI Container.
	Shutdown() error
}

// DependencyShutdowner provides an interface for graceful service shutdown.
type DependencyShutdowner interface {
	// Shutdown makes graceful shutdown of any dependency which implements this interface.
	Shutdown()
}
