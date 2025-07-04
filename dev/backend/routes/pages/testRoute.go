package routes_pages

import (
	"net/http"
	app "placeholder/dev/features/app"
	test_page_mob "placeholder/dev/frontend/desktop/pages/test_page"
	"strings"
)

func init() {
	app.GetInstance().Router.RegisterRoutes("/test/route", "GET", TestPageRoute)
}

func TestPageRoute(w http.ResponseWriter, r *http.Request) {
	// Send the page only if the user agent is mobile
	if strings.Contains(r.UserAgent(), "Mobile") {
		test_page_mob.TestPage("test").Render(r.Context(), w)
		return
	}

	// If the user agent is not mobile, return a 404 error
	http.Error(w, "Page not found", http.StatusNotFound)
}
