package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, header, _ := r.FormFile("file")
		file, _ := header.Open()
		path := fmt.Sprintf("%s", header.Filename)
		fmt.Sprintf("upload: %s\n", path)
		buf, _ := ioutil.ReadAll(file)
		ioutil.WriteFile(path, buf, 0644)
		http.Redirect(w, r, "/"+path, 301)
	} else {
		http.Redirect(w, r, "/", 301)
	}
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
    <head>
        <meta charset="utf-8"></meta>
        <title>Uploader</title>
    </head>
    <body>
        <form action="/upload" method="post" enctype="multipart/form-data">
            <input type="file" id="file" name="file">
            <input type="submit" name="submit" value="upload">
        </form>
    </body>
</html>`)
}
func main() {
	staticServer := http.StripPrefix("/files/", http.FileServer(http.Dir("files/")))
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	http.Handle("/files/", staticServer)
	fmt.Printf("Listen port :5000\n")
	panic(http.ListenAndServe(":5000", nil))
}
