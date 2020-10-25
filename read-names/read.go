package nreader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Name struct for reading from a file
type Name struct {
	fname string
	lname string
}

func trimEndings(s *string) string {
	return strings.TrimRight(strings.TrimRight(*s, "\n"), "\r")

}

func nreader() {

	var names []Name
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Program to read names from a given file.")
	fmt.Println("Assuming each line has single space-separated first and last names (ex: John Smith).")
	fmt.Println("\nEnter a file name:")

	if fileName, err := in.ReadString('\n'); err == nil {
		if file, err := os.Open(trimEndings(&fileName)); err == nil {

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				if name := strings.Split(scanner.Text(), " "); len(name) < 2 {
					log.Fatalf("Line %v should have at least 2 words\n", name)
				} else {
					names = append(names, Name{fname: name[0], lname: name[1]})
				}
			}
			file.Close()
		} else {
			log.Fatalf("Error opening file name - %v\n", err)
		}

	} else {
		log.Fatalf("Error reading file name - %v\n", err)
	}

	fmt.Println("\nfname | lname\n-------------")
	for _, v := range names {
		fmt.Println(v.fname, v.lname)
	}
}
``