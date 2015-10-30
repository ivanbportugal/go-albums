package main

import (
	"./server"
	"github.com/go-martini/martini"
	"net/http"
	"regexp"
	"strings"
)

// The one and only access token! In real-life scenarios, a more complex authentication
// middleware than auth.Basic should be used, obviously.
// const AuthToken = "token"

// The one and only martini instance.
var m *martini.Martini

var (
	hostname     string
	port         int
	topStaticDir string
)

func init() {
	m = martini.New()

	// I could probably just use martini.Classic() instead of configure these manually

	// Static files
	m.Use(martini.Static(`public`))
	// Setup middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())
	// m.Use(auth.Basic(AuthToken, ""))
	m.Use(MapEncoder)

	// Setup routes
	r := martini.NewRouter()
	r.Get(`/albums`, server.GetAlbums)
	r.Get(`/albums/:id`, server.GetAlbum)
	r.Post(`/albums`, server.AddAlbum)
	r.Put(`/albums/:id`, server.UpdateAlbum)
	r.Delete(`/albums/:id`, server.DeleteAlbum)

	// Inject database
	m.MapTo(server.DBInstance, (*server.DB)(nil))
	// Add the router action
	m.Action(r.Handle)
}

// The regex to check for the requested format (allows an optional trailing
// slash).
var rxExt = regexp.MustCompile(`(\.(?:xml|text|json))\/?$`)

// MapEncoder intercepts the request's URL, detects the requested format,
// and injects the correct encoder dependency for this request. It rewrites
// the URL to remove the format extension, so that routes can be defined
// without it.
func MapEncoder(c martini.Context, w http.ResponseWriter, r *http.Request) {
	// Get the format extension
	matches := rxExt.FindStringSubmatch(r.URL.Path)
	ft := ".json"
	if len(matches) > 1 {
		// Rewrite the URL without the format extension
		l := len(r.URL.Path) - len(matches[1])
		if strings.HasSuffix(r.URL.Path, "/") {
			l--
		}
		r.URL.Path = r.URL.Path[:l]
		ft = matches[1]
	}
	// Inject the requested encoder
	switch ft {
	case ".xml":
		c.MapTo(server.XmlEncoder{}, (*server.Encoder)(nil))
		w.Header().Set("Content-Type", "application/xml")
	case ".text":
		c.MapTo(server.TextEncoder{}, (*server.Encoder)(nil))
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	default:
		c.MapTo(server.JsonEncoder{}, (*server.Encoder)(nil))
		w.Header().Set("Content-Type", "application/json")
	}
}

func main() {

	m.Run()
}
