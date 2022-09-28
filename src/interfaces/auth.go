package interfaces

import (
	"github.com/wildanfaz/backendgolang2_week10/src/database/orm/models"
	"github.com/wildanfaz/backendgolang2_week10/src/libs"
)

type AuthService interface {
	Login(body models.User) *libs.Resp
}
