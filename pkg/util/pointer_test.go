package util

import "testing"

// P関数のテストを書く
func TestP(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		v := 42
		p := P(v)
		if *p != v {
			t.Errorf("expected %d, got %d", v, *p)
		}
	})

	t.Run("string", func(t *testing.T) {
		v := "hello"
		p := P(v)
		if *p != v {
			t.Errorf("expected %s, got %s", v, *p)
		}
	})
}

// PorNil 関数のテストを書く
func TestPorNil(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		v := 42
		p := PorNil(v, 0)
		if *p != v {
			t.Errorf("expected %d, got %d", v, *p)
		}
	})

	t.Run("string", func(t *testing.T) {
		v := "hello"
		p := PorNil(v, "")
		if *p != v {
			t.Errorf("expected %s, got %s", v, *p)
		}
	})
	// uint型のテストを追加します
	t.Run("uint", func(t *testing.T) {
		v := uint(42)
		p := PorNil(v, 0)
		if *p != v {
			t.Errorf("expected %d, got %d", v, *p)
		}
	})
}
