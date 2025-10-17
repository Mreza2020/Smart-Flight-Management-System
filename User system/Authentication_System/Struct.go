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
