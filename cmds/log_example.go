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
	var topicName = "log"
	var handle = func(cmd *cobra.Command) *cobra.Command {
		return cmd
	}

	return handle(&cobra.Command{
		Use:   "log_example",
		Short: "http proxy example",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			r := gin.Default()

			r.GET("/topic", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, topicName)
				return
			})

			r.POST("/task", func(ctx *gin.Context) {
				task := &kts.Task{}
				utils.MustNotError(task.DecodeFromReader(ctx.Request.Body))
				utils.P(task)
				task.Status = cnst.TaskStatus.Success
				task.Output = "ok"
				ctx.JSON(http.StatusOK, task)
			})

			return r.Run(":8081")
		},
	})
}
