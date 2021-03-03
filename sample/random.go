package sample

import (
	"math/rand"

	"github.com/aljo242/grpc/pb"
	"github.com/google/uuid"
)

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 0:
		return pb.Keyboard_QWERTY
	case 1:
		return pb.Keyboard_QWERTZ
	case 2:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_UNKNOWN
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD", "ARM")
}

func randomCPUName(brand string) string {
	switch brand {
	case "Intel":
		return randomStringFromSet(
			"Pentium Silver N6005",
			"Pentium Silver N6000",
			"Pentium Gold 7505",
			"Celeron N4500",
			"Celeron N4505",
			"Celeron N5105",
			"Celeron N5100",
			"Celeron 6305",
			"Celeron 6305E",
			"Xeon Platinum 8380HL",
			"Xeon Platinum 8353H",
			"Xeon Platinum 8380H",
			"Xeon Platinum 8354H",
			"Xeon Platinum 8376H",
			"Xeon Platinum 8376HL",
			"Xeon Platinum 8253",
			"Xeon Platinum 8256",
			"Xeon Platinum 8260",
			"Xeon Platinum 8276",
			"Xeon Platinum 8276L",
			"Xeon Gold 6328H",
			"Xeon Gold 6328HL",
			"Xeon Gold 5320H",
			"Xeon Gold 5318H",
			"Xeon Gold 6348H",
			"Xeon Gold 6246R",
			"Xeon Gold 6240R",
			"Xeon Gold 5220R",
			"Xeon Silver 4210R",
			"Xeon Silver 4214R",
			"Xeon Silver 4210T",
			"Xeon Silver 4215R",
			"Xeon Silver 4209T",
			"Xeon Bronze 3206R",
			"Xeon Bronze 3204",
			"Xeon Bronze 3106",
			"Xeon Bronze 3104",
			"Core i7-11375H",
			"Core i7-11370H",
			"Core i5-11300H",
			"Core i5-1145G7",
			"Core i3-1115G4",
			"Core i3-1110G4",
		)
	case "AMD":
		return randomStringFromSet(
			"Ryzen 7 5800X",
			"Ryzen 5 5600X",
			"Ryzen 9 5950X",
			"Ryzen 9 5900X",
			"Ryzen 9 5900",
			"Ryzen 9 3900XT",
			"Ryzen 7 3800XT",
			"Ryzen 5 3600XT",
		)
	case "ARM":
		return randomStringFromSet(
			"Cortex-A78",
			"Cortex-A78C",
			"Cortex-A78AE",
			"Cortex-A77",
			"Cortex-A76",
			"Cortex-A76AE",
		)
	}

	return ""
}

func randomGPUName(brand string) string {
	switch brand {
	case "Intel":
		return randomStringFromSet(
			"iRIS Xe MAX",
			"iRIS Xe",
		)
	case "AMD":
		return randomStringFromSet(
			"Radeon RX 6900XT",
			"Radeon RX 6800XT",
			"Radeon RX 6800",
			"Radeon RX 6700XT",
		)
	case "NVIDIA":
		return randomStringFromSet(
			"RTX 3090",
			"RTX 3080",
			"RTX 3070",
			"RTX 3060 Ti",
			"RTX 3060",
		)
	}

	return ""
}

func randomGPUBrand() string {
	return randomStringFromSet("Intel", "AMD", "NVIDIA")
}

func randomStringFromSet(list ...string) string {
	n := len(list)
	if n == 0 {
		return ""
	}

	return list[rand.Intn(n)]
}

/*
enum Panel {
        UNKNOWN = 0;
        TN = 1;
        VA = 2;
        IPS = 3;
        OLED = 4;
    }
*/
func randomScreenPanel() pb.Screen_Panel {
	switch rand.Intn(5) {
	case 0:
		return pb.Screen_TN
	case 1:
		return pb.Screen_VA
	case 2:
		return pb.Screen_IPS
	case 3:
		return pb.Screen_OLED
	}

	return pb.Screen_UNKNOWN
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16.0 / 9.0

	return &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Dell", "Microsoft", "ASUS", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Dell":
		return randomStringFromSet(
			"Inspiron 15 3000",
			"Inspiron 14 3000",
			"Inspiron 15 5000",
			"XPS 13",
		)
	case "Microsoft":
		return randomStringFromSet(
			"Suface Laptop Go",
			"Surface Laptop 3",
		)
	case "ASUS":
		return randomStringFromSet(
			"M515",
			"M415",
			"L410",
			"L210",
			"L510",
		)
	case "Lenovo":
		return randomStringFromSet(
			"Thinkbook 14",
			"T15",
			"X1 Carbon Gen 8",
			"X1 Carbon Gen 9",
			"IdeaPad 5",
		)
	}

	return ""
}
