package kts

type Articles []Article

type Article struct {
	WebsiteSource string `json:"website_source"`
	WebsiteUrl    string `json:"website_url"`
	SourceUrl     string `json:"source_url"`
	SourceName    string `json:"source_name"`
	Author        string `json:"author"`
	Published     string `json:"published"`
	ClickNum      string `json:"click_num"`
	Title         string `json:"title"`
	CommentNum    string `json:"comment_num"`
	Keywords      string `json:"keywords"`
	Description   string `json:"description"`
	CrawlerTime   string `json:"crawler_time"`
}

func (t *Article) Encode() ([]byte, error) {
	return json.Marshal(t)
}
