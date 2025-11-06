package Authentication_System

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Level    string `json:"level"`
}

var LoginCheckU = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var LoginCheckP = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"

var LevelCheck = []string{
	"Admin",
	"Moderator",
	"Member",
}

type SignForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Level    string `json:"level"`
	Code     string `json:"code"`
	Email    string `json:"email"`
}

type SignForm1 struct {
	Otp string `json:"otp"`
}

var Code = []string{
	"12as12",
	"lo106o5",
}
