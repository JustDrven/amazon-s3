package pkg

import "errors"

func IsPortAvaliable(port int) error {
	if port < 0 {
		return errors.New("The API port can't be below zero!")
	}

	if port > 65535 {
		return errors.New("The API port reached the limit!")
	}

	return nil
}
