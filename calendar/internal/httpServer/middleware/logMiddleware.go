package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LogMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		fmt.Printf("LogMiddleWare: %s\n",r.URL.Path)
		start :=time.Now()
		next.ServeHTTP(w,r)
		fmt.Printf("[%s] %s, %s %s\n\n",
			r.Method,r.RemoteAddr,r.URL.Path,time.Since(start))
	})
}