## Named scene

Named scene must be named from config file


## Syntax for dialogue presented

Simple example:
`"You turned her against me"-"I will do what I must"`

Feature full example:
`#4"You turned her against me"(-2)[1]-"I will do what I must"(3)[3]`

There is a regular expression used to match this string: `^(#\d+)?"[\w\s]+"(\(\-?[0-9]+\))?(\[\d+\])?\-"[\w\s]+"(\(\-?[0-9]+\))?(\[\d+\])?$`

Firstly you can explicitly define movie id if u want by add prefix hash + movie id.

Then, between quote marks (") you must pass the first phrase.
Nextly you can set offset (in seconds) between parentheses.

If there are many phrases engine return error with information about possible phrases, then you can define what phrase you meant by add the number of phrase between square brackets (like in many programming languages).

Then you can do the same with a second phrase.
