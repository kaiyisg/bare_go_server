package api

import (
  "fmt"
)

type apiRoute struct {
	Method string
	Path   string
}

func (r apiRoute) Match(method, path string) bool {
  if r.Method == method && strings.Contains(path, r.Path) {
    return true
  }
  return false
}
