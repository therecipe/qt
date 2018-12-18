// +build !minimal

package speech

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "speech.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtSpeech_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtSpeech_PackedString) []byte {
	if int(s.len) == -1 {
		return []byte(C.GoString(s.data))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}

type QTextToSpeech struct {
	core.QObject
}

type QTextToSpeech_ITF interface {
	core.QObject_ITF
	QTextToSpeech_PTR() *QTextToSpeech
}

func (ptr *QTextToSpeech) QTextToSpeech_PTR() *QTextToSpeech {
	return ptr
}

func (ptr *QTextToSpeech) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeech) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeech(ptr QTextToSpeech_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeech_PTR().Pointer()
	}
	return nil
}

func NewQTextToSpeechFromPointer(ptr unsafe.Pointer) (n *QTextToSpeech) {
	n = new(QTextToSpeech)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QTextToSpeech__State
//QTextToSpeech::State
type QTextToSpeech__State int64

const (
	QTextToSpeech__Ready        QTextToSpeech__State = QTextToSpeech__State(0)
	QTextToSpeech__Speaking     QTextToSpeech__State = QTextToSpeech__State(1)
	QTextToSpeech__Paused       QTextToSpeech__State = QTextToSpeech__State(2)
	QTextToSpeech__BackendError QTextToSpeech__State = QTextToSpeech__State(3)
)

func QTextToSpeech_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QTextToSpeech_QTextToSpeech_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QTextToSpeech) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QTextToSpeech_QTextToSpeech_Tr(sC, cC, C.int(int32(n))))
}

func QTextToSpeech_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QTextToSpeech_QTextToSpeech_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QTextToSpeech) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QTextToSpeech_QTextToSpeech_TrUtf8(sC, cC, C.int(int32(n))))
}

func QTextToSpeech_AvailableEngines() []string {
	return strings.Split(cGoUnpackString(C.QTextToSpeech_QTextToSpeech_AvailableEngines()), "|")
}

func (ptr *QTextToSpeech) AvailableEngines() []string {
	return strings.Split(cGoUnpackString(C.QTextToSpeech_QTextToSpeech_AvailableEngines()), "|")
}

func NewQTextToSpeech(parent core.QObject_ITF) *QTextToSpeech {
	tmpValue := NewQTextToSpeechFromPointer(C.QTextToSpeech_NewQTextToSpeech(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQTextToSpeech2(engine string, parent core.QObject_ITF) *QTextToSpeech {
	var engineC *C.char
	if engine != "" {
		engineC = C.CString(engine)
		defer C.free(unsafe.Pointer(engineC))
	}
	tmpValue := NewQTextToSpeechFromPointer(C.QTextToSpeech_NewQTextToSpeech2(C.struct_QtSpeech_PackedString{data: engineC, len: C.longlong(len(engine))}, core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQTextToSpeech_LocaleChanged
func callbackQTextToSpeech_LocaleChanged(ptr unsafe.Pointer, locale unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "localeChanged"); signal != nil {
		signal.(func(*core.QLocale))(core.NewQLocaleFromPointer(locale))
	}

}

func (ptr *QTextToSpeech) ConnectLocaleChanged(f func(locale *core.QLocale)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "localeChanged") {
			C.QTextToSpeech_ConnectLocaleChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "localeChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "localeChanged", func(locale *core.QLocale) {
				signal.(func(*core.QLocale))(locale)
				f(locale)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "localeChanged", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectLocaleChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectLocaleChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "localeChanged")
	}
}

func (ptr *QTextToSpeech) LocaleChanged(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_LocaleChanged(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

//export callbackQTextToSpeech_Pause
func callbackQTextToSpeech_Pause(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pause"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechFromPointer(ptr).PauseDefault()
	}
}

func (ptr *QTextToSpeech) ConnectPause(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "pause"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pause", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pause", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectPause() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "pause")
	}
}

func (ptr *QTextToSpeech) Pause() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_Pause(ptr.Pointer())
	}
}

func (ptr *QTextToSpeech) PauseDefault() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_PauseDefault(ptr.Pointer())
	}
}

//export callbackQTextToSpeech_PitchChanged
func callbackQTextToSpeech_PitchChanged(ptr unsafe.Pointer, pitch C.double) {
	if signal := qt.GetSignal(ptr, "pitchChanged"); signal != nil {
		signal.(func(float64))(float64(pitch))
	}

}

func (ptr *QTextToSpeech) ConnectPitchChanged(f func(pitch float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pitchChanged") {
			C.QTextToSpeech_ConnectPitchChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pitchChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pitchChanged", func(pitch float64) {
				signal.(func(float64))(pitch)
				f(pitch)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pitchChanged", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectPitchChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectPitchChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pitchChanged")
	}
}

func (ptr *QTextToSpeech) PitchChanged(pitch float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_PitchChanged(ptr.Pointer(), C.double(pitch))
	}
}

//export callbackQTextToSpeech_RateChanged
func callbackQTextToSpeech_RateChanged(ptr unsafe.Pointer, rate C.double) {
	if signal := qt.GetSignal(ptr, "rateChanged"); signal != nil {
		signal.(func(float64))(float64(rate))
	}

}

func (ptr *QTextToSpeech) ConnectRateChanged(f func(rate float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "rateChanged") {
			C.QTextToSpeech_ConnectRateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "rateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rateChanged", func(rate float64) {
				signal.(func(float64))(rate)
				f(rate)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rateChanged", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectRateChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectRateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "rateChanged")
	}
}

func (ptr *QTextToSpeech) RateChanged(rate float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_RateChanged(ptr.Pointer(), C.double(rate))
	}
}

//export callbackQTextToSpeech_Resume
func callbackQTextToSpeech_Resume(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resume"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QTextToSpeech) ConnectResume(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resume"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "resume", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resume", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectResume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resume")
	}
}

func (ptr *QTextToSpeech) Resume() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_Resume(ptr.Pointer())
	}
}

func (ptr *QTextToSpeech) ResumeDefault() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_ResumeDefault(ptr.Pointer())
	}
}

