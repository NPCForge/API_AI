package utils

import (
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

func LogErrorUtils(context string, err error) {
	if err != nil {
		_, file, line, ok := runtime.Caller(1)
		if ok {
			color.Red("❌ Error in %s: %s (at %s:%d)", context, err.Error(), file, line)
		} else {
			color.Red("❌ Error in %s: %s (location unknown)", context, err.Error())
		}
	}
}

func SendResponse(conn *websocket.Conn, initialRoute string, entityChecksum string, fields map[string]interface{}) {
	resp := map[string]interface{}{
		"status":   "success",
		"route":    initialRoute,
		"checksum": entityChecksum,
	}

	for key, value := range fields {
		resp[key] = value
	}

	color.Green("✅ Sending success response:")
	color.Cyan("📤 Payload: %+v", resp)

	conn.WriteJSON(resp)
}

func SendError(conn *websocket.Conn, initialRoute string, entityChecksum string, fields map[string]interface{}) {
	resp := map[string]interface{}{
		"status":   "error",
		"route":    initialRoute,
		"checksum": entityChecksum,
	}

	for key, value := range fields {
		resp[key] = value
	}

	color.Red("🚨 Sending error response:")
	color.Yellow("📤 Payload: %+v", resp)

	conn.WriteJSON(resp)
}

func IsRunningInDocker() bool {
	_, err := os.Stat("/.dockerenv")
	return err == nil
}
