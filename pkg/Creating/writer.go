package creating

import (
	"fmt"
	"os"
)

func (w *WorkWithElements) WriteInfoToFiles(file string, data string) {
	f, err := os.Create(file)
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = f.Write([]byte(data))
	if err != nil {
		fmt.Print(err.Error())
	}
	f.Close()
	fmt.Println("done")
}

func (w *WorkWithElements) WriteInfoForFiles(workWithElems WorkWithElements, myAppName string, path string) {
	for i := 0; i < len(w.files); i++ {
		workWithElems.WriteInfoToFiles(fmt.Sprintf("%s%s", myAppName, w.files[i]), w.data[i])
	}
}
