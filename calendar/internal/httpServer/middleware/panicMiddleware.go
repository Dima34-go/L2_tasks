package middleware

import (
	"L_2/calendar/internal/httpServer"
	"fmt"
	"net/http"
)

func PanicMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		fmt.Printf("PanicMiddleWare: %s\n",r.URL.Path)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recovered",err)
				httpServer.WriteError(w,&httpServer.CustomError{
					Code:    500,
					Message: httpServer.InternalError,
					ErrDesc: fmt.Sprintf("%s",err),
				})
				return
			}
		}()
		next.ServeHTTP(w,r)
	})
}