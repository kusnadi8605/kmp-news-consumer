package datastruct

//NewsJSON struct for news
type NewsJSON struct {
	ID      int32  `json:"id"`
	Author  string `json:"author"`
	Body    string `json:"body"`
	Created string `json:"created"`
}

//SaveNewsJSON ..
type SaveNewsJSON struct {
	ID      int64  `json:"id"`
	Created string `json:"created"`
}
