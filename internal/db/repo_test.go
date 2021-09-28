package db

import (
	_ "image/jpeg"
	"testing"
	"time"

	"github.com/ivis-yoshida/gogs/internal/markup"
	"github.com/stretchr/testify/assert"
)

func TestRepository_ComposeMetas(t *testing.T) {
	repo := &Repository{
		Name: "testrepo",
		Owner: &User{
			Name: "testuser",
		},
		ExternalTrackerFormat: "https://someurl.com/{user}/{repo}/{issue}",
	}

	t.Run("no external tracker is configured", func(t *testing.T) {
		repo.EnableExternalTracker = false

		metas := repo.ComposeMetas()
		assert.Equal(t, metas["repoLink"], repo.Link())

		// Should no format and style if no external tracker is configured
		_, ok := metas["format"]
		assert.False(t, ok)
		_, ok = metas["style"]
		assert.False(t, ok)
	})

	t.Run("an external issue tracker is configured", func(t *testing.T) {
		repo.ExternalMetas = nil
		repo.EnableExternalTracker = true

		// Default to numeric issue style
		assert.Equal(t, markup.ISSUE_NAME_STYLE_NUMERIC, repo.ComposeMetas()["style"])
		repo.ExternalMetas = nil

		repo.ExternalTrackerStyle = markup.ISSUE_NAME_STYLE_NUMERIC
		assert.Equal(t, markup.ISSUE_NAME_STYLE_NUMERIC, repo.ComposeMetas()["style"])
		repo.ExternalMetas = nil

		repo.ExternalTrackerStyle = markup.ISSUE_NAME_STYLE_ALPHANUMERIC
		assert.Equal(t, markup.ISSUE_NAME_STYLE_ALPHANUMERIC, repo.ComposeMetas()["style"])
		repo.ExternalMetas = nil

		metas := repo.ComposeMetas()
		assert.Equal(t, "testuser", metas["user"])
		assert.Equal(t, "testrepo", metas["repo"])
		assert.Equal(t, "https://someurl.com/{user}/{repo}/{issue}", metas["format"])
	})
}

