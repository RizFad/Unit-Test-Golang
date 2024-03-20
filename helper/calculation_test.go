package helper

import (
	"testing"
)

func TestKubus(t *testing.T) {
	t.Run("calculate volume", func(t *testing.T) {
		kubus := Kubus{Sisi: 3}
		expectedVolume := 27.0
		volume := kubus.Volume()
		if volume != expectedVolume {
			t.Errorf("Volume calculation is incorrect, got: %.2f, want: %.2f", volume, expectedVolume)
		}
	})

	t.Run("calculate surface area", func(t *testing.T) {
		kubus := Kubus{Sisi: 3}
		expectedLuas := 54.0
		luas := kubus.Luas()
		if luas != expectedLuas {
			t.Errorf("Surface area calculation is incorrect, got: %.2f, want: %.2f", luas, expectedLuas)
		}
	})

	t.Run("calculate perimeter", func(t *testing.T) {
		kubus := Kubus{Sisi: 3}
		expectedKeliling := 36.0
		keliling := kubus.Keliling()
		if keliling != expectedKeliling {
			t.Errorf("Perimeter calculation is incorrect, got: %.2f, want: %.2f", keliling, expectedKeliling)
		}
	})
}
