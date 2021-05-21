package core

import (
	"log"

	"github.com/google/uuid"
)

type Utility struct {
}

func UseUtil() *Utility {
	return &Utility{}
}

func (u *Utility) UUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatalln(err)
	}
	return id.String()
}
