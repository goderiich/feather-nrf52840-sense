package displayservice

import (
	"machine"
	"time"

	font "github.com/Nondzu/ssd1306_font"
	"tinygo.org/x/drivers/ssd1306"
)

func DisplayInit() (font.Display, ssd1306.Device) {
	//Start Display Init
	time.Sleep(time.Millisecond * 100) // Please wait some time after turning on the device to properly initialize the display
	machine.I2C0.Configure(machine.I2CConfig{Frequency: 400000})

	// Display
	dev := ssd1306.NewI2C(machine.I2C0)
	dev.Configure(ssd1306.Config{
		Address: ssd1306.Address_128_32,
		Width:   128,
		Height:  32,
	})
	dev.ClearBuffer()
	dev.ClearDisplay()

	//font library init
	display := font.NewDisplay(dev)
	display.Configure(font.Config{FontType: font.FONT_7x10}) //set font here

	display.YPos = 0 // set position Y
	display.XPos = 0 // set position X
	//Finished Display Init
	return display, dev
}
