package bucket

type ListBucketResult struct {
	Name     string          `xml:"Bucket"`
	Prefix   string          `xml:"Prefix"`
	FullPath string          `xml:"FullPath"`
	KeyCount int             `xml:"KeyCount"`
	Contents []ObjectContent `xml:"Contents"`
}

type ObjectContent struct {
	Key          string `xml:"Key"`
	ObjectPath   string `xml:"ObjectPath"`
	LastModified string `xml:"LastModified"`
	ETag         string `xml:"ETag"`
	Type         string `xml:"ObjectType"`
	Size         int64  `xml:"Size"`
}
