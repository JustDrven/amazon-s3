package object

type GetObjectResult struct {
	Prefix   string `xml:"Prefix"`
	Name     string `xml:"Name"`
	FilePath string `xml:"FilePath"`
	FileName string `xml:"FileName"`
	Size     int    `xml:"Size"`
	Content  string `xml:"Content"`
}
