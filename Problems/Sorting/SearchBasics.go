package Sorting

import (
	"fmt"
	"sort"
)

/*
	Standard sorting on primitive types
*/
func Sorting() {
	// ========= 1 ===========
	int_arr := []int{4, 6, 2, 1, 9, 5, 3, 7}

	// Similar are possible with Strings and Floats
	sort.Ints(int_arr)
	sort.IntsAreSorted(int_arr)

	pos := sort.SearchInts(int_arr, 5)
	fmt.Printf("Found %d at index %d in %v\n", 5, pos, int_arr)
	// =======================

	// ========== 2 ============
	int_arr = []int{4, 6, 2, 1, 9, 5, 3, 7}
	pos = sort.SearchInts(int_arr, 3)
	fmt.Printf("Found %d at index %d in %v\n", 3, pos, int_arr)
	// ========================

	// ========== 3 =============
	fmt.Println("\n######## Search works in descending order  ########")
	i := sort.Search(len(int_arr), func(i int) bool { return int_arr[i] <= 3 })
	fmt.Printf("Found %d at index %d in %v\n", 3, i, int_arr)
	// ==========================
}

/*
	Sorting on custom data structures
*/

type Mobile struct {
	BrandName string
	Price     int
}

type ByPrice []Mobile

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }

type ByBrandName []Mobile

func (a ByBrandName) Len() int           { return len(a) }
func (a ByBrandName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBrandName) Less(i, j int) bool { return a[i].BrandName < a[j].BrandName }

func SortingByStructs() {
	mobiles := []Mobile{
		{"Nokia", 50},
		{"OnePlus", 100},
		{"Apple", 999},
		{"Samsung", 500},
	}

	sort.Sort(ByPrice(mobiles))
	fmt.Printf("Mobiles sorted in order of their price %+v\n", mobiles)

	sort.Sort(ByBrandName(mobiles))
	fmt.Printf("Mobiles sorted in order of their name %+v\n", mobiles)

	fmt.Println("\nFound mobiles price is sorted :", sort.IsSorted(ByBrandName(mobiles))) // false

}
