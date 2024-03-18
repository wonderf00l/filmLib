package auth

import "time"

type Session struct {
	Key     string
	UserID  int
	Expires time.Time
}
