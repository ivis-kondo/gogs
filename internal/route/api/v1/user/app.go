// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"net/http"

	api "github.com/gogs/go-gogs-client"

	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	"github.com/NII-DG/gogs/internal/form"
	log "unknwon.dev/clog/v2"
)

func ListAccessTokens(c *context.APIContext) {
	tokens, err := db.AccessTokens.List(c.User.ID)
	if err != nil {
		c.Error(err, "list access tokens")
		return
	}

	apiTokens := make([]*api.AccessToken, len(tokens))
	for i := range tokens {
		apiTokens[i] = &api.AccessToken{Name: tokens[i].Name, Sha1: tokens[i].Sha1}
	}
	c.JSONSuccess(&apiTokens)
}

func CreateAccessToken(c *context.APIContext, form api.CreateAccessTokenOption) {

	t, err := db.AccessTokens.Create(c.User.ID, form.Name)
	if err != nil {
		if db.IsErrAccessTokenAlreadyExist(err) {
			c.ErrorStatus(http.StatusUnprocessableEntity, err)
		} else {
			c.Error(err, "new access token")
		}
		return
	}
	c.JSON(http.StatusCreated, &api.AccessToken{Name: t.Name, Sha1: t.Sha1})
}

func DeleteAccessTokenSelf(c *context.APIContext, form form.DeleteAccessTokenOption) {
	q_token := c.Query("token")
	log.Trace("[RCOS DEBUG] q_token : %s", q_token)
	err := db.AccessTokens.DeleteByToken(c.User.ID, form.Token)
	if err != nil {
		if db.IsErrAccessTokenAlreadyExist(err) {
			c.ErrorStatus(http.StatusUnprocessableEntity, err)
		} else {
			c.Error(err, "delete access token")
		}
		return

	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete access token",
	})

}
