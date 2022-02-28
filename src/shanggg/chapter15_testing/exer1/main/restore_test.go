package main

import "testing"

func TestRestore(t *testing.T) {
	// 先创建一个monster
	var monster Monster

	res := monster.Restore()
	if !res {
		t.Fatalf("monster.Store() 错误，希望为=%v 实际为%v", true, res)
	}
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.Restore() 错误，希望为=%v 实际为%v", "红孩儿", res)
	}
	t.Logf("monster.Restore() 测试成功..")
}
