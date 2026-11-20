package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"projectrazlom/models"
)

func SavedAnket(ankets [] models.Anketa) {
	file, err := json.MarshalIndent(ankets, "   ", "")
	if err != nil{
		fmt.Println(err)
		return
	}


	err = os.WriteFile("Ankets.json", file, 0644)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("Запись прошла успешно!")
}