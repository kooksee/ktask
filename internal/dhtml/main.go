package dhtml

import (
	"github.com/gin-gonic/gin"
	"github.com/kooksee/ktask/internal/cnst"
	"github.com/kooksee/ktask/internal/kts"
	"github.com/kooksee/ktask/internal/utils"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/page"
	"math/big"
	"net/http"
	"time"
)

func (t *config) Run(addr string) error {

	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/task", func(ctx *gin.Context) {
		c := t.GetChrome()

		domContent, err := c.Page.DOMContentEventFired(ctx)
		defer domContent.Close()
		utils.MustNotError(err)

		// Enable events on the Page domain, it's often preferrable to create
		// event clients before enabling events so that we don't miss any.

		utils.MustNotError(c.Page.Enable(ctx))

		url, _ := ctx.GetQuery("url")
		_, err = c.Page.Navigate(ctx, page.NewNavigateArgs(url))
		utils.MustNotError(err)

		sleepTime, _ := ctx.GetQuery("sleep_time")
		if sleepTime == "" {
			time.Sleep(time.Second * time.Duration(t.SleepTime))
		} else {
			dt, ok := big.NewInt(0).SetString(sleepTime, 10)
			if ok {
				time.Sleep(time.Second * time.Duration(dt.Int64()))
			} else {
				ctx.String(http.StatusBadRequest, "sleep_time解析失败,正确为[sleep_time=2]")
				return
			}
		}

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

		t := &kts.Task{}
		t.Output = result.OuterHTML
		t.Status = cnst.TaskStatus.Success
		t.Code = "ok"
		ctx.String(http.StatusOK, result.OuterHTML)
		return
	})

	return r.Run(addr)
}
