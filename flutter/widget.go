package flutter

//#include "embedder.h"
import "C"

import (
	"encoding/json"
	"os"
	"runtime"
	"sync"
	"time"
	"unsafe"

	"github.com/bluszcz/cutego/core"
	"github.com/bluszcz/cutego/gui"
	"github.com/bluszcz/cutego/widgets"

	"github.com/bluszcz/cutego/interop"
	_ "github.com/bluszcz/cutego/interop/dart"
)

var (
	USE_ALTERNATIVE_OFFSCREEN_RENDERING = (runtime.GOOS == "windows" && os.Getenv("QT_ALT_RENDERING") != "false") || os.Getenv("QT_ALT_RENDERING") == "true"
	DEBUG                               = DEBUG_HOT || false
	DEBUG_HOT                           = false
	IGNORE_THREADS                      = true //TODO: make resizing thread safe; (race condition with the flutter and the main thread sharing the widgets openglcontext)
	TEST_RESIZING                       = false
)

var (
	win *FlutterWidget

	devicePixelRatio float64

	frameMutex         = new(sync.Mutex)
	offScreenSur       *gui.QOffscreenSurface
	offScreenCtx       *gui.QOpenGLContext
	offScreenCtxShared *gui.QOpenGLContext
	offScreenFBO       *gui.QOpenGLFramebufferObject
	offScreenFBO_Frame *gui.QImage

	textInputHelper *widgets.QPlainTextEdit

	resizePhase int
)

//

func init() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_DontCheckOpenGLContextThreadAffinity, IGNORE_THREADS)
	if runtime.GOOS == "windows" && os.Getenv("COMPUTERNAME") == "VAGRANT-2012" {
		core.QCoreApplication_SetAttribute(core.Qt__AA_UseSoftwareOpenGL, true)
	}

	//TODO: manually set the format for each surface?
	format := gui.QSurfaceFormat_DefaultFormat()
	format.SetVersion(3, 2)
	format.SetProfile(gui.QSurfaceFormat__CoreProfile)
	gui.QSurfaceFormat_SetDefaultFormat(format)
}

type FlutterWidget struct{ *widgets.QOpenGLWidget }

