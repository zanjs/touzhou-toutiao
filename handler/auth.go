package handler

import (
	"time"

	"anla.io/hound/config"
	"anla.io/hound/models"
	"anla.io/hound/response"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
)

var jwtConfig = config.Config.JWT

// PostLogin is login get token
func PostLogin(ctx iris.Context) {

	u := &models.UserLogin{}
	if err := ctx.ReadJSON(u); err != nil {
		response.JSONError(ctx, err.Error())
		return
	}

	if u.Username == "" {
		response.JSONError(ctx, "Username where?")
		return
	}

	if u.Password == "" {
		response.JSONError(ctx, "Password where?")
		return
	}

	user, _ := models.User{}.GetByUsername(u.Username)

	if user.ID == 0 {
		response.JSONError(ctx, "用户名不存在")
		return
	}

	// hashPassword := utils.HashPassword(u.Password)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))

	if err != nil {
		response.JSONError(ctx, "用户名或密码错误")
		return
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtConfig.Secret))
	if err != nil {
		response.JSONError(ctx, "err")
		return
	}

	response.JSON(ctx, t)
}
