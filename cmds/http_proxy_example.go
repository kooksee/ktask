package cmds

import (
	"github.com/gin-gonic/gin"
	"github.com/kooksee/ktask/internal/cnst"
	"github.com/kooksee/ktask/internal/kts"
	"github.com/kooksee/ktask/internal/utils"
	"github.com/spf13/cobra"
	"net/http"
)

func HttpProxyExampleCmd() *cobra.Command {
	var logHttpCmd = func(cmd *cobra.Command) *cobra.Command {
		return cmd
	}
	return logHttpCmd(&cobra.Command{
		Use:   "http_proxy_example",
		Short: "http proxy example",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			r := gin.Default()
			r.POST("/task", func(ctx *gin.Context) {
				task := &kts.Task{}
				utils.MustNotError(task.DecodeFromReader(ctx.Request.Body))
				utils.P(task)
				ctx.JSON(http.StatusOK, &kts.TaskResult{
					TaskID: task.TaskID,
					Status: cnst.TaskStatus.Success,
					Output: "ok",
					Code:   "ok",
				})
			})

			return r.Run(":8081")
		},
	})
}
