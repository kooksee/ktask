package dhtml

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kooksee/ktask/internal/cnst"
	"github.com/kooksee/ktask/internal/kts"
	"github.com/kooksee/ktask/internal/utils"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/rs/zerolog"
	"net/http"
	url2 "net/url"
	"time"
)

const topic = "dhtml"

func (t *config) Run(addr string) error {

	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/topic", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, topic)
		return
	})

	r.POST("/task", func(ctx *gin.Context) {
		task := &kts.Task{}
		utils.MustNotError(task.DecodeFromReader(ctx.Request.Body))
		utils.P(task.TaskID)

		hi := &kts.HTMLItem{}
		utils.MustNotError(hi.Decode([]byte(task.Input)))

		hp := t.GetPattern(hi.PatternName)

		if hi.Static {
			var dt []byte
			var err error
			var ddd []byte

			err = utils.Retry(10, func(l zerolog.Logger) error {
				dt, err = utils.HttpGet(hi.URL, hi.Header)
				if err != nil {
					l.Error().
						Err(err).
						Str("url", hi.URL).
						Str("mth", "utils.HttpGet").
						Interface("task", task).
						Interface("html_item", hi).
						Msg(err.Error())
					return err
				}

				if hi.IsList {
					var ddt []map[string]interface{}
					utils.UnMashallHtml(dt, hp.List.Pattern, &ddt)
					ddd, err = json.Marshal(ddt)
					utils.MustNotError(err)
					if len(ddt) == 0 {
						return fmt.Errorf("url: %s 解析为空", hi.URL)
					}

					for _, d1 := range ddt {
						url := utils.UrlCheck(hi.PUrl, utils.Str(d1["source_url"]))
						// 为空的过滤掉
						if url == "" {
							continue
						}

						if _, err := url2.Parse(url); err != nil {
							l.Error().Err(err).Str("url", url).Msg("url parse error")
							continue
						}

						d1["source_url"] = url

						nhi := &kts.HTMLItem{
							URL:         url,
							PUrl:        hi.PUrl,
							Static:      hp.Detail.Static,
							PatternName: hi.PatternName,
							IsList:      false,
							Extra:       d1,
						}

						ttk := (&kts.Task{}).Mock()
						ttk.TopicName = cnst.Consumer.DHtml
						ttk.ServiceName = "dhtml"
						ttk.Input = string(nhi.Encode())

						utils.P(ttk)
						utils.MustNotError(utils.Retry(10, func(l zerolog.Logger) error {
							return t.TaskPost(ttk.Encode())
						}))

					}
				} else {
					var ddt map[string]interface{}
					utils.UnMashallHtml(dt, hp.Detail.Pattern, &ddt)
					ddd, err = json.Marshal(ddt)
					utils.MustNotError(err)
				}

				task.Output = string(ddd)
				task.Status = cnst.TaskStatus.Success
				ctx.JSON(http.StatusOK, task)

				return nil
			})
			if err != nil {
				task.Status = cnst.TaskStatus.Failed
				task.Output = err.Error()
				ctx.JSON(http.StatusOK, task)
			}
			return
		}

		var err error
		var ddd []byte
		err = utils.Retry(10, func(l zerolog.Logger) error {
			c := t.GetChrome()

			domContent, err := c.Page.DOMContentEventFired(ctx)
			defer domContent.Close()
			utils.MustNotError(err)

			// Enable events on the Page domain, it's often preferrable to create
			// event clients before enabling events so that we don't miss any.

			utils.MustNotError(c.Page.Enable(ctx))

			_, err = c.Page.Navigate(ctx, page.NewNavigateArgs(hi.URL))
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

			dt := []byte(result.OuterHTML)
			hp := t.GetPattern(hi.PatternName)
			if hi.IsList {
				var ddt []map[string]interface{}
				utils.UnMashallHtml(dt, hp.List.Pattern, &ddt)
				ddd, err = json.Marshal(ddt)
				utils.MustNotError(err)
				if len(ddt) == 0 {
					return fmt.Errorf("url: %s 解析为空", hi.URL)
				}

				for _, d1 := range ddt {
					url := utils.UrlCheck(hi.PUrl, utils.Str(d1["source_url"]))
					// 为空的过滤掉
					if url == "" {
						continue
					}

					if _, err := url2.Parse(url); err != nil {
						l.Error().Err(err).Str("url", url).Msg("url parse error")
						continue
					}

					d1["source_url"] = url

					nhi := &kts.HTMLItem{
						URL:         url,
						PUrl:        hi.PUrl,
						Static:      hp.Detail.Static,
						PatternName: hi.PatternName,
						IsList:      false,
						Extra:       d1,
					}

					ttk := (&kts.Task{}).Mock()
					ttk.TopicName = cnst.Consumer.DHtml
					ttk.ServiceName = "dhtml"
					ttk.Input = string(nhi.Encode())

					utils.MustNotError(utils.Retry(10, func(l zerolog.Logger) error {
						return t.TaskPost(ttk.Encode())
					}))

				}
			} else {
				var ddt map[string]interface{}
				utils.UnMashallHtml(dt, hp.Detail.Pattern, &ddt)
				ddd, err = json.Marshal(ddt)
				utils.MustNotError(err)
			}

			return nil
		})

		if err != nil {
			task.Status = cnst.TaskStatus.Failed
			task.Output = err.Error()
			ctx.JSON(http.StatusOK, task)
			return
		}

		task.Output = string(ddd)
		task.Status = cnst.TaskStatus.Success
		task.Code = "ok"
		ctx.JSON(http.StatusOK, task)
		return

		return
	})

	return r.Run(addr)
}
