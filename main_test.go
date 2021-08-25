package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	l := &TreeNode{Val: 1}
	r := &TreeNode{Val: 7}
	in := &TreeNode{Val: 5, Left: l, Right: r}
	out := increasingBST(in)
	want := 1
	if out.Val != want {
		t.Errorf("got %v, want %v", out, want)
	}
}
