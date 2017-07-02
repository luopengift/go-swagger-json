package swagger

import (
    "testing"
    "encoding/json"
    "fmt"
)


func Test_swagger(t *testing.T) {
    sw := InitSwagger()
    b, err := json.Marshal(sw)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(b))
}
