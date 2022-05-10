package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type FileHandler struct {
	HostAddr  string
	UploadDir string
}

func main() {
	uploadHandler := &FileHandler{UploadDir: "./upload", HostAddr: "http://localhost:8080"}
	http.HandleFunc("/upload", uploadHandler.uploadFileHandler())
	http.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir(uploadHandler.UploadDir))))
	http.HandleFunc("/list", uploadHandler.ListFilesHandler())

	log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (h *FileHandler) uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Unable to read file", http.StatusBadRequest)
			return
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Unable to read file", http.StatusBadRequest)
			return
		}
		filePath := h.UploadDir + "/" + header.Filename
		err = ioutil.WriteFile(filePath, data, 0777)
		if err != nil {
			log.Println(err)
			http.Error(w, "Unable to save file", http.StatusInternalServerError)
			return
		}
		fileLink := h.HostAddr + "/" + header.Filename
		fmt.Fprintln(w, fileLink)
		fmt.Fprintf(w, "File %s has been successfully uploaded", header.Filename)
	})
}

func (h *FileHandler) ListFilesHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		files, err := ioutil.ReadDir(h.UploadDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			//fmt.Fprintln(w, "Name: ", f.Name(), " Size: ", f.Size(), "isDir: ", f.IsDir(), "Extension: ", filepath.Ext(string(f.Name())))
			fmt.Fprintf(w, "Name: %15v Size: %8v isDir: %8v Extension: %8v \n", f.Name(), f.Size(), f.IsDir(), filepath.Ext(string(f.Name())))

		}
	})
}
