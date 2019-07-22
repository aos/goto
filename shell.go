package shell

import "fmt"

// Shell is the exported bash script that will integrate with our shell
type Shell string

func main() {
	fmt.Println(`hello
there
today`)
}
