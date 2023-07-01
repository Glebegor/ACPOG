package creating

type WorkWithElements struct {
	name  string
	path  string
	dirs  []string
	files []string
	data  []string
}

func NewWorkWithElements(nameApp string, pathApp string, dirsApp []string, filesApp []string, dataApp []string) *WorkWithElements {
	return &WorkWithElements{
		name:  nameApp,
		path:  pathApp,
		dirs:  dirsApp,
		files: filesApp,
		data:  dataApp,
	}
}
