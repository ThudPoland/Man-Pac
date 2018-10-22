package level

import (
	"bufio"
	"io"
	"os"
)

type Data struct {
	lines []string
}

//Load function for level
func (data Data) Load(path string) {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		lines = nil
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == nil {
			lines = append(lines, line)
			continue
		} else if nil == io.EOF {
			break
		}
	}

	data.lines = lines
}
