package main

import (
	"fmt"
	"regexp"

	"products/router"
)

func main() {
	router.StartRouter()

	fmt.Println(regexp.QuoteMeta(`SELECT id, email, password from users WHERE email = $1`))
}
