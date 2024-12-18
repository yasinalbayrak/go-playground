package main

var global *int

func f1() {
	var x int
	x = 1
	global = &x
}

/*
Here , x must be heap-allocated because it is still reach able from the variable global after f
has returned, despite being declared as a local variable; we say x escapes from f. Conversely,
when g returns, the variable *y becomes unreachable and can be recycled. Since *y do es not
es cap e from g, it’s safe for the compiler to allocate *y on the stack, even though it was allocated with new.
In any case, the notion of escaping is not something that you need to worry
about in order to write correct code, though it’s good to keep in mind during performance
optimization, since each variable that escapes re quires an extra memory allocation.
*/
func f2() {
	y := new(int)
	*y = 1
}
