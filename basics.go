package main

import "fmt"

func main() {
	const availableStockCount int = 1000
	var productName string = "Oneplus 11R"
	var productPrice int = 35000
	var companyName = "Oneplus"

	category := "Smartphone"

	fmt.Println("Product name is", productName, "and its price is", productPrice, "INR.", "available at", companyName, "in the category of", category)

	loops_demo()
	arrays_demo()
	maps_demo()

	x, y := check_odd_even(10)
	fmt.Println("The number is", x, "and the return value is", y)
	demo_pointers()
}

func loops_demo() {
	for i := 0; i < 5; i++ {
		fmt.Println("Loop iteration:", i)
	}

	for i := range 3 {
		fmt.Println("Range loop iteration:", i)
	}

	for i, char := range "Hello" {
		fmt.Println("String range loop iteration:", i, char)
	}

	j := 10

	for j > 0 {
		if j == 5 {
			fmt.Println("Skipping iteration when j is 5")
			j--
			continue
		} else if j == 2 {
			fmt.Println("Skipping iteration when j is 2")
			j--
			continue
		} else {
			fmt.Println("While loop iteration:", j)
			j--
		}
	}
}

func arrays_demo() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	var arr1 []int

	arr1 = append(arr1, 1, 2, 3, 4, 5)

	fmt.Println("Array 1:", arr1)

	for i, value := range arr1 {
		fmt.Println("Index:", i, "Value:", value)
	}
}

func maps_demo() {
	productPrices := map[string]int{
		"Oneplus 11R": 35000,
		"iPhone 14":   80000,
		"Samsung S23": 70000,
	}

	fmt.Print("Product Prices: ", productPrices)

	customMap := map[string]string{}

	fmt.Println("Custom Map:", customMap)

	emptyMap := make(map[string]int)

	fmt.Println("Empty Map:", emptyMap)

	emptyMap["Oneplus 11R"] = 35000
	emptyMap["iPhone 14"] = 80000
	emptyMap["Samsung S23"] = 70000

	fmt.Println("Updated Empty Map:", emptyMap)

	for key, value := range emptyMap {
		fmt.Println("Key:", key, "Value:", value)
	}
}

func check_odd_even(num int) (string, int) {
	if num%2 == 0 {
		return "Even", 1
	} else {
		return "Odd", 0
	}
}

func demo_pointers() {
	i := 120

	var ptr *int = &i // this pointer can store the address of an integer variable

	ptr1 := &i // alternative way to declare a pointer

	fmt.Println("Value of i:", i, "Pointer to i:", ptr, "Pointer1 to i:", ptr1)

	fmt.Println("Value at pointer ptr:", *ptr)
}
