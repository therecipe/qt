package flutter

// #cgo darwin LDFLAGS: -F ${SRCDIR} -framework FlutterEmbedder -Wl,-rpath,${SRCDIR}
// #cgo windows LDFLAGS: -L ${SRCDIR} -l flutter_engine
// #cgo linux LDFLAGS: -L ${SRCDIR} -l flutter_engine -Wl,-rpath,${SRCDIR}
/*
#include "stdlib.h"
#include "embedder.h"

FlutterEngine enginePtr;

extern bool make_current(void*);
extern bool clear_current(void*);
extern bool present(void*);
extern uint32_t fbo_callback(void*);
extern bool make_resource_current(void*);
extern void* gl_proc_resolver(void*, char*);
extern void platform_message_callback(FlutterPlatformMessage*, void*);
extern void data_callback(uint8_t*, size_t, void*);

void RunFlutterEngine(char* assets_path, char* icu_data_path) {

	FlutterRendererConfig config = {};
	config.type = kOpenGL;
	config.open_gl.struct_size = sizeof(config.open_gl);
	config.open_gl.make_current = make_current;
	config.open_gl.clear_current = clear_current;
	config.open_gl.present = present;
	config.open_gl.fbo_callback = fbo_callback;
	config.open_gl.make_resource_current = make_resource_current;
	config.open_gl.gl_proc_resolver = gl_proc_resolver;

	FlutterProjectArgs args = {};
	args.struct_size = sizeof(args);
	args.assets_path = assets_path;
	args.icu_data_path = icu_data_path;
	args.platform_message_callback = platform_message_callback;

	FlutterEngineRun(1, &config, &args, enginePtr, &enginePtr);
}
*/
import "C"
import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"time"
	"unsafe"

	"github.com/bluszcz/cutego/core"
)

func FlutterEngineRun() {

	var (
		assets_path   string
		icu_data_path string
	)

	pwd, _ := os.Getwd()
	for _, path := range []string{pwd, core.QCoreApplication_ApplicationDirPath()} {
		assets_path = filepath.Join(path, "build", "flutter_assets")
		if _, err := os.Stat(assets_path); err != nil {
			assets_path = filepath.Join(path, "flutter_assets")
		}

		//TODO: support flutter-rs aot engine builds as well
		switch runtime.GOOS {
		case "darwin":
			icu_data_path = filepath.Join(path, "FlutterEmbedder.framework", "Versions", "A", "Resources", "icudtl.dat")
			if _, err := os.Stat(icu_data_path); err != nil {
				icu_data_path = filepath.Join(path, "..", "Frameworks", "FlutterEmbedder.framework", "Versions", "A", "Resources", "icudtl.dat")
			}
		case "windows", "linux":
			icu_data_path = filepath.Join(path, "icudtl.dat")
		}

		if _, err := os.Stat(assets_path); err == nil {
			break
		}
	}

	C.RunFlutterEngine(C.CString(assets_path), C.CString(icu_data_path))
	FlutterEngineSendWindowMetricsEvent(win.Width(), win.Height())
}

func FlutterEngineSendWindowMetricsEvent(w, h int) {
	e := C.FlutterWindowMetricsEvent{}
	e.struct_size = C.size_t(unsafe.Sizeof(e))
	e.width = C.size_t(float64(w) * devicePixelRatio)
	e.height = C.size_t(float64(h) * devicePixelRatio)
	e.pixel_ratio = C.double(devicePixelRatio)
	C.FlutterEngineSendWindowMetricsEvent(C.enginePtr, &e)
}

var pointerPhase C.FlutterPointerPhase = C.kCancel

func FlutterEngineSendPointerEvent(phase C.FlutterPointerPhase, x int, y int) {
	switch phase {
	case C.kAdd, C.kHover:

		if !(pointerPhase == C.kCancel || pointerPhase == C.kRemove) {
			if phase == C.kHover && pointerPhase == C.kAdd { //needed for resizing issues
				break
			}
			return
		} else {
			phase = C.kAdd
			pointerPhase = phase
		}

	case C.kRemove:

		if !(pointerPhase == C.kCancel || pointerPhase == C.kAdd) {
			return
		} else {
			pointerPhase = phase
		}

	default:
		if !(pointerPhase == C.kCancel || pointerPhase == C.kAdd) {
			return
		}
	}

	e := C.FlutterPointerEvent{}
	e.struct_size = C.size_t(unsafe.Sizeof(e))
	e.phase = phase
	e.x = C.double(float64(x) * devicePixelRatio)
	e.y = C.double(float64(y) * devicePixelRatio)
	e.timestamp = C.size_t(time.Now().Unix() / 1e3)
	C.FlutterEngineSendPointerEvent(C.enginePtr, &e, 1)
}

