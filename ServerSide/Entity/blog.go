package Entity

import "time"

type Blog struct {
	blogId      string
	title       string
	content     string
	authorId    string
	createdDate time.Time
	updatedDate time.Time
}
