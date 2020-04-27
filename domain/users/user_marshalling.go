package users

import "encoding/json"

type  PrivateUser struct {
	Id int64 `json:"id"`
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	Email string `json:"email"`
	Datecreated string `json:"date_created"`
	Status string `json:"status"`
}

type  PublicUser struct {
	Id int64 `json:"id"`
	Email string `json:"email"`
	Datecreated string `json:"date_created"`
}

func (u *User)Marshal(IsPublic bool) interface{} {
	if IsPublic{
		return PublicUser{
			Id: u.Id	,
			Email: u.Email,
			Datecreated: u.Datecreated,
		}
	}
	// Imp approach
	bytes, _ := json.Marshal(u)
	var user PrivateUser
	json.Unmarshal(bytes, &user)
	return user
}