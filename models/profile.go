package models

import "github.com/Reza-Rayan/twitter-like-app/db"

type Profile struct {
	ID        int64        `json:"id"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
	Avatar    *string      `json:"avatar,omitempty"`
	Followers []PublicUser `json:"followers"`
	Following []PublicUser `json:"following"`
}

func GetUserProfile(userID int64) (*Profile, error) {
	query := `SELECT id, email, COALESCE(username, ''), avatar FROM users WHERE id = ?`

	var user Profile
	err := db.DB.QueryRow(query, userID).Scan(&user.ID, &user.Email, &user.Username, &user.Avatar)
	if err != nil {
		return nil, err
	}

	// Get followers
	followers, err := GetFollowers(userID)
	if err != nil {
		return nil, err
	}
	user.Followers = followers

	// Get following
	following, err := GetFollowing(userID)
	if err != nil {
		return nil, err
	}
	user.Following = following

	return &user, nil
}

func UpdateUserAvatar(userID int64, avatarURL string) error {
	query := `UPDATE users SET avatar = ? WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(avatarURL, userID)
	return err
}
