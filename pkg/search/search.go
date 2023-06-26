package search

import (
	"context"
	"github.com/compression-algorithm-research-lab/go-unzip"
	"github.com/java-sec/class-version-tracker/pkg/domain"
	"github.com/scagogogo/mvn-sdk/pkg/command"
	"github.com/scagogogo/mvn-sdk/pkg/finder"
	"github.com/scagogogo/mvn-sdk/pkg/local_repository"
	"strings"
)

// FindClass 从jar包中寻找类
func FindClass(jarPath string, classFullName string) ([]byte, error) {
	classFullPath := strings.ReplaceAll(classFullName, ".", "/") + ".class"
	options := unzip.NewOptions().SetSourceZipFile(jarPath).SetWorkerNum(1)
	u := unzip.New(options)
	// TODO 2023-6-20 10:33:11 单个线程直接赋值，多个线程的时候怎么办
	var classBytes []byte
	err := u.SafeTraversal(func(file *unzip.File, options *unzip.Options) error {

		// 全路径类名要相同，后面再根据类名搜索
		if file.Name == classFullPath {
			bytes, _ := file.ReadBytes()
			classBytes = bytes
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return classBytes, nil
}

// SearchLocalRepository 从本地的Maven仓库中搜索给定的类
func SearchLocalRepository(ctx context.Context, groupId, artifactId string, versions []string, classFullName string) ([]*domain.Class, error) {
	maven, err := finder.FindMaven()
	if err != nil {
		return nil, err
	}

	directory, err := command.GetLocalRepositoryDirectory(maven)
	if err != nil {
		return nil, err
	}

	r := make([]*domain.Class, len(versions))
	for i, version := range versions {
		jar, err := local_repository.FindJar(directory, groupId, artifactId, version)
		if err != nil {
			r[i] = &domain.Class{
				Version: version,
			}
			continue
		}
		classBytes, err := FindClass(jar, classFullName)
		if err != nil {
			r[i] = &domain.Class{
				Version: version,
			}
			continue
		}
		r[i] = &domain.Class{
			Version: version,
			Bytes:   classBytes,
		}
	}
	return r, nil
}
