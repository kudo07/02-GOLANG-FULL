package main

import "fmt"

// 1. GENERAL TYPES
// GENERAL TYPES
var (
	floatVar32 float32 = 0.12

	// FLAOT32 IS SMALLER THAN FLOAT64 ALWAYS USE FLOAT 64

	floatVar64 float64 = 0.12

	// STIRNG
	name string = "fevq"
	// INTEGER
	intVar32 = 3
	//
	intVar64 int64 = -34
	intVar8  int8  = -128 // 127
	// -128 to 127 ELSE GIVES AN ERROR

	// UNSIGNED INTEGER
	uintVar32 uint32 = 234
	//
	uintVar64 uint64 = 234
	uintVar8  uint8  = 42
	// 423 GIVES AN ERROR AS UINT TAKE ONLY 2**8-1 NUMBER ON THE POSITIVE SIDE

	byteVar byte = 0x2
	runVar  rune = 'f'
)

// 2. STRUCTS

type Player struct {
	name        string
	health      int
	attackPower float64
}

// NAME OF THE VARIBALE AND TYPE OF THE VARIABLE
func getHealth(player Player) int {
	return player.health
}

func (player Player) getHealthOne() int {
	return player.health
}

// 3. MAPS

// 6. CUSTOM TYPES
type Vehicle string

func getVehicle(vehicle Vehicle) string {
	// return vehicle
	// VEHICLE IS THE  TYPE OF VEHICLE WHICH IS STRING STILL GIVES ERROR CUSTOME TYPES
	return string(vehicle)
}

var vehicle string = "car"

func main() {
	// 2. STRUCTS
	player := Player{
		name:        "CAP Luffy",
		health:      100,
		attackPower: 87.4,
	} // INFER THE TYPE
	fmt.Println(player)
	fmt.Printf("this is the player:%v\n", player)
	// SUPER VERBOSE
	fmt.Printf("this is the player:%+v\n", player)
	// { 0 0}
	// this is the player:{ 0 0}
	// ATTACH STRUCT TO THE FUNCTION
	fmt.Printf("this is the player:%d\n", getHealth(player))
	fmt.Printf("this is the player:%d\n", player.getHealthOne())
	fmt.Printf("this is the player:%d\n", player.health)

	// 3. MAPS
	// INITIALISE THE EMPTY MAP
	users := map[string]int{
		"OP": 32,
	}
	// THIS IS CALLED AN EMPTY MAP
	users["foo"] = 21
	// usersOne := make(map[string]int)
	// EXACTLY SAME
	fmt.Printf("this is the player:%v\n", users)
	// ACCESS THE MAP
	age := users["foo"]
	fmt.Printf("a:%d\n", age)
	// IF THE KEY EXISTS OR NOT
	ageeERROR, ok := users["kEYNOTEXIST"]
	// OK IS THE BOOLEAN
	if !ok {
		fmt.Println("keynotexist is not in the map")
	} else {
		fmt.Println("keynotexist in the map:", ageeERROR)
	}
	fmt.Printf("a:%d\n", ageeERROR)

	// DELETE IN THE MAP
	// DELETE KEYWORD IS RESERVED FOR THE MAP ONLY
	fmt.Printf("after delate:%v\n", users)
	delete(users, "foo")

	// d. RANGE OVER MAPS
	users["addsecond"] = 3241
	for k, v := range users {
		fmt.Printf("the key %s and the value %d\n", k, v)
	}

	// 4. SLICES
	// PREFFERED TO MAKE SLICE LIKE THIS
	numbers := []int{1, 2, 3, 4, 5}
	// UNSUAL
	oddNumbers := make([]int, 2, 23)
	fmt.Println(numbers)
	fmt.Println(oddNumbers)
	// SLICE CAN DYNAMICALLY GROW
	numbers = append(numbers, 23, 34)
	fmt.Println(numbers)
	// BUT ARRAYS CAN

	// 5. ARRAYS
	// SLICE CAN INCREASE OR DECREASE
	arrayEx := [2]int{}
	// arrayEx = append(arrayEx, 1)
	// CANNOT DO THIS AS THIS IS PRE FIXED LENGTH WHICH IS ARRAY
	arrayEx[0] = 1
	arrayEx[1] = 2
	fmt.Println(arrayEx)

	// 6. CUSTOME TYPES
}
