package model

// 定义一个结构体
type student struct {
	Name  string
	score float64 // 小写s
}

// 因为student结构体首字母是小写，因此是只能在model中使用
// 我们通过工厂模式来解决

func NewStudent(n string, s float64) *student {
	// 返回一个指针
	return &student{
		Name:  n,
		score: s,
	}
}

// 如果score字段首字母小写，则在其他包不可以直接访问，我们可以提供一个方法来解决
func (s *student) GetScore() float64 {
	return s.score // ok 可以访问，在本包
}
