// Surface computes an SVG rendering of a 3-D surface function.

package main

import (
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/hashicorp/go-multierror"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

//make enum of plot types egg, moguls, and saddle
type Plot string

const (
	egg    = Plot("egg")
	moguls = Plot("moguls")
	saddle = Plot("saddle")
)

var plot Plot

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// get param plot from query string
	plot = Plot(r.URL.Query().Get("plot"))

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var errs error
			ax, ay, err := corner(i+1, j)
			if err != nil {
				errs = multierror.Append(errs, err)
			}
			bx, by, err := corner(i, j)
			if err != nil {
				errs = multierror.Append(errs, err)

			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				errs = multierror.Append(errs, err)

			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				errs = multierror.Append(errs, err)

			}

			if errs != nil {
				fmt.Printf("plot = %s  %v \n", plot, errs)
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, fmt.Errorf("invalid value %f", z)
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, nil
}

func f(x, y float64) float64 {
	var z float64

	switch plot {
	case egg:
		z = math.Pow(2, math.Sin(x)) * math.Pow(2, math.Sin(y)) / 12
	case moguls:
		z = math.Sin(x*y/10) / 10
	case saddle:
		r := math.Hypot(x, y) // distance from (0,0)
		z = math.Sin(-x) * math.Pow(1.5, -r)
	default:
		r := math.Hypot(x, y) // distance from (0,0)
		z = math.Sin(r) / r
	}

	return z
}
