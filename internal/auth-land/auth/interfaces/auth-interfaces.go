package interfaces

import (
	citf "auth-demo/internal/common/interfaces"
	d "auth-demo/internal/auth-land/auth/domain"

	"github.com/gin-gonic/gin"
)

type AuthServiceITF interface {
	citf.Common[d.Auth]
	Comparate(gctx *gin.Context, data *d.Auth) error
}

type AuthRepositoryITF interface {
	citf.Common[d.Auth]
	GetByEmail(gctx *gin.Context, data *d.Auth) error
}
