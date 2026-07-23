package y2026

import "fmt"


func RansomNotes(magazine []string, note []string){
    m := make(map[string]int32, len(magazine))

    for _, w := range magazine {
        m[w] ++
    }

    for _, w := range note {
        if m[w] <= 0 {
            fmt.Println("No")
            return 
        } 
        m[w] --
    }

    fmt.Println("Yes")
}
