package main

import (
  "fmt"
  "flag"
  "time"
  "io/ioutil"
  "encoding/xml"
)

// Activities -> Activity -> Lap [props] -> Track -> Trackpoint
type Trackpoint struct {
  Time string
  HeartRateBpm int `xml:"HeartRateBpm>Value"`
  Latitude float64 `xml:"Position>LatitudeDegrees"`
  Longitude float64 `xml:"Position>LongitudeDegrees"`
}

type Lap struct {
  Calories int
  Intensity string
  Trackpoints []Trackpoint `xml:"Track>Trackpoint"`
}

type Activity struct {
  Id string  `xml:"Activities>Activity>Id"`
  Laps []Lap `xml:"Activities>Activity>Lap"`
}

func main() {
  // parse cli args <== like a boss.
  var file_name string
  flag.StringVar(&file_name, "file", "", "some usage here.")
  flag.Parse()
  if file_name == "" {
    panic("GO FIGURE: I cannot function without a filename.")
  }

  // read file contents <== I'm doing it like a boss.
  bs, err := ioutil.ReadFile(file_name)
  if err != nil {
    panic(fmt.Sprintf("GO FIGURE: I can't read the file `%s': `%v'.", file_name, err))
  }

  // unmarshal the xml out of it <== like a boss.
  activity := Activity{}
  err = xml.Unmarshal(bs, &activity)
  if err != nil {
    panic(fmt.Sprintf("GO FIGURE: I cannot unmarshall like a boss: `%v'", err))
  }

  k := 0

  fmt.Println("date\thr")

  for i := 0; i < len(activity.Laps); i++ {
    // inception... a loop inside a loop.
    for j:= 0; j < len(activity.Laps[i].Trackpoints); j++ {
      time, err := time.Parse(time.RFC3339, activity.Laps[i].Trackpoints[j].Time)
      if err != nil {
        fmt.Println("Most likely we failed to parse the time.", err)
      }
      fmt.Printf("%d\t%d\n", time.Unix(), activity.Laps[i].Trackpoints[j].HeartRateBpm)
      k += 1
      if k > 200 {
        return
      }
    }
  }
}
