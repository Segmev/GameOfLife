package main

import (
  "time"
  "fmt"
  "math/rand"
  "bytes"
)

var maxX, maxY = 77, 30

func display(tab [][]int) {
  var buffer bytes.Buffer
  
  for j := 0; j < maxY; j++ {
    buffer.WriteString("   ")
    for i := 0; i < maxX; i++ {
      if tab[j][i] == 0 {
        buffer.WriteString("  ")
      } else if tab[j][i] < 3  {
        buffer.WriteString("¤ ")
      } else if tab[j][i] < 6  {
        buffer.WriteString("* ")
      } else if tab[j][i] < 10  {
        buffer.WriteString("% ")
      } else {
        buffer.WriteString("# ")
      }
    }
    buffer.WriteString("\n")
  }
  fmt.Println(buffer.String())
  fmt.Println()
  fmt.Println()
}

func updateTab(tab [][]int) [][]int {
  buff := make([][]int, maxY)
  for i := 0; i < maxY; i++ {
    buff[i] = make([]int, maxX)
  }

  for j := 0; j < maxY; j++ {
    for i := 0; i < maxX; i++ {
      c := 0

      for a := -1; a <= 1; a++ {
        for b := -1; b <= 1; b++ {
          if (!(a == 0 && b == 0)) {
            if (j + a >= 0 && j + a < maxY && i + b >= 0 && i + b < maxX) {
              if (tab[j + a][i + b] > 0) {
                c += 1
              }
            }
          }
        }
      }

      if (tab[j][i] > 15) {
          for a := 0; a <= 1; a++ {
      	     for b := -1; b <= 0; b++ {
              	if (!(a == 0 && b == 0)) {
            	   if (j + a >= 0 && j + a < maxY && i + b >= 0 && i + b < maxX) {
	             buff[j + a][i + b] += 1
	        }
              }
           }
         }
      }

      if (c == 3) {
        buff[j][i] = tab[j][i] + 1
      } else if (tab[j][i] > 0 && c == 2) {
        buff[j][i] = tab[j][i] + 1
      }
    }
  }
  return buff
}

func fill_tab(tab [][]int) {
    rand.Seed(time.Now().UTC().UnixNano())
    for z := 0; z < 1000; z++ {
      tab[rand.Intn(maxY)][rand.Intn(maxX)] = 1
    }
    display(tab)
}

func main() {
  test := make([][]int, maxY)
  for i := 0; i < maxY; i++ {
    test[i] = make([]int, maxX)
  }


  fill_tab(test)

  for (true) {
    test = updateTab(test)
    display(test[:])
    time.Sleep(120 * time.Millisecond)
  }
}
