package pkg

import (
	"fmt"
	"os"
)

// DisplayType represents a display type with a name, color, and icon associated.
type DisplayType struct {
	Name  string
	Color string
	Icon  string
}

// Constants defining the different display types.
var (
	Debug   = DisplayType{Name: "debug", Color: ColorCodes.Orange, Icon: "🧰"}
	Update  = DisplayType{Name: "update", Color: ColorCodes.Green, Icon: "🔄"}
	File    = DisplayType{Name: "file", Color: ColorCodes.LightGreen, Icon: "--- "}
	Default = DisplayType{Name: "default", Color: ColorCodes.Blue, Icon: "⚙️"}
	Error   = DisplayType{Name: "error", Color: ColorCodes.Red, Icon: "❌"}
)

// DisplayContext displays a message in the terminal with contextual information.
// Usage examples:
// DisplayContext("Hello world!", Default)  // Display a default message
// DisplayContext("Hello world!", Error, err, true)  // Display an error message
// DisplayContext("Hello world!", Update)  // Display an update message
// DisplayContext("Hello world!", Debug)  // Display a debug message
// Syntax: DisplayContext("message", type, error (optional), exit (boolean, optional))
func DisplayContext(context string, displayType DisplayType, args ...any) {
	// Display the main message with the associated color and icon
	fmt.Printf("%s%s: %s%s\n", displayType.Color, displayType.Icon, context, ColorCodes.Reset)

	for _, arg := range args {
		switch v := arg.(type) {
		case bool:
			if v {
				os.Exit(0) // Exit successfully (no error)
			}
		case error:
			// If an error is passed, display the error and exit with an error code
			fmt.Printf("%s%s%s\n", Error.Color, v.Error(), ColorCodes.Reset)
		}
	}
}
