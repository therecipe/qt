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
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
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

func NewQTextToSpeechFromPointer(ptr unsafe.Pointer) *QTextToSpeech {
	var n = new(QTextToSpeech)
	n.SetPointer(ptr)
	return n
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

func QTextToSpeech_AvailableEngines() []string {
	return strings.Split(cGoUnpackString(C.QTextToSpeech_QTextToSpeech_AvailableEngines()), "|")
}

func (ptr *QTextToSpeech) AvailableEngines() []string {
	return strings.Split(cGoUnpackString(C.QTextToSpeech_QTextToSpeech_AvailableEngines()), "|")
}

func NewQTextToSpeech(parent core.QObject_ITF) *QTextToSpeech {
	var tmpValue = NewQTextToSpeechFromPointer(C.QTextToSpeech_NewQTextToSpeech(core.PointerFromQObject(parent)))
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
	var tmpValue = NewQTextToSpeechFromPointer(C.QTextToSpeech_NewQTextToSpeech2(C.struct_QtSpeech_PackedString{data: engineC, len: C.longlong(len(engine))}, core.PointerFromQObject(parent)))
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
func callbackQTextToSpeech_VolumeChanged(ptr unsafe.Pointer, volume C.int) {
	if signal := qt.GetSignal(ptr, "volumeChanged"); signal != nil {
		signal.(func(int))(int(int32(volume)))
	}

}

func (ptr *QTextToSpeech) ConnectVolumeChanged(f func(volume int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "volumeChanged") {
			C.QTextToSpeech_ConnectVolumeChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "volumeChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "volumeChanged", func(volume int) {
				signal.(func(int))(volume)
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

func (ptr *QTextToSpeech) VolumeChanged(volume int) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_VolumeChanged(ptr.Pointer(), C.int(int32(volume)))
	}
}

func (ptr *QTextToSpeech) Locale() *core.QLocale {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQLocaleFromPointer(C.QTextToSpeech_Locale(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QLocale).DestroyQLocale)
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeech) AvailableLocales() []*core.QLocale {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSpeech_PackedList) []*core.QLocale {
			var out = make([]*core.QLocale, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQTextToSpeechFromPointer(l.data).__availableLocales_atList(i)
			}
			return out
		}(C.QTextToSpeech_AvailableLocales(ptr.Pointer()))
	}
	return make([]*core.QLocale, 0)
}

func (ptr *QTextToSpeech) AvailableVoices() []*QVoice {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSpeech_PackedList) []*QVoice {
			var out = make([]*QVoice, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQTextToSpeechFromPointer(l.data).__availableVoices_atList(i)
			}
			return out
		}(C.QTextToSpeech_AvailableVoices(ptr.Pointer()))
	}
	return make([]*QVoice, 0)
}

func (ptr *QTextToSpeech) Voice() *QVoice {
	if ptr.Pointer() != nil {
		var tmpValue = NewQVoiceFromPointer(C.QTextToSpeech_Voice(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QVoice).DestroyQVoice)
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
		var tmpValue = core.NewQLocaleFromPointer(C.QTextToSpeech___availableLocales_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QTextToSpeech___availableLocales_newList(ptr.Pointer()))
}

func (ptr *QTextToSpeech) __availableVoices_atList(i int) *QVoice {
	if ptr.Pointer() != nil {
		var tmpValue = NewQVoiceFromPointer(C.QTextToSpeech___availableVoices_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QTextToSpeech___availableVoices_newList(ptr.Pointer()))
}

func (ptr *QTextToSpeech) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QTextToSpeech___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QTextToSpeech___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QTextToSpeech) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QTextToSpeech___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QTextToSpeech___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QTextToSpeech) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QTextToSpeech___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QTextToSpeech___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QTextToSpeech) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QTextToSpeech___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QTextToSpeech___findChildren_newList(ptr.Pointer()))
}

func (ptr *QTextToSpeech) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QTextToSpeech___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return unsafe.Pointer(C.QTextToSpeech___children_newList(ptr.Pointer()))
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
		return C.QTextToSpeech_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
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
		return C.QTextToSpeech_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
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

func NewQTextToSpeechPluginFromPointer(ptr unsafe.Pointer) *QTextToSpeechPlugin {
	var n = new(QTextToSpeechPlugin)
	n.SetPointer(ptr)
	return n
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

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		var tmpValue = core.NewQVariantFromPointer(C.QTextToSpeechPlugin___createTextToSpeechEngine_parameters_atList(ptr.Pointer(), C.struct_QtSpeech_PackedString{data: iC, len: C.longlong(len(i))}))
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
	return unsafe.Pointer(C.QTextToSpeechPlugin___createTextToSpeechEngine_parameters_newList(ptr.Pointer()))
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_keyList() []string {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSpeech_PackedList) []string {
			var out = make([]string, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQTextToSpeechPluginFromPointer(l.data).____createTextToSpeechEngine_keyList_atList(i)
			}
			return out
		}(C.QTextToSpeechPlugin___createTextToSpeechEngine_keyList(ptr.Pointer()))
	}
	return make([]string, 0)
}

func (ptr *QTextToSpeechPlugin) ____createTextToSpeechEngine_keyList_atList(i int) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return ""
}

func (ptr *QTextToSpeechPlugin) ____createTextToSpeechEngine_keyList_setList(i string) {
	if ptr.Pointer() != nil {
		var iC *C.char
		if i != "" {
			iC = C.CString(i)
			defer C.free(unsafe.Pointer(iC))
		}
		C.QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_setList(ptr.Pointer(), C.struct_QtSpeech_PackedString{data: iC, len: C.longlong(len(i))})
	}
}

func (ptr *QTextToSpeechPlugin) ____createTextToSpeechEngine_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_newList(ptr.Pointer()))
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

func NewQVoiceFromPointer(ptr unsafe.Pointer) *QVoice {
	var n = new(QVoice)
	n.SetPointer(ptr)
	return n
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
	var tmpValue = NewQVoiceFromPointer(C.QVoice_NewQVoice())
	runtime.SetFinalizer(tmpValue, (*QVoice).DestroyQVoice)
	return tmpValue
}

func NewQVoice2(other QVoice_ITF) *QVoice {
	var tmpValue = NewQVoiceFromPointer(C.QVoice_NewQVoice2(PointerFromQVoice(other)))
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

func (ptr *QVoice) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QVoice_Name(ptr.Pointer()))
	}
	return ""
}
