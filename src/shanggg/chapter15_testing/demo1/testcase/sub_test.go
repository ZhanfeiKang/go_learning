package cal

import (
	"testing"
) // 引入go的testing框架包

// 编写测试用例，去测试addUpper是否正确
func TestGetSub(t *testing.T) {
	// 调用
	res := getsub(10, 3)
	if res != 7 {
		t.Fatalf("getsub(10,3) 错误 期望值=%v 实际值=%v\n", 7, res)
	}

	// 如果正确，输出日志
	t.Logf("getsub(10,3) 执行正确...")
}
