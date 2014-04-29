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
    defer fmt.Printf("\n")
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

    tor_info, ok := datb["info"].(map[string]interface{})
    if !ok {

    }

    // first look for files key
    files, ok := tor_info["files"].([]interface{})
    if ok {
        for _, v := range files {
            file := v.(map[string]interface{})
            path := file["path"].([]interface{})[0].(string)
            fmt.Printf("%s\n", path)
        }

        return // finished printing multifile torrent
    }

    // second look for name key
    file, ok := tor_info["name"].(string)
    if ok {
        fmt.Printf("%s\n", file)
        return
    }

    fmt.Printf("couldnt find files in torrent: %s", path)
}
