package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseUrl string, pathPrefix string) http.HandlerFunc {
	// create url object from targetBaseUrl string
	target, err := url.Parse(targetBaseUrl)
	if err != nil {
		fmt.Println("Error in parsing the url:", err)
		return nil
	}
	// returns a proxy object
	proxy := httputil.NewSingleHostReverseProxy(target)

	// use director to manipulate the url

	orginalDirector := proxy.Director // store the orginal director

	proxy.Director = func(r *http.Request) {

		orginalDirector(r)
		//proxy.Director(r) ❌ calling yourself will keep you infinite recursion so we are stroring in variable and then calling it
		
		//Start manipulating the request
		r.Host = target.Host
		r.URL.Path = strings.TrimPrefix(r.URL.Path, pathPrefix)

		if userId, ok := r.Context().Value("userId").(string); ok {
			r.Header.Set("X-User-ID", userId)
		}
	}
	return proxy.ServeHTTP
}
