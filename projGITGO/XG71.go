// Go program to illustrate the
// concept of multi-dimension array
package main
 
    import & quot;
fmt& quot;
 
func main()
{
 
// Creating and initializing
// 2-dimensional array
// Using shorthand declaration
// Here the (,) Comma is necessary
arr:
    = [ 3 ][3] string{ { "
    C #& quot;
    , "
    C& quot;
    , "
    Python& quot;
}
,
{
    "
    Java& quot;
    , "
    Scala& quot;
    , "
    Perl& quot;
}
,
{
    "
    C++& quot;
    , "
    Go& quot;
    , "
    HTML& quot;
}
,
}
 
// Accessing the values of the
// array Using for loop
fmt.Println("Elements of Array 1")
for x:= 0;
x& lt;
3;
x++
{
for
y:
    = 0;
y& lt;
3;
y++ { fmt.Println(arr[x][y]) }
}
 
// Creating a 2-dimensional
// array using var keyword
// and initializing a multi
// -dimensional array using index
var arr1 [2][2] int
arr1[0][0] = 100
arr1[0][1] = 200
arr1[1][0] = 300
arr1[1][1] = 400
 
// Accessing the values of the array
fmt.Println("Elements of array 2")
for p:= 0;
p& lt;
2;
p++
{
for
q:
    = 0;
q& lt;
2;
q++ { fmt.Println(arr1[p][q]) }
}
}