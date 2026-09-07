package main

import "fmt"

// -Задача 1 
func task1() {
	fmt.Println("\n--- Задача 1 ---")
	rent := 95000
	increase := 10 // процентов
	newRent := rent + rent*increase/100
	fmt.Printf("Старая аренда: %d руб.\n", rent)
	fmt.Printf("Новая аренда (с повышением на %d%%): %d руб.\n", increase, newRent)
}

// -Задача 2 
func task2() {
	fmt.Println("\n--- Задача 2 ---")
	laptops := 6
	laptopPrice := 55480
	monitors := 3
	monitorPrice := 21830
	mice := 11
	mousePrice := 890
	keyboards := 5
	keyboardPrice := 1560

	total := laptops*laptopPrice + monitors*monitorPrice + mice*mousePrice + keyboards*keyboardPrice
	fmt.Printf("Ноутбуки: %d шт. по %d руб. = %d руб.\n", laptops, laptopPrice, laptops*laptopPrice)
	fmt.Printf("Мониторы: %d шт. по %d руб. = %d руб.\n", monitors, monitorPrice, monitors*monitorPrice)
	fmt.Printf("Мышки: %d шт. по %d руб. = %d руб.\n", mice, mousePrice, mice*mousePrice)
	fmt.Printf("Клавиатуры: %d шт. по %d руб. = %d руб.\n", keyboards, keyboardPrice, keyboards*keyboardPrice)
	fmt.Printf("Общая сумма: %d руб.\n", total)
}

// -Задача 3 
func task3() {
	fmt.Println("\n--- Задача 3 ---")
	storage := 5000
	fileSize := 256
	filesCount := storage / fileSize
	remainder := storage % fileSize
	fmt.Printf("Объём хранилища: %d Гб\n", storage)
	fmt.Printf("Размер одного файла: %d Гб\n", fileSize)
	fmt.Printf("Поместится файлов: %d\n", filesCount)
	fmt.Printf("Свободное место: %d Гб\n", remainder)
}

// -Задача 4
func task4() {
	fmt.Println("\n--- Задача 4 ---")
	var fahrenheit float64
	fmt.Print("Введите температуру в градусах Фаренгейта: ")
	fmt.Scan(&fahrenheit)
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, celsius)
}

// -Задача 9
func task9() {
	fmt.Println("\n--- Задача 9 ---")
	var price float64
	fmt.Print("Введите сумму покупки: ")
	fmt.Scan(&price)
	discounted := price * 0.8
	fmt.Println("Сумма со скидкой 20%:", discounted)
}

// Заглушки для вызываемых ниже функций (допиши сюда их код, когда сделаешь)
func task11() { fmt.Println("\n--- Задача 11 (В разработке) ---") }
func task12() { fmt.Println("\n--- Задача 12 (В разработке) ---") }

// - Задача 13
type Lesson struct {
	Subject string
	Room    int
	Teacher string
}

func task13() {
	fmt.Println("\n--- Задача 13 ---")
	const (
		Monday    = "Понедельник"
		Tuesday   = "Вторник"
		Wednesday = "Среда"
		Thursday  = "Четверг"
		Friday    = "Пятница"
		Saturday  = "Суббота"
		Sunday    = "Воскресенье"
	)

	schedule := make(map[string][]Lesson)

	schedule[Monday] = []Lesson{
		{"История", 250, "Кузнецова"},
		{"Астрономия", 121, "Корчагин"},
		{"Программирование", 303, "Антропов"},
	}

	schedule[Wednesday] = []Lesson{
		{"Математика", 292, "Козова"},
		{"Базы данных", 562, "Цыганков"},
	}

	var day string
	fmt.Print("Введите день недели (например, Понедельник): ")
	fmt.Scan(&day)

	if lessons, ok := schedule[day]; ok {
		fmt.Printf("Расписание на %s:\n", day)
		for i, les := range lessons {
			fmt.Printf("%d. %s, ауд. %d, преподаватель: %s\n", i+1, les.Subject, les.Room, les.Teacher)
		}
	} else {
		fmt.Printf("На %s занятий нет или день не найден.\n", day)
	}
}


func main() {
	fmt.Println("Практическая работа №1")

	task1()
	task2()
	task3()
	task4()
	task9() // Добавил вызов задачи 9
	task11()
	task12()
	task13()

	fmt.Println("\nВсе задачи выполнены.")
}
