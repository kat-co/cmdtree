
package main

import (
    "github.com/katco-/cmdtree"
    "fmt"
)

func main() {
    const delimiter = " "
    cmd := cmdtree.NewCmd(delimiter, "help", func(arg string) error {
        fmt.Printf(`You requested help for "%s".`, arg)
        fmt.Println()
        return nil
    })

    cmd.Execute("help cmdtree")
}
