package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	autorisationHeader = "Authorization"
	userCtx            = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(autorisationHeader)
	if header == "" {
		newErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	//parse token
	userId, err := h.services.Authorisation.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}
