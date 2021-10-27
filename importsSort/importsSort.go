package importsSort

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func SortImports(path string) {
	// opening file, receiving contents as byte slice
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// splitting content by lines, parsing imports into slice and index of lines at start and end of imports
	var imports []string
	importsStarted := false
	importsStart, importsEnd := 0, 0
	lines := strings.Split(string(content), "\n")
	for index, line := range lines {
		if importsStarted {
			if line == ")" {
				importsStarted = false
				importsEnd = index - 1
				break
			} else {
				imports = append(imports, line)
			}
		}
		if strings.Contains(line, "import") {
			importsStarted = true
			importsStart = index + 1
		}
	}

	// sorting imports alphabetically
	if len(imports) == 0 {
		return
	}
	sort.Strings(imports)

	// replacing imports with sorted
	for i, j := importsStart, 0; i <= importsEnd; i, j = i + 1, j + 1 {
		lines[i] = imports[j]
	}

	// writing edited lines to the file
	formatted := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(formatted), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

