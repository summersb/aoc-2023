package main

import (
    "fmt"
    "bufio"
    "os"
    "log"
    "regexp"
    "strconv"
    "aoc2023/support"
)

func main() {
    fmt.Println("Day 01 2023")
    part1();
    part2();
}

func part1() {

    fmt.Println("Part 1")
    file, err := os.Open("./data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)


    total := 0

    re := regexp.MustCompile("[0-9]")
    for scanner.Scan() {
        txt := scanner.Text()

        nums := re.FindAllString(txt, -1)

        coord := 0
        if len(nums) == 1 {
            first, err := strconv.Atoi(nums[0])
            if err != nil {
                log.Fatal(err)
            }
            coord = first * 10 + first;
        } else {
            size := len(nums)
            first, err := strconv.Atoi(nums[0])
            if err != nil {
                log.Fatal(err)
            }
            second, err := strconv.Atoi(nums[size-1])
            coord = first * 10 + second
        }

        //fmt.Println(nums, coord)
        total = total + coord

    }
    fmt.Println("Total", total)

}

func part2() {
    fmt.Println("Part 2")
    file, err := os.Open("./data.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)


    total := 0

    for scanner.Scan() {
        txt := scanner.Text()
        
        first := support.FindFirstNumber(txt)
        second := support.FindLastNumber(txt)

        total = total + first * 10 + second
        fmt.Println("Number", first, second, txt, total)
    }

    fmt.Println("Total", total)
}


func findSecondNumber(txt string) int {
    return 1
}
