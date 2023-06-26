package domain

// Report 追踪结果报告
type Report struct {

	// 包的ga
	GroupId    string `json:"group_id"`
	ArtifactId string `json:"artifact_id"`

	Class string `json:"class"`

	// 版本分组信息
	Groups []*Group `json:"groups"`
}
