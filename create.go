package main

import (
    "os"
    "image"
    "image/color"
    "image/png"
    "fmt"
)

var width int
var height int

func createImage(number int) {

width = 200
height = 200

upLeft := image.Point{0, 0}
lowRight := image.Point{width, height}

img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

// Colors are defined by Red, Green, Blue, Alpha uint8 values.
cyan := color.RGBA{100, 200, 200, 0xff}

// Set Background for each pixel.
for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
            img.Set(x, y, cyan)
    }
}
_, err := os.Stat("images")

if os.IsNotExist(err) {
    errDir := os.MkdirAll("images", 0755)
    if errDir != nil {
        fmt.Println(err)
    }
}

var count int = 0
for {
    if number == 0 {
        count = 0
        break
    }
    rem := number % 10
    number = number / 10
    // Getting or setting the number
    setDigit(rem, img)
    // Name of the image
image_name := fmt.Sprintf("%d_%d.png", count, rem)
// Encode as PNG.
f, _ := os.Create("images/" + image_name)
png.Encode(f, img)
count++
}
}

func setDigit(number int, img *image.RGBA){
    // Set Background for each pixel.
for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
            img.Set(x, y, color.RGBA{100, 200, 200, 0xff})
    }
}
    switch number {
    case 1:
        setOne(img)
    case 2:
        setTwo(img)
    case 3:
        setThree(img)
    case 4:
        setFour(img)
    case 5:
        setFive(img)
    case 6:
        setSix(img)
    case 7:
        setSeven(img)
    case 8:
        setEight(img)
    case 9:
        setNine(img)
    case 0:
        setZero(img)
    }
}

func setZero(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch  {
            // Top Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y > 20 && y < 110:
                img.Set(x, y, color.Black)
            // Bottom Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            // Top Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 20 && y < 30):
                img.Set(x, y, color.Black)
            // Bottom Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 180 && y < 190):
                img.Set(x, y, color.Black)
            // Top Left Side Bar
            case x >= width/2 - 40 && x <= width/2 - 30 && y > 20 && y < 110:
                img.Set(x, y, color.Black)
            // Bottom Left Side Bar
            case x >= width/2 - 40 && x <= width/2 - 30 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            }
        }
    }
}

func setOne(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch  {
            // Top Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y > 20 && y < 110:
                img.Set(x, y, color.Black)
            // Bottom Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            }
        }
    }
}

func setTwo(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch  {
            // Top Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y > 20 && y < 100:
                img.Set(x, y, color.Black)
            // Top Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 20 && y < 30):
                img.Set(x, y, color.Black)
            // Bottom Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 180 && y < 190):
                img.Set(x, y, color.Black)
            // Center Bar
            case x >= width/2 - 40 && x <= width/2 + 40 && y >= 100 && y < 110:
                img.Set(x, y, color.Black)
            // Bottom Left Side Bar
            case x >= width/2 - 40 && x <= width/2 - 30 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            }
        }
    }
}

func setThree(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch  {
            // Top Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y > 20 && y < 100:
                img.Set(x, y, color.Black)
            // Bottom Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            // Top Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 20 && y < 30):
                img.Set(x, y, color.Black)
            // Bottom Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 180 && y < 190):
                img.Set(x, y, color.Black)
            // Center Bar
            case x >= width/2 - 40 && x <= width/2 + 40 && y >= 100 && y < 110:
                img.Set(x, y, color.Black)
            }
        }
    }
}

func setFour(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch  {
            // Top Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y > 20 && y < 100:
                img.Set(x, y, color.Black)
            // Bottom Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            // Center Bar
            case x >= width/2 - 40 && x <= width/2 + 40 && y >= 100 && y < 110:
                img.Set(x, y, color.Black)
            // Top Left Side Bar
            case x >= width/2 - 40 && x <= width/2 - 30 && y > 20 && y < 100:
                img.Set(x, y, color.Black)
            }
        }
    }
}

func setFive(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch  {
            // Bottom Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            // Top Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 20 && y < 30):
                img.Set(x, y, color.Black)
            // Bottom Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 180 && y < 190):
                img.Set(x, y, color.Black)
            // Center Bar
            case x >= width/2 - 40 && x <= width/2 + 40 && y >= 100 && y < 110:
                img.Set(x, y, color.Black)
            // Top Left Side Bar
            case x >= width/2 - 40 && x <= width/2 - 30 && y > 20 && y < 100:
                img.Set(x, y, color.Black)
            }
        }
    }
}

func setSix(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch  {
            // Bottom Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            // Top Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 20 && y < 30):
                img.Set(x, y, color.Black)
            // Bottom Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 180 && y < 190):
                img.Set(x, y, color.Black)
            // Center Bar
            case x >= width/2 - 40 && x <= width/2 + 40 && y >= 100 && y < 110:
                img.Set(x, y, color.Black)
            // Top Left Side Bar
            case x >= width/2 - 40 && x <= width/2 - 30 && y > 20 && y < 100:
                img.Set(x, y, color.Black)
            // Bottom Left Side Bar
            case x >= width/2 - 40 && x <= width/2 - 30 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            }
        }
    }
}


func setSeven(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch  {
            // Top Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y > 20 && y < 110:
                img.Set(x, y, color.Black)
            // Bottom Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            // Top Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 20 && y < 30):
                img.Set(x, y, color.Black)
            }
        }
    }
}

func setEight(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch  {
            // Top Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y > 20 && y < 100:
                img.Set(x, y, color.Black)
            // Bottom Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            // Top Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 20 && y < 30):
                img.Set(x, y, color.Black)
            // Bottom Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 180 && y < 190):
                img.Set(x, y, color.Black)
            // Center Bar
            case x >= width/2 - 40 && x <= width/2 + 40 && y >= 100 && y < 110:
                img.Set(x, y, color.Black)
            // Top Left Side Bar
            case x >= width/2 - 40 && x <= width/2 - 30 && y > 20 && y < 100:
                img.Set(x, y, color.Black)
            // Bottom Left Side Bar
            case x >= width/2 - 40 && x <= width/2 - 30 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            }
        }
    }
}


func setNine(img *image.RGBA) {
    for x := 0; x < width; x++ {
        for y := 0; y < height; y++ {
            switch  {
            // Top Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y > 20 && y < 100:
                img.Set(x, y, color.Black)
            // Bottom Right Side Bar
            case x >= width/2 + 30 && x <= width/2 + 40 && y >= 110 && y < 190:
                img.Set(x, y, color.Black)
            // Top Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 20 && y < 30):
                img.Set(x, y, color.Black)
            // Bottom Bar
            case (x >= width/2 - 40 && x <= width/2 + 40 && y > 180 && y < 190):
                img.Set(x, y, color.Black)
            // Center Bar
            case x >= width/2 - 40 && x <= width/2 + 40 && y >= 100 && y < 110:
                img.Set(x, y, color.Black)
            // Top Left Side Bar
            case x >= width/2 - 40 && x <= width/2 - 30 && y > 20 && y < 100:
                img.Set(x, y, color.Black)
            }
        }
    }
}

