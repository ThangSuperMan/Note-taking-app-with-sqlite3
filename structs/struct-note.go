package structs

import (
	"time"
)

type Note struct {
	title     string
	content   string
	updateAt  time.Time
	id_author string
}
