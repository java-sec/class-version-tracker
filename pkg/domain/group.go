package domain

// Group 类的字节码分组信息
type Group struct {

	// 这一组的md5
	MD5 string `json:"md5"`

	// 都有哪些版本属于此版本
	Versions []string `json:"versions"`

	// 类的字节码
	Bytes []byte `json:"-"`
}

// IsEmpty 这一组是个空组
func (x *Group) IsEmpty() bool {
	return x.MD5 == "d41d8cd98f00b204e9800998ecf8427e"
}