//export callbackQTextToSpeech_Say
func callbackQTextToSpeech_Say(ptr unsafe.Pointer, text C.struct_QtSpeech_PackedString) {
	if signal := qt.GetSignal(ptr, "say"); signal != nil {
		signal.(func(string))(cGoUnpackString(text))
	} else {
		NewQTextToSpeechFromPointer(ptr).SayDefault(cGoUnpackString(text))
	}
}

func (ptr *QTextToSpeech) ConnectSay(f func(text string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "say"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "say", func(text string) {
				signal.(func(string))(text)
				f(text)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "say", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectSay() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "say")
	}
}

func (ptr *QTextToSpeech) Say(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextToSpeech_Say(ptr.Pointer(), C.struct_QtSpeech_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

func (ptr *QTextToSpeech) SayDefault(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextToSpeech_SayDefault(ptr.Pointer(), C.struct_QtSpeech_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQTextToSpeech_SetLocale
func callbackQTextToSpeech_SetLocale(ptr unsafe.Pointer, locale unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setLocale"); signal != nil {
		signal.(func(*core.QLocale))(core.NewQLocaleFromPointer(locale))
	} else {
		NewQTextToSpeechFromPointer(ptr).SetLocaleDefault(core.NewQLocaleFromPointer(locale))
	}
}

func (ptr *QTextToSpeech) ConnectSetLocale(f func(locale *core.QLocale)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLocale"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLocale", func(locale *core.QLocale) {
				signal.(func(*core.QLocale))(locale)
				f(locale)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLocale", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectSetLocale() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLocale")
	}
}

func (ptr *QTextToSpeech) SetLocale(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

func (ptr *QTextToSpeech) SetLocaleDefault(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_SetLocaleDefault(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

//export callbackQTextToSpeech_SetPitch
func callbackQTextToSpeech_SetPitch(ptr unsafe.Pointer, pitch C.double) {
	if signal := qt.GetSignal(ptr, "setPitch"); signal != nil {
		signal.(func(float64))(float64(pitch))
	} else {
		NewQTextToSpeechFromPointer(ptr).SetPitchDefault(float64(pitch))
	}
}

func (ptr *QTextToSpeech) ConnectSetPitch(f func(pitch float64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPitch"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPitch", func(pitch float64) {
				signal.(func(float64))(pitch)
				f(pitch)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPitch", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectSetPitch() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPitch")
	}
}

func (ptr *QTextToSpeech) SetPitch(pitch float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_SetPitch(ptr.Pointer(), C.double(pitch))
	}
}

func (ptr *QTextToSpeech) SetPitchDefault(pitch float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_SetPitchDefault(ptr.Pointer(), C.double(pitch))
	}
}

//export callbackQTextToSpeech_SetRate
func callbackQTextToSpeech_SetRate(ptr unsafe.Pointer, rate C.double) {
	if signal := qt.GetSignal(ptr, "setRate"); signal != nil {
		signal.(func(float64))(float64(rate))
	} else {
		NewQTextToSpeechFromPointer(ptr).SetRateDefault(float64(rate))
	}
}

func (ptr *QTextToSpeech) ConnectSetRate(f func(rate float64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRate", func(rate float64) {
				signal.(func(float64))(rate)
				f(rate)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRate", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectSetRate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRate")
	}
}

func (ptr *QTextToSpeech) SetRate(rate float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_SetRate(ptr.Pointer(), C.double(rate))
	}
}

func (ptr *QTextToSpeech) SetRateDefault(rate float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_SetRateDefault(ptr.Pointer(), C.double(rate))
	}
}

//export callbackQTextToSpeech_SetVoice
func callbackQTextToSpeech_SetVoice(ptr unsafe.Pointer, voice unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setVoice"); signal != nil {
		signal.(func(*QVoice))(NewQVoiceFromPointer(voice))
	} else {
		NewQTextToSpeechFromPointer(ptr).SetVoiceDefault(NewQVoiceFromPointer(voice))
	}
}

func (ptr *QTextToSpeech) ConnectSetVoice(f func(voice *QVoice)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setVoice"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setVoice", func(voice *QVoice) {
				signal.(func(*QVoice))(voice)
				f(voice)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setVoice", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectSetVoice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setVoice")
	}
}

func (ptr *QTextToSpeech) SetVoice(voice QVoice_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_SetVoice(ptr.Pointer(), PointerFromQVoice(voice))
	}
}

func (ptr *QTextToSpeech) SetVoiceDefault(voice QVoice_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_SetVoiceDefault(ptr.Pointer(), PointerFromQVoice(voice))
	}
}

//export callbackQTextToSpeech_SetVolume
func callbackQTextToSpeech_SetVolume(ptr unsafe.Pointer, volume C.double) {
	if signal := qt.GetSignal(ptr, "setVolume"); signal != nil {
		signal.(func(float64))(float64(volume))
	} else {
		NewQTextToSpeechFromPointer(ptr).SetVolumeDefault(float64(volume))
	}
}

func (ptr *QTextToSpeech) ConnectSetVolume(f func(volume float64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setVolume"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setVolume", func(volume float64) {
				signal.(func(float64))(volume)
				f(volume)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setVolume", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectSetVolume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setVolume")
	}
}

func (ptr *QTextToSpeech) SetVolume(volume float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_SetVolume(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QTextToSpeech) SetVolumeDefault(volume float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_SetVolumeDefault(ptr.Pointer(), C.double(volume))
	}
}

//export callbackQTextToSpeech_StateChanged
func callbackQTextToSpeech_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(QTextToSpeech__State))(QTextToSpeech__State(state))
	}

}

func (ptr *QTextToSpeech) ConnectStateChanged(f func(state QTextToSpeech__State)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QTextToSpeech_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(state QTextToSpeech__State) {
				signal.(func(QTextToSpeech__State))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QTextToSpeech) StateChanged(state QTextToSpeech__State) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQTextToSpeech_Stop
func callbackQTextToSpeech_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechFromPointer(ptr).StopDefault()
	}
}

func (ptr *QTextToSpeech) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stop", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QTextToSpeech) Stop() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_Stop(ptr.Pointer())
	}
}

func (ptr *QTextToSpeech) StopDefault() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_StopDefault(ptr.Pointer())
	}
}

//export callbackQTextToSpeech_VolumeChanged
func callbackQTextToSpeech_VolumeChanged(ptr unsafe.Pointer, volume C.double) {
	if signal := qt.GetSignal(ptr, "volumeChanged"); signal != nil {
		signal.(func(float64))(float64(volume))
	}

}

