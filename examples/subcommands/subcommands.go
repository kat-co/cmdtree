
package main

import (
    "fmt"
    "github.com/katco-/cmdtree"
)

func main() {
    const delimiter = " "
    cmd := cmdtree.NewCmd(delimiter, "help", func(arg string) error {
        fmt.Printf(`You requested help for "%s".`, arg)
        return nil
    })

    cmd.SubCmd("deep", func(arg string) error {
        fmt.Printf(`You requested DEEP help for "%s".`, arg)
        return nil
    })

    cmd.Execute("help cmdtree")
    fmt.Println()
    cmd.Execute("help deep cmdtree internals")
}
