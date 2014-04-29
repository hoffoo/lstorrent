// Utility to list torrent file contents
package main

import (
    "fmt"
    "github.com/marksamman/bencode"
    "os"
)

func main() {

    if len(os.Args) < 2 {
        fmt.Printf("Usage: lstorrent [files]")
        return
    }

    for _, f := range os.Args[1:] {
        lstorrent(f)
    }
}

func lstorrent(path string) {
    fmt.Printf("%s\n", path)
    defer fmt.Printf("\n\n")
    f, err := os.OpenFile(path, os.O_RDONLY, 0600)


    if err != nil {
        fmt.Printf("Failed opening file: ", err)
        return
    }
    defer f.Close()

    datb, err := bencode.Decode(f)
    if err != nil {
        fmt.Printf("Failed decoding:", err)
        return
    }

    files := datb["info"].(map[string]interface{})["files"].([]interface{})

    for _, v := range files {
        file := v.(map[string]interface{})
        path := file["path"].([]interface{})[0].(string)
        fmt.Printf("%s\n", path)
    }
}
