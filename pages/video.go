package pages

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var file = "/movie.mp4"
var defaultChunkSize int64 = 32554430 //less than 32MB
var fileSize int64 = 0

func init() {
	info, err := os.Stat(file)
	if err != nil {
		fmt.Println("Unable to access the video file")
		os.Exit(1)
	}
	fileSize = info.Size()
}

func init() {
	http.HandleFunc("/video", VideoStream)
}

//Serve the video file in chucks that are smaller than 32MB starting at the offset determined by the client
func VideoStream(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodHead {
		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Content-Type", "video/mp4")
		w.WriteHeader(http.StatusOK)
	} else if r.Method == http.MethodGet {
		rangeHeader := r.Header.Get("range")
		if rangeHeader == "" {
			rangeHeader = "bytes=0-"
		}

		parts := strings.Split(strings.Replace(rangeHeader, "bytes=", "", 1), "-")
		start, _ := strconv.ParseInt(parts[0], 10, 64)
		end := start + defaultChunkSize
		if end > fileSize {
			end = fileSize - 1
		}

		stream, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		defer func() { _ = stream.Close() }()

		chunkSize := (end - start) + 1
		contentRange := fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize)
		section := io.NewSectionReader(stream, start, chunkSize)

		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Content-Range", contentRange)
		w.Header().Set("Content-Length", strconv.FormatInt(chunkSize, 10))
		w.Header().Set("Content-Type", "video/mp4")
		w.WriteHeader(206)

		_, _ = io.Copy(w, section)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}