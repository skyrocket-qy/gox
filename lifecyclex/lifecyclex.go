package lifecyclex

import (
	"context"

	"github.com/skyrocket-qy/erx"
	"github.com/skyrocket-qy/gox/dsa/graph/topologicalsort"
)

type Closer func(context.Context) error

type SimpleLifecycle struct {
	closers []Closer
}

func NewSimpleLifecycle() *SimpleLifecycle {
	return &SimpleLifecycle{}
}

func (l *SimpleLifecycle) Add(fn Closer) {
	l.closers = append(l.closers, fn)
}

func (l *SimpleLifecycle) Shutdown(c context.Context) error {
	for i := len(l.closers) - 1; i >= 0; i-- {
		if err := l.closers[i](c); err != nil {
			return erx.W(err)
		}
	}

	return nil
}

// ... SimpleLifecycle remains the same ...

type ConcurrentLifecycle struct {
	// appUpstreams maps an app to the components it depends on (its upstreams).
	appUpstreams map[any][]any
	// appCloser stores the shutdown function for each app.
	appCloser map[any]Closer
}

func NewConcurrentLifecycle() *ConcurrentLifecycle {
	return &ConcurrentLifecycle{
		appUpstreams: make(map[any][]any),
		appCloser:    make(map[any]Closer),
	}
}

func (l *ConcurrentLifecycle) Add(app any, fn Closer, deps ...any) {
	l.appUpstreams[app] = deps
	l.appCloser[app] = fn
}

func (l *ConcurrentLifecycle) Shutdown(c context.Context) error {
	// Build the reverse graph for the topological sort.
	// The shutdown process requires dependents to be shut down before their dependencies.
	// So, an edge A -> B means A must shut down before B.
	// Our `appUpstreams` map is `A -> {B, C}` which means A depends on B and C.
	// So for shutdown, we need to process things that nothing depends on first.

	// The graph for topsort should be: dependent -> {list of its dependencies}
	// Your `appUpstreams` map is already in this format: app -> {list of its upstreams/dependencies}
	// However, the shutdown logic requires closing `app` before closing its `deps`.
	// This means `app` is a "dependent" and must be processed first.
	// Kahn's algorithm processes nodes with an in-degree of 0 first.
	// To shut down `app` first, it must have an in-degree of 0.
	// This requires inverting the graph: dependency -> {list of dependents}

	shutdownGraph := make(map[any][]any)
	allNodes := make(map[any]struct{})

	for app, deps := range l.appUpstreams {
		allNodes[app] = struct{}{}
		if _, ok := shutdownGraph[app]; !ok {
			shutdownGraph[app] = []any{}
		}
		for _, dep := range deps {
			allNodes[dep] = struct{}{}
			shutdownGraph[dep] = append(shutdownGraph[dep], app)
		}
	}

	// Define the worker function that will be executed by the topological sort.
	worker := func(ctx context.Context, node any) error {
		// Only execute the closer if one was registered for the node.
		if closer, ok := l.appCloser[node]; ok && closer != nil {
			return closer(ctx)
		}
		return nil
	}

	// Delegate the entire concurrent shutdown process to the generic function.
	return topologicalsort.ConcurrentTopologicalSort(c, shutdownGraph, worker)
}
