package main

import(
"fmt"
"runtime"
"strconv"
"sync"
)

func main(){
runtime.GOMAXPROCS(4);
var wait sync.WaitGroup;
wait.Add(4);
var str1 string;
var str2 string;
var str3 string;
var str4 string;
var answer int;
fmt.Scanf("%d",&answer);
if(answer>100000){
return;
}
if(answer<=0){
return;
}
if(answer<=3){
for i:=1;i<=answer;i++{
fmt.Println(i);
}
return;
}

go func(start int,end int,str *string){
defer wait.Done();
for i:=start;i<end;i++{
*str=*str+strconv.Itoa(i)+"\n";
}
}(1,(answer/4),&str1);
go func(start int,end int,str *string){
defer wait.Done();
for i:=start;i<end;i++{
*str=*str+strconv.Itoa(i)+"\n";
}
}((answer/4),(answer/4)*2,&str2);
go func(start int,end int,str *string){
defer wait.Done();
for i:=start;i<end;i++{
*str=*str+strconv.Itoa(i)+"\n";
}
}((answer/4)*2,(answer/4)*3,&str3);
go func(start int,end int,str *string){
defer wait.Done();
for i:=start;i<end;i++{
*str=*str+strconv.Itoa(i)+"\n";
}
}((answer/4)*3,answer+1,&str4);
wait.Wait();
fmt.Printf("%s",str1+str2+str3+str4);
}//main

