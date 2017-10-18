package app

import (
	"net/http"

	"github.com/PennTex/pet-whisperer/src/imageservice"
)

func init() {
	http.Handle("/", imageservice.New())
}
