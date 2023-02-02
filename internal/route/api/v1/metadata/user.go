package metadata

import (
	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	log "unknwon.dev/clog/v2"
)

func SearchUser(c *context.APIContext) {
	userName := c.Params(":username")
	u, err := db.GetUserByName(userName)
	if err != nil {
		c.NotFoundOrError(err, "get user by name")
		return
	}

	org := UserOrgMetadata{
		Name:      "東京大学",
		Url:       "https://u_tokyo",
		AliasName: "U-Tokyo",
	}
	user := UserMatadata{
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

func SearchUsers(c *context.APIContext, form UserNameList) {
	req_user := c.User
	log.Trace("user : %s", req_user.ID)
	log.Trace("user : %s", req_user.FullName)
	log.Trace("user : %s", req_user.Email)

	users := UsersMatadata{}
	for _, userName := range form.UsersName {
		u, err := db.GetUserByName(userName)
		if err != nil {
			c.NotFoundOrError(err, "get user by name")
			return
		}
		org := UserOrgMetadata{
			Name:      "東京大学",
			Url:       "https://u_tokyo",
			AliasName: "U-Tokyo",
		}
		user := UserMatadata{
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
