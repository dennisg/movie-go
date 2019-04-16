package pages

import "net/http"

var index = []byte(`<html style="border:0">
<video id="videoPlayer" controls autoplay style="width: 100%">
	<source src="/video" type="video/mp4">
</video>
</html>
`)

func init() {
	http.HandleFunc("/", IndexPage)
}

func IndexPage(w http.ResponseWriter, r *http.Request)  {
	_, _ = w.Write(index)
}