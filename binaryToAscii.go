package main
import (
	"fmt"
    "bufio"
//    "strings"
    "os"
    "strconv"
)

func main() {
//    numbers := [] string {}
	binaryNumbers, err := os.Open("bin.txt")
    fmt.Print(err)

    scanner := bufio.NewScanner(binaryNumbers)
//    b := make([]byte, 0)
    for scanner.Scan() {
//        fmt.Println(scanner.Text())
        //convert binary text to ascii
//        binString = fmt.Sprintf("%s%b",scanner.Text(), c)
        n, _ := strconv.ParseUint(scanner.Text(), 2, 8)
//        b = append(b, byte(n))
//        i := strconv.Itoa(int(n))
        fmt.Printf("%c", n)
    }
}
