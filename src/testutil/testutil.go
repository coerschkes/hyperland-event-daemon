package testutil

import (
	"testing"
)

type MethodCall struct {
	MethodName string
	CallNumber int
	Params     any
}

type Mock struct {
	MethodCalls     map[string][]MethodCall
	MethodResponses map[string]any
}

func NewMock(methodResponses map[string]any) *Mock {
	return &Mock{MethodResponses: methodResponses, MethodCalls: make(map[string][]MethodCall)}
}

func (m *Mock) AddMethodCall(methodName string, params ...any) {
	m.MethodCalls[methodName] = append(m.MethodCalls[methodName],
		MethodCall{
			MethodName: methodName,
			CallNumber: len(m.MethodCalls[methodName]) + 1,
			Params:     params,
		},
	)
}

func AssertFail(t *testing.T, name string, input any, got any, want any, reason string) {
	t.Helper()
	t.Fatalf("\nTest: %s \n\n\tinput: %v \n\tgot: %v \n\twant: %v \n\n\treason: %s\n", name, input, got, want, reason)
}
