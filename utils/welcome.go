package utils

import (
	"github.com/fatih/color"
)

// WelcomeArt returns the ASCII art for the welcome message.
func WelcomeArt() {
	violet := color.New(color.FgHiMagenta)
	violet.Println(`
                +                   ****     ***   *****************       **     *** ************  
               ****                 ****    **     ***          ****      ***     *** ****       ** 
              +****+                +++*  ++       +++           +*+     +++*+   +++  ++++        + 
             +*+++**+               ++++++++       +++            +++   ++ ++++ ++++  ++++      +++ 
            ++++  ++++              +++++++++      ++++++++++++   +++   +   + + +++   +++++++++++   
           ++++    ++++             ++ ++ ++++     +++             +   ++   +++++++   ++++     ++   
         ++++       ++++            ++++   +++++   +++             ++ ++     +++++    ++++      ++  
        ++++         ++++           ++++     ++++  +++              + ++      +++     ++++       ++ 
       ++++           ++++          ++++      ++++ ++         ++    +++        ++     +++        ++ 
      ++++             +++++                                                                        
     ++++    ++++++++   +++++       ********    ***    ********   ***   ********     **     ***     
    +++++++++++++++++++++ +++       +++*  +++   *++   +++    +*   +++     +++       +++     +++     
   ++++++++++       +++++++++++     ++++   +++  +++  +++    ++++  +++     +++      +++++    +++     
  +++++++                +++++++    ++++   +++  +++  +++    ++++  +++     +++     +++ +++   +++     
++++                          +++   ++++  ++++  +++   +++    +++  +++     +++    +++++++++  +++     
                                    ++++++++    +++    ++++++++   +++     +++   ++++    +++ +++++++    
	`)
}