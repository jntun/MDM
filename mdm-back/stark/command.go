package stark

import (
    "bufio"
    "fmt"
    "os"
)

type Command struct {
        cmd   string
        vargs []string
        act   func(vargs []string) error
}

func PopInterpreter() {
    for true {
        cmd, err := getNextInput()
        if err != nil {
            StarkError("[Intptr]: %s", err)
        }
        StarkLog("[Intptr]: %s", cmd)
    }
}

func getNextInput() (*Command, error) {
        reader := bufio.NewReader(os.Stdin)
        for true {
            fmt.Print(">")
            text, err := reader.ReadString('\n')
            if err != nil {
                StarkError("[Input]: %s", err)
            }
            StarkLog("[Input] %s", text)
        }

        var cmd *Command
        return  cmd, nil
}
