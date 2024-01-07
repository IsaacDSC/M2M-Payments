package shared

import "net/http"

type Path string
type Handler func(w http.ResponseWriter, r *http.Request)
type TypeRouters map[Path]Handler
