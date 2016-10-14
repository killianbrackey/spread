package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
)

func GetAuth(c *gin.Context) {
	redirectURL := `http://localhost:8080/callback`
	const (
		redirectURI  = `http://localhost:8080/callback`
		clientID     = "6481b67311a646f4a4f49f43012fc39e"
		clientSecret = "a6e657975b4a49cdbbd6dcec81d0d4ea"
		state        = "abc123"
	)
	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate)
	auth.SetAuthInfo(clientID, clientSecret)

	url := auth.AuthURL("")
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	client := <-ch

	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You are logged in as:", user.ID)

}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}
