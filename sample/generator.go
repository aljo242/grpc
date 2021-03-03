package sample

import (
	"github.com/aljo242/grpc/pb"
)

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	kb := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return kb
}

// NewCPU returns a new sample CPU
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	numberCores := randomInt(2, 16)
	numberThreads := randomInt(numberCores, 32)

	minGHz := randomFloat64(2.0, 3.5)
	maxGHz := randomFloat64(minGHz, 5.0)

	cpu := &pb.CPU{
		Brand:      brand,
		Name:       randomCPUName(brand),
		NumCores:   uint32(numberCores),
		NumThreads: uint32(numberThreads),
		MinGhz:     minGHz,
		MaxGhz:     maxGHz,
	}

	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()

	minGHz := randomFloat64(1.0, 1.5)
	maxGHz := randomFloat64(minGHz, 2.0)

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   randomGPUName(brand),
		MinGhz: minGHz,
		MaxGhz: maxGHz,
		Memory: &pb.Memory{
			Value: uint64(randomInt(2, 6)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return gpu
}

// NewRAM returns a new sample RAM
func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(4, 128)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new sample SSD
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new sample HDD
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen returns a new sample Screen
func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   randomFloat32(13, 27),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop returns a new sample Laptop
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightLb{
			WeightLb: randomFloat64(2, 6),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2021)),
	}

	return laptop
}
