package storage

import "mini-task/models"


var Users = []models.User{}
var lastId = 0

func CreateUser(user models.User) models.User {
	lastId++
	user.ID = lastId 
	Users = append(Users, user)
	return  user
}

func GetUsersFiltered(afterID, limit int) ([]models.User, int) {
	filtered := []models.User{}
	for _, u := range Users {
		if u.ID > afterID {
			filtered = append(filtered, u)
		}
	}

	if len(filtered) > limit {
		filtered = filtered[:limit]
	}

	nextAfter := 0
	if len(filtered) > 0 {
		nextAfter = filtered[len(filtered)-1].ID
	}

	return filtered, nextAfter
}

func UpdateUser(id int, userData models.User) (models.User, bool) {
	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = userData.Name
			Users[i].Email = userData.Email
			Users[i].Age = userData.Age
			return Users[i], true
		}
	}
	return models.User{}, false
}
func DeleteUser(id int) bool {
	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			return true
		}
	}
	return false
}