package main

import (
	"fmt"
	"os"
)

type NameFolders struct {
	Name string
}

var folders []NameFolders
var intFolders []NameFolders

func main() {
	folders = append(folders, NameFolders{Name: "api"})
	folders = append(folders, NameFolders{Name: "build"})
	folders = append(folders, NameFolders{Name: "cmd"})
	folders = append(folders, NameFolders{Name: "configs"})
	folders = append(folders, NameFolders{Name: "deployments"})
	folders = append(folders, NameFolders{Name: "docs"})
	folders = append(folders, NameFolders{Name: "internal"})
	folders = append(folders, NameFolders{Name: "pkg"})

	for _, folder := range folders {
		dirPath := fmt.Sprintf("../my-new-project/%s", folder.Name)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			fmt.Printf("Ошибка при создании папки: %v", err)
			return
		}

		fmt.Println("Вложенные папки успешно созданы по пути :", dirPath)

	}

	dirPath := "../my-new-project/cmd/app"
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Printf("Ошибка при создании папки: %v", err)
		return
	}

	intFolders = append(intFolders, NameFolders{Name: "app"})
	intFolders = append(intFolders, NameFolders{Name: "config"})
	intFolders = append(intFolders, NameFolders{Name: "database"})
	intFolders = append(intFolders, NameFolders{Name: "models"})
	intFolders = append(intFolders, NameFolders{Name: "services"})
	intFolders = append(intFolders, NameFolders{Name: "transport"})

	for _, folder := range intFolders {
		dirPath := fmt.Sprintf("../my-new-project/internal/%s", folder.Name)
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			fmt.Printf("Ошибка при создании папки: %v\n", err)
			return
		}

		fmt.Println("Папка успешно создана по пути:", dirPath)
	}

	fileName := "../my-new-project/README.md"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Ошибка при создании файла %v", err)
	}
	defer file.Close()

	fmt.Printf("Файл %s успешно создан\n", fileName)
}
