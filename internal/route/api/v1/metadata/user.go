package metadata

import (
	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	log "unknwon.dev/clog/v2"
)

func GetUser(c *context.APIContext) {
	req_user := c.User
	log.Trace("API to get Users Metadata has been done by User[ID : %d]", req_user.ID)

	userName := c.Params(":username")
	u, err := db.GetUserByName(userName)
	if err != nil {
		c.NotFoundOrError(err, "get user by name")
		return
	}

	org := UserOrgMetadata{
		Name:        u.Affiliation,
		Url:         u.AffiliationURL,
		AliasName:   u.AffiliationAlias,
		Description: u.AffiliationDescription,
	}
	user := UserMatadata{
		UserName:    u.Name,
		Url:         u.PersonalURL,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		AliasName:   u.AliasName,
		EMail:       u.Email,
		Telephone:   u.Telephone,
		ERadNumber:  u.ERadResearcherNumber,
		Affiliation: org,
	}
	c.JSONSuccess(user)
}

func GetUsers(c *context.APIContext, form UserNameList) {
	req_user := c.User
	log.Trace("API to get Users Metadata has been done by User[ID : %d]", req_user.ID)

	users := UsersMatadata{}
	for _, userName := range form.UsersName {
		u, err := db.GetUserByName(userName)
		if err != nil {
			c.NotFoundOrError(err, "get user by name")
			return
		}
		org := UserOrgMetadata{
			Name:        u.Affiliation,
			Url:         u.AffiliationURL,
			AliasName:   u.AffiliationAlias,
			Description: u.AffiliationDescription,
		}
		user := UserMatadata{
			UserName:    u.Name,
			Url:         u.PersonalURL,
			FirstName:   u.FirstName,
			LastName:    u.LastName,
			AliasName:   "",
			EMail:       u.Email,
			Telephone:   u.Telephone,
			ERadNumber:  u.ERadResearcherNumber,
			Affiliation: org,
		}
		users.Users = append(users.Users, user)
	}
	c.JSONSuccess(users)
}
