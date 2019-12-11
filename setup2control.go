package main

import(
  "os"
  "strings"
//  "io/ioutil"
)


//Method used to catch error
func check(e error) {
  if e != nil {
      panic(e)
  }
}


//Method to remove a byte into a byte slice
func deleteByte(input string, rmbyte byte) string {
  var out = []byte{}
    for _, byte := range []byte(input){
      if byte != rmbyte {
        out = append(out, byte)
    }
  }
  return string(out)
}



//Main Code
func main() {
  //Create Control file
  f, err := os.Create("control")
  check(err)
  defer f.Close()
  //Add immutable fields
  _, err = f.WriteString("Architecture: amd64\n")
  check(err)
  _, err = f.WriteString("Essential: no\n")
  check(err)
  _, err = f.WriteString("Priority: optional\n")
  check(err)
    _, err = f.WriteString("Depends: python3-pip, python3-venv, systemd, mosquitto\n")
  check(err)
  //Open setup.py file
  file, err := os.Open("setup.py")
  check(err)
  //Read data from setup.py
  data := make([]byte, 3200)
  _, err = file.Read(data)
  check(err)
  //Splits data at each ","
  splits := strings.Split(string(data), "(")
  //Use data on the right of the (
  splits = strings.Split(string(splits[1]), ",")
  for _,value := range(splits) {
    //Splits data at each "="
    var splitvalue = strings.Split(string(value), "=")
    //Clean output
    out := deleteByte(deleteByte(splitvalue[0],32),10)
    switch string(out) {
    case "name" :
      _, err := f.WriteString("Package: "+deleteByte(splitvalue[1],39)+"\n")
      check(err)
    case "version" :
      _, err := f.WriteString("Version: "+deleteByte(splitvalue[1],39)+"\n")
      check(err)
    case "long_description" :
      _, err := f.WriteString("Description: "+deleteByte(splitvalue[1],39)+"\n")
      check(err)
    case "author" :
      _, err := f.WriteString("Maintainer: "+deleteByte(splitvalue[1],39)+"\n")
      check(err)
    }
  }
}
