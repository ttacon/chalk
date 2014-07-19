chalk
=============

Chalk is a go package for styling console/terminal output.

Check out the godoc for some example usage:
http://godoc.org/github.com/ttacon/chalk

The api is pretty clean, they are default Colors and TextStyles
which can be mixed to create more intense Styles. Styles and Colors
can just be printed in normal strings (i.e. ```fmt.Sprintf(chalk.Red)```), but
Styles, Colors and TextStyles are more meant to be used to style specific
text segments (i.e. fmt.Println(chalk.Red.Color("this is red")) or
fmt.Println(myStyle.Style("this is blue text that is underlined"))).

WARNING
=============

This package should be pretty stable, but I'm not making any promises :)