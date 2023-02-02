package metadata

import (
	"github.com/NII-DG/gogs/internal/context"
	"github.com/NII-DG/gogs/internal/db"
	ds "github.com/NII-DG/gogs/internal/route/api/v1/metadata/data_structure"
)

func Search(c *context.APIContext) {
	userName := c.Params(":username")
	u, err := db.GetUserByName(userName)
	if err != nil {
		c.NotFoundOrError(err, "get user by name")
		return
	}
	user := ds.UserMatadata{
		UserName:   u.FullName,
		Url:        "https://sample",
		FirstName:  "sam",
		LastName:   "ple",
		AliasName:  "sp",
		EMail:      "sample@gmail.com",
		Telephone:  "090-1111-22222",
		ERadNumber: "12345678",
	}
	c.JSONSuccess(user)
}
