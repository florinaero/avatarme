package avt

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"os"
	"strconv"
)

type coord struct {
	x int
	y int
}

var map_sqr = make(map[int]coord)
const WIDTH int = 512
const HEIGHT int = 512
const SQUARES int  = 16
const W_SQR = WIDTH/SQUARES
const H_SQR = HEIGHT/SQUARES

func GenerateHash(name *string, alg crypto.Hash) []byte {
	var alg_name string
	var res []byte
	input := []byte(*name)
	
	switch alg{
	case crypto.MD5:
		alg_name = "md5"
		r := md5.Sum(input)
		res = r[:]
	case crypto.SHA256:
		alg_name = "sha256"
		r := sha256.Sum256(input)
		res = r[:]
	default:
		return nil
	}
	
	fmt.Printf("\n%s = %x\n",alg_name,res)
	GenerateImage(*name+"_"+alg_name, res)
	return res
}

func GenerateImage(name string, hash []byte) {
	white := color.RGBA{255,255,255,0xff}
	upLeft := image.Point{0, 0}
	lowRight := image.Point{WIDTH, HEIGHT}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	
	color_background(img,lowRight,white)
	file, err := os.Create(name+".txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i:=0;i<len(hash);i++ {
		file.WriteString(strconv.Itoa(int(hash[i]))+"\n")
		color_sqr(img, int(hash[i]))
	}
	file.Close()
	

	f, _ := os.Create(name +".png")
	png.Encode(f,img)
}

func color_background(img *image.RGBA, size image.Point, clr color.RGBA) {
	for x:=0;x<size.X;x++ {
		for y:=0;y<size.Y;y++ {
			img.Set(x,y,clr)
		}
	}
}

func color_sqr(img *image.RGBA,sqr_no int) {
	red,_ := rand.Int(rand.Reader, big.NewInt(256))
	blue,_ := rand.Int(rand.Reader, big.NewInt(256))
	green,_ := rand.Int(rand.Reader, big.NewInt(256))
	r := red.Uint64()
	b := blue.Uint64()
	g := green.Uint64()

	clr := color.RGBA{uint8(r),uint8(b),uint8(g),0xff}

	x := map_sqr[sqr_no].x
	y := map_sqr[sqr_no].y

	for i:=x;i<x+W_SQR;i++ {
		for j:=y;j<y+H_SQR;j++ {
			img.Set(i,j,clr)
		}
	}
}

func Init() {
	map_img()
}

func map_img() {
	var count int

	for i:=0;i<SQUARES;i++ {
		for j:=0;j<SQUARES;j++ {
			map_sqr[count] = coord{j*W_SQR,i*H_SQR}
			count++
		}
	} 
}