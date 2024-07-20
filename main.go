package main

import (
  "fmt"
  "os"
)



func dir_walker( path *string , filename *string){


  files, err := os.ReadDir(*path);
  if err != nil {
    fmt.Println(err)
    return
  }

  for _, file := range files {
    real_path := *path + "/" + file.Name()
    _, err := os.ReadDir(real_path)
    if err != nil {

      if file.Name() == *filename{
        fmt.Println("FOUND File : " , real_path);
      }
    }else {
      if *filename == file.Name(){
        fmt.Println("FOUND Dir : " , real_path);
      }


      dir_walker(&real_path , filename);

    }
  }

}

func main() {
  var dir string
  fmt.Println("insert dir path : ");
  fmt.Scanln(&dir);

  var filename string
  fmt.Println("insert name of the file : ")
  fmt.Scanln(&filename)




  dir_walker( &dir , &filename);



}
