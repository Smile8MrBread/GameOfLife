package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Universe struct {
	World [][]bool
}


func (u *Universe) newWorld(width int, height int) {
	u.World = make([][]bool, height)
	
	for i := 0; i < height; i++ {
		u.World[i] = make([]bool, width)
	}
}

func (u *Universe) Show() {
	var t string

	for i := range u.World {
		for j := range u.World[0] {	
			if u.World[i][j] {
				// t += "#"
				t += "\xF0\x9F\x9F\xA9"
				} else {
				// t += " "
				t += "\xF0\x9F\x9F\xAB"
			}
		}
		t += "\n"
	}

	fmt.Println(t)
}

func (u *Universe) Seed() {
	for i := range u.World {
		for j := range u.World[0] {
			r := rand.Intn(100)
			if r <= 8 {
				u.World[i][j] = true
				} else {
				u.World[i][j] = false
			}
		}
	}
}

func (u *Universe) Neighbors(x, y int) int {
	count := 0

    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            if i == 0 && j == 0 {
                continue // Пропускаем текущую клетку
            }
            
            // Вычисляем соседние клетки с учетом тора
            neighborX := (x + i + len(u.World[0])) % len(u.World[0])
            neighborY := (y + j + len(u.World)) % len(u.World)

            // Если соседняя клетка живая, увеличиваем счетчик
			if u.World[neighborY][neighborX] {
                count++
            }
        }
    }
    return count
}

func (u *Universe) Next(x, y int) bool{
	neighbor := u.Neighbors(x, y)
	if u.World[y][x] {	
		if neighbor == 2 || neighbor == 3 {
			return true
			} else {
				return false
			}
			} else {
				if neighbor == 3 {
			return true
		} else {
			return false
		}
	}
}

func Step(a, b Universe) {
	b = a
	for i := range b.World {
		for j := range b.World[0] {
			b.World[i][j] = b.Next(j, i)
		}
	}

	a = b
}

func main() {
	const (
		width = 50
		height = 50
	)

	World := Universe{}
	SecondWorld := Universe{}

	World.newWorld(width, height)
	World.Seed()
	
	fmt.Println("\033[2J")
	for {
		Step(World, SecondWorld)
		fmt.Print("\033[H")
		World.Show()
		time.Sleep(time.Second / 15)
	}
}