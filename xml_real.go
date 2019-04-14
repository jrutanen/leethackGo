package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

//struct for songs
type Songs struct {
	XMLName xml.Name `xml:"songs"`
	Songs   []Song   `xml:"song"`
}

//struct for one song
type Song struct {
	XMLName xml.Name `xml:"song"`
	Type    bool     `xml:"willPlay,attr"`
	Name    string   `xml:",chardata"`
}

func main() {
	//create a slice for songNames
	songNames := []string{}
	xmlFile, err := os.Open("songs.xml")
	fmt.Print(err)
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	//put values from xml file to the struct
	var songs Songs
	xml.Unmarshal(byteValue, &songs)

	//create map so we can keep track of individual songs and
	//how many times they appear in the list
	var m map[string]int
	m = make(map[string]int)

	for i := 0; i < len(songs.Songs); i++ {
		//if willPlay is true
		if songs.Songs[i].Type == true {
			//check if song is already added to the map
			_, ok := m[songs.Songs[i].Name]
			if ok {
				//already in the map, increase number of occurences
				m[songs.Songs[i].Name] = m[songs.Songs[i].Name] + 1
			} else {
				//not in the map, add to map and set number of occurences to 1
				m[songs.Songs[i].Name] = 1
			}
		}
	}
	//add songs in the map to the songNames slice
	for key, value := range m {
		songNames = append(songNames, key+","+strconv.Itoa(value)+"\n")
	}
	//sort the slice of songs and ignore case
	IgnoreCase := func(i, j int) bool {
		return strings.ToLower(songNames[i]) < strings.ToLower(songNames[j])
	}

	sort.SliceStable(songNames, IgnoreCase)
	fmt.Println(songNames)
}