func (ptr *QTextToSpeech) ConnectVolumeChanged(f func(volume float64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "volumeChanged") {
			C.QTextToSpeech_ConnectVolumeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "volumeChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "volumeChanged", func(volume float64) {
				signal.(func(float64))(volume)
				f(volume)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "volumeChanged", f)
		}
	}
}

func (ptr *QTextToSpeech) DisconnectVolumeChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "volumeChanged")
	}
}

func (ptr *QTextToSpeech) VolumeChanged(volume float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_VolumeChanged(ptr.Pointer(), C.double(volume))
	}
}

func (ptr *QTextToSpeech) Locale() *core.QLocale {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLocaleFromPointer(C.QTextToSpeech_Locale(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeech) State() QTextToSpeech__State {
	if ptr.Pointer() != nil {
		return QTextToSpeech__State(C.QTextToSpeech_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextToSpeech) AvailableLocales() []*core.QLocale {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSpeech_PackedList) []*core.QLocale {
			out := make([]*core.QLocale, int(l.len))
			tmpList := NewQTextToSpeechFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__availableLocales_atList(i)
			}
			return out
		}(C.QTextToSpeech_AvailableLocales(ptr.Pointer()))
	}
	return make([]*core.QLocale, 0)
}

func (ptr *QTextToSpeech) AvailableVoices() []*QVoice {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSpeech_PackedList) []*QVoice {
			out := make([]*QVoice, int(l.len))
			tmpList := NewQTextToSpeechFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__availableVoices_atList(i)
			}
			return out
		}(C.QTextToSpeech_AvailableVoices(ptr.Pointer()))
	}
	return make([]*QVoice, 0)
}

func (ptr *QTextToSpeech) Voice() *QVoice {
	if ptr.Pointer() != nil {
		tmpValue := NewQVoiceFromPointer(C.QTextToSpeech_Voice(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QVoice).DestroyQVoice)
		return tmpValue
	}
	return nil
}

//export callbackQTextToSpeech_MetaObject
func callbackQTextToSpeech_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQTextToSpeechFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTextToSpeech) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTextToSpeech_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextToSpeech) Pitch() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextToSpeech_Pitch(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextToSpeech) Rate() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextToSpeech_Rate(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextToSpeech) Volume() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextToSpeech_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextToSpeech) __availableLocales_atList(i int) *core.QLocale {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLocaleFromPointer(C.QTextToSpeech___availableLocales_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeech) __availableLocales_setList(i core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech___availableLocales_setList(ptr.Pointer(), core.PointerFromQLocale(i))
	}
}

func (ptr *QTextToSpeech) __availableLocales_newList() unsafe.Pointer {
	return C.QTextToSpeech___availableLocales_newList(ptr.Pointer())
}

func (ptr *QTextToSpeech) __availableVoices_atList(i int) *QVoice {
	if ptr.Pointer() != nil {
		tmpValue := NewQVoiceFromPointer(C.QTextToSpeech___availableVoices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QVoice).DestroyQVoice)
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeech) __availableVoices_setList(i QVoice_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech___availableVoices_setList(ptr.Pointer(), PointerFromQVoice(i))
	}
}

func (ptr *QTextToSpeech) __availableVoices_newList() unsafe.Pointer {
	return C.QTextToSpeech___availableVoices_newList(ptr.Pointer())
}

func (ptr *QTextToSpeech) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QTextToSpeech___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeech) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QTextToSpeech) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QTextToSpeech___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QTextToSpeech) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextToSpeech___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeech) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextToSpeech) __findChildren_newList2() unsafe.Pointer {
	return C.QTextToSpeech___findChildren_newList2(ptr.Pointer())
}

func (ptr *QTextToSpeech) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextToSpeech___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeech) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextToSpeech) __findChildren_newList3() unsafe.Pointer {
	return C.QTextToSpeech___findChildren_newList3(ptr.Pointer())
}

func (ptr *QTextToSpeech) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextToSpeech___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeech) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextToSpeech) __findChildren_newList() unsafe.Pointer {
	return C.QTextToSpeech___findChildren_newList(ptr.Pointer())
}

func (ptr *QTextToSpeech) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextToSpeech___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeech) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextToSpeech) __children_newList() unsafe.Pointer {
	return C.QTextToSpeech___children_newList(ptr.Pointer())
}

//export callbackQTextToSpeech_Event
func callbackQTextToSpeech_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTextToSpeechFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTextToSpeech) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextToSpeech_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQTextToSpeech_EventFilter
func callbackQTextToSpeech_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTextToSpeechFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTextToSpeech) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextToSpeech_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQTextToSpeech_ChildEvent
func callbackQTextToSpeech_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTextToSpeechFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTextToSpeech) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTextToSpeech_ConnectNotify
func callbackQTextToSpeech_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTextToSpeechFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTextToSpeech) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTextToSpeech_CustomEvent
func callbackQTextToSpeech_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTextToSpeechFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextToSpeech) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTextToSpeech_DeleteLater
func callbackQTextToSpeech_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTextToSpeech) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQTextToSpeech_Destroyed
func callbackQTextToSpeech_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQTextToSpeech_DisconnectNotify
func callbackQTextToSpeech_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTextToSpeechFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTextToSpeech) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTextToSpeech_ObjectNameChanged
func callbackQTextToSpeech_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSpeech_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQTextToSpeech_TimerEvent
func callbackQTextToSpeech_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTextToSpeechFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTextToSpeech) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QTextToSpeechEngine struct {
	core.QObject
}

type QTextToSpeechEngine_ITF interface {
	core.QObject_ITF
	QTextToSpeechEngine_PTR() *QTextToSpeechEngine
}

func (ptr *QTextToSpeechEngine) QTextToSpeechEngine_PTR() *QTextToSpeechEngine {
	return ptr
}

func (ptr *QTextToSpeechEngine) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechEngine) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechEngine(ptr QTextToSpeechEngine_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngine_PTR().Pointer()
	}
	return nil
}

func NewQTextToSpeechEngineFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechEngine) {
	n = new(QTextToSpeechEngine)
	n.SetPointer(ptr)
	return
}
func QTextToSpeechEngine_Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QTextToSpeechEngine_QTextToSpeechEngine_Tr(sC, cC, C.int(int32(n))))
}

