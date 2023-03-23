/*
 * @Author: kingford
 * @Date: 2023-03-20 16:00:39
 * @LastEditTime: 2023-03-23 10:15:38
 */
package cmd

import (
	"errors"
	"fmt"
	"go-gin-template/cmd/server"
	"go-gin-template/common/global"

	"go-gin-template/pkg"

	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "go-gin-template",
	Short:        "go-gin-template",
	SilenceUsage: true,
	Long:         `go-gin-template`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`go-gin-template `+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	usageStr1 := `也可以参考 http://127.0.0.1/swagger 的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	rootCmd.AddCommand(server.GinCmd)
	// rootCmd.AddCommand(StartServerCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
