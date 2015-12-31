//runs on port localhost:8000
package main

import (
	"math"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	width, height 	= 600,320	//canvas size in pixels
	cells		= 100			//number of grid cells
	xyrange 	= 30.0 					// axis ranges (-xyrange..+xyrange)
	xyscale		= width / 2 / xyrange	//pixels per x or y unit
	zscale		= height * 0.4 			//pixels per z unit
	angle		= math.Pi / 6 			//angle of x,y axes(=30 degrees)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //sin(30 degrees), cos(30 degrees)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/svg", svgMake);
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path= %q\n",r.URL.Path)
}
func svgMake(w http.ResponseWriter, r *http.Request) {
	
	//set the header content type to svg
	w.Header().Set("Content-Type","image/svg+xml")	

	if r.FormValue("cells") != "" {
		cells, err := strconv.Atoi(r.FormValue("cells"))
		if err != nil {
			log.Fatal("error in cells",err)
		}
		
		fmt.Fprintf(w,"<svg xmlns='http://www.w3.org/2000/svg' " +
			"style='stroke: grey; fill: white; stroke-width: 0.7' " +
			"width='%d' height='%d'>", width, height)
		
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, colorRedA := corner(i+1, j)
				bx, by, colorRedB := corner(i, j)
				cx, cy, colorRedC := corner(i, j+1)
				dx, dy, colorRedD := corner(i+1, j+1)
				if colorRedA && colorRedB && colorRedC && colorRedD {
					fmt.Fprintf(w, "<polygon points='%g %g %g %g %g %g %g %g' style='fill:#ff0000; stroke:black;'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				} else {
					fmt.Fprintf(w, "<polygon points='%g %g %g %g %g %g %g %g' style='fill:#0000ff; stroke:black;'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				}
			}
		}	
		fmt.Fprintf(w, "</svg>\n")
	} else {
		fmt.Fprintf(w,"<svg xmlns='http://www.w3.org/2000/svg' " +
			"style='stroke: grey; fill: white; stroke-width: 0.7' " +
			"width='%d' height='%d'>", width, height)
		
		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, colorRedA := corner(i+1, j)
				bx, by, colorRedB := corner(i, j)
				cx, cy, colorRedC := corner(i, j+1)
				dx, dy, colorRedD := corner(i+1, j+1)
				if colorRedA && colorRedB && colorRedC && colorRedD {
					fmt.Fprintf(w, "<polygon points='%g %g %g %g %g %g %g %g' style='fill:#ff0000; stroke:black;'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				} else {
					fmt.Fprintf(w, "<polygon points='%g %g %g %g %g %g %g %g' style='fill:#0000ff; stroke:black;'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
				}
			}
		}
		fmt.Fprintf(w, "</svg>\n")
	}
	
}

func corner(i, j int) (float64, float64, bool) {

	//find point (x,y) at corner of cell (i,j).
	isHigh := false
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	//Compute the surface height z
	z := f(x, y)
	if z > 0 {
		isHigh = true
	}
	//Project (x,y,z ) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale	

	return sx, sy, isHigh
}

func f(x, y float64) float64 {
	
	//distance from (0,0))
	r := math.Hypot(x, y)
	num := math.Sin(r) / r

	if !math.IsInf(num, 1) && !math.IsInf(num, -1) {
		return math.Sin(r) / r
	} else {
		return 0.0
	}
}