func (ptr *QTextToSpeechEngine) Tr(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QTextToSpeechEngine_QTextToSpeechEngine_Tr(sC, cC, C.int(int32(n))))
}

func QTextToSpeechEngine_TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QTextToSpeechEngine_QTextToSpeechEngine_TrUtf8(sC, cC, C.int(int32(n))))
}

func (ptr *QTextToSpeechEngine) TrUtf8(s string, c string, n int) string {
	var sC *C.char
	if s != "" {
		sC = C.CString(s)
		defer C.free(unsafe.Pointer(sC))
	}
	var cC *C.char
	if c != "" {
		cC = C.CString(c)
		defer C.free(unsafe.Pointer(cC))
	}
	return cGoUnpackString(C.QTextToSpeechEngine_QTextToSpeechEngine_TrUtf8(sC, cC, C.int(int32(n))))
}

func NewQTextToSpeechEngine(parent core.QObject_ITF) *QTextToSpeechEngine {
	tmpValue := NewQTextToSpeechEngineFromPointer(C.QTextToSpeechEngine_NewQTextToSpeechEngine(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QTextToSpeechEngine_VoiceData(voice QVoice_ITF) *core.QVariant {
	tmpValue := core.NewQVariantFromPointer(C.QTextToSpeechEngine_QTextToSpeechEngine_VoiceData(PointerFromQVoice(voice)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func (ptr *QTextToSpeechEngine) VoiceData(voice QVoice_ITF) *core.QVariant {
	tmpValue := core.NewQVariantFromPointer(C.QTextToSpeechEngine_QTextToSpeechEngine_VoiceData(PointerFromQVoice(voice)))
	runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
	return tmpValue
}

func QTextToSpeechEngine_CreateVoice(name string, gender QVoice__Gender, age QVoice__Age, data core.QVariant_ITF) *QVoice {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQVoiceFromPointer(C.QTextToSpeechEngine_QTextToSpeechEngine_CreateVoice(C.struct_QtSpeech_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(gender), C.longlong(age), core.PointerFromQVariant(data)))
	runtime.SetFinalizer(tmpValue, (*QVoice).DestroyQVoice)
	return tmpValue
}

func (ptr *QTextToSpeechEngine) CreateVoice(name string, gender QVoice__Gender, age QVoice__Age, data core.QVariant_ITF) *QVoice {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQVoiceFromPointer(C.QTextToSpeechEngine_QTextToSpeechEngine_CreateVoice(C.struct_QtSpeech_PackedString{data: nameC, len: C.longlong(len(name))}, C.longlong(gender), C.longlong(age), core.PointerFromQVariant(data)))
	runtime.SetFinalizer(tmpValue, (*QVoice).DestroyQVoice)
	return tmpValue
}

//export callbackQTextToSpeechEngine_SetLocale
func callbackQTextToSpeechEngine_SetLocale(ptr unsafe.Pointer, locale unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "setLocale"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QLocale) bool)(core.NewQLocaleFromPointer(locale)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QTextToSpeechEngine) ConnectSetLocale(f func(locale *core.QLocale) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setLocale"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setLocale", func(locale *core.QLocale) bool {
				signal.(func(*core.QLocale) bool)(locale)
				return f(locale)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setLocale", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectSetLocale() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setLocale")
	}
}

func (ptr *QTextToSpeechEngine) SetLocale(locale core.QLocale_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextToSpeechEngine_SetLocale(ptr.Pointer(), core.PointerFromQLocale(locale))) != 0
	}
	return false
}

//export callbackQTextToSpeechEngine_SetPitch
func callbackQTextToSpeechEngine_SetPitch(ptr unsafe.Pointer, pitch C.double) C.char {
	if signal := qt.GetSignal(ptr, "setPitch"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(float64) bool)(float64(pitch)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QTextToSpeechEngine) ConnectSetPitch(f func(pitch float64) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setPitch"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setPitch", func(pitch float64) bool {
				signal.(func(float64) bool)(pitch)
				return f(pitch)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setPitch", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectSetPitch() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setPitch")
	}
}

func (ptr *QTextToSpeechEngine) SetPitch(pitch float64) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextToSpeechEngine_SetPitch(ptr.Pointer(), C.double(pitch))) != 0
	}
	return false
}

//export callbackQTextToSpeechEngine_SetRate
func callbackQTextToSpeechEngine_SetRate(ptr unsafe.Pointer, rate C.double) C.char {
	if signal := qt.GetSignal(ptr, "setRate"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(float64) bool)(float64(rate)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QTextToSpeechEngine) ConnectSetRate(f func(rate float64) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setRate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setRate", func(rate float64) bool {
				signal.(func(float64) bool)(rate)
				return f(rate)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setRate", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectSetRate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setRate")
	}
}

func (ptr *QTextToSpeechEngine) SetRate(rate float64) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextToSpeechEngine_SetRate(ptr.Pointer(), C.double(rate))) != 0
	}
	return false
}

//export callbackQTextToSpeechEngine_SetVoice
func callbackQTextToSpeechEngine_SetVoice(ptr unsafe.Pointer, voice unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "setVoice"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QVoice) bool)(NewQVoiceFromPointer(voice)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QTextToSpeechEngine) ConnectSetVoice(f func(voice *QVoice) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setVoice"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setVoice", func(voice *QVoice) bool {
				signal.(func(*QVoice) bool)(voice)
				return f(voice)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setVoice", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectSetVoice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setVoice")
	}
}

func (ptr *QTextToSpeechEngine) SetVoice(voice QVoice_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextToSpeechEngine_SetVoice(ptr.Pointer(), PointerFromQVoice(voice))) != 0
	}
	return false
}

//export callbackQTextToSpeechEngine_SetVolume
func callbackQTextToSpeechEngine_SetVolume(ptr unsafe.Pointer, volume C.double) C.char {
	if signal := qt.GetSignal(ptr, "setVolume"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(float64) bool)(float64(volume)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QTextToSpeechEngine) ConnectSetVolume(f func(volume float64) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setVolume"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setVolume", func(volume float64) bool {
				signal.(func(float64) bool)(volume)
				return f(volume)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setVolume", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectSetVolume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setVolume")
	}
}

