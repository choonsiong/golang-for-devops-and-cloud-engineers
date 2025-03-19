package main

func main() {
	arr := []string{} // this array is initialized with 0 length!!!
	//arr[0] = "a"      // error! because the length of arr is 0!!!
	arr = append(arr, "a") // ok, because append will create a new array!!!
	print(arr[0])
}
