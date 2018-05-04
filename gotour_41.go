package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m:=make(map[string]int);
    str:=strings.Fields(s);
    for i:=0;i<len(str);i++{
        _,ok:=m[str[i]];
     if ok==false{
     m[str[i]]=1;
     }else{
     m[str[i]]=m[str[i]]+1;
     }
        
    }
    return m;
}

func main() {
    wc.Test(WordCount)
}
