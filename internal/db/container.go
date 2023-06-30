// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package db

import (
	"time"

	"github.com/NII-DG/gogs/internal/db/errors"
	"xorm.io/xorm"
)

type ContainerOptions struct {
	RepoID            int64  `json:"repo_id" binding:"Required"`
	UserID            int64  `json:"user_id" binding:"Required"`
	ExperimentPackage string `json:"experiment_package" binding:"MaxSize(255)"`
	ServerName        string `json:"server_name" binding:"Required;MaxSize(255)"`
	Url               string `binding:"Url"`
}

type JupyterContainer struct {
	ID                int64
	RepoID            int64  `xorm:"NOT NULL" gorm:"NOT NULL"`
	UserID            int64  `xorm:"NOT NULL" gorm:"NOT NULL"`
	ExperimentPackage string `xorm:"NOT NULL" gorm:"NOT NULL"`
	ServerName        string `xorm:"UNIQUE NOT NULL" gorm:"UNIQUE"`
	Url               string
	IsDelete          bool      `xorm:"NOT NULL DEFAULT false" gorm:"NOT NULL;DEFAULT:FALSE"`
	Created           time.Time `xorm:"-" gorm:"-" json:"-"`
	CreatedUnix       int64
	Updated           time.Time `xorm:"-" gorm:"-" json:"-"`
	UpdatedUnix       int64
}

func (j *JupyterContainer) BeforeInsert() {
	j.CreatedUnix = time.Now().Unix()
	j.UpdatedUnix = time.Now().Unix()
}

func (j *JupyterContainer) BeforeUpdate() {
	j.UpdatedUnix = time.Now().Unix()
}

func (j *JupyterContainer) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_unix":
		j.Created = time.Unix(j.CreatedUnix, 0).Local()
	case "updated_unix":
		j.Updated = time.Unix(j.UpdatedUnix, 0).Local()
	}
}

func AddJupyterContainer(container *JupyterContainer) (err error) {

	sess := x.NewSession()
	defer sess.Close()

	if err = sess.Begin(); err != nil {
		return err
	}

	_, err = sess.InsertOne(container)
	if err != nil {
		return err
	}

	return sess.Commit()
}

func GetJupyterContainer(RepoID int64, UserID int64) ([]*JupyterContainer, error) {
	containers := make([]*JupyterContainer, 0)
	return containers, x.Where("repo_id=?", RepoID).And("user_id=?", UserID).Find(&containers)
}

func UpdateJupyterContainer(container *JupyterContainer) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}
	var oldContainer JupyterContainer
	res, err := sess.Where("server_name = ?", container.ServerName).Get(&oldContainer)

	if !res {
		return errors.New("container not found")
	} else {
		oldContainer.IsDelete = container.IsDelete
		_, err = sess.ID(oldContainer.ID).AllCols().Update(&oldContainer)
		if err != nil {
			return err
		}
	}
	return sess.Commit()
}
