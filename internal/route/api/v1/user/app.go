// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"fmt"
	"net/http"
	"strings"

	api "github.com/gogs/go-gogs-client"

	"github.com/NII-DG/gogs/internal/conf"
	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	"github.com/NII-DG/gogs/internal/form"
	"github.com/NII-DG/gogs/internal/utils/const_utils"
	gen "github.com/NII-DG/gogs/internal/utils/generator"
)

func ListAccessTokens(c *context.APIContext) {
	tokens, err := db.AccessTokens.List(c.User.ID)
	if err != nil {
		c.Error(err, "list access tokens")
		return
	}
	tokens_without_build_token := []*db.AccessToken{}
	for i := range tokens {
		token_name := tokens[i].Name
		if strings.HasPrefix(token_name, const_utils.Get_BUILD_TOKEN()) {
			continue
		} else {
			tokens_without_build_token = append(tokens_without_build_token, tokens[i])
		}
	}

	apiTokens := make([]*api.AccessToken, len(tokens_without_build_token))
	for i := range tokens_without_build_token {
		apiTokens[i] = &api.AccessToken{Name: tokens_without_build_token[i].Name, Sha1: tokens_without_build_token[i].Sha1}

	}
	c.JSONSuccess(&apiTokens)
}

func CreateAccessToken(c *context.APIContext, form form.CreateAccessTokenOption) {

	t, err := db.AccessTokens.Create(c.User.ID, form.Name, form.ExpireMinutes)
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

func DeleteAccessTokenSelf(c *context.APIContext) {
	req_token := c.Query("token")
	err := db.AccessTokens.DeleteByToken(c.User.ID, req_token)
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

func CreateAccessTokenForLaunch(c *context.APIContext) {
	// create new token for building jupyter container
	randStr, err := gen.MakeRandomStrByAlphabetDigit(7)
	if err != nil {
		c.Error(err, "Failed to generate random string")
	}
	token_name := fmt.Sprintf("%s-%s", const_utils.Get_BUILD_TOKEN(), randStr)

	t, err := db.AccessTokens.Create(c.User.ID, token_name, conf.DG.BuildAccessTokenExpireMinutes)
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