func FlutterEngineSendPointerScrollEvent(x int, y int, xd int, yd int) {
	if !(pointerPhase == C.kCancel || pointerPhase == C.kAdd) {
		return
	}

	e := C.FlutterPointerEvent{}
	e.struct_size = C.size_t(unsafe.Sizeof(e))
	e.phase = C.kHover
	e.x = C.double(float64(x) * devicePixelRatio)
	e.y = C.double(float64(y) * devicePixelRatio)
	e.scroll_delta_x = C.double(xd)
	e.scroll_delta_y = C.double(yd)
	e.signal_kind = C.kFlutterPointerSignalKindScroll
	e.timestamp = C.size_t(time.Now().Unix() / 1e3)
	C.FlutterEngineSendPointerEvent(C.enginePtr, &e, 1)
}

func FlutterEngineSendPlatformMessage(ch string, msg []byte) {
	var handle *C.FlutterPlatformMessageResponseHandle
	C.FlutterPlatformMessageCreateResponseHandle(C.enginePtr, (*[0]byte)(C.data_callback), unsafe.Pointer(C.enginePtr), &handle)

	channel := C.CString(ch)
	message := C.CBytes(msg)

	cPlatformMessage := C.FlutterPlatformMessage{}
	cPlatformMessage.struct_size = C.size_t(unsafe.Sizeof(cPlatformMessage))
	cPlatformMessage.channel = channel
	cPlatformMessage.message = (*C.uint8_t)(message)
	cPlatformMessage.message_size = C.size_t(len(msg))
	cPlatformMessage.response_handle = (*C.FlutterPlatformMessageResponseHandle)(unsafe.Pointer(handle))
	C.FlutterEngineSendPlatformMessage(C.enginePtr, &cPlatformMessage)

	C.free(unsafe.Pointer(channel))
	C.free(unsafe.Pointer(message))
	C.FlutterPlatformMessageReleaseResponseHandle(C.enginePtr, handle)
}

func handleKeyEvent(key int, typeKey string, scancode uint) {

	type MethodCall struct {
		Method string      `json:"method"`
		Args   interface{} `json:"args"`
	}

	switch core.Qt__Key(key) {
	case core.Qt__Key_Escape:
		msg, _ := json.Marshal(MethodCall{Method: "popRoute"})
		FlutterEngineSendPlatformMessage("flutter/navigation", msg)

	case core.Qt__Key_Enter, core.Qt__Key_Return:
		msg, _ := json.Marshal(MethodCall{Method: "TextInputClient.performAction", Args: []interface{}{textInputHelper.Property("clientID").ToInterface(), textInputHelper.Property("action").ToInterface()}})
		FlutterEngineSendPlatformMessage("flutter/textinput", msg)

	default:

		km := runtime.GOOS
		if km == "darwin" {
			km = "macos"
		}

		event := struct {
			KeyCode   int    `json:"keyCode"`
			Type      string `json:"type"`
			ScanCode  int    `json:"scanCode"`
			Modifiers int    `json:"modifiers"`
			Keymap    string `json:"keymap"`
		}{
			KeyCode:   int(key),
			Type:      typeKey,
			ScanCode:  int(scancode),
			Modifiers: int(0), //TODO:
			Keymap:    km,
		}

		msg, _ := json.Marshal(event)
		FlutterEngineSendPlatformMessage("flutter/keyevent", msg)

		if textInputHelper != nil && textInputHelper.Property("clientID").ToDouble(nil) == 0 {
			return
		}

		state := struct {
			Text              string `json:"text"`
			SelectionBase     int    `json:"selectionBase"`
			SelectionExtent   int    `json:"selectionExtent"`
			SelectionAffinity string `json:"selectionAffinity"`
		}{
			Text:              textInputHelper.ToPlainText(),
			SelectionBase:     textInputHelper.TextCursor().SelectionStart(),
			SelectionExtent:   textInputHelper.TextCursor().SelectionEnd(),
			SelectionAffinity: "TextAffinity.downstream",
		}

		msg, _ = json.Marshal(MethodCall{Method: "TextInputClient.updateEditingState", Args: []interface{}{textInputHelper.Property("clientID").ToInterface(), state}})
		FlutterEngineSendPlatformMessage("flutter/textinput", msg)
	}
}
