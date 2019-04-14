package main
import (
	"encoding/xml"
	"fmt"
    "io/ioutil"
    "sort"
//    "bufio"
//    "strings"
    "os"
//    "strconv"
)

// our struct which contains the complete
// array of all Users in the file
type Songs struct {
    XMLName xml.Name `xml:"songs"`
    Songs   []Song   `xml:"song"`
}

// the user struct, this contains our
// Type attribute, our user's name and
// a social struct which will contain all
// our social links
type Song struct {
    XMLName xml.Name `xml:"song"`
    Type    bool  `xml:"willPlay,attr"`
    Name    string `xml:",chardata"`
}

func main() {
    songNames := [] string {}
	xmlFile, err := os.Open("songs.xml")
    fmt.Print(xmlFile)
	byteValue, _ := ioutil.ReadAll(xmlFile)
    fmt.Print(err)

	var songs Songs
	xml.Unmarshal(byteValue, &songs)

    var m map[string]int
    m = make(map[string]int)

	for i := 0; i < len(songs.Songs); i++ {
//        fmt.Println("Song Type: " + songs.Song[i].Type)
        if songs.Songs[i].Type == true {
//            fmt.Println(songs.Songs[i])
            _, ok := m[songs.Songs[i].Name]
            if ok {
                //already in the list
                m[songs.Songs[i].Name] = m[songs.Songs[i].Name] + 1
            } else {
                //no key
                m[songs.Songs[i].Name] = 1
            }
            songNames = append(songNames, songs.Songs[i].Name + "\n")
        }
    }
    sort.Strings(songNames)

    fmt.Println(m)
}
