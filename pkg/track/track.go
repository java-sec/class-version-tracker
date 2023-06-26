package track

import (
	"context"
	"errors"
	"fmt"
	"github.com/java-sec/class-version-tracker/pkg/analyze"
	"github.com/java-sec/class-version-tracker/pkg/domain"
	"github.com/java-sec/class-version-tracker/pkg/search"
	"github.com/scagogogo/maven-crawler/pkg/crawler"
	"github.com/scagogogo/mvn-sdk/pkg/command"
	"github.com/scagogogo/mvn-sdk/pkg/finder"
	"github.com/scagogogo/mvn-sdk/pkg/local_repository"
)

// Track 追踪分析给定的类的变化
func Track(ctx context.Context, groupId, artifactId string, classFullName string) (*domain.Report, error) {

	// 参数检查
	if groupId == "" {
		return nil, errors.New("group id can not empty")
	} else if artifactId == "" {
		return nil, errors.New("artifact id can not empty")
	} else if classFullName == "" {
		return nil, errors.New("class full name can not empty")
	}

	// 获取都有哪些版本
	versions, err := getVersions(ctx, groupId, artifactId)
	if err != nil {
		return nil, err
	}

	// 把这些版本下载到本地
	err = downloadVersions(ctx, groupId, artifactId, versions)
	if err != nil {
		return nil, err
	}

	// 在每个版本中寻找目标类
	classSlice, err := search.SearchLocalRepository(ctx, groupId, artifactId, versions, classFullName)
	if err != nil {
		return nil, err
	}

	groups, err := analyze.Analyze(ctx, classSlice)
	if err != nil {
		return nil, err
	}
	return &domain.Report{
		GroupId:    groupId,
		ArtifactId: artifactId,
		Class:      classFullName,
		Groups:     groups,
	}, nil
}

// 下载给定版本到本地
// TODO 2023-6-20 11:28:46 修改为多线程下载
// TODO 2023-6-20 12:08:04 可能会因为镜像仓库的坑比而导致各种下载失败
func downloadVersions(ctx context.Context, groupId, artifactId string, versions []string) error {
	maven, err := finder.FindMaven()
	if err != nil {
		return err
	}

	localRepositoryDirectory, err := command.GetLocalRepositoryDirectory(maven)
	if err != nil {
		return err
	}

	for _, version := range versions {

		// 已经存在的话就不重复下载了
		jar, err := local_repository.FindJar(localRepositoryDirectory, groupId, artifactId, version)
		if err == nil && jar != "" {
			continue
		}

		// 跳过下载依赖
		stdout, err := command.ExecForStdout(maven, "dependency:get", "-DgroupId="+groupId, "-DartifactId="+artifactId, "-Dversion="+version, "-Dtransitive=false")
		if err != nil {
			return err
		}
		fmt.Println(stdout)
	}
	return nil
}

// 获取所有的版本
func getVersions(ctx context.Context, groupId, artifactId string) ([]string, error) {
	mavenCrawler, err := crawler.NewMavenCrawler()
	if err != nil {
		return nil, err
	}
	xml, err := mavenCrawler.GetMavenMetadataXml(ctx, groupId, artifactId)
	if err != nil {
		return nil, err
	}
	versions := make([]string, 0)
	for _, x := range xml.Versioning.Versions.Version {
		versions = append(versions, x.Text)
	}
	return versions, nil
}
