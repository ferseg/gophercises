package urlshort

import (
	"gopkg.in/yaml.v3"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this.
	return func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		path, ok := pathsToUrls[uri]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, path, http.StatusPermanentRedirect)
	}
}

type ShortenerInfo struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
  var paths []ShortenerInfo
  if err := yaml.Unmarshal(yml, &paths); err != nil {
    return nil, err
  }
  pathsMap := make(map[string]string, len(paths))
  for _, path := range paths {
    pathsMap[path.Path] = path.URL
  }
	return MapHandler(pathsMap, fallback), nil
}
