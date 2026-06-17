package bucket

type ListBucketResult struct {
	Name     string          `xml:"Bucket"`
	Prefix   string          `xml:"Prefix"`
	KeyCount int32           `xml:"KeyCount"`
	Contents []ObjectContent `xml:"Contents"`
}

type ObjectContent struct {
	Key          string `xml:"Key"`
	LastModified string `xml:"LastModified"`
	ETag         string `xml:"ETag"`
	Size         int32  `xml:"Size"`
}
