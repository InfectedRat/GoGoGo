package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Terminal")
	fmt.Println("---------------------")

	for {
		fmt.Print(">> ")
		// Чтение команды от пользователя
		input, _ := reader.ReadString('\n')

		// Удаление лишних пробелов и символов новой строки
		input = strings.TrimSpace(input)

		// Если введена команда "exit", выходим из программы
		if input == "exit" {
			fmt.Println("Exiting...")
			break
		}

		// Разделяем команду и аргументы
		args := strings.Fields(input)
		cmd := exec.Command(args[0], args[1:]...)

		// Установка стандартных потоков ввода/вывода для команды
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Выполнение команды
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Ошибка при выполнении команды: %v\n", err)
		}
	}
}
