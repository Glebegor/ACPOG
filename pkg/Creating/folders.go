package creating

import (
	"fmt"
	"os"
)

func (w *WorkWithElements) CreateFolder(name string) {
	if err := os.Mkdir(name, os.ModePerm); err != nil {
		fmt.Println(err.Error())
	}
}
func (w *WorkWithElements) CreateArch(workWithElems WorkWithElements, myAppName string, path string) {
	for i := 0; i < len(w.dirs); i++ {
		workWithElems.CreateFolder(fmt.Sprintf("%s%s", myAppName, w.dirs[i]))
	}
}
