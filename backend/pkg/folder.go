package pkg

import "strings"

func FixFolderPath(path *string) {

	if !strings.HasSuffix(*path, "/") {
		*path += "/"
	}

}
