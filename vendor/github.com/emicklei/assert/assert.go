package assert

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// testingT defines the api that is used from testing.T
// this exists for testing Assert using a mock.
type testingT interface {
	Fatalf(string, ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

// Fatalf calls Fatalf on a test instance t.
// You can inject your own implementation of assert.testingT
var Fatalf = func(t testingT, format string, args ...interface{}) {
	t.Fatalf("%s", Scolorf(AssertFatalColorSyntaxCode, format, args...))
}

// Log calls Log on a test instance t.
// You can inject your own implementation of assert.testingT
var Log = func(t testingT, args ...interface{}) {
	t.Log(args...)
}

// Logf calls Log on a test instance t.
// You can inject your own implementation of assert.testingT
var Logf = func(t testingT, format string, args ...interface{}) {
	t.Logf("%s", Scolorf(AssertSuccessColorSyntaxCode, format, args...))
}

// testingA decorates a *testing.T to create an Operand using That(..) and do error logging
type testingA struct {
	t testingT
}

// That creates an Operand on the value we have got and describes the variable that is being testing.
func (a testingA) That(label string, got interface{}) Operand {
	return Operand{a, label, got, equals{}}
}

// That creates an Operand on the value we have got and describes the variable that is being testing.
func That(t testingT, label string, got interface{}) Operand {
	return Operand{testingA{t}, label, got, equals{}}
}

// Assert creates an Operand on a value that needs to be checked.
func Assert(t testingT, label string, value interface{}) Operand {
	return testingA{t}.That(label, value)
}
