package profile

import (
	"fmt"
	"hexerent/backend/microservices/api/home"
	"hexerent/backend/models"
)

// GetProfileInformation processes all the information to be rendered on user's
// profile page.
func GetProfileInformation() models.Profile {

	userID := home.GetUserIdentification()
	fmt.Println("this is userID in api/profile--", userID)
	profile := models.FindProfile(userID)
	//profileData := models.NewProfile(0, userID, 0, 0, "", "", "", "", "")
	return profile
}
