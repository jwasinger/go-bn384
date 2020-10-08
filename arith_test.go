package go_evm384

import (
	"testing"
)

func TestMulMod_BLS12381(t *testing.T) {
	x := Element{0xb1f598e5f390298f, 0x6b3088c3a380f4b8, 0x4d10c051c1fa23c0, 0x2945981a13aec13, 0x3bcea128c5c8d172, 0xdaa35e7a880a2ca}
	y := Element{0x4c64af08c847d3ec, 0xf47665551a973a7a, 0x4f0090b4b602e334, 0x670a33daa7a418b4, 0x8b9b1631a9ecad43, 0x15e1e13af71de992}
	expected := Element{0x20b39e434f6b7627, 0xe3b9585c3bc798c3, 0xd601841435360731, 0x592efb881d54c66d, 0x8ba6599731e3b7f3, 0x8e7724179630faa}
	mod := Element{0xb9feffffffffaaab, 0x1eabfffeb153ffff, 0x6730d2a0f6b0f624, 0x64774b84f38512bf, 0x4b1ba7b6434bacd7, 0x1a0111ea397fe69a}
	var inv uint64
	inv = 0x89f3fffcfffcfffd
	out := Element{0,0,0,0,0,0}

	MulMod(&out, &x, &y, &mod, inv)

	for i := 0; i < 6; i++ {
		if out[i] != expected[i] {
			t.Fatalf("invalid result %x (expected) != %x", expected[i], out[i])
		}
	}
}

func TestAddMod_BLS12381(t *testing.T) {
	mod := Element{0xb9feffffffffaaab, 0x1eabfffeb153ffff, 0x6730d2a0f6b0f624, 0x64774b84f38512bf, 0x4b1ba7b6434bacd7, 0x1a0111ea397fe69a}
	out := Element{0,0,0,0,0,0}

	x := Element{0x20b39e434f6b7627, 0xe3b9585c3bc798c3, 0xd601841435360731, 0x592efb881d54c66d, 0x8ba6599731e3b7f3, 0x08e7724179630faa}
	y := Element{0xd2f66b13d3e3cc9e, 0xc4ad7d09d3b8497d, 0xfc3bcaaeef9fd81e, 0x55ff24e182d1d704, 0x0c05276fe4bb9c8f, 0x12c30706122c2df5}
	expected := Element{0x39ab0957234f981a, 0x89bad5675e2be241, 0x6b0c7c222e24e92c, 0x4ab6d4e4aca18ab3, 0x4c8fd950d353a7ab, 0x01a9675d520f5705}
        AddMod(&out, &x, &y, &mod)

	for i := 0; i < 6; i++ {
		if out[i] != expected[i] {
			t.Fatalf("invalid result %x (expected) != %x", expected[i], out[i])
		}
	}
}

func TestSubMod_BLS12381(t *testing.T) {
	mod := Element{0xb9feffffffffaaab, 0x1eabfffeb153ffff, 0x6730d2a0f6b0f624, 0x64774b84f38512bf, 0x4b1ba7b6434bacd7, 0x1a0111ea397fe69a}
	out := Element{0,0,0,0,0,0}

	x := Element{0xa584d3cbca05d33a, 0x355e0a6052a33624, 0x09144f6bf33b19ed, 0x706873aa9aa513e5, 0xc1cdbc99214d26ea, 0x0673294630c0dcf5}
	y := Element{0x07bc332f7b875434, 0x3db7db5119634f45, 0x40f68c063c472537, 0x67a7222b8e080228, 0xcabcd9dd9073c83b, 0x10257d25a0b6c84f}
	expected := Element{0x57c7a09c4e7e29b1, 0x16522f0dea93e6df, 0x2f4e9606ada4eada, 0x6d389d040022247c, 0x422c8a71d4250b86, 0x104ebe0ac989fb40}
        SubMod(&out, &x, &y, &mod)

	for i := 0; i < 6; i++ {
		if out[i] != expected[i] {
			t.Fatalf("invalid result %x (expected) != %x", expected[i], out[i])
		}
	}
}

func BenchmarkMulMod_BLS12381(b *testing.B) {
	x := Element{0xb1f598e5f390298f, 0x6b3088c3a380f4b8, 0x4d10c051c1fa23c0, 0x2945981a13aec13, 0x3bcea128c5c8d172, 0xdaa35e7a880a2ca}
	y := Element{0x4c64af08c847d3ec, 0xf47665551a973a7a, 0x4f0090b4b602e334, 0x670a33daa7a418b4, 0x8b9b1631a9ecad43, 0x15e1e13af71de992}
	mod := Element{0xb9feffffffffaaab, 0x1eabfffeb153ffff, 0x6730d2a0f6b0f624, 0x64774b84f38512bf, 0x4b1ba7b6434bacd7, 0x1a0111ea397fe69a}
	var inv uint64
	inv = 0x89f3fffcfffcfffd

	for n := 0; n < b.N; n++ {
		MulMod(&x, &x, &y, &mod, inv)
	}
}
