package main

import (
     "strings"
     "fmt"
)

var txt = string("There have been some complaints about the fact that, whether it is true or not, we first look at the driver's identity, 0.2% who has been registered for several years. Then also registered the Hitch, completed more than 3,000 singles a year was complained, four times, it seems that more than three thousand single inside nearly a year behind four is not particularly many, this driver is OK, So let's judge, alas, this is the first person who matches the car, and the driver is also a better driver, no problem, that is What? What happened to this case? What are we analyzing about the passengers? The passenger may have a point that he is a victim of a crowd, why? Because, first of all, it was a case we had, and then we got a lot of information, and he was a drunk passenger, and he went alone, hit a car, and then, Actually, the passenger was a woman, but our system identified it backstage, he was a man, and then, uh, he was probably, uh, you know, you see his usual list, That is, he has completed a total of fifteen orders. Most of them are fixed to a place called the International Mansion in France around six in the evening, and then what?")


func main () {
   fmt.Println("raw: ", txt)

   strs := strings.SplitAfter(txt, ".")
   for i, s := range strs {
       fmt.Println(i, ":", s)
   }

   fmt.Println()
   strs = strings.SplitAfter(txt, ". ")
   for i, s := range strs {
       fmt.Println(i, ":", s)
   }

}

/*

raw:  There have been some complaints about the fact that, whether it is true or not, we first look at the driver's identity, 0.2% who has been registered for several years. Then also registered the Hitch, completed more than 3,000 singles a year was complained, four times, it seems that more than three thousand single inside nearly a year behind four is not particularly many, this driver is OK, So let's judge, alas, this is the first person who matches the car, and the driver is also a better driver, no problem, that is What? What happened to this case? What are we analyzing about the passengers? The passenger may have a point that he is a victim of a crowd, why? Because, first of all, it was a case we had, and then we got a lot of information, and he was a drunk passenger, and he went alone, hit a car, and then, Actually, the passenger was a woman, but our system identified it backstage, he was a man, and then, uh, he was probably, uh, you know, you see his usual list, That is, he has completed a total of f
ifteen orders. Most of them are fixed to a place called the International Mansion in France around six in the evening, and then what?
0 : There have been some complaints about the fact that, whether it is true or not, we first look at the driver's identity, 0.
1 : 2% who has been registered for several years.
2 :  Then also registered the Hitch, completed more than 3,000 singles a year was complained, four times, it seems that more than three thousand single inside nearly a year behind four is not particularly many, this driver is OK, So let's judge, alas, this is the first person who matches the car, and the driver is also a better driver, no problem, that is What? What happened to this case? What are we analyzing about the passengers? The passenger may have a point that he is a victim of a crowd, why? Because, first of all, it was a case we had, and then we got a lot of information, and he was a drunk passenger, and he went alone, hit a car, and then, Actually, the passenger was a woman, but our system identified it backstage, he was a man, and then, uh, he was probably, uh, you know, you see his usual list, That is, he has completed a total of fifteen orders.
3 :  Most of them are fixed to a place called the International Mansion in France around six in the evening, and then what?

0 : There have been some complaints about the fact that, whether it is true or not, we first look at the driver's identity, 0.2% who has been registered for several years. 
1 : Then also registered the Hitch, completed more than 3,000 singles a year was complained, four times, it seems that more than three thousand single inside nearly a year behind four is not particularly many, this driver is OK, So let's judge, alas, this is the first person who matches the car, and the driver is also a better driver, no problem, that is What? What happened to this case? What are we analyzing about the passengers? The passenger may have a point that he is a victim of a crowd, why? Because, first of all, it was a case we had, and then we got a lot of information, and he was a drunk passenger, and he went alone, hit a car, and then, Actually, the passenger was a woman, but our system identified it backstage, he was a man, and then, uh, he was probably, uh, you know, you see his usual list, That is, he has completed a total of fifteen orders. 
2 : Most of them are fixed to a place called the International Mansion in France around six in the evening, and then what?

*/
