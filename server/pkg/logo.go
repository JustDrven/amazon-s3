package pkg

import (
	"fmt"
)

func PrintSpace(count int) {

	for i := 0; i < count; i++ {
		fmt.Println()
	}

}

func PrintLogo() {
	count := 4

	PrintSpace(count)

	fmt.Println("      ▄▄▄▄   ▄▄▄▄  ▄▄▄  ▄▄▄▄  ▄▄▄▄▄▄▄        ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄")
	fmt.Println("    ▄██▀▀██▄ ▀███  ███  ███▀ █████▀▀▀       █████▀▀▀ ▀▀▀▀████ ")
	fmt.Println("    ███  ███  ███  ███  ███   ▀████▄         ▀████▄    ▄▄██▀  ")
	fmt.Println("    ███▀▀███  ███▄▄███▄▄███     ▀████          ▀████     ███▄ ")
	fmt.Println("    ███  ███   ▀████▀████▀   ███████▀       ███████▀ ███████▀ ")

	PrintSpace(count)

}
