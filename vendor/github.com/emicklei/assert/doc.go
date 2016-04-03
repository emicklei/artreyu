/*
Package assert, for writing checks in unit tests

This package provides functions to reduce the amount of code needed to write simple assertions.
It implements the best practice pattern where the output of a failure explains what the check "got" and what it wanted.
The assert functions are defined such that writing requires less code but still are easy to understand.
It works by decorating the standard testing.T type in your test and report (Fatal) the offending asserting call if a check fails.

Example

	import (
		"testing.T"
		"github.com/emicklei/assert"
	)

	func TestShoeFits(t *testing.T) {
		shoeSize := 42
		assert.That(t,"shoeSize",shoeSize).Equals(46)
	}

which will report

	got [42] (int) for "shoeSize" but want [46] (int)


Examples: (using the dot import)

	Assert(t,"err",err).IsNil()
	Assert(t,"isOffline",isOffline).IsTrue()
	Assert(t,"country",country).Equals("NL")
	Assert(t,"job",job).IsKindOf(new(Job))
	Assert(t,"names", []string{}).Len(0)

	// you can negate a check
	Assert(t,"isOnline",isOnline).Not().IsTrue()

You can create and use your own checks by implementing the RelationalOperator.

	type caseInsensitiveStringEquals struct{}

	func (c caseInsensitiveStringEquals) Apply(left, right interface{}) bool {
		s_left, ok := left.(string)
		if !ok {
			return false
		}
		s_right, ok := right.(string)
		if !ok {
			return false
		}
		return strings.EqualFold(s_left, s_right)
	}

	func TestCompareUsing(t *testing.T) {
		Assert(t, "insensitive", "ABC").With(caseInsensitiveStringEquals{}).Equals("abc")
	}

*/
package assert
