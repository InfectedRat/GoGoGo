package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	// Создаем функцию-обработчик для кнопки
	callback := js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		// Пишем "Hello, World!" в документ
		js.Global().Get("document").Call("getElementById", "output").Set("innerHTML", "Hello, World!")
		return nil
	})

	// Регистрируем функцию-обработчик
	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", callback)

	<-c // Ожидаем завершения (бесконечно блокируем выполнение)
}
