package y2026


func AlternatingCharacters(s string) int32 {
     
    var prev rune
    del := int32(0)
    for _, c := range s {
        if c == prev {
            del ++
        }  
        prev = c
    }
    return del
}
