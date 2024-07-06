package main

func PassThePillowSolution(n int, time int) int {
    number_passes := time / (n-1)
    remaider := time % (n -1 )
    if (number_passes % 2 == 0 ){
        return remaider + 1
    }
    return n - remaider
}
