package public

// User 用户
type User struct {
	UserID     uint64 `json:"user_id,string" db:"user_id"`
	UserName   string `json:"username" db:"username"`
	Password   string `json:"password" db:"password"`
	NickName   string `json:"nickname" db:"nickname"`
	Picture    string `json:"picture" db:"picture"`
	Content    string `json:"content" db:"content"`
	CreateTime string `json:"create_time" db:"create_time"`
	UpdateTime string `json:"update_time" db:"update_time"`
	IsDelete   bool   `json:"is_delete" db:"is_delete"`
}
