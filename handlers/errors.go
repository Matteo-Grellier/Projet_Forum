package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func Error404(w http.ResponseWriter, req *http.Request) bool {
	fmt.Println(req.URL.Path)

	arr := []string{"/", "/connexion", "/likedPosts", "/oneCategory", "/postsActivity", "/topic", "/inscription", "/test", "/categories"}
	compteurURL := 0
	for i := 0; i < len(arr); i++ {
		if req.URL.Path != arr[i] {
			compteurURL++
		} else if req.URL.Path == arr[i] {
			break
		}
	}
	if compteurURL == len(arr) {
		t, _ := template.ParseFiles("./templates/layouts/error404.html", "./templates/layouts/header.html", "./templates/layouts/sidebar.html")
		Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Page404'")
		t.Execute(w, nil)
		return false
	}
	return true
}

func NoItemsError(w http.ResponseWriter) {
	t, _ := template.ParseFiles("./templates/layouts/error404.html", "./templates/layouts/header.html", "./templates/layouts/sidebar.html")
	Color(1, "[SERVER_INFO_PAGE] : 🟢 Page 'Page404'")
	t.Execute(w, nil)
}

func Error500(w http.ResponseWriter, req *http.Request, err error) {
	Color(3, "[SERVER_INFO_PAGE] : 🟠 Template execution : ")
	fmt.Println(err)
	t, _ := template.ParseFiles("./templates/layouts/error500.html")
	t.Execute(w, nil)
}

/*// indexHandler - Serve INDEX.HTML page for 'HTTP GET /' request
//
func indexHandler(w http.ResponseWriter, r *http.Request) {

	// HTTP/404 - Not Found management
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		logger.Printf("[HTTP/404] - Page not found URL='%s'\n", r.URL.Path)
		return
	}

	// Load banner list from COMMON declaration
	tmpBannerList := make([]string, len(bannerSupported))
	indexBanner := 0
	for key := range bannerSupported {
		tmpBannerList[indexBanner] = key
		indexBanner++
	}

	// Initialize dynamic content of the page
	htmlData := HTMLData{
		Title:       "YTrack - Ascii-art-web",
		Description: "Simple Golang webserver with built in logging, tracing and health check",
		Keywords:    "golang web server ascii art",
		Author:      "Cedric OBEJERO, <cedric.obejero@tanooki.fr>",
		BannerFiles: tmpBannerList,
		ASCIIText:   template.HTML(``),
	}
}*/
