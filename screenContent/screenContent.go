package screenContent

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ScreenContentBySeparator(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	form := r.Form
	path := form.Get("path")
	savePath := form.Get("savePath")
	separator := make([]string, 2)
	separator[0] = form.Get("firstSeparator")
	separator[1] = form.Get("secondSeparator")
	err = screenContent(path, savePath, separator)
	if err != nil {
		_, _ = fmt.Fprintf(w, "文件错误")
		w.(http.Flusher).Flush()
		_, _ = fmt.Fprintf(w, "文件错误")
		return
	}
	_, err = io.WriteString(w, "正在生成文件请稍等")
	if err != nil {
		log.Fatal(err)
	}

}

func screenContent(path string, savePath string, separator []string) error {
	directory, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for i := 0; i < len(directory); i++ {
		if directory[i].IsDir() {
			savePath += string(filepath.Separator) + directory[i].Name()
			path += string(filepath.Separator) + directory[i].Name()
			_ = os.Mkdir(savePath, os.ModePerm)
			_ = screenContent(path, savePath, separator)
		} else {
			file, err := os.Open(path + string(filepath.Separator) + directory[i].Name())
			if err != nil {
				return err
			}
			content, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}

			fmt.Println(screenContentByTwoSeparator(string(content), separator[0]))
		}
	}
	return nil
}

func screenContentByOneSeparator(content string, separator string) string {
	contentArr := strings.Split(content, separator)
	return contentArr[1]
}

func screenContentByTwoSeparator(content string, separator string) *string {

	contentArr := strings.Split(content, separator)
	for _, contentTmp := range contentArr {
		if strings.Contains(contentTmp, separator) {
			return &strings.Split(contentTmp, separator)[0]
		}
	}

	return nil

}
