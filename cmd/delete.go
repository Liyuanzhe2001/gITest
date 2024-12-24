/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	tmpl "gITest/template"
	"gITest/utils"
	"io/ioutil"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Send DELETE request to URL",
	Run: func(cmd *cobra.Command, args []string) {
		// 获取参数
		url, err := cmd.Flags().GetString("url")
		utils.Exception.ErrorHandle(err)

		output, err := cmd.Flags().GetString("output")
		utils.Exception.ErrorHandle(err)

		headers, err := cmd.Flags().GetStringSlice("headers")

		utils.Exception.ErrorHandle(err)
		if (len(headers) != 0 || output != "") && url == "" {
			utils.Exception.ErrorPrint("Unable to find URL")
		}
		if url == "" {
			cmd.HelpFunc()(cmd, []string{})
			return
		}

		// 参数检查没问题 发送DELETE请求
		out, err := utils.Net.DELETE(url, headers)
		utils.Exception.ErrorHandle(err)

		fmt.Println(out.Response)
		// 需要保存结果
		if output != "" {
			parse, _ := template.New("out").Parse(tmpl.GetTemplate)
			var b strings.Builder
			if err := parse.Execute(&b, out); err != nil {
				utils.Exception.AddPrefixErrorHandle("Error parsing template", err)
			}
			err = ioutil.WriteFile(output, []byte(b.String()), 0644)
			if err != nil {
				utils.Exception.AddPrefixErrorHandle("Error writing file", err)
			}
		}
	},
}

func init() {
	deleteCmd.Flags().StringP("url", "u", "", "url")
	deleteCmd.Flags().StringSliceP("headers", "h", []string{}, "请求头，K1:V1,K2:V2...的格式")
	deleteCmd.Flags().StringP("output", "o", "", "本次测试保存位置")
	// delete方法的参数在url中
	//deleteCmd.Flags().StringP("param", "p", "{}", "请求参数")
	deleteCmd.Flags().Bool("help", false, "显示帮助信息")
	deleteCmd.Flags().MarkHidden("help")
	rootCmd.AddCommand(deleteCmd)
}
