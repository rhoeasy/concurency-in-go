package pool

import "testing"

func TestNewPool(t *testing.T) {
	type poolConfig struct {
		size int
		len  int
	}
	tc := []struct {
		name string
		in   poolConfig
		want poolConfig
	}{
		{
			name: "test case 1",
			in: poolConfig{
				size: 5,
				len:  4,
			},
			want: poolConfig{
				size: 5,
				len:  4,
			},
		},
	}
	for _, tt := range tc {
		p := NewPool(tt.in.size, tt.in.len)
		if len(p.works) != tt.want.size {
			t.Errorf("test case %v failed, want size: %v, got size: %v", tt.name, tt.want.size, len(p.works))
		}

		for i := 0; i < tt.want.len; i++ {
			p.works[0] <- "a"
		}
		if len(p.works[0]) != tt.want.len {
			t.Errorf("test case %v failed, want queue len: %v, got queue len: %v", tt.name, tt.want.len, len(p.works[0]))
		}
	}
}
