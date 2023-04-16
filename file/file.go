package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)



func ReadJson  [T any](file string, obj T) T{

    jsonFile, err := os.Open(file)

    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf( "Successfully Opened %s", file)
    fmt.Println()

    defer jsonFile.Close()


    byteValue, _ := ioutil.ReadAll(jsonFile)

    json.Unmarshal(byteValue, &obj)

    return obj
}