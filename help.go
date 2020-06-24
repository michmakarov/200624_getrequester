// help
package main

var help = `
This (getrequester) is console program that sends a get request
to  URL given by its first argument.
If it is run without arguments it shows this text and ends its work.
The prog shows the answer (if any) in various forms accordingly of the mode.
The mode is string that may be set by third argument and it has next values
aheaders - (answer headers) it shows answer headers only
head - if content type is text/http then it shows answer headers plus head tag else it equal the aheaders.
all - it shows whole answer (headers and body)
If third argument is ommited or not equal the listed above its value is aheaders.
`
