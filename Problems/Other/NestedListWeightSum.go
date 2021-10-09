package Problems

type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool { return false }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int { return 0 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger { return nil }

func depthSumInverse(nl []*NestedInteger) int {
	return f2(nl, f1(nl, 1))
}
func f2(l []*NestedInteger, d int) int {
	r := 0
	for i := range l {
		if !l[i].IsInteger() {
			r += f2(l[i].GetList(), d-1)
		} else {
			r += (l[i].GetInteger() * d)
		}
	}
	return r
}
func f1(l []*NestedInteger, d int) int {
	r := d
	for i := range l {
		if !l[i].IsInteger() {
			v := f1(l[i].GetList(), d+1)
			if v > r {
				r = v
			}
		}
	}
	return r
}
