package metadata

import (
	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	ds "github.com/NII-DG/gogs/internal/route/api/v1/metadata/data_structure"
	log "unknwon.dev/clog/v2"
)

func Search(c *context.APIContext) {
	userName := c.Params(":username")
	u, err := db.GetUserByName(userName)
	if err != nil {
		c.NotFoundOrError(err, "get user by name")
		return
	}

	org := ds.UserOrgMetadata{
		Name:      "東京大学",
		Url:       "https://u_tokyo",
		AliasName: "U-Tokyo",
	}
	user := ds.UserMatadata{
		UserName:    u.FullName,
		Url:         "https://sample",
		FirstName:   "sam",
		LastName:    "ple",
		AliasName:   "sp",
		EMail:       "sample@gmail.com",
		Telephone:   "090-1111-22222",
		Affiliation: org,
	}
	c.JSONSuccess(user)
}

func SearchUsers(c *context.APIContext, form ds.UserNameList) {
	h := c.Header()
	he := h.Get("Authorization")
	log.Trace("HTTP Header %s", he)

	users := ds.UsersMatadata{}
	for _, userName := range form.UsersName {
		u, err := db.GetUserByName(userName)
		if err != nil {
			c.NotFoundOrError(err, "get user by name")
			return
		}
		org := ds.UserOrgMetadata{
			Name:      "東京大学",
			Url:       "https://u_tokyo",
			AliasName: "U-Tokyo",
		}
		user := ds.UserMatadata{
			UserName:    u.FullName,
			Url:         "https://sample",
			FirstName:   "sam",
			LastName:    "ple",
			AliasName:   "sp",
			EMail:       "sample@gmail.com",
			Telephone:   "090-1111-22222",
			Affiliation: org,
		}
		users.Users = append(users.Users, user)
	}
	c.JSONSuccess(users)
}
