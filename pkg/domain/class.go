package domain

// Class 表示一个类的字节码
type Class struct {

	// 这个类是属于哪个版本的
	Version string `json:"version"`

	// 类的字节码
	Bytes []byte `json:"bytes"`
}

// IsValid 此类是否可用，需要有对应的字节码
func (x *Class) IsValid() bool {
	return len(x.Bytes) != 0
}
