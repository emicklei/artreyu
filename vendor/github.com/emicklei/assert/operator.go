package assert

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// RelationalOperator specifies the function that can operate on two values
type RelationalOperator interface {
	// Return the result of applying the two values left and right
	Apply(left, right interface{}) bool
}

// equals compares two values using ==
type equals struct{}

// Apply returns the result of comparing left and right using ==
func (e equals) Apply(left, right interface{}) bool {
	return left == right
}

// Not is to negate the result of a RelationalOperator
type not struct {
	r RelationalOperator
}

// Apply returns the opposite boolean result of applying left and right
func (n not) Apply(left, right interface{}) bool {
	return !n.r.Apply(left, right)
}
