// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/davidkhala/protoutil"
)

type Policy struct {
	EvaluateSignedDataStub        func([]*protoutil.SignedData) error
	evaluateSignedDataMutex       sync.RWMutex
	evaluateSignedDataArgsForCall []struct {
		arg1 []*protoutil.SignedData
	}
	evaluateSignedDataReturns struct {
		result1 error
	}
	evaluateSignedDataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Policy) EvaluateSignedData(arg1 []*protoutil.SignedData) error {
	var arg1Copy []*protoutil.SignedData
	if arg1 != nil {
		arg1Copy = make([]*protoutil.SignedData, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.evaluateSignedDataMutex.Lock()
	ret, specificReturn := fake.evaluateSignedDataReturnsOnCall[len(fake.evaluateSignedDataArgsForCall)]
	fake.evaluateSignedDataArgsForCall = append(fake.evaluateSignedDataArgsForCall, struct {
		arg1 []*protoutil.SignedData
	}{arg1Copy})
	stub := fake.EvaluateSignedDataStub
	fakeReturns := fake.evaluateSignedDataReturns
	fake.recordInvocation("EvaluateSignedData", []interface{}{arg1Copy})
	fake.evaluateSignedDataMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Policy) EvaluateSignedDataCallCount() int {
	fake.evaluateSignedDataMutex.RLock()
	defer fake.evaluateSignedDataMutex.RUnlock()
	return len(fake.evaluateSignedDataArgsForCall)
}

func (fake *Policy) EvaluateSignedDataCalls(stub func([]*protoutil.SignedData) error) {
	fake.evaluateSignedDataMutex.Lock()
	defer fake.evaluateSignedDataMutex.Unlock()
	fake.EvaluateSignedDataStub = stub
}

func (fake *Policy) EvaluateSignedDataArgsForCall(i int) []*protoutil.SignedData {
	fake.evaluateSignedDataMutex.RLock()
	defer fake.evaluateSignedDataMutex.RUnlock()
	argsForCall := fake.evaluateSignedDataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Policy) EvaluateSignedDataReturns(result1 error) {
	fake.evaluateSignedDataMutex.Lock()
	defer fake.evaluateSignedDataMutex.Unlock()
	fake.EvaluateSignedDataStub = nil
	fake.evaluateSignedDataReturns = struct {
		result1 error
	}{result1}
}

func (fake *Policy) EvaluateSignedDataReturnsOnCall(i int, result1 error) {
	fake.evaluateSignedDataMutex.Lock()
	defer fake.evaluateSignedDataMutex.Unlock()
	fake.EvaluateSignedDataStub = nil
	if fake.evaluateSignedDataReturnsOnCall == nil {
		fake.evaluateSignedDataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.evaluateSignedDataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Policy) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.evaluateSignedDataMutex.RLock()
	defer fake.evaluateSignedDataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Policy) recordInvocation(key string, args []interface{}) {
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
