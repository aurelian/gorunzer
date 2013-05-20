package main

import (
  "fmt"
  "flag"
  "io/ioutil"
  "encoding/xml"
  "encoding/json"
)

// Activities -> Activity -> Lap [props] -> Track -> Trackpoint
type Point struct {
  Time      string
  HeartRate int     `xml:"HeartRateBpm>Value"`
  Altitude  float64 `xml:"AltitudeMeters"`
  Distance  float64 `xml:"DistanceMeters"`
  Latitude  float64 `xml:"Position>LatitudeDegrees"`
  Longitude float64 `xml:"Position>LongitudeDegrees"`
}

type Lap struct {
  Calories int
  Intensity string
  Points []Point `xml:"Track>Trackpoint"`
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

  // and marshall it to json.
  json, err := json.Marshal(activity)
  if err != nil {
    panic(fmt.Sprintf("GO FIGURE: I cannot marshall like a boss: `%v'", err))
  }

  fmt.Printf("%s", json)

}
