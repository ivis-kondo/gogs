// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package container

import (
	"net/http"

	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
)

type ContainerOptions struct {
	RepoID            int64  `json:"repo_id" binding:"Required"`
	UserID            int64  `json:"user_id" binding:"Required"`
	ExperimentPackage string `json:"experiment_package" binding:"MaxSize(255)"`
	ServerName        string `json:"server_name" binding:"Required;MaxSize(255)"`
	Url               string `binding:"Url"`
}

func AddJupyterContainer(c *context.APIContext, opts ContainerOptions) {

	if opts.UserID != c.UserID() {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"ok": false,
		})
		return
	}

	container := &db.JupyterContainer{
		UserID:            opts.UserID,
		RepoID:            opts.RepoID,
		ExperimentPackage: opts.ExperimentPackage,
		ServerName:        opts.ServerName,
		Url:               opts.Url,
	}

	err := db.AddJupyterContainer(container)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"ok":    false,
			"error": err.Error(),
		})
		return
	} else {
		c.JSONSuccess(map[string]interface{}{
			"ok": true,
		})
	}
}

func UpdateJupyterContainer(c *context.APIContext) {

	ServerName := c.Query("server_name")
	UserID := c.QueryInt64("user_id")

	if UserID != c.UserID() {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"ok": false,
		})
		return
	}

	err := db.UpdateJupyterContainer(&db.JupyterContainer{
		ServerName: ServerName,
		UserID:     UserID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"ok":    false,
			"error": err.Error(),
		})
		return
	} else {
		c.JSONSuccess(map[string]interface{}{
			"ok": true,
		})
	}
}

func DeleteJupyterContainer(c *context.APIContext) {

	ServerName := c.Query("server_name")
	UserID := c.QueryInt64("user_id")

	if UserID != c.UserID() {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"ok": false,
		})
		return
	}

	err := db.UpdateJupyterContainer(&db.JupyterContainer{
		ServerName: ServerName,
		UserID:     UserID,
		IsDelete:   true,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"ok":    false,
			"error": err.Error(),
		})
		return
	} else {
		c.JSONSuccess(map[string]interface{}{
			"ok": true,
		})
	}
}
