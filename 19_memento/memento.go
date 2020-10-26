// Package memento 备忘录模式
// 下面这个例子采用原课程的例子，一个输入程序
// 如果输入 :list 则显示当前保存的内容
// 如果输入 :undo 则删除上一次的输入
// 如果输入其他的内容则追加保存
package memento

// InputText 用于保存数据
type InputText struct {
	content string
}

// Append 追加数据
func (in *InputText) Append(content string) {
	in.content += content
}

// GetText 获取数据
func (in *InputText) GetText() string {
	return in.content
}

// Snapshot 创建快照
func (in *InputText) Snapshot() *Snapshot {
	return &Snapshot{content: in.content}
}

// Restore 从快照中恢复
func (in *InputText) Restore(s *Snapshot) {
	in.content = s.GetText()
}

// Snapshot 快照，用于存储数据快照
// 对于快照来说，只能不能被外部（不同包）修改，只能获取数据，满足封装的特性
type Snapshot struct {
	content string
}

// GetText GetText
func (s *Snapshot) GetText() string {
	return s.content
}
