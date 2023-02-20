package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Введите математическое выражение, разделяя оператор и числа пробелами:")

    for scanner.Scan() {
        input := strings.Split(scanner.Text(), " ")

        if len(input) != 3 {
            fmt.Println("Ошибка: некорректное количество аргументов")
            os.Exit(1)
        }

        var a, b int
        var err error

        switch input[0] {
        case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":
            a, err = romanToArabic(input[0])
            if err != nil {
                fmt.Println("Ошибка: некорректное римское число")
                os.Exit(1)
            }
        default:
            a, err = strconv.Atoi(input[0])
            if err != nil {
                fmt.Println("Ошибка: некорректное число")
                os.Exit(1)
            }
        }

        switch input[2] {
        case "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X":
            b, err = romanToArabic(input[2])
            if err != nil {
                fmt.Println("Ошибка: некорректное римское число")
                os.Exit(1)
            }
        default:
            b, err = strconv.Atoi(input[2])
            if err != nil {
                fmt.Println("Ошибка: некорректное число")
                os.Exit(1)
            }
        }

        switch input[1] {
        case "+":
            fmt.Println(a + b)
        case "-":
            fmt.Println(a - b)
        case "*":
            fmt.Println(a * b)
        case "/":
            if b == 0 {
                fmt.Println("Ошибка: нельзя делить на ноль")
                os.Exit(1)
            }
            fmt.Println(a / b)
        default:
            fmt.Println("Ошибка: некорректный оператор")
            os.Exit(1)
        }
    }
}

func romanToArabic(roman string) (int, error) {
    values := map[rune]int{'I': 1, 'V': 5, 'X': 10}

    var result int
    var prevValue int

    for _, r := range roman {
        value, ok := values[r]
        if !ok {
            return 0, fmt.Errorf("некорректный символ в римском числе: %v", r)
        }
        if value > prevValue {
            result += value - 2*prevValue
        } else {
            result += value
        }
        prevValue = value
    }

    if result > 10 {
        return 0, fmt.Errorf("результат превышает 10")
    }

    return result, nil
}
