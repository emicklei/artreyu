package assert

// Copyright 2015 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"reflect"
	"time"
)

// Operand represent a value
type Operand struct {
	// this reference is used to report a test failure
	a testingA
	// description of the value, typically a variable or field name
	label string
	// actual value of any type
	value interface{}
	// used to operate on two values
	operator RelationalOperator
}

// With returns a copy Operand that will use the RelationalOperator.
func (o Operand) With(r RelationalOperator) Operand {
	return Operand{o.a, o.label, o.value, r}
}

// Equals checks whether the value we have got is equal to the value we want.
func (o Operand) Equals(want interface{}) {
	if !o.operator.Apply(o.value, want) {
		if reflect.DeepEqual(o.value, want) {
			return
		}
		logCall(o.a.t, "Equals")
		Fatalf(o.a.t, "\ngot [%v] (%T) for \"%s\" but want [%v] (%T)",
			o.value, o.value,
			o.label,
			want, want)
	} else {
		Logf(o.a.t, "%s = %v", o.label, o.value)
	}
}

// Before checks whether the value we have got is before a moment.
func (o Operand) Before(moment time.Time) {
	left, ok := o.value.(time.Time)
	if !ok {
		logCall(o.a.t, "Before")
		Fatalf(o.a.t, "got [%v](%T) for \"%s\" but want a Time", o.value, o.value, o.label)
	}
	if left.Before(moment) {
		Logf(o.a.t, "%s is before %v", o.label, moment)
	} else {
		logCall(o.a.t, "Before")
		Fatalf(o.a.t, "got [%v] for \"%s\" but want it before [%v]", o.value, o.label, moment)
	}
}

// After checks whether the value we have got is after a moment.
func (o Operand) After(moment time.Time) {
	left, ok := o.value.(time.Time)
	if !ok {
		logCall(o.a.t, "After")
		Fatalf(o.a.t, "got [%v](%T) for \"%s\" but want a Time", o.value, o.value, o.label)
	}
	if left.After(moment) {
		Logf(o.a.t, "%s is after %v", o.label, moment)
	} else {
		logCall(o.a.t, "After")
		Fatalf(o.a.t, "got [%v] for \"%s\" but want it after [%v]", o.value, o.label, moment)
	}
}

// IsKindOf checks whether the values are of the same type
func (o Operand) IsKindOf(v interface{}) {
	leftType := reflect.TypeOf(o.value)
	rightType := reflect.TypeOf(v)
	if leftType != rightType {
		logCall(o.a.t, "IsKindOf")
		Fatalf(o.a.t, "got [%v] for \"%s\" but want [%v]", leftType, o.label, rightType)
	} else {
		Logf(o.a.t, "%s is kind of %v", o.label, rightType)
	}
}

// IsNil checks whether the value is nil
func (o Operand) IsNil() {
	if o.operator.Apply(o.value, nil) {
		Logf(o.a.t, "%s is nil", o.label)
		return
	} else {
		// from github.com/go-check/check
		switch v := reflect.ValueOf(o.value); v.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			if v.IsNil() {
				return
			}
		}
	}
	logCall(o.a.t, "IsNil")
	Fatalf(o.a.t, "got [%v] for \"%s\" but want [nil]", o.value, o.label)
}

// IsNotNil checks whether the value is nil
func (o Operand) IsNotNil() {
	if o.operator.Apply(o.value, nil) {
		logCall(o.a.t, "IsNotNil")
		Fatalf(o.a.t, "got unexpected [%v] for \"%s\"", o.value, o.label)
	} else {
		Logf(o.a.t, "%s is not nil", o.value)
	}
}

// IsTrue checks whether the value is true
func (o Operand) IsTrue() {
	if o.operator.Apply(o.value, false) { // i.e fail if !true
		logCall(o.a.t, "IsTrue")
		Fatalf(o.a.t, "got [%v] for \"%s\" but want [true]", o.value, o.label)
	} else {
		Logf(o.a.t, "%s is true", o.label)
	}
}

// IsFalse checks whether the value is false
func (o Operand) IsFalse() {
	if o.operator.Apply(o.value, true) { // i.e fail if !false
		logCall(o.a.t, "IsFalse")
		Fatalf(o.a.t, "got [%v] for \"%s\" but want [false]", o.value, o.label)
	} else {
		Logf(o.a.t, "%s is false", o.label)
	}
}

// Not creates a new Operand with a negated version of its comparator.
func (o Operand) Not() Operand {
	return Operand{o.a, o.label, o.value, not{o.operator}}
}

// Len checks that len(value) or value.Len() is equals to the given length.
// It operates on Array, Chan, Map, Slice, or String and objects that implement Len() int.
func (o Operand) Len(want int) {
	// panic catcher
	defer func() {
		if err := recover(); err != nil {
			// try calling Len
			rt := reflect.TypeOf(o.value)
			rf, ok := rt.MethodByName("Len")
			if !ok {
				logCall(o.a.t, "Len")
				Fatalf(o.a.t, "got [%v] for \"%s\" but it does not implement Len() int", o.value, o.label)
				return
			}
			rv := reflect.ValueOf(o.value)
			gotvs := rf.Func.Call([]reflect.Value{rv})
			got := int(gotvs[0].Int())
			if !o.operator.Apply(got, want) {
				logCall(o.a.t, "Len")
				Fatalf(o.a.t, "got [%v] for \"%s\" but want [%d]", got, o.label, want)
			}
		}
	}()
	rv := reflect.ValueOf(o.value)
	got := rv.Len()
	if !o.operator.Apply(got, want) {
		logCall(o.a.t, "Len")
		Fatalf(o.a.t, "got [%v] for \"%s\" but want [%d]", got, o.label, want)
	}
}

// IsEmpty checks that len(value) or value.Len() is zero.
// It operates on Array, Chan, Map, Slice, or String and objects that implement Len() int.
func (o Operand) IsEmpty() {
	// panic catcher
	defer func() {
		if err := recover(); err != nil {
			// try calling Len
			rt := reflect.TypeOf(o.value)
			rf, ok := rt.MethodByName("Len")
			if !ok {
				logCall(o.a.t, "IsEmpty")
				Fatalf(o.a.t, "got [%v] for \"%s\" but it does not implement Len() int", o.value, o.label)
				return
			}
			rv := reflect.ValueOf(o.value)
			gotvs := rf.Func.Call([]reflect.Value{rv})
			got := int(gotvs[0].Int())
			if !o.operator.Apply(got, 0) {
				logCall(o.a.t, "IsEmpty")
				Fatalf(o.a.t, "got [%v] for len(\"%s\") but want > 0", got, o.label)
			}
		}
	}()
	rv := reflect.ValueOf(o.value)
	got := rv.Len()
	if !o.operator.Apply(got, 0) {
		logCall(o.a.t, "IsEmpty")
		Fatalf(o.a.t, "got [%v] for len(\"%s\") but want > 0", got, o.label)
	}
}
