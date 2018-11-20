package kts

type Images []ImageMetadata

type PicInfo struct {
	AssetFormat string    `json:"asset_format"`
	Brand       string    `json:"brand"`
	BrandId     int64     `json:"brandId"`
	Cameraman   string    `json:"cameraman"`
	Caption     string    `json:"caption"`
	Copyright   string    `json:"copyright"`
	FileType    string    `json:"file_type"`
	Height      int64     `json:"height"`
	ID          int64     `json:"id"`
	ImgDate     string    `json:"img_date"`
	License     string    `json:"license"`
	Oss176      string    `json:"oss176"`
	Restrict    string    `json:"restrict"`
	Size        string    `json:"size"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Width       int64     `json:"width"`
}

type ImageMetadata struct {
	Name            string `json:"name"`
	ID              string `json:"ID"`
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	FileSize        int    `json:"file_size"`
	License         string `json:"license"`
	Format          string `json:"format"`
	FileType        string `json:"file_type"`
	Cameraman       string `json:"cameraman"`
	Copyright       string `json:"copyright"`
	URL             string `json:"url"`
	URL176          string `json:"url176"`
	Keywords        string `json:"keywords"`
	Size            string `json:"size"`
	ImgDate         string `json:"img_date"`
	Price           string `json:"price"`
	Title           string `json:"title"`
	ResID           string `json:"res_id"`
	Restrict        string `json:"restrict"`
	DHash           string `json:"dhash"`
	URLSource       string `json:"url_source"`
	ImageSource     string `json:"image_source"`
	ImageAuthor     string `json:"image_author"`
	ImageOrg        string `json:"image_org"`
	Brand           string `json:"brand"`
	DciID           string `json:"dci_id"`
	DciName         string `json:"dci_name"`
	DciAuthor       string `json:"dci_author"`
	DciAuthorized   string `json:"dci_authorized"`
	DciType         string `json:"dci_type"`
	DciPublishTime  string `json:"dci_publish_time"`
	DciCreateTime   string `json:"dci_create_time"`
	DciRegisterTime string `json:"dci_register_time"`
}

