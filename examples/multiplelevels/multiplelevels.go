
package main

import (
    "bytes"
    "fmt"
    "github.com/katco-/cmdtree"
)

func main() {
    const delimiter = " "
    help := cmdtree.NewCmd(delimiter, "help", nil)

    var with cmdtree.Command
    with = help.SubCmd("with", func(string) error {
        var buff bytes.Buffer
        fmt.Fprint(&buff, "Available options:")
        fmt.Fprintln(&buff)
        fmt.Fprintln(&buff, with.String())

        return fmt.Errorf(buff.String())
    })

    with.SubCmd("cmdtree", func(string) error {
        fmt.Println("It makes cupcakes! ...I think.")
        return nil
    })

    with.SubCmd("life", func(string) error {
        fmt.Println("Whoa. A code example is the wrong place for that, friend.")
        return nil
    })

    with.SubCmd("sleep", func(string) error {
        fmt.Println("Is that what I'm supposed to be doing right now?")
        return nil
    })

    if err := help.Execute("help with"); err != nil {
        fmt.Println(err)
    }

    fmt.Println("What does cmdtree do for me?")
    help.Execute("help with cmdtree")
}
