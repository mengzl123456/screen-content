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
	separator := make([]string, 1, 2)
	separator[0] = form.Get("firstSeparator")
	secondSeparator := form.Get("secondSeparator")
	if secondSeparator != "" {
		separator = append(separator, secondSeparator)
	}
	_, err = io.WriteString(w, "<a href='/'>返回</a>")
	w.(http.Flusher).Flush()
	_, err = io.WriteString(w, "<br>正在处理中请稍等</br>")
	w.(http.Flusher).Flush()
	if err != nil {
		log.Fatal(err)
	}
	err = screenContent(path, savePath, separator)
	if err != nil {
		_, _ = fmt.Fprintf(w, "<br>文件错误</br>")
		w.(http.Flusher).Flush()
		log.Fatal(err)
		return
	}
	_, err = io.WriteString(w, "<br>处理完成</br>")
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
			savePathTemp := savePath + string(filepath.Separator) + directory[i].Name()
			pathTemp := path + string(filepath.Separator) + directory[i].Name()
			_ = os.Mkdir(savePathTemp, os.ModePerm)
			err = screenContent(pathTemp, savePathTemp, separator)
			if err != nil {
				return err
			}
		} else {
			file, err := os.Open(path + string(filepath.Separator) + directory[i].Name())
			if err != nil {
				return err
			}
			content, err := ioutil.ReadAll(file)
			if err != nil {
				return err
			}
			if len(separator) == 1 {
				fmt.Println(screenContentByOneSeparator(string(content), separator[0]))
				err = ioutil.WriteFile(savePath+string(filepath.Separator)+directory[i].Name(), []byte(screenContentByOneSeparator(string(content), separator[0])), 0777)
				if err != nil {
					return err
				}
			} else {
				fmt.Println(screenContentByTwoSeparator(string(content), separator))
				err = ioutil.WriteFile(savePath+string(filepath.Separator)+directory[i].Name(), []byte(screenContentByTwoSeparator(string(content), separator)), 0777)
				if err != nil {
					return err
				}
			}
			err = file.Close()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func screenContentByOneSeparator(content string, separator string) string {
	contentArr := strings.Split(content, separator)
	return contentArr[1]
}

func screenContentByTwoSeparator(content string, separator []string) string {

	contentArr := strings.Split(content, separator[0])
	for _, contentTmp := range contentArr {
		if strings.Contains(contentTmp, separator[1]) {
			return strings.Split(contentTmp, separator[1])[0]
		}
	}

	return ""
}
