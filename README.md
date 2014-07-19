chalk
=============

Chalk is a go package for styling console/terminal output.

The api is pretty clean, they are default Colors and TextStyles
which can be mixed to create more intense Styles. Styles and Colors
can just be printed in normal strings (i.e. fmt.Sprintf(chalk.Red)), but
Styles, Colors and TextStyles are more meant to be used to style specific
text segments (i.e. fmt.Println(chalk.Red.Color("this is red")) or
fmt.Println(myStyle.Style("this is blue text that is underlined"))).