package main

import "fmt"

// Bash is the exported bash script that will integrate with our shell
const Bash = `# To make the binary work, add the following lines of code
# to your ~/.bash_profile or ~/.bash_rc
#
# eval $(goto --init)
#
# It will autogenerate this text and below to make the magic happen.
# "test -d"


g() {
  local dir="$(goto $@)"
  test -d "$dir" && "cd $dir" || echo "$dir"
}`

// PrintShellIntegration prints out the shell integration string
func PrintShellIntegration(b string) {
	fmt.Println(b)
}
