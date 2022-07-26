package main

import ("fmt")

func main() {

  adj := [2]string{"big", "tasty"}
  
  fruits := [3]string{"apple", "orange", "banana"}
  
  for i:=0; i < 2; i++ { //len(adj)
   fmt.Println(adj[i])
   // for j:=0; j <3 ; j++ { //len(fruits)
    
      //fmt.Println(adj[i],fruits[j])
   // }
   
  }
  for j:=0; j <3 ; j++{
  
   fmt.Println(fruits[j])
  }
} 
