package dhtml

import (
	"github.com/gin-gonic/gin"
	"github.com/kooksee/ktask/internal/cnst"
	"github.com/kooksee/ktask/internal/kts"
	"github.com/kooksee/ktask/internal/utils"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/page"
	"net/http"
	"time"
)

func (t *config) Run(addr string) error {

	r := gin.New()
	r.Use(gin.Logger())
	r.POST("/task", func(ctx *gin.Context) {
		c := t.GetChrome()

		domContent, err := c.Page.DOMContentEventFired(ctx)
		defer domContent.Close()
		utils.MustNotError(err)

		// Enable events on the Page domain, it's often preferrable to create
		// event clients before enabling events so that we don't miss any.

		utils.MustNotError(c.Page.Enable(ctx))

		task := &kts.Task{}
		utils.MustNotError(task.DecodeFromReader(ctx.Request.Body))

		hi := &kts.HTMLItem{}
		utils.MustNotError(hi.Decode([]byte(task.Input)))

		url := hi.URL
		_, err = c.Page.Navigate(ctx, page.NewNavigateArgs(url))
		utils.MustNotError(err)

		time.Sleep(time.Second * time.Duration(t.SleepTime))

		// Wait until we have a DOMContentEventFired event.
		_, err = domContent.Recv()
		utils.MustNotError(err)

		// Fetch the document root node. We can pass nil here
		// since this method only takes optional arguments.
		doc, err := c.DOM.GetDocument(ctx, nil)
		utils.MustNotError(err)

		// Get the outer HTML for the page.
		result, err := c.DOM.GetOuterHTML(ctx, &dom.GetOuterHTMLArgs{NodeID: &doc.Root.NodeID})
		utils.MustNotError(err)

		hp := t.GetPattern(hi.PatternName)

		var ddt []map[string]interface{}
		utils.UnMashallHtml([]byte(result.OuterHTML), hp.List.Pattern, &ddt)

		utils.P(ddt)

		task.Output = result.OuterHTML
		task.Status = cnst.TaskStatus.Success
		task.Code = "ok"

		utils.P(task)
		ctx.JSON(http.StatusOK, task)
		return
	})

	return r.Run(addr)
}
