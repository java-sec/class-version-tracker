package cmd

import (
	"encoding/json"
	"github.com/java-sec/class-version-tracker/pkg/domain"
	"github.com/java-sec/class-version-tracker/pkg/track"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
)

var (

	// 要追踪的组ID
	groupId string

	// 要追踪的文档ID
	artifactId string

	// 要追踪的类
	class string

	// 结果的输出位置
	output string
)

func init() {

	trackCmd.Flags().StringVarP(&groupId, "groupId", "g", "", "Group ID for track, example: org.apache.dubbo")
	trackCmd.Flags().StringVarP(&artifactId, "artifactId", "a", "", "Artifact ID for track, example: dubbo")
	trackCmd.Flags().StringVarP(&class, "class", "c", "", "Class for track, example: org.springframework.web.servlet.mvc.method.RequestMappingInfo")
	trackCmd.Flags().StringVarP(&output, "output", "o", "", "result output directory, default ./output")

	rootCmd.AddCommand(trackCmd)
}

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "track for class",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		if output == "" {
			output = "./output"
		}

		report, err := track.Track(cmd.Context(), groupId, artifactId, class)
		if err != nil {
			return err
		}
		return saveReport(output, report)
	},
}

// 把报告保存到给定的文件夹
func saveReport(directory string, report *domain.Report) error {

	// 确保输出目录存在
	classDumpDirectory := filepath.Join(directory, strings.ReplaceAll(report.GroupId, ".", "."), report.ArtifactId, report.Class)
	err := os.MkdirAll(classDumpDirectory, os.ModePerm)
	if err != nil {
		return err
	}

	// 把报告导出一份JSON格式的
	marshal, err := json.Marshal(report)
	if err != nil {
		return err
	}
	reportPath := filepath.Join(classDumpDirectory, "report.json")
	err = os.WriteFile(reportPath, marshal, 0644)
	if err != nil {
		return err
	}

	// 把每一组的类导出
	for _, group := range report.Groups {
		classFileName := filepath.Join(classDumpDirectory, group.Versions[0]+".class")
		err := os.WriteFile(classFileName, group.Bytes, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