func (ptr *QTextToSpeechEngine) SetVolume(volume float64) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextToSpeechEngine_SetVolume(ptr.Pointer(), C.double(volume))) != 0
	}
	return false
}

//export callbackQTextToSpeechEngine_Pause
func callbackQTextToSpeechEngine_Pause(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "pause"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextToSpeechEngine) ConnectPause(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "pause"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pause", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pause", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectPause() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "pause")
	}
}

func (ptr *QTextToSpeechEngine) Pause() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_Pause(ptr.Pointer())
	}
}

//export callbackQTextToSpeechEngine_Resume
func callbackQTextToSpeechEngine_Resume(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "resume"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextToSpeechEngine) ConnectResume(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "resume"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "resume", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "resume", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectResume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "resume")
	}
}

func (ptr *QTextToSpeechEngine) Resume() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_Resume(ptr.Pointer())
	}
}

//export callbackQTextToSpeechEngine_Say
func callbackQTextToSpeechEngine_Say(ptr unsafe.Pointer, text C.struct_QtSpeech_PackedString) {
	if signal := qt.GetSignal(ptr, "say"); signal != nil {
		signal.(func(string))(cGoUnpackString(text))
	}

}

func (ptr *QTextToSpeechEngine) ConnectSay(f func(text string)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "say"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "say", func(text string) {
				signal.(func(string))(text)
				f(text)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "say", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectSay() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "say")
	}
}

func (ptr *QTextToSpeechEngine) Say(text string) {
	if ptr.Pointer() != nil {
		var textC *C.char
		if text != "" {
			textC = C.CString(text)
			defer C.free(unsafe.Pointer(textC))
		}
		C.QTextToSpeechEngine_Say(ptr.Pointer(), C.struct_QtSpeech_PackedString{data: textC, len: C.longlong(len(text))})
	}
}

//export callbackQTextToSpeechEngine_StateChanged
func callbackQTextToSpeechEngine_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(QTextToSpeech__State))(QTextToSpeech__State(state))
	}

}

func (ptr *QTextToSpeechEngine) ConnectStateChanged(f func(state QTextToSpeech__State)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QTextToSpeechEngine_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(state QTextToSpeech__State) {
				signal.(func(QTextToSpeech__State))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QTextToSpeechEngine) StateChanged(state QTextToSpeech__State) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQTextToSpeechEngine_Stop
func callbackQTextToSpeechEngine_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QTextToSpeechEngine) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stop", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QTextToSpeechEngine) Stop() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_Stop(ptr.Pointer())
	}
}

//export callbackQTextToSpeechEngine_DestroyQTextToSpeechEngine
func callbackQTextToSpeechEngine_DestroyQTextToSpeechEngine(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTextToSpeechEngine"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechEngineFromPointer(ptr).DestroyQTextToSpeechEngineDefault()
	}
}

func (ptr *QTextToSpeechEngine) ConnectDestroyQTextToSpeechEngine(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTextToSpeechEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QTextToSpeechEngine", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTextToSpeechEngine", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectDestroyQTextToSpeechEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTextToSpeechEngine")
	}
}

func (ptr *QTextToSpeechEngine) DestroyQTextToSpeechEngine() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_DestroyQTextToSpeechEngine(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QTextToSpeechEngine) DestroyQTextToSpeechEngineDefault() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_DestroyQTextToSpeechEngineDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQTextToSpeechEngine_Locale
func callbackQTextToSpeechEngine_Locale(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "locale"); signal != nil {
		return core.PointerFromQLocale(signal.(func() *core.QLocale)())
	}

	return core.PointerFromQLocale(core.NewQLocale())
}

func (ptr *QTextToSpeechEngine) ConnectLocale(f func() *core.QLocale) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "locale"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "locale", func() *core.QLocale {
				signal.(func() *core.QLocale)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "locale", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectLocale() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "locale")
	}
}

func (ptr *QTextToSpeechEngine) Locale() *core.QLocale {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLocaleFromPointer(C.QTextToSpeechEngine_Locale(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

//export callbackQTextToSpeechEngine_State
func callbackQTextToSpeechEngine_State(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "state"); signal != nil {
		return C.longlong(signal.(func() QTextToSpeech__State)())
	}

	return C.longlong(0)
}

func (ptr *QTextToSpeechEngine) ConnectState(f func() QTextToSpeech__State) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "state"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "state", func() QTextToSpeech__State {
				signal.(func() QTextToSpeech__State)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "state", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectState() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "state")
	}
}

func (ptr *QTextToSpeechEngine) State() QTextToSpeech__State {
	if ptr.Pointer() != nil {
		return QTextToSpeech__State(C.QTextToSpeechEngine_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQTextToSpeechEngine_AvailableLocales
func callbackQTextToSpeechEngine_AvailableLocales(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "availableLocales"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQTextToSpeechEngineFromPointer(NewQTextToSpeechEngineFromPointer(nil).__availableLocales_newList())
			for _, v := range signal.(func() []*core.QLocale)() {
				tmpList.__availableLocales_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQTextToSpeechEngineFromPointer(NewQTextToSpeechEngineFromPointer(nil).__availableLocales_newList())
		for _, v := range make([]*core.QLocale, 0) {
			tmpList.__availableLocales_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QTextToSpeechEngine) ConnectAvailableLocales(f func() []*core.QLocale) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "availableLocales"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "availableLocales", func() []*core.QLocale {
				signal.(func() []*core.QLocale)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "availableLocales", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectAvailableLocales() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "availableLocales")
	}
}

func (ptr *QTextToSpeechEngine) AvailableLocales() []*core.QLocale {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSpeech_PackedList) []*core.QLocale {
			out := make([]*core.QLocale, int(l.len))
			tmpList := NewQTextToSpeechEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__availableLocales_atList(i)
			}
			return out
		}(C.QTextToSpeechEngine_AvailableLocales(ptr.Pointer()))
	}
	return make([]*core.QLocale, 0)
}

