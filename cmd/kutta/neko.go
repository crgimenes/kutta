package main

import (
	"kutta/foil"
	"kutta/scene"
)

// nekoScene is the easter-egg neko (cat) — a side-view silhouette treated as an
// "airfoil". It began life as a cow, but everyone agreed it looks like a cat, so
// a cat it is. The generic object model handles any closed shape, so the solver
// happily computes the flow around it (spoiler: cats are terrible airfoils). It
// faces upstream (nose at low x), with two legs, ears and a tail; loop 0, so it
// just sits in the wind. Trigger it by typing "neko" in the NACA field or by
// opening examples/neko.afoil.
func nekoScene() *scene.Scene {
	p := func(x, y float64) foil.Point { return foil.Point{X: x, Y: y} }
	pts := []foil.Point{
		p(118, 100), // nose tip
		p(124, 110), // muzzle top
		p(130, 114), // forehead
		p(133, 126), // left ear
		p(138, 116), // between ears
		p(142, 128), // right ear
		p(147, 116), // neck top
		p(160, 124), // withers
		p(200, 126), // back
		p(224, 122), // rump
		p(232, 120), // tail base top
		p(240, 132), // tail tip
		p(236, 108), // tail mid
		p(231, 100), // tail base bottom
		p(228, 98),  // hind thigh back
		p(226, 72),  // hind foot back
		p(219, 72),  // hind foot front
		p(217, 96),  // hind thigh front
		p(200, 92),  // belly rear
		p(152, 92),  // belly front
		p(150, 72),  // front foot back
		p(143, 72),  // front foot front
		p(141, 95),  // chest bottom
		p(132, 94),  // lower neck
		p(126, 92),  // chin
	}
	return &scene.Scene{Objects: []*scene.Object{{
		Name:  "neko",
		Shape: pts,
		Pivot: foil.Point{X: 180, Y: 100},
	}}}
}
