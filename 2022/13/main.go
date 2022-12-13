package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const LIST = "LIST"
const VALUE = "VALUE"

type packet interface {
	Type() string
}

type list struct {
	items []packet
}

func (p list) Type() string {
	return LIST
}

type single struct {
	v int
}

func (p single) Type() string {
	return VALUE
}

type pair struct {
	first  list
	second list
}

func toInt(raw string) int {
	section, err := strconv.Atoi(raw)

	if err != nil {
		panic(err)
	}

	return section
}

func createList(raw string, i *int) list {
	list := list{}
	readingNumber := false
	for *i < len(raw) {
		char := raw[*i : *i+1]
		*i = *i + 1
		if char == "[" {
			list.items = append(list.items, createList(raw, i))
			readingNumber = false
		} else if char == "]" {
			return list
		} else if char == "," {
			readingNumber = false
		} else {
			if readingNumber {
				lastIndex := len(list.items) - 1
				value := list.items[lastIndex].(single).v
				list.items = list.items[:lastIndex]
				newValue := strconv.Itoa(value) + char
				list.items = append(list.items, single{v: toInt(newValue)})
			} else {
				list.items = append(list.items, single{v: toInt(char)})
			}
			readingNumber = true
		}
	}
	return list
}

func createPacket(raw string) list {
	i := 1
	return createList(raw, &i)
}

func createPair(raw1, raw2 string) pair {
	return pair{createPacket(raw1), createPacket(raw2)}
}

func loadPairs(filename string) []pair {
	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	pairs := []pair{}
	for fileScanner.Scan() {
		p1 := fileScanner.Text()
		if p1 != "" {
			fileScanner.Scan()
			p2 := fileScanner.Text()

			pairs = append(pairs, createPair(p1, p2))
		}
	}

	return pairs
}

func printList(l list) {
	fmt.Print("[")
	for _, i := range l.items {
		if i.Type() == LIST {
			printList(i.(list))
		} else {
			fmt.Print(i.(single).v)
		}
		fmt.Print(",")
	}
	fmt.Print("]")
}

func printPairs(pairs []pair) {
	for _, p := range pairs {
		printList(p.first)
		fmt.Println()
		printList(p.second)
		fmt.Println()
		fmt.Println()
	}
}

// 1 -> ordered, -1 -> unordered, 0 -> continue
func isOrdered(p pair) int {
	i := 0
	for {
		if i >= len(p.first.items) && i < len(p.second.items) {
			return 1
		} else if i < len(p.first.items) && i >= len(p.second.items) {
			return -1
		} else if i >= len(p.first.items) && i >= len(p.second.items) {
			return 0
		}

		lItem, rItem := p.first.items[i], p.second.items[i]
		if lItem.Type() == VALUE && rItem.Type() == VALUE {
			if lItem.(single).v < rItem.(single).v {
				return 1
			} else if lItem.(single).v > rItem.(single).v {
				return -1
			}
		} else {
			lList, rList := list{}, list{}

			if lItem.Type() == LIST {
				lList = lItem.(list)
			} else {
				lList.items = append(lList.items, lItem.(single))
			}

			if rItem.Type() == LIST {
				rList = rItem.(list)
			} else {
				rList.items = append(rList.items, rItem.(single))
			}

			order := isOrdered(pair{lList, rList})
			if order != 0 {
				return order
			}
		}
		i++
	}
}

func getOrderedPackets(pairs []pair) int {
	sum := 0

	for i, p := range pairs {
		if isOrdered(p) == 1 {
			sum += i + 1
		}
	}

	return sum
}

func flat(pairs []pair) []list {
	res := []list{}

	for _, p := range pairs {
		res = append(res, p.first)
		res = append(res, p.second)
	}

	return res
}

func findPacket(packets []list, p list) int {
	for i, v := range packets {
		if isOrdered(pair{p, v}) == 0 {
			return i
		}
	}
	panic("Not found")
}

func getDecoderKey(pairs []pair) int {
	packets := flat(pairs)

	p1, p2 := createPacket("[[2]]"), createPacket("[[6]]")
	packets = append(packets, p1)
	packets = append(packets, p2)

	sort.Slice(packets, func(i, j int) bool {
		return isOrdered(pair{packets[i], packets[j]}) == 1
	})

	k1Idx, k2Idx := findPacket(packets, p1), findPacket(packets, p2)

	return (k1Idx + 1) * (k2Idx + 1)
}

func main() {
	input := os.Args[1]

	pairs := loadPairs(input)

	sum := getOrderedPackets(pairs)
	fmt.Println("1:", sum)

	key := getDecoderKey(pairs)
	fmt.Println("2:", key)
}
