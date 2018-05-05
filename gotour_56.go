package main

import (
    "fmt"
    "math"
)
type ErrNegativeSqrt float64
    
func (f ErrNegativeSqrt) Error() string{
    return fmt.Sprintf("%s:%f","cannot Sqrt negative number",f);
}
func Sqrt(f float64) (float64, error) {
    var err error;
    var result float64;
    if(f<0){
        result=f
        err=ErrNegativeSqrt(f);  
    }else{
        result=math.Sqrt(f)
        err=nil
    }
    return result,err
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