func NewFlutterWidget(parent widgets.QWidget_ITF) *FlutterWidget {

	if USE_ALTERNATIVE_OFFSCREEN_RENDERING {
		devicePixelRatio = 1
	} else {
		devicePixelRatio = gui.QGuiApplication_PrimaryScreen().DevicePixelRatio()
	}

	win = &FlutterWidget{widgets.NewQOpenGLWidget(parent, 0)}
	win.SetMouseTracking(true)
	win.Resize2(400, 300)

	offScreenCtx = gui.NewQOpenGLContext(nil)
	ok := offScreenCtx.Create()
	if !ok {
		println("couldn't create valid offScreenCtx")
	}

	offScreenCtxShared = gui.NewQOpenGLContext(nil)
	if USE_ALTERNATIVE_OFFSCREEN_RENDERING {
		offScreenCtxShared.SetShareContext(offScreenCtx)
	} else {
		offScreenCtxShared.SetShareContext(win.Context())
	}
	ok = offScreenCtxShared.Create()
	if !ok {
		println("couldn't create valid offScreenCtxShared")
	}

	offScreenSur = gui.NewQOffscreenSurface2(nil)
	offScreenSur.Create()

	//

	//paint events

	win.ConnectPaintEvent(func(event *gui.QPaintEvent) {
		if !USE_ALTERNATIVE_OFFSCREEN_RENDERING || offScreenFBO_Frame.Pointer() == nil {
			return
		}
		painter := gui.NewQPainter2(win)
		painter.DrawImage6(core.NewQRect4(0, 0, int(win.Width()), int(win.Height())), offScreenFBO_Frame)
		painter.DestroyQPainter()
	})

	win.ConnectAboutToCompose(func() {
		frameMutex.Lock()
		//win.MakeCurrent() //TODO: needed for windows 10 with def rendering ?
	})

	win.ConnectFrameSwapped(func() {
		//win.DoneCurrent() //TODO: needed for windows 10 with def rendering ?
		resizePhase = 0
		frameMutex.Unlock()
	})

	win.ConnectResizeEvent(func(event *gui.QResizeEvent) {
		frameMutex.Lock()
		defer frameMutex.Unlock()

		if resizePhase != 0 {
			return
		}
		resizePhase = 1

		if win.ParentWidget().Pointer() != nil {
			win.ParentWidget().SetUpdatesEnabled(false)
		}
		win.SetUpdatesEnabled(false)
		win.ResizeEventDefault(event)

		FlutterEngineSendWindowMetricsEvent(win.Width(), win.Height())
	})

	//mouse events

	win.ConnectMousePressEvent(func(event *gui.QMouseEvent) {
		FlutterEngineSendPointerEvent(C.kDown, event.Pos().X(), event.Pos().Y())
	})

	win.ConnectMouseReleaseEvent(func(event *gui.QMouseEvent) {
		FlutterEngineSendPointerEvent(C.kUp, event.Pos().X(), event.Pos().Y())
	})

	win.ConnectEnterEvent(func(event *core.QEvent) {
		if gui.QGuiApplication_MouseButtons() != core.Qt__NoButton {
			return
		}

		e := gui.NewQEnterEventFromPointer(event.Pointer())
		FlutterEngineSendPointerEvent(C.kAdd, e.Pos().X(), e.Pos().Y())
	})

	win.ConnectLeaveEvent(func(event *core.QEvent) {
		if gui.QGuiApplication_MouseButtons() != core.Qt__NoButton {
			return
		}

		FlutterEngineSendPointerEvent(C.kRemove, 0, 0)
	})

	win.ConnectMouseMoveEvent(func(event *gui.QMouseEvent) {
		if event.Buttons() != core.Qt__NoButton {
			FlutterEngineSendPointerEvent(C.kMove, event.Pos().X(), event.Pos().Y())
		} else {
			FlutterEngineSendPointerEvent(C.kHover, event.Pos().X(), event.Pos().Y())
		}
	})

	win.ConnectWheelEvent(func(event *gui.QWheelEvent) {
		angleDelta := event.AngleDelta()
		deltaX := angleDelta.X()
		deltaY := angleDelta.Y()
		if event.Inverted() { //for natural scrolling of the touchpad on macOS
			deltaY *= -1
			deltaX *= -1
		}
		FlutterEngineSendPointerScrollEvent(event.Pos().X(), event.Pos().Y(), deltaX, deltaY)
	})

	//keyboard events

	win.ConnectKeyPressEvent(func(event *gui.QKeyEvent) {
		handleKeyEvent(event.Key(), "keydown", event.NativeVirtualKey())
	})

	win.ConnectKeyReleaseEvent(func(event *gui.QKeyEvent) {
		handleKeyEvent(event.Key(), "keyup", event.NativeVirtualKey())
	})

	//simple resize test

	if TEST_RESIZING {
		go func() {
			time.Sleep(5 * time.Second)
			for range time.NewTicker(time.Second / 15).C {
				RunOnMainBlocking(func() {
					if win.Width() < widgets.QApplication_Desktop().Screen(0).Width() {
						win.Resize2(win.Width()+1, win.Height()+1)
					}
					if win.Width() >= widgets.QApplication_Desktop().Screen(0).Width() {
						win.Resize2(400, 300)
					}
				})
			}
		}()
	}

	win.Show()        //TODO: make sure opengl context (or paintable surface?) is initialized before running the flutter engine without using Show
	win.DoneCurrent() //needed for windows 10 with def rendering

	go func() {
		runtime.LockOSThread()
		FlutterEngineRun()

		//flutter event loop hack using a deprecated api
		//TODO: replace with custom_task_runners
		for range time.NewTicker(time.Second / 60).C {
			C.__FlutterEngineFlushPendingTasksNow()
		}
	}()

	return win
}

//

func MoveToMainThread(ctx *gui.QOpenGLContext) {
	if IGNORE_THREADS {
		return
	}

	ctxThread := ctx.Thread().Pointer()
	mainThread := core.QCoreApplication_Instance().Thread()

	if ctxThread == mainThread.Pointer() {
		return
	}

	if ctxThread != core.QThread_CurrentThread().Pointer() {
		panic("COULDN'T MOVE CONTEXT FROM CURRENT THREAD TO MAIN THREAD")
		return
	}

	ctx.MoveToThread(mainThread)
}

