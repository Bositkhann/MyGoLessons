package main 

import "fmt"

func main(){
var h string = "*"
var j string = "="
	fmt.Printf("    %s       %s        %s          %s %s %s      %s      %s %s %s           %s                         \n", h,h,h,h,h,h,h,h,h,h,h                             )
	fmt.Printf("    %s %s     %s       %s  %s        %s               %s                %s  %s                       \n",h,h,h,h,h,h,h,h,h                                  )
	fmt.Printf("    %s  %s    %s      %s    %s       %s          %s    %s               %s    %s                      \n",h,h, h,h,h,h,h,h,h,h                                  )
	fmt.Printf("    %s   %s   %s     %s      %s      %s %s %s      %s     %s %s           %s      %s                     \n",h,h,h, h,h,h,h,h,h,h,h,h,h                                )
	fmt.Printf("    %s    %s  %s    %s  %s  %s  %s     %s          %s         %s        %s  %s  %s  %s                    \n",h,h,h,h,h, h,h,h,h,h,h,h,h,h                                  )
	fmt.Printf("    %s     %s %s   %s          %s    %s          %s         %s       %s          %s                   \n",h ,h,h, h,h,h,h,h,h,h,                                 )
	fmt.Printf("    %s       %s  %s            %s   %s          %s    %s %s %s       %s            %s                  \n", h,h,h,h,h,h,h,h,h,h,h                                 )
	
	
 m  := "Naf"
 n  := "isa"

fmt.Printf("%d %s %d %s %d", m ,h,n,j,m + n)
fmt.Println( m + n)

}