//export callbackQTextToSpeechEngine_AvailableVoices
func callbackQTextToSpeechEngine_AvailableVoices(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "availableVoices"); signal != nil {
		return func() unsafe.Pointer {
			tmpList := NewQTextToSpeechEngineFromPointer(NewQTextToSpeechEngineFromPointer(nil).__availableVoices_newList())
			for _, v := range signal.(func() []*QVoice)() {
				tmpList.__availableVoices_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		tmpList := NewQTextToSpeechEngineFromPointer(NewQTextToSpeechEngineFromPointer(nil).__availableVoices_newList())
		for _, v := range make([]*QVoice, 0) {
			tmpList.__availableVoices_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QTextToSpeechEngine) ConnectAvailableVoices(f func() []*QVoice) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "availableVoices"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "availableVoices", func() []*QVoice {
				signal.(func() []*QVoice)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "availableVoices", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectAvailableVoices() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "availableVoices")
	}
}

func (ptr *QTextToSpeechEngine) AvailableVoices() []*QVoice {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSpeech_PackedList) []*QVoice {
			out := make([]*QVoice, int(l.len))
			tmpList := NewQTextToSpeechEngineFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__availableVoices_atList(i)
			}
			return out
		}(C.QTextToSpeechEngine_AvailableVoices(ptr.Pointer()))
	}
	return make([]*QVoice, 0)
}

//export callbackQTextToSpeechEngine_Voice
func callbackQTextToSpeechEngine_Voice(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "voice"); signal != nil {
		return PointerFromQVoice(signal.(func() *QVoice)())
	}

	return PointerFromQVoice(NewQVoice())
}

func (ptr *QTextToSpeechEngine) ConnectVoice(f func() *QVoice) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "voice"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "voice", func() *QVoice {
				signal.(func() *QVoice)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "voice", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectVoice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "voice")
	}
}

func (ptr *QTextToSpeechEngine) Voice() *QVoice {
	if ptr.Pointer() != nil {
		tmpValue := NewQVoiceFromPointer(C.QTextToSpeechEngine_Voice(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QVoice).DestroyQVoice)
		return tmpValue
	}
	return nil
}

//export callbackQTextToSpeechEngine_MetaObject
func callbackQTextToSpeechEngine_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQTextToSpeechEngineFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QTextToSpeechEngine) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QTextToSpeechEngine_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQTextToSpeechEngine_Pitch
func callbackQTextToSpeechEngine_Pitch(ptr unsafe.Pointer) C.double {
	if signal := qt.GetSignal(ptr, "pitch"); signal != nil {
		return C.double(signal.(func() float64)())
	}

	return C.double(0)
}

func (ptr *QTextToSpeechEngine) ConnectPitch(f func() float64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "pitch"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pitch", func() float64 {
				signal.(func() float64)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pitch", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectPitch() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "pitch")
	}
}

func (ptr *QTextToSpeechEngine) Pitch() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextToSpeechEngine_Pitch(ptr.Pointer()))
	}
	return 0
}

//export callbackQTextToSpeechEngine_Rate
func callbackQTextToSpeechEngine_Rate(ptr unsafe.Pointer) C.double {
	if signal := qt.GetSignal(ptr, "rate"); signal != nil {
		return C.double(signal.(func() float64)())
	}

	return C.double(0)
}

func (ptr *QTextToSpeechEngine) ConnectRate(f func() float64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "rate"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "rate", func() float64 {
				signal.(func() float64)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "rate", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectRate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "rate")
	}
}

func (ptr *QTextToSpeechEngine) Rate() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextToSpeechEngine_Rate(ptr.Pointer()))
	}
	return 0
}

//export callbackQTextToSpeechEngine_Volume
func callbackQTextToSpeechEngine_Volume(ptr unsafe.Pointer) C.double {
	if signal := qt.GetSignal(ptr, "volume"); signal != nil {
		return C.double(signal.(func() float64)())
	}

	return C.double(0)
}

func (ptr *QTextToSpeechEngine) ConnectVolume(f func() float64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "volume"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "volume", func() float64 {
				signal.(func() float64)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "volume", f)
		}
	}
}

func (ptr *QTextToSpeechEngine) DisconnectVolume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "volume")
	}
}

func (ptr *QTextToSpeechEngine) Volume() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QTextToSpeechEngine_Volume(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextToSpeechEngine) __availableLocales_atList(i int) *core.QLocale {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQLocaleFromPointer(C.QTextToSpeechEngine___availableLocales_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechEngine) __availableLocales_setList(i core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine___availableLocales_setList(ptr.Pointer(), core.PointerFromQLocale(i))
	}
}

func (ptr *QTextToSpeechEngine) __availableLocales_newList() unsafe.Pointer {
	return C.QTextToSpeechEngine___availableLocales_newList(ptr.Pointer())
}

func (ptr *QTextToSpeechEngine) __availableVoices_atList(i int) *QVoice {
	if ptr.Pointer() != nil {
		tmpValue := NewQVoiceFromPointer(C.QTextToSpeechEngine___availableVoices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QVoice).DestroyQVoice)
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechEngine) __availableVoices_setList(i QVoice_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine___availableVoices_setList(ptr.Pointer(), PointerFromQVoice(i))
	}
}

func (ptr *QTextToSpeechEngine) __availableVoices_newList() unsafe.Pointer {
	return C.QTextToSpeechEngine___availableVoices_newList(ptr.Pointer())
}

func (ptr *QTextToSpeechEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QTextToSpeechEngine___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QTextToSpeechEngine) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.QTextToSpeechEngine___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QTextToSpeechEngine) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextToSpeechEngine___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechEngine) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextToSpeechEngine) __findChildren_newList2() unsafe.Pointer {
	return C.QTextToSpeechEngine___findChildren_newList2(ptr.Pointer())
}

func (ptr *QTextToSpeechEngine) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextToSpeechEngine___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechEngine) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextToSpeechEngine) __findChildren_newList3() unsafe.Pointer {
	return C.QTextToSpeechEngine___findChildren_newList3(ptr.Pointer())
}

func (ptr *QTextToSpeechEngine) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextToSpeechEngine___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechEngine) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextToSpeechEngine) __findChildren_newList() unsafe.Pointer {
	return C.QTextToSpeechEngine___findChildren_newList(ptr.Pointer())
}

func (ptr *QTextToSpeechEngine) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QTextToSpeechEngine___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechEngine) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QTextToSpeechEngine) __children_newList() unsafe.Pointer {
	return C.QTextToSpeechEngine___children_newList(ptr.Pointer())
}

