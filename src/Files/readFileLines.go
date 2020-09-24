package Files

import _ "fmt"
import "os"
import "bufio"

func main() {
	fin, err := os.Open("temp.dat")
	if err != nil { panic(err) }
	defer fin.Close()

	fout, err := os.Create("temp3.dat")
	if err != nil { panic(err) }
	defer fout.Close()

	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		s := scanner.Text()
		fout.Write([]byte(s + "\n"))
	}

	//f, err := os.OpenFile("temp4.dat", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
}
