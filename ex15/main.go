package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100])) // Копирование первых 100 символов в новую строку
}

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func main() {
	someFunc()
}

/*
В Go строки являются неизменяемыми, но подсрез строки не создает новую строку.
Вместо этого он создает новую структуру, которая указывает на тот же массив байт,
что и оригинальная строка. В результате, даже если срез включает только первые 1
00 символов, в памяти будет оставаться весь исходный массив байт, что может привести
к утечке памяти, особенно если createHugeString возвращает очень большую строку.
*/
