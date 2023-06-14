package route

import "github.com/NII-DG/gogs/internal/db"

func filterUnlistedRepos(repos []*db.Repository) []*db.Repository {
	// Filter out Unlisted repositories
	var showRep []*db.Repository
	for _, repo := range repos {
		if !repo.IsUnlisted {
			showRep = append(showRep, repo)
		}
	}
	return showRep
}

func filterRepos(userID int64, repos []*db.Repository) []*db.Repository {
	var showRep []*db.Repository
	for _, repo := range repos {
		if repo.Owner.IsOrganization() {
			if repo.Owner.IsOrgMember(userID) {
				showRep = append(showRep, repo)
			}
		} else {
			// make repo private temporary
			tmpIsPrivate := repo.IsPrivate
			repo.IsPrivate = true
			// Authorize
			if db.Perms.Authorize(userID, repo, db.AccessModeRead) {
				repo.IsPrivate = tmpIsPrivate
				showRep = append(showRep, repo)
			}
		}
	}
	return showRep
}

func filterUsers(userID int64, users []*db.User) []*db.User {
	var showUser []*db.User
	for _, user := range users {
		if user.IsOrganization() {
			// show only belonging organization
			if user.IsOrgMember(userID) {
				showUser = append(showUser, user)
			}
		} else {
			// show only user itself
			if user.ID == userID {
				showUser = append(showUser, user)
			}
		}
	}
	return showUser
}
