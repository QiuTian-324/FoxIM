package dto

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// 添加好友请求
type AddFriendRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

type AddFriendResponse struct {
	FriendID uint `json:"friend_id"`
}

// 好友列表响应
type FriendListResponse struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
	Bio       string `json:"bio"`
	Online    int    `json:"online"`
}
