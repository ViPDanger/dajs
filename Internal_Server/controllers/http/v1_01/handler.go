package v1_01

import (
	"github.com/gin-gonic/gin"
)

func SetupHandlers(r *gin.Engine) {

}

var users = map[string]string{} // Имитация БД: username -> hashedPassword
