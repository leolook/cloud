package fileserver

import "net/http"

type FileServer struct {
	Upload string
	Remove string
}

func (f *FileServer) DoUpload(w http.ResponseWriter, req *http.Request) {
	Upload(w, req)
}

func (f *FileServer) DoRemove(w http.ResponseWriter, req *http.Request) {
	Remove(w, req)
}