func MoveToCurrentThread(ctx *gui.QOpenGLContext) {
	if IGNORE_THREADS {
		return
	}

	ctxThread := ctx.Thread().Pointer()
	curThread := core.QThread_CurrentThread()

	if ctxThread == curThread.Pointer() {
		return
	}

	if ctxThread != core.QCoreApplication_Instance().Thread().Pointer() {
		panic("COULDN'T MOVE CONTEXT FROM MAIN THREAD TO CURRENT THREAD")
		return
	}

	RunOnMainBlocking(func() {
		ctx.MoveToThread(curThread)
	})
}

func RunOnMainBlocking(f func()) {
	if core.QThread_CurrentThread().Pointer() == core.QCoreApplication_Instance().Thread().Pointer() {
		f()
		return
	}

	done := make(chan bool, 0)
	interop.MainThreadHelper.RunOnMainThread(func() { f(); done <- true })
	<-done
}

//
//paint hot path functions
//

//export make_current
func make_current(unsafe.Pointer) bool {
	if DEBUG_HOT {
		println("make_current")
	}

	frameMutex.Lock()
	defer frameMutex.Unlock()

	/* TODO: mark these contexts as done for windows 10 ?
	if gui.QOpenGLContext_CurrentContext().Pointer() != nil {
		gui.QOpenGLContext_CurrentContext().DoneCurrent()
	}
	*/

	if !USE_ALTERNATIVE_OFFSCREEN_RENDERING {
		MoveToCurrentThread(win.Context())
		win.MakeCurrent()
		return true
	}

	//

	MoveToCurrentThread(offScreenCtx)
	offScreenCtx.MakeCurrent(offScreenSur)

	if win.Width() != offScreenFBO.Width() || win.Height() != offScreenFBO.Height() {
		if offScreenFBO.Pointer() != nil {
			offScreenFBO.DestroyQOpenGLFramebufferObjectDefault()
		}
		offScreenFBO = gui.NewQOpenGLFramebufferObject2(win.Width(), win.Height(), 3553) //GL_TEXTURE_2D
		return true
	}
	return offScreenFBO.Bind()
}

//export present
func present(unsafe.Pointer) bool {
	if DEBUG_HOT {
		println("present")
	}

	frameMutex.Lock()
	defer frameMutex.Unlock()

	switch resizePhase {
	case 0, 3:

		if !USE_ALTERNATIVE_OFFSCREEN_RENDERING {
			win.Context().Functions().GlFlush()
			//win.DoneCurrent() //TODO: needed for windows 10 with def rendering ?
			MoveToMainThread(win.Context())
		} else {
			offScreenFBO_Frame = offScreenFBO.ToImage2()
		}

		win.SetUpdatesEnabled(true)
		win.Update()
		if win.ParentWidget().Pointer() != nil {
			win.ParentWidget().SetUpdatesEnabled(true)
			win.ParentWidget().Update()
		}
	case 1:
		resizePhase = 2
		FlutterEngineSendPointerEvent(C.kAdd, 0, 0)
	case 2:
		resizePhase = 3
		FlutterEngineSendPointerEvent(C.kRemove, 0, 0)
	}

	return true
}

//
//general painting helper functions
//

//export clear_current
func clear_current(unsafe.Pointer) bool {
	if DEBUG {
		println("clear_current")
	}

	frameMutex.Lock()
	defer frameMutex.Unlock()

	if !USE_ALTERNATIVE_OFFSCREEN_RENDERING {
		win.DoneCurrent()
		MoveToMainThread(win.Context())
	} else {
		offScreenCtx.DoneCurrent()
		MoveToMainThread(offScreenCtx)
	}

	return true
}

//export fbo_callback
func fbo_callback(unsafe.Pointer) uint32 {
	if DEBUG {
		println("fbo_callback")
	}

	if !USE_ALTERNATIVE_OFFSCREEN_RENDERING {
		return uint32(0) //uint32(win.Context().DefaultFramebufferObject())
	}

	return uint32(offScreenFBO.Handle())
}

//export gl_proc_resolver
func gl_proc_resolver(_ unsafe.Pointer, procName *C.char) unsafe.Pointer {
	if DEBUG {
		println("gl_proc_resolver")
	}

	return gui.QOpenGLContext_CurrentContext().GetProcAddress2(C.GoString(procName))
}

//export make_resource_current
func make_resource_current(unsafe.Pointer) bool {
	if DEBUG {
		println("make_resource_current")
	}

	MoveToCurrentThread(offScreenCtxShared)
	return offScreenCtxShared.MakeCurrent(offScreenSur)
}

