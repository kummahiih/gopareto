//
// @copyright: 2022 by Pauli Rikula <pauli.rikula@gmail.com>
// @license: MIT <http://www.opensource.org/licenses/mit-license.php>
//

package gopareto

// This file contains the interface definitions of the gopareto module for
// pareto front searching.
// Done porting quite directly the https://github.com/kummahiih/pypareto into golang
// and probably needs some refactoring again and again.
// For usage skip this file and see the usage from the unit tests instead.
// It will be probably easier to use than understand

// Item is the type wrapper
type Item interface {
	IsGreaterThan(b Item) (bool, error)
	// IsEmpty makes it clear did you cast properly the value
	// or was the value empty. This just prevents me from confusing myself
	IsEmpty() bool
	// Empty return an empty Item of the same subtype
	Empty() Item
	// String returns the string representation of the item. Handy for unit tests
	String() string
}

type MaxMin int

const (
	SKIP = iota
	MAX
	MIN
)

type Domination int

const (
	ERROR = iota
	GREATER
	EQUAL
	LESS
)

type MaxMinList interface {
	AsArray() []MaxMin
	GetDim() int
	IsNoneGood() bool
}

type WithTargets interface {
	GetTargets() MaxMinList
	SetTargets(t MaxMinList)
}

type Comparer func(a Item, b Item) Domination

type WithComparer interface {
	GetComparer() Comparer
	SetComparer(v Comparer)
}

type Comparable interface {
	Compare(a Item, b Item) Domination
}

type Cmp interface {
	Comparable
	IsPareto() bool
	IsGroup() bool
	GetGroup(a Item) int
	AndThen(c Cmp) Cmp
	AsChain() ComparisonChain
}

type Comparison interface {
	Cmp
	WithComparer
	WithTargets
}

type ComparisonChain interface {
	Comparable
	AsArray() []Comparison
	AndThen(c Cmp) Cmp
	SplitByPareto(value []Item) [][]Item
}

type GroupEmpy interface {
	Cmp
	WithTargets
}

func ByEmpty(a Item, b Item) (Domination, error) {
	if a.IsEmpty() && !b.IsEmpty() {
		return GREATER, nil
	}
	if b.IsEmpty() && !a.IsEmpty() {
		return LESS, nil
	}
	return EQUAL, nil
}

func ByValue(a Item, b Item) (Domination, error) {
	if c, err := a.IsGreaterThan(b); err != nil {
		if c {
			return GREATER, nil
		}
	} else {
		return ERROR, err
	}

	if c, err := b.IsGreaterThan(a); err != nil {
		if c {
			return LESS, nil
		}
	} else {
		return ERROR, err
	}

	return EQUAL, nil
}

//type ByValue
//type CmpToTarget
//type Dominates
