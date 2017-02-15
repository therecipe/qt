// +build !minimal

package speech

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "speech.h"
import "C"
import (
	"fmt"
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
	ptr unsafe.Pointer
}

type QTextToSpeech_ITF interface {
	QTextToSpeech_PTR() *QTextToSpeech
}

func (ptr *QTextToSpeech) QTextToSpeech_PTR() *QTextToSpeech {
	return ptr
}

func (ptr *QTextToSpeech) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QTextToSpeech) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

func (ptr *QTextToSpeech) DestroyQTextToSpeech() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
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
	return NewQTextToSpeechFromPointer(C.QTextToSpeech_NewQTextToSpeech(core.PointerFromQObject(parent)))
}

func NewQTextToSpeech2(engine string, parent core.QObject_ITF) *QTextToSpeech {
	var engineC = C.CString(engine)
	defer C.free(unsafe.Pointer(engineC))
	return NewQTextToSpeechFromPointer(C.QTextToSpeech_NewQTextToSpeech2(engineC, core.PointerFromQObject(parent)))
}

//export callbackQTextToSpeech_LocaleChanged
func callbackQTextToSpeech_LocaleChanged(ptr unsafe.Pointer, locale unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "localeChanged"); signal != nil {
		signal.(func(*core.QLocale))(core.NewQLocaleFromPointer(locale))
	}

}

func (ptr *QTextToSpeech) ConnectLocaleChanged(f func(locale *core.QLocale)) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_ConnectLocaleChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "localeChanged", f)
	}
}

func (ptr *QTextToSpeech) DisconnectLocaleChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectLocaleChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "localeChanged")
	}
}

func (ptr *QTextToSpeech) LocaleChanged(locale core.QLocale_ITF) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_LocaleChanged(ptr.Pointer(), core.PointerFromQLocale(locale))
	}
}

//export callbackQTextToSpeech_Pause
func callbackQTextToSpeech_Pause(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "pause"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechFromPointer(ptr).PauseDefault()
	}
}

func (ptr *QTextToSpeech) ConnectPause(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "pause", f)
	}
}

func (ptr *QTextToSpeech) DisconnectPause() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "pause")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "pitchChanged"); signal != nil {
		signal.(func(float64))(float64(pitch))
	}

}

func (ptr *QTextToSpeech) ConnectPitchChanged(f func(pitch float64)) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_ConnectPitchChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "pitchChanged", f)
	}
}

func (ptr *QTextToSpeech) DisconnectPitchChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectPitchChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "pitchChanged")
	}
}

func (ptr *QTextToSpeech) PitchChanged(pitch float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_PitchChanged(ptr.Pointer(), C.double(pitch))
	}
}

//export callbackQTextToSpeech_RateChanged
func callbackQTextToSpeech_RateChanged(ptr unsafe.Pointer, rate C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "rateChanged"); signal != nil {
		signal.(func(float64))(float64(rate))
	}

}

func (ptr *QTextToSpeech) ConnectRateChanged(f func(rate float64)) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_ConnectRateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "rateChanged", f)
	}
}

func (ptr *QTextToSpeech) DisconnectRateChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectRateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "rateChanged")
	}
}

func (ptr *QTextToSpeech) RateChanged(rate float64) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_RateChanged(ptr.Pointer(), C.double(rate))
	}
}

//export callbackQTextToSpeech_Resume
func callbackQTextToSpeech_Resume(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "resume"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechFromPointer(ptr).ResumeDefault()
	}
}

func (ptr *QTextToSpeech) ConnectResume(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "resume", f)
	}
}

func (ptr *QTextToSpeech) DisconnectResume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "resume")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "say"); signal != nil {
		signal.(func(string))(cGoUnpackString(text))
	} else {
		NewQTextToSpeechFromPointer(ptr).SayDefault(cGoUnpackString(text))
	}
}

func (ptr *QTextToSpeech) ConnectSay(f func(text string)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "say", f)
	}
}

func (ptr *QTextToSpeech) DisconnectSay() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "say")
	}
}

func (ptr *QTextToSpeech) Say(text string) {
	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		C.QTextToSpeech_Say(ptr.Pointer(), textC)
	}
}

func (ptr *QTextToSpeech) SayDefault(text string) {
	if ptr.Pointer() != nil {
		var textC = C.CString(text)
		defer C.free(unsafe.Pointer(textC))
		C.QTextToSpeech_SayDefault(ptr.Pointer(), textC)
	}
}

//export callbackQTextToSpeech_SetLocale
func callbackQTextToSpeech_SetLocale(ptr unsafe.Pointer, locale unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setLocale"); signal != nil {
		signal.(func(*core.QLocale))(core.NewQLocaleFromPointer(locale))
	} else {
		NewQTextToSpeechFromPointer(ptr).SetLocaleDefault(core.NewQLocaleFromPointer(locale))
	}
}

