http.Handler is an interface with the ServeHTTP method

http.HandlerFunc is function type that accepts the same args as ServeHTTP method.
It also implements http.Handler

http.Handle is a function which takes a pattern or string and a http.Handler

http.Handle("/", http.Handler)
http.HandleFunc("/", pathHandler)