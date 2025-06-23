package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Title line
	fmt.Println("========================================")
	fmt.Println("   GO DATA TYPES DEMONSTRATION")
	fmt.Println("========================================")
	fmt.Println()

	// Declare variables using all primitive data types
	var (
		// Boolean type
		isActive bool = true
		
		// String type
		greeting string = "Welcome to Go programming!"
		
		// Integer types
		age int = 25
		smallNum int8 = 127
		mediumNum int16 = 32767
		bigNum int32 = 2147483647
		hugeNum int64 = 9223372036854775807
		
		// Unsigned integer types
		count uint = 42
		tinyCount uint8 = 255
		smallCount uint16 = 65535
		mediumCount uint32 = 4294967295
		largeCount uint64 = 18446744073709551615
		
		// Byte type (alias for uint8)
		byteVal byte = 'A'
		
		// Rune type (alias for int32, represents Unicode code points)
		runeVal rune = '🚀'
		
		// Floating-point types
		temperature float32 = 98.6
		precision float64 = 3.14159265359
		
		// Complex types
		complexNum32 complex64 = 3 + 4i
		complexNum64 complex128 = 5.5 + 2.3i
		
		// Pointer type
		ptr *int = &age
	)

	// Display all primitive data types
	fmt.Println("📊 PRIMITIVE DATA TYPES IN GO:")
	fmt.Println("-------------------------------")
	fmt.Printf("Boolean (bool): %v (type: %T)\n", isActive, isActive)
	fmt.Printf("String: %q (type: %T)\n", greeting, greeting)
	fmt.Printf("Integer (int): %d (type: %T)\n", age, age)
	fmt.Printf("Int8: %d (type: %T)\n", smallNum, smallNum)
	fmt.Printf("Int16: %d (type: %T)\n", mediumNum, mediumNum)
	fmt.Printf("Int32: %d (type: %T)\n", bigNum, bigNum)
	fmt.Printf("Int64: %d (type: %T)\n", hugeNum, hugeNum)
	fmt.Printf("Unsigned int (uint): %d (type: %T)\n", count, count)
	fmt.Printf("Uint8: %d (type: %T)\n", tinyCount, tinyCount)
	fmt.Printf("Uint16: %d (type: %T)\n", smallCount, smallCount)
	fmt.Printf("Uint32: %d (type: %T)\n", mediumCount, mediumCount)
	fmt.Printf("Uint64: %d (type: %T)\n", largeCount, largeCount)
	fmt.Printf("Byte: %c (%d) (type: %T)\n", byteVal, byteVal, byteVal)
	fmt.Printf("Rune: %c (%d) (type: %T)\n", runeVal, runeVal, runeVal)
	fmt.Printf("Float32: %.2f (type: %T)\n", temperature, temperature)
	fmt.Printf("Float64: %.11f (type: %T)\n", precision, precision)
	fmt.Printf("Complex64: %v (type: %T)\n", complexNum32, complexNum32)
	fmt.Printf("Complex128: %v (type: %T)\n", complexNum64, complexNum64)
	fmt.Printf("Pointer (*int): %p -> %d (type: %T)\n", ptr, *ptr, ptr)
	fmt.Println()

	// Interactive section with input validation and retry logic
	fmt.Println("🎯 INTERACTIVE SECTION:")
	fmt.Println("-----------------------")
	
	scanner := bufio.NewScanner(os.Stdin)
	
	// Get user's name with retry
	var userName string
	for {
		fmt.Print("Please enter your name: ")
		if scanner.Scan() {
			userName = strings.TrimSpace(scanner.Text())
			if userName != "" {
				break
			}
			fmt.Println("❌ Error: Name cannot be empty. Please try again.")
		} else {
			fmt.Println("❌ Error reading input. Please try again.")
		}
	}

	// Get user's age with validation and retry
	var userAge int
	for {
		fmt.Print("Please enter your age: ")
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			if parsedAge, err := strconv.Atoi(input); err != nil {
				fmt.Printf("❌ Error: '%s' is not a valid number. Please enter a valid age.\n", input)
			} else if parsedAge < 0 {
				fmt.Println("❌ Error: Age cannot be negative. Please enter a valid age.")
			} else if parsedAge > 150 {
				fmt.Println("❌ Error: Age seems unrealistic (>150). Please enter a valid age.")
			} else {
				userAge = parsedAge
				break
			}
		} else {
			fmt.Println("❌ Error reading input. Please try again.")
		}
	}

	// Get user's favorite number (float) with validation and retry
	var favoriteNumber float64
	for {
		fmt.Print("Please enter your favorite number (can be decimal): ")
		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())
			if parsedNum, err := strconv.ParseFloat(input, 64); err != nil {
				fmt.Printf("❌ Error: '%s' is not a valid number. Please enter a valid number.\n", input)
			} else {
				favoriteNumber = parsedNum
				break
			}
		} else {
			fmt.Println("❌ Error reading input. Please try again.")
		}
	}

	// Get yes/no response with validation and retry
	var likesGo bool
	for {
		fmt.Print("Do you like Go programming? (yes/no or y/n): ")
		if scanner.Scan() {
			input := strings.ToLower(strings.TrimSpace(scanner.Text()))
			switch input {
			case "yes", "y", "true", "1":
				likesGo = true
				goto continueProgram
			case "no", "n", "false", "0":
				likesGo = false
				goto continueProgram
			default:
				fmt.Printf("❌ Error: '%s' is not a valid yes/no response. Please enter 'yes' or 'no'.\n", input)
			}
		} else {
			fmt.Println("❌ Error reading input. Please try again.")
		}
	}

continueProgram:
	// Process and display results
	fmt.Println()
	fmt.Println("✅ INPUT PROCESSING COMPLETE!")
	fmt.Println("==============================")
	
	// Calculate some values using the input
	ageInMonths := userAge * 12
	isAdult := userAge >= 18
	numberSquared := favoriteNumber * favoriteNumber
	
	// Create a personalized response
	fmt.Printf("Hello, %s! 👋\n", userName)
	fmt.Printf("You are %d years old, which is %d months!\n", userAge, ageInMonths)
	fmt.Printf("Adult status: %v\n", isAdult)
	fmt.Printf("Your favorite number %.2f squared is %.4f\n", favoriteNumber, numberSquared)
	
	if likesGo {
		fmt.Println("Great! Go is an awesome programming language! 🎉")
	} else {
		fmt.Println("That's okay, maybe you'll grow to like Go! 😊")
	}

	// Demonstrate type conversions
	fmt.Println()
	fmt.Println("🔄 TYPE CONVERSION EXAMPLES:")
	fmt.Println("----------------------------")
	userAgeFloat := float64(userAge)
	favoriteNumberInt := int(favoriteNumber)
	fmt.Printf("Your age as float64: %.2f\n", userAgeFloat)
	fmt.Printf("Your favorite number as int: %d\n", favoriteNumberInt)
	
	// Show some calculations with different types
	result := userAgeFloat + favoriteNumber
	fmt.Printf("Age + Favorite Number: %.2f\n", result)

	// Memory addresses demonstration
	fmt.Println()
	fmt.Println("💾 MEMORY ADDRESSES:")
	fmt.Println("--------------------")
	fmt.Printf("Address of userName: %p\n", &userName)
	fmt.Printf("Address of userAge: %p\n", &userAge)
	fmt.Printf("Address of favoriteNumber: %p\n", &favoriteNumber)

	// Ending line
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("   PROGRAM COMPLETED SUCCESSFULLY!")
	fmt.Println("   Thank you for exploring Go data types!")
	fmt.Println("========================================")
}