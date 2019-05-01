// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/db"
)

type FakeResourceFactory struct {
	VisibleResourcesStub        func([]string) ([]db.Resource, error)
	visibleResourcesMutex       sync.RWMutex
	visibleResourcesArgsForCall []struct {
		arg1 []string
	}
	visibleResourcesReturns struct {
		result1 []db.Resource
		result2 error
	}
	visibleResourcesReturnsOnCall map[int]struct {
		result1 []db.Resource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceFactory) VisibleResources(arg1 []string) ([]db.Resource, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.visibleResourcesMutex.Lock()
	ret, specificReturn := fake.visibleResourcesReturnsOnCall[len(fake.visibleResourcesArgsForCall)]
	fake.visibleResourcesArgsForCall = append(fake.visibleResourcesArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("VisibleResources", []interface{}{arg1Copy})
	fake.visibleResourcesMutex.Unlock()
	if fake.VisibleResourcesStub != nil {
		return fake.VisibleResourcesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.visibleResourcesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceFactory) VisibleResourcesCallCount() int {
	fake.visibleResourcesMutex.RLock()
	defer fake.visibleResourcesMutex.RUnlock()
	return len(fake.visibleResourcesArgsForCall)
}

func (fake *FakeResourceFactory) VisibleResourcesCalls(stub func([]string) ([]db.Resource, error)) {
	fake.visibleResourcesMutex.Lock()
	defer fake.visibleResourcesMutex.Unlock()
	fake.VisibleResourcesStub = stub
}

func (fake *FakeResourceFactory) VisibleResourcesArgsForCall(i int) []string {
	fake.visibleResourcesMutex.RLock()
	defer fake.visibleResourcesMutex.RUnlock()
	argsForCall := fake.visibleResourcesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResourceFactory) VisibleResourcesReturns(result1 []db.Resource, result2 error) {
	fake.visibleResourcesMutex.Lock()
	defer fake.visibleResourcesMutex.Unlock()
	fake.VisibleResourcesStub = nil
	fake.visibleResourcesReturns = struct {
		result1 []db.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) VisibleResourcesReturnsOnCall(i int, result1 []db.Resource, result2 error) {
	fake.visibleResourcesMutex.Lock()
	defer fake.visibleResourcesMutex.Unlock()
	fake.VisibleResourcesStub = nil
	if fake.visibleResourcesReturnsOnCall == nil {
		fake.visibleResourcesReturnsOnCall = make(map[int]struct {
			result1 []db.Resource
			result2 error
		})
	}
	fake.visibleResourcesReturnsOnCall[i] = struct {
		result1 []db.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.visibleResourcesMutex.RLock()
	defer fake.visibleResourcesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.ResourceFactory = new(FakeResourceFactory)