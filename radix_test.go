package radix

import (
	"fmt"
	"testing"
)

func printit(r *Radix, level int) {
	for i := 0; i < level; i++ {
		fmt.Print("\t")
	}
	fmt.Printf("'%v'  value: %v    parent %p\n", r.key, r.Value, r.parent)
	for _, child := range r.children {
		printit(child, level+1)
	}
}

func radixsimpletree() *Radix {
	r := New()
	r.Insert("a", nil)
	r.Insert("b", nil)
	r.Insert("c", nil)
	return r
}

func radixsimpletree2() *Radix {
	r := New()
	r.Insert("aa", nil)
	r.Insert("bb", nil)
	r.Insert("cc", nil)
	r.Insert("cd", nil)
	return r
}

func radixtree() *Radix {
	r := New()
	r.Insert("test", nil)
	r.Insert("slow", nil)
	r.Insert("water", nil)
	r.Insert("watsol", nil)
	r.Insert("tester", nil)
	r.Insert("testering", nil)
	r.Insert("rewater", nil)
	r.Insert("waterrat", nil)
	return r
}

// None, of the childeren must have a prefix incommon with r.key
func validate(r *Radix) bool {
	for _, child := range r.children {
		_, i := longestCommonPrefix(r.key, child.key)
		if i != 0 {
			return false
		}
		validate(child)
	}
	return true
}

func TestPrint(t *testing.T) {
	r := radixtree()
	printit(r, 0)
}

func TestNextSimple(t *testing.T) {
	r := radixsimpletree()
	printit(r, 0)
	println("a", r.Find("a").Next().key)
	println("b", r.Find("b").Next().key)
	println("c", r.Find("c").Next().key)
	println("c", r.Find("c").Next().Next().key)
}

func TestNextSimple2(t *testing.T) {
	r := radixsimpletree2()
	printit(r, 0)
	println("aa", r.Find("aa").Next().key)
	println("bb", r.Find("bb").Next().key)
	println("cc", r.Find("cc").Next().key)
	println("cd", r.Find("cd").Next().key)
}

func TestInsert(t *testing.T) {
	r := New()
	if !validate(r) {
		t.Log("Tree does not validate")
		t.Fail()
	}
	if r.Len() != 0 {
		t.Log("Len should be 0", r.Len())
	}
	r.Insert("test", nil)
	r.Insert("slow", nil)
	r.Insert("water", nil)
	r.Insert("tester", nil)
	r.Insert("testering", nil)
	r.Insert("rewater", nil)
	r.Insert("waterrat", nil)
	if !validate(r) {
		t.Log("Tree does not validate")
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	r := New()
	r.Insert("test", "aa")
	r.Insert("slow", "bb")

	if k := r.Remove("slow").Value; k != "bb" {
		t.Log("should be bb", k)
		t.Fail()
	}

	if r.Remove("slow") != nil {
		t.Log("should be nil")
		t.Fail()
	}
}

func ExampleFind() {
	r := New()
	r.Insert("tester", nil)
	r.Insert("testering", nil)
	r.Insert("te", nil)
	r.Insert("testeringandmore", nil)
	f := r.Find("tester")
	iter(f, f.Prefix("tester"))
	// Output:
	// prefix tester
	// prefix testering
	// prefix testeringandmore
}

func BenchmarkFind(b *testing.B) {
	b.StopTimer()
	r := radixtree()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		_ = r.Find("tester")
	}
}

func iter(r *Radix, prefix string) {
	fmt.Printf("prefix %s\n", prefix+r.Key())
	for _, child := range r.Children() {
		iter(child, prefix+r.Key())
	}
}
