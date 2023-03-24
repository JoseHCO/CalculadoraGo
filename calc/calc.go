package calc

func somar(a, b float32) float32 {
	return a + b
}

func subtrair(a, b float32) float32 {
	return a - b
}

func multiplicar(a, b float32) float32 {
	return a * b
}

func dividir(a, b float32) float32 {
	return a / b
}

func ValidarCalculo(a, b float32, c string) interface{} {
	var response float32
	if c == "+" {
		response = somar(a, b)
	} else if c == "-" {
		response = subtrair(a, b)
	} else if c == "/" {
		response = dividir(a, b)
	} else if c == "*" {
		response = multiplicar(a, b)
	} else {
		return false
	}
	return response
}
