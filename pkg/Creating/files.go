package creating

import (
	"fmt"
	"os"
)

func (w *WorkWithElements) CreateFile(name string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(f.Name())
	f.Close()
}
func (w *WorkWithElements) CreateFiles(workWithElems WorkWithElements, myAppName string, path string) {
	for i := 0; i < len(w.files); i++ {
		workWithElems.CreateFile(fmt.Sprintf("%s%s", myAppName, w.files[i]))
	}
}
