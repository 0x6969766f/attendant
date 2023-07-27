package game

import (
	"fmt"
	"net/http"

	"github.com/0x6969766f/attendant/api/router/middleman"
)

func (a *API) GetGames(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("getting games...")
	games, _ := a.List()
	fmt.Println("got games", games)
	return middleman.JSONResponse(w, http.StatusOK, games)
}

// func (a *API) GetGames(w http.ResponseWriter, r *http.Request) {
// 	games, _ := a.ListGames()
// 	fmt.Println(games)
// 	render.JSON(w, r, games)
// 	// if err != nil {
// 	// 	fmt.Println("error")
// 	// 	return
// 	// }

// 	// if games == nil {
// 	// 	fmt.Println("[]")
// 	// 	return
// 	// }

// 	// if err := json.NewEncoder(w).Encode(games); err != nil {
// 	// 	fmt.Println("error encoding json")
// 	// 	return
// 	// }
// }
