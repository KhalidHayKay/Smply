package model

import (
	"fmt"
	"snipfyi/config"
	"time"
)

type Url struct {
	Id       int64     `json:"id"`
	Original string    `json:"original"`
	Short    string    `json:"short"`
	Visited  int64     `json:"visited"`
	Created  time.Time `json:"created"`
}

func (u *Url) ShortToUrl() {
	u.Short = fmt.Sprintf("%s/r/%s", config.Env.AppUrl, u.Short)
}
