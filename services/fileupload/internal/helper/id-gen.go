package helper

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/teris-io/shortid"
)

func GenerateID() string {
	sid, err := shortid.Generate()
	if err != nil {
		return fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Int63())
	}
	return sid
}
