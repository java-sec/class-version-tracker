package analyze

import (
	"context"
	"github.com/java-sec/class-version-tracker/pkg/domain"
	"github.com/java-sec/class-version-tracker/pkg/utils"
	"github.com/scagogogo/versions"
	"sort"
)

// Analyze 分析给定的字节码，对其进行分组
func Analyze(ctx context.Context, classSlice []*domain.Class) ([]*domain.Group, error) {
	groupMap := make(map[string]*domain.Group, 0)
	for _, class := range classSlice {
		md5, err := utils.MD5(class.Bytes)
		if err != nil {
			return nil, err
		}
		group, exists := groupMap[md5]
		if !exists {
			group = &domain.Group{
				MD5:   md5,
				Bytes: class.Bytes,
			}
			groupMap[md5] = group
		}
		group.Versions = append(group.Versions, class.Version)
	}

	groupSlice := make([]*domain.Group, 0)
	for _, group := range groupMap {
		// 每个分组内的版本排序
		group.Versions = versions.SortVersionStringSlice(group.Versions)
		groupSlice = append(groupSlice, group)
	}

	// 返回的分组排序
	sort.Slice(groupSlice, func(i, j int) bool {
		v1 := versions.NewVersion(groupSlice[i].Versions[0])
		v2 := versions.NewVersion(groupSlice[j].Versions[0])
		return v1.CompareTo(v2) < 0
	})

	return groupSlice, nil
}
