package main import(_"embed""fmt""time")var s string var x,y,i,dx int var
 d[8]int var f=[]int{31599,19812,14479,31207,23524,29411,29679,30866,31727,31719
,1040}func p(ch byte){i=x/2/(3+2)dx=x/2%(3+2)if i<8&&y/2<5&&dx<3&&(f[d[i]]>>((5-
y/2-1)*3+dx))&1==1{fmt.Printf("\033[1;41;30m%c\033[0m",ch)}else{fmt.Printf("%c",
ch)}if ch=='\n' {y+=1 x=0}else{x+=1}}func gd(){x=0 y=0 t:=time.Now()d[0]=t.Hour(
)/10 d[1]=t.Hour()%10 d[2]=10 d[3]=t.Minute()/10 d[4]=t.Minute()%10 d[5]=10 d[6]
=t.Second()/10 d[7]=t.Second()%10}func main(){for gd();;gd(){for so:=0;so<len(s)
;so++{if s[so]==63{for si:=0;si<len(s);si++{switch s[si]{case'\n':p('\\')p('n')p
('"')p('\n')p('"')case'"':p('\\')p('"')case'\\':p('\\')p('\\')default:p(s[si])}}
}else{p(s[so])}}fmt.Printf("\n\033[%dA\033[%dD",y+1,x)time.Sleep(1*time.Second)}
}