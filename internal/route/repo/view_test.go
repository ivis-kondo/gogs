// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"testing"

	"github.com/ivis-yoshida/gogs/internal/context"
)

func Test_bidingDmpSchemaList(t *testing.T) {
	type args struct {
		c       *context.Context
		dirPath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bidingDmpSchemaList(tt.args.c, tt.args.dirPath)
		})
	}
}

func Test_fetchDmpSchema(t *testing.T) {
	type args struct {
		c        *context.Context
		filePath string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fetchDmpSchema(tt.args.c, tt.args.filePath)
		})
	}
}
