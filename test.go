package main

import ( 
    "fmt"
    "math/rand"
    "flag"
    "time"
)

var userNumb = flag.Int("start", 0, "?")
var userCount int

var randd = rand.Intn(100)

func main() {
    
    flag.Parse() 
    Sorter(randd, userNumb)
    Ggame(userCount, *userNumb) 
}

func Ggame(userCount int, userNumb int) {
    your_time := time.Now()
    for i := 0; i < userNumb; i++ {     
    fmt.Scan(&userCount)
        if randd == userCount {
            fmt.Println("вы угадали число")
            fmt.Println(time.Since(your_time))
            break
        } else { 
              fmt.Println("Вы не угадали число")
    }
}
}

func Sorter(rand_numb int) {
    un := userNumb
    if un == 3 {
    start := (rand_numb / 10) * 10
    end := + 9
    fmt.Println(start, "-", end)
    }
}
