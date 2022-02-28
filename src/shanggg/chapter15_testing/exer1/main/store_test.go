package main

import "testing"

func TestStore(t *testing.T) {
	// 先创建一个monster
	monster := &Monster{
		Name:  "红孩儿",
		Age:   10,
		Skill: "三昧真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误，希望为=%v 实际为%v", true, res)
	}
	t.Logf("monster.Store() 测试成功..")
}
