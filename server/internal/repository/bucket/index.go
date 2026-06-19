package bucket

import (
	"encoding/base64"
	"errors"
	"os"
	"strings"
)

func computeObjectPath(bucket string, key string, name string) string {
	// TODO: Improve this function, this is big a piece of dog shit
	path := bucket + "/" + key + "/" + name

	return strings.ReplaceAll(path, "//", "/")
}

func GetListObjects(bucket string, key string, path string) (*ListBucketResult, error) {

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, errors.New("failed to open bucket folder")
	}

	contents := make([]ObjectContent, len(entries))

	for i := range entries {
		entry := entries[i]

		name := entry.Name()
		data := []byte(name + entry.Type().String())
		lastModified := "unknown"
		objectType := "FILE"

		info, infoErr := entry.Info()
		size := int64(-1)
		objectPath := computeObjectPath(bucket, key, name)

		if infoErr == nil {
			if info.IsDir() {
				objectType = "FOLDER"
			}
			lastModified = info.ModTime().String()
			size = info.Size()
		}

		contents[i] = ObjectContent{
			Key:          name,
			ObjectPath:   objectPath,
			LastModified: lastModified,
			Size:         size,
			Type:         objectType,
			ETag:         base64.StdEncoding.EncodeToString(data),
		}
	}
	toReturn := &ListBucketResult{
		Name:     key,
		Prefix:   bucket,
		FullPath: path,

		KeyCount: len(contents),
		Contents: contents,
	}

	return toReturn, nil
}
