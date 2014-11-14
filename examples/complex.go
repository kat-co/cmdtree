
package main

import (
    "fmt"
    "github.com/katco-/cmdtree"
    "strconv"
)

type User struct {
    Name                  string
    LevelOfLoveForCmdtree int
}

func main() {

    var currentUser *User
    users := []*User{
        &User{"Wirt. Just Wirt.", 0},
        &User{"Greg the frog catcher", 100},
        &User{"Beatrice the Bluebird", 5},
    }

    root := cmdtree.Root(" ")

    set := root.SubCmd("set", nil)
    set.SubCmd("love", setLoveForUserFn(&currentUser))

    print := root.SubCmd("print", nil)
    print.SubCmd("users", func (name string) error{
        fmt.Println("Users:")
        for _, user := range users {
            if name != "" && user.Name != name {
                continue
            }

            fmt.Printf(`"%s" loves cmdtree %d%%!`, user.Name, user.LevelOfLoveForCmdtree)
            fmt.Println()
        }
        fmt.Println()
        return nil
    })

    root.Execute("print users")

    for _, user := range users {
        currentUser = user
        root.Execute("set love 100")
    }

    root.Execute("print users Wirt. Just Wirt.")
    root.Execute("print users")
}

func setLoveForUserFn(user **User) cmdtree.CommandExecutor {
    return func(level string) error {
        numericLevel, err := strconv.Atoi(level)
        if err != nil {
            return err
        } else if numericLevel < 0 {
            return fmt.Errorf("I'm sorry %s, I can't do that.", (*user).Name)
        }

        (*user).LevelOfLoveForCmdtree = numericLevel
        return nil
    }
}
