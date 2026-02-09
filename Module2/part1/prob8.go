package part1

func swap(a, b *int){
	c := *a
	*a = *b
	*b = c
}

func Prob8(){
	val1 := 1
	val2 := 2
	swap(&val1, &val2)
	println(val1)
	println(val2)
}