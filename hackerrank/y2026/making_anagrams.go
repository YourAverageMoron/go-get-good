package y2026


func MakeAnagram(a string, b string) int32 {

    ma := make(map[rune]float32, len(a))
    mb := make(map[rune]float32, len(b))

    for _, c := range a {
        ma[c] ++
    }

    counter := int32(0)
    for _, c := range b {
        mb[c] ++
        if ma[c] > 0 {
            ma[c] --
        } else {
            counter ++
        }
    }

    for _, c := range a {
        if mb[c] > 0 {
            mb[c] --
        } else {
            counter ++
        }
    }

    return counter
}
