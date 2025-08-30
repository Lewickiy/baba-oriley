package main

import (
	"baba-oriley/internal/audio"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Генерирую пустой wav файл")

	err := audio.CreateSilenceWav("silence", 44100, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Файл out/silence.wav создан")
}
