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
	http.HandleFunc("/filetype", uploadHandler.TypeFilesHandler())

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
		detectedFileType := http.DetectContentType(data)
		switch detectedFileType {
		case "image/jpeg", "image/jpg":
			fmt.Fprintf(w, "File has detectedFileType: %v, and can not be uploaded on server", detectedFileType)
			return
		case "text/plain", "text/plain; charset=utf-8":
			fmt.Fprintln(w, "File has detectedFileType:", detectedFileType)
		}
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
		fileLink := h.HostAddr + "/files/" + header.Filename
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
			content, _ := ioutil.ReadFile(h.UploadDir + "/" + f.Name())
			detectedFileType := http.DetectContentType(content)
			// fmt.Fprintln(w, "Name: ", f.Name(), " Size: ", f.Size(), "isDir: ", f.IsDir(), "Extension: ", filepath.Ext(string(f.Name())))
			fmt.Fprintf(w, "Name: %15v;	Size: %3v;	isDir: %5v;	Extension: %5v; detectedFileType: %3v; \n", f.Name(), f.Size(), f.IsDir(), filepath.Ext(string(f.Name())), detectedFileType)

		}
	})
}

func (h *FileHandler) TypeFilesHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		files, err := ioutil.ReadDir(h.UploadDir)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			content, _ := ioutil.ReadFile(h.UploadDir + "/" + f.Name())
			detectedFileType := http.DetectContentType(content)
			// fmt.Fprintln(w, "Name: ", f.Name(), " Size: ", f.Size(), "isDir: ", f.IsDir(), "Extension: ", filepath.Ext(string(f.Name())))
			fmt.Fprintf(w, "Name: %15v;	detectedFileType: %3v; \n", f.Name(), detectedFileType)

		}
	})
}
