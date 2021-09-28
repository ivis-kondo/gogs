// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
	"testing"

	"github.com/ivis-yoshida/gogs/internal/context"
	"github.com/ivis-yoshida/gogs/internal/db"
)

func TestExploreMetadata(t *testing.T) {
	type args struct {
		c *context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sample",
			args: args{
				c: &context.Context{
					User: &db.User{Name: "owner"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExploreMetadata(tt.args.c)
		})
	}
}

func TestDmpBrowsing(t *testing.T) {
	type args struct {
		c *context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DmpBrowsing(tt.args.c)
		})
	}
}

func Test_isContained(t *testing.T) {
	type args struct {
		bufStr      string
		selectedKey string
		keyword     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isContained(tt.args.bufStr, tt.args.selectedKey, tt.args.keyword); got != tt.want {
				t.Errorf("isContained() = %v, want %v", got, tt.want)
			}
		})
	}
}
