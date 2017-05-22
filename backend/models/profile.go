package models

import (
	"fmt"
	"hexerent/backend/database"
)

// Profile stores various information on a user's profile
type Profile struct {
	ProfileID         uint64
	UserID            uint64
	UniversityID      uint64
	CompanyID         uint64
	ProfilePicture    string `json:"profilePicture"`
	BackgroundPicture string `json:"backgroundPicture"`
	About             string `json:"about"`
	Interests         string `json:"interests"`
	ProfileType       string `json:"profileType"`
}

// NewProfile acts as a constructure to create an instance of a new Profile
func NewProfile(userID, universityID, companyID uint64, profilePicure, backgroundPicture, about, interests, profileType string) *Profile {
	profile := new(Profile)
	//profile.ProfileID = profileID
	profile.UserID = userID
	profile.UniversityID = universityID
	profile.CompanyID = companyID
	profile.ProfilePicture = profilePicure
	profile.BackgroundPicture = backgroundPicture
	profile.About = about
	profile.Interests = interests
	profile.ProfileType = profileType

	return profile
}

// FindAllProfile does stuff
func FindAllProfile() {

}

// FindProfile does stuff
func FindProfile(id uint64) Profile {

	var universityID, companyID uint64
	var profilePictureURL, backgroundPictureURL, about, interests, profileType string

	DB, err := database.NewOpen()

	rows, err := DB.Prepare("SELECT UniversityID, CompanyID, ProfilePicture, BackgroundPicture, About, Interests, ProfileType FROM profile WHERE UserID = ?")
	if err != nil {
		fmt.Println(err)
	}

	rows.QueryRow(id).Scan(&universityID, &companyID, &profilePictureURL, &backgroundPictureURL, &about, &interests, &profileType)

	if err != nil {
		fmt.Println(err)
	}

	profile := NewProfile(id, universityID, companyID, profilePictureURL, backgroundPictureURL, about, interests, profileType)

	DB.Close()

	fmt.Println("this is profile from  models: ", profile)

	return *profile

}

// CreateProfile does stuff
func CreateProfile() {

}

// UpdateProfile does stuff
func UpdateProfile() {

}

// DeleteProfile does stuff
func DeleteProfile() {

}
