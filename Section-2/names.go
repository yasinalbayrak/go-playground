package main

/*
The names of Go function, variables, constants, typ es, statement lab els, and packages fol low a
simple rule: a name beg ins wit h a letter (that is, anything that Unicode deems a letter) or an
underscore and may have any number of addition al letters, dig its, and underscores.
*/

/*
break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var
*/

/*
Constants: true false iota nil
Types: int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr float32 float64 complex128 complex64 bool byte rune string error
Functions: make len cap new append copy close delete complex real imag panic recover
*/

/*
If an entity is declared within a function, it is local to that function. If declared outside a
function, however, it is visible in all files of the package to which it belongs. The cas e of the
firs t letter of a name determines its visibility across package bound aries. If the name beg ins
with an upp er-case letter, it is exported, which means that it is visible and accessible outside of
its own package and may be referred to by other parts of the program, as with Printf in the
fmtpackage. Package names themselves are always in lower case.
*/
func main() {

}
