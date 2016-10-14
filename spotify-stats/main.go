package main

import (
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/zmb3/spotify"
)

func init() {
	os.Setenv("SPOTIFY_ID", clientID)
	os.Setenv("SPOTIFY_SECRET", clientSecret)
}

func main() {
	r := gin.Default()
	// SetAPIRoutes(r)
	r.StaticFile("/index", "./index.html")
	r.POST("/callback", handlecallback)
	r.GET("/callback", handlecallback)
	r.GET("/auth", GetAuth)
	// r.GET("/playlists", HandleGetPlaylists)
	// GetAuth()

	r.Run(":8080")
}

func SetAPIRoutes(router *gin.Engine) {

}

func handlecallback(c *gin.Context) {
	txt, _ := httputil.DumpRequest(c.Request, true)
	glog.Info(string(txt))
	c.String(http.StatusAccepted, "")
}

// HandleGetPlaylists
func HandleGetPlaylists(c *gin.Context) {

}

const (
	redirectURI  = "http://localhost:8080/callback"
	clientID     = "6481b67311a646f4a4f49f43012fc39e"
	clientSecret = "a6e657975b4a49cdbbd6dcec81d0d4ea"
	token        = "AQDHlLC5fsaqz1cUJiNef0OoTzy6h2mCBJh_gu635AVyXudm0bAtJ5aZXdc50PzCYz-rCuEBp3btzwboREDyd6-ZRI_C2Hv7xOcv2JM73gRWW5kTTqWUubWx0XlpJ9OG39xGI4I24BLdZApTvacKExkK8a4sXY5bqRGiC8zPDMUWmp7IYxtzY4fy0mJAqA5olt3ZQJOpTwv0nUmK_mLfkfwmQFYkdag"
)

var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate)
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

// func redirectHandler(w http.ResopnseWriter, r *http.Request) {
func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	clientID := "6481b67311a646f4a4f49f43012fc39e"
	clientSecret := "a6e657975b4a49cdbbd6dcec81d0d4ea"
	localURI := "http://localhost:8080"
	auth := spotify.NewAuthenticator(localURI, spotify.ScopeUserReadPrivate)

	url := auth.AuthURL("")
	glog.Info(url)

	auth.SetAuthInfo(clientID, clientSecret)
	// token, err := auth.Token(state, r)
	// if err != nil {
	// 	glog.Info(err)
	// 	http.Error(w, "couldn't get token", http.StatusNotFound)
	// 	return
	// }

}

// func genereateClient() spotify.Client {

// }