func TestRepository_loadAttributes(t *testing.T) {
	type fields struct {
		ID                    int64
		OwnerID               int64
		Owner                 *User
		LowerName             string
		Name                  string
		Description           string
		Website               string
		DefaultBranch         string
		Size                  int64
		UseCustomAvatar       bool
		NumWatches            int
		NumStars              int
		NumForks              int
		NumIssues             int
		NumClosedIssues       int
		NumOpenIssues         int
		NumPulls              int
		NumClosedPulls        int
		NumOpenPulls          int
		NumMilestones         int
		NumClosedMilestones   int
		NumOpenMilestones     int
		NumTags               int
		IsPrivate             bool
		IsUnlisted            bool
		IsBare                bool
		HasMetadata           bool
		IsMirror              bool
		Mirror                *Mirror
		EnableWiki            bool
		AllowPublicWiki       bool
		EnableExternalWiki    bool
		ExternalWikiURL       string
		EnableIssues          bool
		AllowPublicIssues     bool
		EnableExternalTracker bool
		ExternalTrackerURL    string
		ExternalTrackerFormat string
		ExternalTrackerStyle  string
		ExternalMetas         map[string]string
		EnablePulls           bool
		PullsIgnoreWhitespace bool
		PullsAllowRebase      bool
		IsFork                bool
		ForkID                int64
		BaseRepo              *Repository
		Created               time.Time
		CreatedUnix           int64
		Updated               time.Time
		UpdatedUnix           int64
		Downloaded            uint64
	}
	type args struct {
		e Engine
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				ID:                    tt.fields.ID,
				OwnerID:               tt.fields.OwnerID,
				Owner:                 tt.fields.Owner,
				LowerName:             tt.fields.LowerName,
				Name:                  tt.fields.Name,
				Description:           tt.fields.Description,
				Website:               tt.fields.Website,
				DefaultBranch:         tt.fields.DefaultBranch,
				Size:                  tt.fields.Size,
				UseCustomAvatar:       tt.fields.UseCustomAvatar,
				NumWatches:            tt.fields.NumWatches,
				NumStars:              tt.fields.NumStars,
				NumForks:              tt.fields.NumForks,
				NumIssues:             tt.fields.NumIssues,
				NumClosedIssues:       tt.fields.NumClosedIssues,
				NumOpenIssues:         tt.fields.NumOpenIssues,
				NumPulls:              tt.fields.NumPulls,
				NumClosedPulls:        tt.fields.NumClosedPulls,
				NumOpenPulls:          tt.fields.NumOpenPulls,
				NumMilestones:         tt.fields.NumMilestones,
				NumClosedMilestones:   tt.fields.NumClosedMilestones,
				NumOpenMilestones:     tt.fields.NumOpenMilestones,
				NumTags:               tt.fields.NumTags,
				IsPrivate:             tt.fields.IsPrivate,
				IsUnlisted:            tt.fields.IsUnlisted,
				IsBare:                tt.fields.IsBare,
				HasMetadata:           tt.fields.HasMetadata,
				IsMirror:              tt.fields.IsMirror,
				Mirror:                tt.fields.Mirror,
				EnableWiki:            tt.fields.EnableWiki,
				AllowPublicWiki:       tt.fields.AllowPublicWiki,
				EnableExternalWiki:    tt.fields.EnableExternalWiki,
				ExternalWikiURL:       tt.fields.ExternalWikiURL,
				EnableIssues:          tt.fields.EnableIssues,
				AllowPublicIssues:     tt.fields.AllowPublicIssues,
				EnableExternalTracker: tt.fields.EnableExternalTracker,
				ExternalTrackerURL:    tt.fields.ExternalTrackerURL,
				ExternalTrackerFormat: tt.fields.ExternalTrackerFormat,
				ExternalTrackerStyle:  tt.fields.ExternalTrackerStyle,
				ExternalMetas:         tt.fields.ExternalMetas,
				EnablePulls:           tt.fields.EnablePulls,
				PullsIgnoreWhitespace: tt.fields.PullsIgnoreWhitespace,
				PullsAllowRebase:      tt.fields.PullsAllowRebase,
				IsFork:                tt.fields.IsFork,
				ForkID:                tt.fields.ForkID,
				BaseRepo:              tt.fields.BaseRepo,
				Created:               tt.fields.Created,
				CreatedUnix:           tt.fields.CreatedUnix,
				Updated:               tt.fields.Updated,
				UpdatedUnix:           tt.fields.UpdatedUnix,
				Downloaded:            tt.fields.Downloaded,
			}
			if err := repo.loadAttributes(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("Repository.loadAttributes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_LoadAttributes(t *testing.T) {
	type fields struct {
		ID                    int64
		OwnerID               int64
		Owner                 *User
		LowerName             string
		Name                  string
		Description           string
		Website               string
		DefaultBranch         string
		Size                  int64
		UseCustomAvatar       bool
		NumWatches            int
		NumStars              int
		NumForks              int
		NumIssues             int
		NumClosedIssues       int
		NumOpenIssues         int
		NumPulls              int
		NumClosedPulls        int
		NumOpenPulls          int
		NumMilestones         int
		NumClosedMilestones   int
		NumOpenMilestones     int
		NumTags               int
		IsPrivate             bool
		IsUnlisted            bool
		IsBare                bool
		HasMetadata           bool
		IsMirror              bool
		Mirror                *Mirror
		EnableWiki            bool
		AllowPublicWiki       bool
		EnableExternalWiki    bool
		ExternalWikiURL       string
		EnableIssues          bool
		AllowPublicIssues     bool
		EnableExternalTracker bool
		ExternalTrackerURL    string
		ExternalTrackerFormat string
		ExternalTrackerStyle  string
		ExternalMetas         map[string]string
		EnablePulls           bool
		PullsIgnoreWhitespace bool
		PullsAllowRebase      bool
		IsFork                bool
		ForkID                int64
		BaseRepo              *Repository
		Created               time.Time
		CreatedUnix           int64
		Updated               time.Time
		UpdatedUnix           int64
		Downloaded            uint64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &Repository{
				ID:                    tt.fields.ID,
				OwnerID:               tt.fields.OwnerID,
				Owner:                 tt.fields.Owner,
				LowerName:             tt.fields.LowerName,
				Name:                  tt.fields.Name,
				Description:           tt.fields.Description,
				Website:               tt.fields.Website,
				DefaultBranch:         tt.fields.DefaultBranch,
				Size:                  tt.fields.Size,
				UseCustomAvatar:       tt.fields.UseCustomAvatar,
				NumWatches:            tt.fields.NumWatches,
				NumStars:              tt.fields.NumStars,
				NumForks:              tt.fields.NumForks,
				NumIssues:             tt.fields.NumIssues,
				NumClosedIssues:       tt.fields.NumClosedIssues,
				NumOpenIssues:         tt.fields.NumOpenIssues,
				NumPulls:              tt.fields.NumPulls,
				NumClosedPulls:        tt.fields.NumClosedPulls,
				NumOpenPulls:          tt.fields.NumOpenPulls,
				NumMilestones:         tt.fields.NumMilestones,
				NumClosedMilestones:   tt.fields.NumClosedMilestones,
				NumOpenMilestones:     tt.fields.NumOpenMilestones,
				NumTags:               tt.fields.NumTags,
				IsPrivate:             tt.fields.IsPrivate,
				IsUnlisted:            tt.fields.IsUnlisted,
				IsBare:                tt.fields.IsBare,
				HasMetadata:           tt.fields.HasMetadata,
				IsMirror:              tt.fields.IsMirror,
				Mirror:                tt.fields.Mirror,
				EnableWiki:            tt.fields.EnableWiki,
				AllowPublicWiki:       tt.fields.AllowPublicWiki,
				EnableExternalWiki:    tt.fields.EnableExternalWiki,
				ExternalWikiURL:       tt.fields.ExternalWikiURL,
				EnableIssues:          tt.fields.EnableIssues,
				AllowPublicIssues:     tt.fields.AllowPublicIssues,
				EnableExternalTracker: tt.fields.EnableExternalTracker,
				ExternalTrackerURL:    tt.fields.ExternalTrackerURL,
				ExternalTrackerFormat: tt.fields.ExternalTrackerFormat,
				ExternalTrackerStyle:  tt.fields.ExternalTrackerStyle,
				ExternalMetas:         tt.fields.ExternalMetas,
				EnablePulls:           tt.fields.EnablePulls,
				PullsIgnoreWhitespace: tt.fields.PullsIgnoreWhitespace,
				PullsAllowRebase:      tt.fields.PullsAllowRebase,
				IsFork:                tt.fields.IsFork,
				ForkID:                tt.fields.ForkID,
				BaseRepo:              tt.fields.BaseRepo,
				Created:               tt.fields.Created,
				CreatedUnix:           tt.fields.CreatedUnix,
				Updated:               tt.fields.Updated,
				UpdatedUnix:           tt.fields.UpdatedUnix,
				Downloaded:            tt.fields.Downloaded,
			}
			if err := repo.LoadAttributes(); (err != nil) != tt.wantErr {
				t.Errorf("Repository.LoadAttributes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
