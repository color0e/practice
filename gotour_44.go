package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    count:=0;
    m:=make(map[int]int);
    return func( ) int{
        if(count<=1){
            m[count]=1;
        }else{
            m[count]=m[count-2]+m[count-1];            
        }
        result:=m[count];
        count=count+1;
       return result;
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
