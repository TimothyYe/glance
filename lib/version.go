package lib

import "fmt"

var (
	// Logo
	Logo = `
#####                                     
#     # #        ##   #    #  ####  ###### 
#       #       #  #  ##   # #    # #      
#  #### #      #    # # #  # #      #####  
#     # #      ###### #  # # #      #      
#     # #      #    # #   ## #    # #      
 #####  ###### #    # #    #  ####  ###### 
Glance V%s
https://github.com/TimothyYe/glance
`
)

func Display(version string) {
	fmt.Printf(Logo, version)
}
