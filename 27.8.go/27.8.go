package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

func (s Student) info() string {
	return s.name + " " + strconv.Itoa(s.age) + " " + strconv.Itoa(s.grade)
}
func (s *Student) put(m map[string]*Student) {
	m[s.name] = s
}
func (s Student) get(m map[string]*Student) bool {
	_, x := m[s.name]
	return x
}
func main() {
	mapStudents := make(map[string]*Student, 0)
	fmt.Println("Введите имя, возраст и курс студента:")
	var stdin = bufio.NewReader(os.Stdin)

	for i, err := stdin.ReadString('\n'); err != io.EOF; {
		sliceStruct := strings.Fields(i)
		studentName := sliceStruct[0]
		studentAge, errAge := strconv.Atoi(sliceStruct[1])
		studentGrade, errGrade := strconv.Atoi(sliceStruct[2])

		if len(sliceStruct) < 3 {
			fmt.Println("Что-то пошло не так. Ошибка. Попробуйте еще раз")
			continue
		}

		if errAge != nil || errGrade != nil {
			fmt.Println("произошла ошибка в обработке данных в Возрасте или Группе студента. Попробуйте еще раз")
			continue
		}
		student := Student{studentName, studentAge, studentGrade}
		if err := student.get(mapStudents); err != false {
			student.put(mapStudents)
		} else {
			fmt.Println("Студент с таким именем уже существует")
			break
		}
		for _, value := range mapStudents {
			fmt.Println("Вывожу всех студентов:")
			fmt.Println(value.info())
		}
	}
}
