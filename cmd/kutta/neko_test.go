package main

import (
	"os"
	"strings"
	"testing"

	"kutta/sceneio"
)

func TestMenuIsStableAndSignatureTracksContext(t *testing.T) {
	// The top-level bar is stable across contexts so it never changes shape
	// underfoot; items enable/disable instead.
	want := []string{"File", "Edit", "View", "Foil", "Animate"}
	g := editGame()
	g.scn = nil // interactive foil
	for _, w := range want {
		if !strings.Contains(menuTitles(g), w) {
			t.Errorf("foil-mode menu missing %q (got %q)", w, menuTitles(g))
		}
	}
	g2 := editGame() // scene loaded, editor open, animate mode
	g2.editing = true
	g2.editMode = emAnimate
	for _, w := range want {
		if !strings.Contains(menuTitles(g2), w) {
			t.Errorf("editor-animate menu missing %q (got %q)", w, menuTitles(g2))
		}
	}
	// The signature tracks context so the menu rebuilds when state changes.
	s1 := g.menuSignature()
	g.mode = modeVorticity
	if g.menuSignature() == s1 {
		t.Error("menu signature did not change with the field mode")
	}
}

func menuTitles(g *Game) string {
	var b strings.Builder
	for _, it := range g.menuItems() {
		b.WriteString(it.Title + " ")
	}
	return b.String()
}

func TestNekoScene(t *testing.T) {
	sc := nekoScene()
	if len(sc.Objects) != 1 || len(sc.Objects[0].Shape) != 25 {
		t.Fatalf("neko shape: %d objects, %d points", len(sc.Objects), len(sc.Objects[0].Shape))
	}
	solid := 0
	for _, b := range sc.Mask(0, gridW, gridH) {
		if b {
			solid++
		}
	}
	if solid == 0 {
		t.Error("neko rasterized to nothing")
	}
	if solid > gridW*gridH/2 {
		t.Errorf("neko fills too much of the domain: %d cells", solid)
	}
}

// TestNekoFileMatches guards the examples/neko.afoil demo against drifting from
// the Go nekoScene (the easter-egg source).
func TestNekoFileMatches(t *testing.T) {
	b, err := os.ReadFile("../../examples/neko.afoil")
	if err != nil {
		t.Fatal(err)
	}
	sc, err := sceneio.Load(string(b))
	if err != nil {
		t.Fatalf("neko.afoil does not parse: %v", err)
	}
	got, want := len(sc.Objects[0].Shape), len(nekoScene().Objects[0].Shape)
	if got != want {
		t.Errorf("neko.afoil has %d points, nekoScene has %d", got, want)
	}
}