//
//platform plugin functions
//

//export data_callback
func data_callback(*C.char, C.size_t, unsafe.Pointer) {}

//export platform_message_callback
func platform_message_callback(msgC *C.FlutterPlatformMessage, _ unsafe.Pointer) {
	channel := C.GoString(msgC.channel)
	message := C.GoBytes(unsafe.Pointer(msgC.message), C.int(msgC.message_size))
	var msg struct {
		Method string        `json:"method"`
		Args   []interface{} `json:"args"`
	}
	json.Unmarshal(message, &msg)
	method := msg.Method

	switch channel {
	case "flutter/platform":

		switch method {
		case "SystemChrome.setSystemUIOverlayStyle", "HapticFeedback.vibrate", "SystemSound.play", "SystemNavigator.pop":
			//TODO:

		case "Clipboard.setData", "Clipboard.getData":
			//TODO:

		case "SystemChrome.setApplicationSwitcherDescription":
			var m struct {
				Args struct {
					Label        string `json:"label"`
					PrimaryColor int64  `json:"primaryColor"`
				} `json:"args"`
			}
			json.Unmarshal(message, &m)

			RunOnMainBlocking(func() {
				win.SetWindowTitle(m.Args.Label)
			})

		default:
			println("TODO implement:", channel, method, string(message))
		}

	case "flutter/textinput":

		switch method {
		case "TextInput.show", "TextInput.hide", "TextInput.setStyle", "TextInput.setEditableSizeAndTransform", "TextInput.requestAutofill":
		//TODO:

		case "TextInput.setClient":
			var m struct {
				Args []struct {
					InputType struct {
						Name    string      `json:"name"`
						Signed  interface{} `json:"signed"`
						Decimal interface{} `json:"decimal"`
					} `json:"inputType"`
					ObscureText        bool        `json:"obscureText"`
					Autocorrect        bool        `json:"autocorrect"`
					SmartDashesType    string      `json:"smartDashesType"`
					SmartQuotesType    string      `json:"smartQuotesType"`
					EnableSuggestions  bool        `json:"enableSuggestions"`
					ActionLabel        interface{} `json:"actionLabel"`
					InputAction        string      `json:"inputAction"`
					TextCapitalization string      `json:"textCapitalization"`
					KeyboardAppearance string      `json:"keyboardAppearance"`
				} `json:"args"`
			}
			json.Unmarshal(message, &m)

			RunOnMainBlocking(func() {
				textInputHelper = widgets.NewQPlainTextEdit(nil)
				textInputHelper.SetProperty("clientID", core.NewQVariant1(msg.Args[0].(float64)))
				textInputHelper.SetProperty("action", core.NewQVariant1(m.Args[1].InputAction))

				textInputHelper.GrabKeyboard()
				textInputHelper.ConnectKeyPressEvent(func(event *gui.QKeyEvent) {
					textInputHelper.KeyPressEventDefault(event)
					win.KeyPressEvent(event)
				})
				textInputHelper.ConnectKeyReleaseEvent(func(event *gui.QKeyEvent) {
					textInputHelper.KeyReleaseEventDefault(event)
					win.KeyReleaseEvent(event)
				})
			})

		case "TextInput.clearClient":
			textInputHelper.DeleteLater()

		case "TextInput.setEditingState":
			var m struct {
				Args struct {
					Text                   string `json:"text"`
					SelectionBase          int    `json:"selectionBase"`
					SelectionExtent        int    `json:"selectionExtent"`
					SelectionAffinity      string `json:"selectionAffinity"`
					SelectionIsDirectional bool   `json:"selectionIsDirectional"`
					ComposingBase          int    `json:"composingBase"`
					ComposingExtent        int    `json:"composingExtent"`
				} `json:"args"`
			}
			json.Unmarshal(message, &m)

			RunOnMainBlocking(func() {
				textInputHelper.SetPlainTextDefault(m.Args.Text)
				if m.Args.SelectionBase == -1 {
					return
				}
				tc := textInputHelper.TextCursor()
				tc.SetPosition(m.Args.SelectionBase, 0)
				tc.SetPosition(m.Args.SelectionExtent, 1)
				textInputHelper.SetTextCursor(tc)
			})

		default:
			println("TODO implement:", channel, method, string(message))
		}

	case "flutter/isolate":
		//TODO:

	default:
		println("TODO implement:", channel, method, string(message))
	}
}