//export callbackQTextToSpeechEngine_Event
func callbackQTextToSpeechEngine_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTextToSpeechEngineFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QTextToSpeechEngine) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextToSpeechEngine_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQTextToSpeechEngine_EventFilter
func callbackQTextToSpeechEngine_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQTextToSpeechEngineFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QTextToSpeechEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QTextToSpeechEngine_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQTextToSpeechEngine_ChildEvent
func callbackQTextToSpeechEngine_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQTextToSpeechEngineFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QTextToSpeechEngine) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQTextToSpeechEngine_ConnectNotify
func callbackQTextToSpeechEngine_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTextToSpeechEngineFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTextToSpeechEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTextToSpeechEngine_CustomEvent
func callbackQTextToSpeechEngine_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQTextToSpeechEngineFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QTextToSpeechEngine) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQTextToSpeechEngine_DeleteLater
func callbackQTextToSpeechEngine_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechEngineFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QTextToSpeechEngine) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQTextToSpeechEngine_Destroyed
func callbackQTextToSpeechEngine_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQTextToSpeechEngine_DisconnectNotify
func callbackQTextToSpeechEngine_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQTextToSpeechEngineFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QTextToSpeechEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQTextToSpeechEngine_ObjectNameChanged
func callbackQTextToSpeechEngine_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSpeech_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQTextToSpeechEngine_TimerEvent
func callbackQTextToSpeechEngine_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQTextToSpeechEngineFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QTextToSpeechEngine) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeechEngine_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

type QTextToSpeechPlugin struct {
	ptr unsafe.Pointer
}

type QTextToSpeechPlugin_ITF interface {
	QTextToSpeechPlugin_PTR() *QTextToSpeechPlugin
}

func (ptr *QTextToSpeechPlugin) QTextToSpeechPlugin_PTR() *QTextToSpeechPlugin {
	return ptr
}

func (ptr *QTextToSpeechPlugin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextToSpeechPlugin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQTextToSpeechPlugin(ptr QTextToSpeechPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechPlugin_PTR().Pointer()
	}
	return nil
}

func NewQTextToSpeechPluginFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechPlugin) {
	n = new(QTextToSpeechPlugin)
	n.SetPointer(ptr)
	return
}

//export callbackQTextToSpeechPlugin_DestroyQTextToSpeechPlugin
func callbackQTextToSpeechPlugin_DestroyQTextToSpeechPlugin(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QTextToSpeechPlugin"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechPluginFromPointer(ptr).DestroyQTextToSpeechPluginDefault()
	}
}

func (ptr *QTextToSpeechPlugin) ConnectDestroyQTextToSpeechPlugin(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QTextToSpeechPlugin"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QTextToSpeechPlugin", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QTextToSpeechPlugin", f)
		}
	}
}

func (ptr *QTextToSpeechPlugin) DisconnectDestroyQTextToSpeechPlugin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QTextToSpeechPlugin")
	}
}

func (ptr *QTextToSpeechPlugin) DestroyQTextToSpeechPlugin() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechPlugin_DestroyQTextToSpeechPlugin(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QTextToSpeechPlugin) DestroyQTextToSpeechPluginDefault() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechPlugin_DestroyQTextToSpeechPluginDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQTextToSpeechPlugin_CreateTextToSpeechEngine
func callbackQTextToSpeechPlugin_CreateTextToSpeechEngine(ptr unsafe.Pointer, parameters C.struct_QtSpeech_PackedList, parent unsafe.Pointer, errorString C.struct_QtSpeech_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createTextToSpeechEngine"); signal != nil {
		return PointerFromQTextToSpeechEngine(signal.(func(map[string]*core.QVariant, *core.QObject, string) *QTextToSpeechEngine)(func(l C.struct_QtSpeech_PackedList) map[string]*core.QVariant {
			out := make(map[string]*core.QVariant, int(l.len))
			tmpList := NewQTextToSpeechPluginFromPointer(l.data)
			for i, v := range tmpList.__createTextToSpeechEngine_parameters_keyList() {
				out[v] = tmpList.__createTextToSpeechEngine_parameters_atList(v, i)
			}
			return out
		}(parameters), core.NewQObjectFromPointer(parent), cGoUnpackString(errorString)))
	}

	return PointerFromQTextToSpeechEngine(NewQTextToSpeechPluginFromPointer(ptr).CreateTextToSpeechEngineDefault(func(l C.struct_QtSpeech_PackedList) map[string]*core.QVariant {
		out := make(map[string]*core.QVariant, int(l.len))
		tmpList := NewQTextToSpeechPluginFromPointer(l.data)
		for i, v := range tmpList.__createTextToSpeechEngine_parameters_keyList() {
			out[v] = tmpList.__createTextToSpeechEngine_parameters_atList(v, i)
		}
		return out
	}(parameters), core.NewQObjectFromPointer(parent), cGoUnpackString(errorString)))
}

func (ptr *QTextToSpeechPlugin) ConnectCreateTextToSpeechEngine(f func(parameters map[string]*core.QVariant, parent *core.QObject, errorString string) *QTextToSpeechEngine) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createTextToSpeechEngine"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createTextToSpeechEngine", func(parameters map[string]*core.QVariant, parent *core.QObject, errorString string) *QTextToSpeechEngine {
				signal.(func(map[string]*core.QVariant, *core.QObject, string) *QTextToSpeechEngine)(parameters, parent, errorString)
				return f(parameters, parent, errorString)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createTextToSpeechEngine", f)
		}
	}
}

func (ptr *QTextToSpeechPlugin) DisconnectCreateTextToSpeechEngine() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createTextToSpeechEngine")
	}
}

