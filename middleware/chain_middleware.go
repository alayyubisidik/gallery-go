package middleware

import "github.com/julienschmidt/httprouter"

func ChainMiddleware(handler httprouter.Handle, middlewares ...func(httprouter.Handle) httprouter.Handle) httprouter.Handle {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}