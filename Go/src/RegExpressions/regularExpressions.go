package RegExpressions

import "fmt"
import "regexp"

func Examples() {
	var pat string
	s := "hello hi 1234 ok fine 56789 done"
	fmt.Print("enter some string:")
	fmt.Scan(&pat)
	cpat, _ := regexp.Compile(pat)
	if cpat.MatchString(s) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found")
	}

	result := cpat.FindAllString(s, -1)
	n := len(result)
	fmt.Println("result =", result)
	fmt.Println("total number of items found =", n)

	s1 := cpat.ReplaceAllString(s, "ABC")
	fmt.Println("s1 =", s1)

}


/*
Regular Expressions: to search, extract, and replace
----------------------------------------------------
List of Metacharacters:
-----------------------
. 			--->	any one character
? 			--->	zero or one
+ 			--->	one or more
* 			--->	zero or more
[abc] 	 	--->	any one of a b c
(abc)      	--->	group of abc
{m}			--->	'm' times
{m, n}		--->	at least 'm' times, at most 'n' times
^ 			--->	at the begining of the string
$			--->	at the end of the string
|			--->	or
\			--->	escape sequence character
\d 			---> 	a digit character	<===>	[0-9]
\w			--->	a word character	<===>	[a-zA-Z0-9_]
\s 			--->	a space
\b 			--->	a word boundary
---------------------------------------------------------

\d 			---> 	a single digit number(0 to 9)
\d\d 		--->	a two digit number (00 to 99)
\d\d\d 		--->	a three digit number (000 to 999)
------------------------------------------------------------
NOTE: ?, +, *, {}  are used as Quantifiers (to represent quantity)
-------------------------------------------------------------
\d{3}		--->	a three digit number
\d{3, 5}	--->

hell?o      ---> 	helo | hello
hell+o 		--->	hello | helllo | hellllo | ...
hell*o 		--->	helo | hello | helllo | ...
he(ll)+o 	---> 	hello | hellllo | hellllllo | ...

s = "hi how are you hello"

hello 		--->    YES
^hello 		--->	NO
hello$		--->	YES
---------------------------------
s1 = "hello how are you hello"

^hello$		--->  NO

^hello .* hello$

[bB]angalore   --->		bangalore | Bangalore
[bB]anga?lore? 	--->

\d{3}[13579] 	---> 	a 4-digit ODD number

[78]\d{3, 5}[123]  --->	

\d+ 			--->	a number
[0123456789]+	--->	a number
[0-9]+			--->	a number

\d+\s\w+ 		--->	a number followed by a space followed by a word
\d+\s[a-z]+		--->	a number followed by a space followed by a word of lower case alphabets only

\d+\s[a-zA-K786]{5, 10}  --->

(ABC)?\d{3}:\d{4}[XYZ]?

*/