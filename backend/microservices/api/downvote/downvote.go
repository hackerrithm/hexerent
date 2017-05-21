package downvote

import (
	"fmt"
	"hexerent/backend/controllers/home"
	"hexerent/backend/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Create stuff
func Create(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/user/home", http.StatusSeeOther)
	} else if r.Method == http.MethodPost {

		var newID uint64
		vars := mux.Vars(r)
		postID := vars["postID"]
		fmt.Println("Post show:", postID)
		newID, _ = strconv.ParseUint(postID, 10, 0)

		userID := home.GetUserInformation()
		downvote := models.NewDownvote(0, userID, newID, true)

		models.CreateDownvote(downvote, newID, userID)
		http.Redirect(w, r, "/user/home", http.StatusSeeOther)
	}
}