func (ptr *QTextToSpeech) ConnectSetLocale(f func(locale *core.QLocale)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "setLocale", f)
	}
}

func (ptr *QTextToSpeech) DisconnectSetLocale() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "setLocale")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setPitch"); signal != nil {
		signal.(func(float64))(float64(pitch))
	} else {
		NewQTextToSpeechFromPointer(ptr).SetPitchDefault(float64(pitch))
	}
}

func (ptr *QTextToSpeech) ConnectSetPitch(f func(pitch float64)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "setPitch", f)
	}
}

func (ptr *QTextToSpeech) DisconnectSetPitch() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "setPitch")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setRate"); signal != nil {
		signal.(func(float64))(float64(rate))
	} else {
		NewQTextToSpeechFromPointer(ptr).SetRateDefault(float64(rate))
	}
}

func (ptr *QTextToSpeech) ConnectSetRate(f func(rate float64)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "setRate", f)
	}
}

func (ptr *QTextToSpeech) DisconnectSetRate() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "setRate")
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

//export callbackQTextToSpeech_SetVolume
func callbackQTextToSpeech_SetVolume(ptr unsafe.Pointer, volume C.double) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "setVolume"); signal != nil {
		signal.(func(float64))(float64(volume))
	} else {
		NewQTextToSpeechFromPointer(ptr).SetVolumeDefault(float64(volume))
	}
}

func (ptr *QTextToSpeech) ConnectSetVolume(f func(volume float64)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "setVolume", f)
	}
}

func (ptr *QTextToSpeech) DisconnectSetVolume() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "setVolume")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "stateChanged"); signal != nil {
		signal.(func(QTextToSpeech__State))(QTextToSpeech__State(state))
	}

}

func (ptr *QTextToSpeech) ConnectStateChanged(f func(state QTextToSpeech__State)) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "stateChanged", f)
	}
}

func (ptr *QTextToSpeech) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "stateChanged")
	}
}

func (ptr *QTextToSpeech) StateChanged(state QTextToSpeech__State) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQTextToSpeech_Stop
func callbackQTextToSpeech_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "stop"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechFromPointer(ptr).StopDefault()
	}
}

func (ptr *QTextToSpeech) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "stop", f)
	}
}

func (ptr *QTextToSpeech) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "stop")
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "volumeChanged"); signal != nil {
		signal.(func(int))(int(int32(volume)))
	}

}

func (ptr *QTextToSpeech) ConnectVolumeChanged(f func(volume int)) {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_ConnectVolumeChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "volumeChanged", f)
	}
}

func (ptr *QTextToSpeech) DisconnectVolumeChanged() {
	if ptr.Pointer() != nil {
		C.QTextToSpeech_DisconnectVolumeChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "volumeChanged")
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

func (ptr *QTextToSpeech) __availableVoices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QTextToSpeech___availableVoices_newList(ptr.Pointer()))
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
	if signal := qt.GetSignal(fmt.Sprint(ptr), "~QTextToSpeechPlugin"); signal != nil {
		signal.(func())()
	} else {
		NewQTextToSpeechPluginFromPointer(ptr).DestroyQTextToSpeechPluginDefault()
	}
}

func (ptr *QTextToSpeechPlugin) ConnectDestroyQTextToSpeechPlugin(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "~QTextToSpeechPlugin", f)
	}
}

func (ptr *QTextToSpeechPlugin) DisconnectDestroyQTextToSpeechPlugin() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "~QTextToSpeechPlugin")
	}
}

func (ptr *QTextToSpeechPlugin) DestroyQTextToSpeechPlugin() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechPlugin_DestroyQTextToSpeechPlugin(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextToSpeechPlugin) DestroyQTextToSpeechPluginDefault() {
	if ptr.Pointer() != nil {
		C.QTextToSpeechPlugin_DestroyQTextToSpeechPluginDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_atList(i string) *core.QVariant {
	if ptr.Pointer() != nil {
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		var tmpValue = core.NewQVariantFromPointer(C.QTextToSpeechPlugin___createTextToSpeechEngine_parameters_atList(ptr.Pointer(), iC))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_setList(key string, i core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		var keyC = C.CString(key)
		defer C.free(unsafe.Pointer(keyC))
		C.QTextToSpeechPlugin___createTextToSpeechEngine_parameters_setList(ptr.Pointer(), keyC, core.PointerFromQVariant(i))
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
		var iC = C.CString(i)
		defer C.free(unsafe.Pointer(iC))
		C.QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_setList(ptr.Pointer(), iC)
	}
}

func (ptr *QTextToSpeechPlugin) ____createTextToSpeechEngine_keyList_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_newList(ptr.Pointer()))
}
