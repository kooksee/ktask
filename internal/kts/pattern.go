package kts

type HtmlPattern struct {
	Detail struct {
		Pattern string `json:"pattern"`
		Static  bool   `json:"static"`
	} `json:"detail"`
	IsList  bool   `json:"is_list"`
	Crontab string `json:"crontab"`
	List    struct {
		Pattern string   `json:"pattern"`
		Static  bool     `json:"static"`
		URL     string   `json:"url"`
		URLS    []string `json:"urls"`
	} `json:"list"`
	Type        string `json:"type"`
	WebsiteName string `json:"website_name"`
	WebsiteURL  string `json:"website_url"`
}

func (t *HtmlPattern) Decode(data []byte) error {
	return json.Unmarshal(data, t)
}
