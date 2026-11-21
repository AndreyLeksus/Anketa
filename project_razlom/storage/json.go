package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"projectrazlom/models"
)

func SavedAnket(ankets []models.Anketa) {
	path := "Ankets.json"

	file, err := json.MarshalIndent(ankets, "   ", "")
	if err != nil {
		fmt.Println(err)
		return
	}
		

		

		err = os.WriteFile(path, file, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}


		
		

		fmt.Println("Запись прошла успешно!")
	}

	


func LoadAnkets() ([]models.Anketa, error) {

	var ankets []models.Anketa

	data, err := os.ReadFile("Ankets.json")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println(err)

		}
		fmt.Println("Файл не  создан")
	}

	err = json.Unmarshal(data, &ankets)
	if err != nil {
		return nil, err
	}

	fmt.Println("Файл успешно прочитан")
	return ankets, err
}
