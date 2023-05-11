package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) + 1
	fmt.Println("Введите число или end чтобы завершить программу")
	number := bufio.NewReader(os.Stdin)
	for i := 10; i >= 0; i-- {
		input, err := number.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		grate, err := strconv.Atoi(input)
		if grate == n {
			fmt.Println("Число верное. Это:", grate)
			break
		} else if grate < n || grate > n {
			fmt.Println("Число не верно,\n Повторите еще раз. Число попыток \n", i, "раз")
		}
		if i == 0 {
			fmt.Println("Число попыток превышено")
			return
		}
		if input == "end" {
			os.Exit(1)
		}
	}

}
