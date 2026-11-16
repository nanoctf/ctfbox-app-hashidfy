package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"hash-identify/detector"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// If a hash is supplied as an argument, process it first.
	first := ""
	if len(os.Args) > 1 {
		first = os.Args[1]
	}

	for {
		fmt.Println(strings.Repeat("-", 50))

		hash := first
		first = ""

		if hash == "" {
			fmt.Print(" HASH: ")
			line, err := reader.ReadString('\n')
			if err != nil && len(line) == 0 {
				return // EOF (Ctrl+D)
			}
			hash = strings.TrimSpace(line)
		}

		if hash == "" {
			continue
		}

		var ids []string
		detector.Detect(hash, &ids)

		if len(ids) == 0 {
			fmt.Println("\n Not Found.")
			continue
		}

		// Match Python's jerar.sort(): IDs are sorted as strings.
		sort.Strings(ids)

		if len(ids) > 2 {
			fmt.Println("\nPossible Hashs:")
			fmt.Println("[+] " + detector.NameForID(ids[0]))
			fmt.Println("[+] " + detector.NameForID(ids[1]))
			fmt.Println("\nLeast Possible Hashs:")
			for _, id := range ids[2:] {
				fmt.Println("[+] " + detector.NameForID(id))
			}
		} else {
			fmt.Println("\nPossible Hashs:")
			for _, id := range ids {
				fmt.Println("[+] " + detector.NameForID(id))
			}
		}
	}
}