func (ptr *QTextToSpeechPlugin) CreateTextToSpeechEngine(parameters map[string]*core.QVariant, parent core.QObject_ITF, errorString string) *QTextToSpeechEngine {
	if ptr.Pointer() != nil {
		var errorStringC *C.char
		if errorString != "" {
			errorStringC = C.CString(errorString)
			defer C.free(unsafe.Pointer(errorStringC))
		}
		tmpValue := NewQTextToSpeechEngineFromPointer(C.QTextToSpeechPlugin_CreateTextToSpeechEngine(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQTextToSpeechPluginFromPointer(NewQTextToSpeechPluginFromPointer(nil).__createTextToSpeechEngine_parameters_newList())
			for k, v := range parameters {
				tmpList.__createTextToSpeechEngine_parameters_setList(k, v)
			}
			return tmpList.Pointer()
		}(), core.PointerFromQObject(parent), C.struct_QtSpeech_PackedString{data: errorStringC, len: C.longlong(len(errorString))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechPlugin) CreateTextToSpeechEngineDefault(parameters map[string]*core.QVariant, parent core.QObject_ITF, errorString string) *QTextToSpeechEngine {
	if ptr.Pointer() != nil {
		var errorStringC *C.char
		if errorString != "" {
			errorStringC = C.CString(errorString)
			defer C.free(unsafe.Pointer(errorStringC))
		}
		tmpValue := NewQTextToSpeechEngineFromPointer(C.QTextToSpeechPlugin_CreateTextToSpeechEngineDefault(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQTextToSpeechPluginFromPointer(NewQTextToSpeechPluginFromPointer(nil).__createTextToSpeechEngine_parameters_newList())
			for k, v := range parameters {
				tmpList.__createTextToSpeechEngine_parameters_setList(k, v)
			}
			return tmpList.Pointer()
		}(), core.PointerFromQObject(parent), C.struct_QtSpeech_PackedString{data: errorStringC, len: C.longlong(len(errorString))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_atList(v string, i int) *core.QVariant {
	if ptr.Pointer() != nil {
		var vC *C.char
		if v != "" {
			vC = C.CString(v)
			defer C.free(unsafe.Pointer(vC))
		}
		tmpValue := core.NewQVariantFromPointer(C.QTextToSpeechPlugin___createTextToSpeechEngine_parameters_atList(ptr.Pointer(), C.struct_QtSpeech_PackedString{data: vC, len: C.longlong(len(v))}, C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC *C.char
		if key != "" {
			keyC = C.CString(key)
			defer C.free(unsafe.Pointer(keyC))
		}
		C.QTextToSpeechPlugin___createTextToSpeechEngine_parameters_setList(ptr.Pointer(), C.struct_QtSpeech_PackedString{data: keyC, len: C.longlong(len(key))}, core.PointerFromQVariant(i))
	}
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_newList() unsafe.Pointer {
	return C.QTextToSpeechPlugin___createTextToSpeechEngine_parameters_newList(ptr.Pointer())
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSpeech_PackedList) []string {
			out := make([]string, int(l.len))
			tmpList := NewQTextToSpeechPluginFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.____createTextToSpeechEngine_parameters_keyList_atList(i)
			}
			return out
		}(C.QTextToSpeechPlugin___createTextToSpeechEngine_parameters_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QTextToSpeechPlugin) ____createTextToSpeechEngine_parameters_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextToSpeechPlugin_____createTextToSpeechEngine_parameters_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QTextToSpeechPlugin) ____createTextToSpeechEngine_parameters_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QTextToSpeechPlugin_____createTextToSpeechEngine_parameters_keyList_setList(ptr.Pointer(), C.struct_QtSpeech_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QTextToSpeechPlugin) ____createTextToSpeechEngine_parameters_keyList_newList() unsafe.Pointer {
	return C.QTextToSpeechPlugin_____createTextToSpeechEngine_parameters_keyList_newList(ptr.Pointer())
}

type QVoice struct {
	ptr unsafe.Pointer
}

type QVoice_ITF interface {
	QVoice_PTR() *QVoice
}

func (ptr *QVoice) QVoice_PTR() *QVoice {
	return ptr
}

func (ptr *QVoice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QVoice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQVoice(ptr QVoice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVoice_PTR().Pointer()
	}
	return nil
}

func NewQVoiceFromPointer(ptr unsafe.Pointer) (n *QVoice) {
	n = new(QVoice)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QVoice__Age
//QVoice::Age
type QVoice__Age int64

const (
	QVoice__Child    QVoice__Age = QVoice__Age(0)
	QVoice__Teenager QVoice__Age = QVoice__Age(1)
	QVoice__Adult    QVoice__Age = QVoice__Age(2)
	QVoice__Senior   QVoice__Age = QVoice__Age(3)
	QVoice__Other    QVoice__Age = QVoice__Age(4)
)

//go:generate stringer -type=QVoice__Gender
//QVoice::Gender
type QVoice__Gender int64

const (
	QVoice__Male    QVoice__Gender = QVoice__Gender(0)
	QVoice__Female  QVoice__Gender = QVoice__Gender(1)
	QVoice__Unknown QVoice__Gender = QVoice__Gender(2)
)

func QVoice_AgeName(age QVoice__Age) string {
	return cGoUnpackString(C.QVoice_QVoice_AgeName(C.longlong(age)))
}

func (ptr *QVoice) AgeName(age QVoice__Age) string {
	return cGoUnpackString(C.QVoice_QVoice_AgeName(C.longlong(age)))
}

func QVoice_GenderName(gender QVoice__Gender) string {
	return cGoUnpackString(C.QVoice_QVoice_GenderName(C.longlong(gender)))
}

func (ptr *QVoice) GenderName(gender QVoice__Gender) string {
	return cGoUnpackString(C.QVoice_QVoice_GenderName(C.longlong(gender)))
}

func NewQVoice() *QVoice {
	tmpValue := NewQVoiceFromPointer(C.QVoice_NewQVoice())
	runtime.SetFinalizer(tmpValue, (*QVoice).DestroyQVoice)
	return tmpValue
}

func NewQVoice2(other QVoice_ITF) *QVoice {
	tmpValue := NewQVoiceFromPointer(C.QVoice_NewQVoice2(PointerFromQVoice(other)))
	runtime.SetFinalizer(tmpValue, (*QVoice).DestroyQVoice)
	return tmpValue
}

func (ptr *QVoice) DestroyQVoice() {
	if ptr.Pointer() != nil {
		C.QVoice_DestroyQVoice(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QVoice) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVoice_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QVoice) Age() QVoice__Age {
	if ptr.Pointer() != nil {
		return QVoice__Age(C.QVoice_Age(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVoice) Gender() QVoice__Gender {
	if ptr.Pointer() != nil {
		return QVoice__Gender(C.QVoice_Gender(ptr.Pointer()))
	}
	return 0
}
