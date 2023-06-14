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

// 自分の所有しているリポジトリのみ表示する
func filterRepos(userID int64, repos []*db.Repository) []*db.Repository {
	var showRep []*db.Repository
	for _, repo := range repos {
		if repo.OwnerID == userID {
			showRep = append(showRep, repo)
		}
	}
	return showRep
}

func filterUsers(userID int64, users []*db.User) []*db.User {
	var showUser []*db.User
	for _, user := range users {
		// 組織の場合、自分が所属していれば表示する
		if user.IsOrganization() {
			if user.IsOrgMember(userID) {
				showUser = append(showUser, user)
			}
		} else { // ユーザーの場合
			isShow := map[int64]bool{}
			isShow[userID] = true

			repos, _ := db.GetUserAndCollaborativeRepositories(userID)
			for _, repo := range repos {
				// 自分が所有するリポジトリの共同編集者は表示する
				if repo.OwnerID == userID {
					collaborators, _ := repo.GetCollaborators()
					for _, collaborator := range collaborators {
						isShow[collaborator.User.ID] = true
					}
				} else { // 自分が共同編集しているリポジトリの所有者は表示する
					isShow[repo.OwnerID] = true
				}
			}

			// 自分と同じ組織に所属しているユーザーは表示する
			orgs, _ := db.GetOrgsByUserID(userID, true)
			for _, org := range orgs {
				if org.IsOrgMember(user.ID) {
					isShow[user.ID] = true
				}
			}
			if isShow[user.ID] == true {
				showUser = append(showUser, user)
			}
		}
	}
	return showUser
}
