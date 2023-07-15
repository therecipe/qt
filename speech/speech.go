//go:build !minimal
// +build !minimal

package speech

import (
	"unsafe"

	"github.com/akiyosi/qt/core"
	"github.com/akiyosi/qt/internal"
)

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

func (n *QTextToSpeech) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeech) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQTextToSpeechFromPointer(ptr unsafe.Pointer) (n *QTextToSpeech) {
	n = new(QTextToSpeech)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeech")
	return
}

func (ptr *QTextToSpeech) DestroyQTextToSpeech() {
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

func NewQTextToSpeech(parent core.QObject_ITF) *QTextToSpeech {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.NewQTextToSpeech", "", parent}).(*QTextToSpeech)
}

func NewQTextToSpeech2(engine string, parent core.QObject_ITF) *QTextToSpeech {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.NewQTextToSpeech2", "", engine, parent}).(*QTextToSpeech)
}

func QTextToSpeech_AvailableEngines() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.QTextToSpeech_AvailableEngines", ""}).([]string)
}

func (ptr *QTextToSpeech) AvailableEngines() []string {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.QTextToSpeech_AvailableEngines", ""}).([]string)
}

func (ptr *QTextToSpeech) AvailableLocales() []*core.QLocale {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableLocales"}).([]*core.QLocale)
}

func (ptr *QTextToSpeech) AvailableVoices() []*QVoice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableVoices"}).([]*QVoice)
}

func (ptr *QTextToSpeech) Locale() *core.QLocale {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Locale"}).(*core.QLocale)
}

func (ptr *QTextToSpeech) ConnectLocaleChanged(f func(locale *core.QLocale)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLocaleChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectLocaleChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLocaleChanged"})
}

func (ptr *QTextToSpeech) LocaleChanged(locale core.QLocale_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "LocaleChanged", locale})
}

func (ptr *QTextToSpeech) ConnectPause(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPause", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectPause() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPause"})
}

func (ptr *QTextToSpeech) Pause() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pause"})
}

func (ptr *QTextToSpeech) PauseDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PauseDefault"})
}

func (ptr *QTextToSpeech) Pitch() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pitch"}).(float64)
}

func (ptr *QTextToSpeech) ConnectPitchChanged(f func(pitch float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPitchChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectPitchChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPitchChanged"})
}

func (ptr *QTextToSpeech) PitchChanged(pitch float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "PitchChanged", pitch})
}

func (ptr *QTextToSpeech) Rate() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rate"}).(float64)
}

func (ptr *QTextToSpeech) ConnectRateChanged(f func(rate float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectRateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRateChanged"})
}

func (ptr *QTextToSpeech) RateChanged(rate float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "RateChanged", rate})
}

func (ptr *QTextToSpeech) ConnectResume(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectResume", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectResume() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectResume"})
}

func (ptr *QTextToSpeech) Resume() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Resume"})
}

func (ptr *QTextToSpeech) ResumeDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ResumeDefault"})
}

func (ptr *QTextToSpeech) ConnectSay(f func(text string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSay", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectSay() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSay"})
}

func (ptr *QTextToSpeech) Say(text string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Say", text})
}

func (ptr *QTextToSpeech) SayDefault(text string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SayDefault", text})
}

func (ptr *QTextToSpeech) ConnectSetLocale(f func(locale *core.QLocale)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetLocale", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectSetLocale() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetLocale"})
}

func (ptr *QTextToSpeech) SetLocale(locale core.QLocale_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocale", locale})
}

func (ptr *QTextToSpeech) SetLocaleDefault(locale core.QLocale_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocaleDefault", locale})
}

func (ptr *QTextToSpeech) ConnectSetPitch(f func(pitch float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetPitch", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectSetPitch() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetPitch"})
}

func (ptr *QTextToSpeech) SetPitch(pitch float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPitch", pitch})
}

func (ptr *QTextToSpeech) SetPitchDefault(pitch float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPitchDefault", pitch})
}

func (ptr *QTextToSpeech) ConnectSetRate(f func(rate float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetRate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectSetRate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetRate"})
}

func (ptr *QTextToSpeech) SetRate(rate float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRate", rate})
}

func (ptr *QTextToSpeech) SetRateDefault(rate float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRateDefault", rate})
}

func (ptr *QTextToSpeech) ConnectSetVoice(f func(voice *QVoice)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetVoice", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectSetVoice() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetVoice"})
}

func (ptr *QTextToSpeech) SetVoice(voice QVoice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVoice", voice})
}

func (ptr *QTextToSpeech) SetVoiceDefault(voice QVoice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVoiceDefault", voice})
}

func (ptr *QTextToSpeech) ConnectSetVolume(f func(volume float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetVolume", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectSetVolume() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetVolume"})
}

func (ptr *QTextToSpeech) SetVolume(volume float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVolume", volume})
}

func (ptr *QTextToSpeech) SetVolumeDefault(volume float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVolumeDefault", volume})
}

func (ptr *QTextToSpeech) State() QTextToSpeech__State {

	return QTextToSpeech__State(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QTextToSpeech) ConnectStateChanged(f func(state QTextToSpeech__State)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QTextToSpeech) StateChanged(state QTextToSpeech__State) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", state})
}

func (ptr *QTextToSpeech) ConnectStop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStop"})
}

func (ptr *QTextToSpeech) Stop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Stop"})
}

func (ptr *QTextToSpeech) StopDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StopDefault"})
}

func (ptr *QTextToSpeech) Voice() *QVoice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Voice"}).(*QVoice)
}

func (ptr *QTextToSpeech) Volume() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Volume"}).(float64)
}

func (ptr *QTextToSpeech) ConnectVolumeChanged(f func(volume int)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVolumeChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectVolumeChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVolumeChanged"})
}

func (ptr *QTextToSpeech) VolumeChanged(volume int) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VolumeChanged", volume})
}

func (ptr *QTextToSpeech) ConnectVolumeChanged2(f func(volume float64)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVolumeChanged2", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeech) DisconnectVolumeChanged2() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVolumeChanged2"})
}

func (ptr *QTextToSpeech) VolumeChanged2(volume float64) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "VolumeChanged2", volume})
}

func (ptr *QTextToSpeech) __availableLocales_atList(i int) *core.QLocale {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableLocales_atList", i}).(*core.QLocale)
}

func (ptr *QTextToSpeech) __availableLocales_setList(i core.QLocale_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableLocales_setList", i})
}

func (ptr *QTextToSpeech) __availableLocales_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableLocales_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeech) __availableVoices_atList(i int) *QVoice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableVoices_atList", i}).(*QVoice)
}

func (ptr *QTextToSpeech) __availableVoices_setList(i QVoice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableVoices_setList", i})
}

func (ptr *QTextToSpeech) __availableVoices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableVoices_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeech) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QTextToSpeech) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QTextToSpeech) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeech) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QTextToSpeech) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QTextToSpeech) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeech) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QTextToSpeech) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QTextToSpeech) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeech) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QTextToSpeech) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QTextToSpeech) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeech) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QTextToSpeech) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QTextToSpeech) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QTextToSpeech) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QTextToSpeech) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QTextToSpeech) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QTextToSpeech) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QTextToSpeech) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QTextToSpeech) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
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

func (n *QTextToSpeechEngine) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechEngine) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQTextToSpeechEngineFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechEngine) {
	n = new(QTextToSpeechEngine)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechEngine")
	return
}

func (ptr *QTextToSpeechEngine) DestroyQTextToSpeechEngine() {
}

func NewQTextToSpeechEngine(parent core.QObject_ITF) *QTextToSpeechEngine {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.NewQTextToSpeechEngine", "", parent}).(*QTextToSpeechEngine)
}

func (ptr *QTextToSpeechEngine) ConnectAvailableLocales(f func() []*core.QLocale) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAvailableLocales", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectAvailableLocales() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAvailableLocales"})
}

func (ptr *QTextToSpeechEngine) AvailableLocales() []*core.QLocale {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableLocales"}).([]*core.QLocale)
}

func (ptr *QTextToSpeechEngine) ConnectAvailableVoices(f func() []*QVoice) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectAvailableVoices", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectAvailableVoices() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectAvailableVoices"})
}

func (ptr *QTextToSpeechEngine) AvailableVoices() []*QVoice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "AvailableVoices"}).([]*QVoice)
}

func QTextToSpeechEngine_CreateVoice(name string, gender QVoice__Gender, age QVoice__Age, data core.QVariant_ITF) *QVoice {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.QTextToSpeechEngine_CreateVoice", "", name, gender, age, data}).(*QVoice)
}

func (ptr *QTextToSpeechEngine) CreateVoice(name string, gender QVoice__Gender, age QVoice__Age, data core.QVariant_ITF) *QVoice {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.QTextToSpeechEngine_CreateVoice", "", name, gender, age, data}).(*QVoice)
}

func (ptr *QTextToSpeechEngine) ConnectLocale(f func() *core.QLocale) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectLocale", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectLocale() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectLocale"})
}

func (ptr *QTextToSpeechEngine) Locale() *core.QLocale {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Locale"}).(*core.QLocale)
}

func (ptr *QTextToSpeechEngine) ConnectPause(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPause", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectPause() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPause"})
}

func (ptr *QTextToSpeechEngine) Pause() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pause"})
}

func (ptr *QTextToSpeechEngine) ConnectPitch(f func() float64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectPitch", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectPitch() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectPitch"})
}

func (ptr *QTextToSpeechEngine) Pitch() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Pitch"}).(float64)
}

func (ptr *QTextToSpeechEngine) ConnectRate(f func() float64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectRate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectRate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectRate"})
}

func (ptr *QTextToSpeechEngine) Rate() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Rate"}).(float64)
}

func (ptr *QTextToSpeechEngine) ConnectResume(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectResume", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectResume() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectResume"})
}

func (ptr *QTextToSpeechEngine) Resume() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Resume"})
}

func (ptr *QTextToSpeechEngine) ConnectSay(f func(text string)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSay", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectSay() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSay"})
}

func (ptr *QTextToSpeechEngine) Say(text string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Say", text})
}

func (ptr *QTextToSpeechEngine) ConnectSetLocale(f func(locale *core.QLocale) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetLocale", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectSetLocale() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetLocale"})
}

func (ptr *QTextToSpeechEngine) SetLocale(locale core.QLocale_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetLocale", locale}).(bool)
}

func (ptr *QTextToSpeechEngine) ConnectSetPitch(f func(pitch float64) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetPitch", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectSetPitch() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetPitch"})
}

func (ptr *QTextToSpeechEngine) SetPitch(pitch float64) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetPitch", pitch}).(bool)
}

func (ptr *QTextToSpeechEngine) ConnectSetRate(f func(rate float64) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetRate", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectSetRate() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetRate"})
}

func (ptr *QTextToSpeechEngine) SetRate(rate float64) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetRate", rate}).(bool)
}

func (ptr *QTextToSpeechEngine) ConnectSetVoice(f func(voice *QVoice) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetVoice", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectSetVoice() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetVoice"})
}

func (ptr *QTextToSpeechEngine) SetVoice(voice QVoice_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVoice", voice}).(bool)
}

func (ptr *QTextToSpeechEngine) ConnectSetVolume(f func(volume float64) bool) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectSetVolume", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectSetVolume() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectSetVolume"})
}

func (ptr *QTextToSpeechEngine) SetVolume(volume float64) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "SetVolume", volume}).(bool)
}

func (ptr *QTextToSpeechEngine) ConnectState(f func() QTextToSpeech__State) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectState", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectState() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectState"})
}

func (ptr *QTextToSpeechEngine) State() QTextToSpeech__State {

	return QTextToSpeech__State(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "State"}).(float64))
}

func (ptr *QTextToSpeechEngine) ConnectStateChanged(f func(state QTextToSpeech__State)) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStateChanged", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectStateChanged() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStateChanged"})
}

func (ptr *QTextToSpeechEngine) StateChanged(state QTextToSpeech__State) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "StateChanged", state})
}

func (ptr *QTextToSpeechEngine) ConnectStop(f func()) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectStop", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectStop() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectStop"})
}

func (ptr *QTextToSpeechEngine) Stop() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Stop"})
}

func (ptr *QTextToSpeechEngine) ConnectVoice(f func() *QVoice) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVoice", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectVoice() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVoice"})
}

func (ptr *QTextToSpeechEngine) Voice() *QVoice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Voice"}).(*QVoice)
}

func QTextToSpeechEngine_VoiceData(voice QVoice_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.QTextToSpeechEngine_VoiceData", "", voice}).(*core.QVariant)
}

func (ptr *QTextToSpeechEngine) VoiceData(voice QVoice_ITF) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.QTextToSpeechEngine_VoiceData", "", voice}).(*core.QVariant)
}

func (ptr *QTextToSpeechEngine) ConnectVolume(f func() float64) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectVolume", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechEngine) DisconnectVolume() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectVolume"})
}

func (ptr *QTextToSpeechEngine) Volume() float64 {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Volume"}).(float64)
}

func (ptr *QTextToSpeechEngine) __availableLocales_atList(i int) *core.QLocale {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableLocales_atList", i}).(*core.QLocale)
}

func (ptr *QTextToSpeechEngine) __availableLocales_setList(i core.QLocale_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableLocales_setList", i})
}

func (ptr *QTextToSpeechEngine) __availableLocales_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableLocales_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeechEngine) __availableVoices_atList(i int) *QVoice {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableVoices_atList", i}).(*QVoice)
}

func (ptr *QTextToSpeechEngine) __availableVoices_setList(i QVoice_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableVoices_setList", i})
}

func (ptr *QTextToSpeechEngine) __availableVoices_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__availableVoices_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeechEngine) __children_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_atList", i}).(*core.QObject)
}

func (ptr *QTextToSpeechEngine) __children_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_setList", i})
}

func (ptr *QTextToSpeechEngine) __children_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__children_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeechEngine) __dynamicPropertyNames_atList(i int) *core.QByteArray {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_atList", i}).(*core.QByteArray)
}

func (ptr *QTextToSpeechEngine) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_setList", i})
}

func (ptr *QTextToSpeechEngine) __dynamicPropertyNames_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__dynamicPropertyNames_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeechEngine) __findChildren_atList(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList", i}).(*core.QObject)
}

func (ptr *QTextToSpeechEngine) __findChildren_setList(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList", i})
}

func (ptr *QTextToSpeechEngine) __findChildren_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeechEngine) __findChildren_atList3(i int) *core.QObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_atList3", i}).(*core.QObject)
}

func (ptr *QTextToSpeechEngine) __findChildren_setList3(i core.QObject_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_setList3", i})
}

func (ptr *QTextToSpeechEngine) __findChildren_newList3() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__findChildren_newList3"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeechEngine) ChildEventDefault(event core.QChildEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ChildEventDefault", event})
}

func (ptr *QTextToSpeechEngine) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectNotifyDefault", sign})
}

func (ptr *QTextToSpeechEngine) CustomEventDefault(event core.QEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CustomEventDefault", event})
}

func (ptr *QTextToSpeechEngine) DeleteLaterDefault() {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DeleteLaterDefault"})
}

func (ptr *QTextToSpeechEngine) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectNotifyDefault", sign})
}

func (ptr *QTextToSpeechEngine) EventDefault(e core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventDefault", e}).(bool)
}

func (ptr *QTextToSpeechEngine) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "EventFilterDefault", watched, event}).(bool)
}

func (ptr *QTextToSpeechEngine) MetaObjectDefault() *core.QMetaObject {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "MetaObjectDefault"}).(*core.QMetaObject)
}

func (ptr *QTextToSpeechEngine) TimerEventDefault(event core.QTimerEvent_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "TimerEventDefault", event})
}

type QTextToSpeechEngineAndroid struct {
	QTextToSpeechEngine
}

type QTextToSpeechEngineAndroid_ITF interface {
	QTextToSpeechEngine_ITF
	QTextToSpeechEngineAndroid_PTR() *QTextToSpeechEngineAndroid
}

func (ptr *QTextToSpeechEngineAndroid) QTextToSpeechEngineAndroid_PTR() *QTextToSpeechEngineAndroid {
	return ptr
}

func (ptr *QTextToSpeechEngineAndroid) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechEngineAndroid) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextToSpeechEngine_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechEngineAndroid(ptr QTextToSpeechEngineAndroid_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngineAndroid_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechEngineAndroid) InitFromInternal(ptr uintptr, name string) {
	n.QTextToSpeechEngine_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechEngineAndroid) ClassNameInternalF() string {
	return n.QTextToSpeechEngine_PTR().ClassNameInternalF()
}

func NewQTextToSpeechEngineAndroidFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechEngineAndroid) {
	n = new(QTextToSpeechEngineAndroid)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechEngineAndroid")
	return
}

func (ptr *QTextToSpeechEngineAndroid) DestroyQTextToSpeechEngineAndroid() {
}

type QTextToSpeechEngineFlite struct {
	QTextToSpeechEngine
}

type QTextToSpeechEngineFlite_ITF interface {
	QTextToSpeechEngine_ITF
	QTextToSpeechEngineFlite_PTR() *QTextToSpeechEngineFlite
}

func (ptr *QTextToSpeechEngineFlite) QTextToSpeechEngineFlite_PTR() *QTextToSpeechEngineFlite {
	return ptr
}

func (ptr *QTextToSpeechEngineFlite) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechEngineFlite) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextToSpeechEngine_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechEngineFlite(ptr QTextToSpeechEngineFlite_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngineFlite_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechEngineFlite) InitFromInternal(ptr uintptr, name string) {
	n.QTextToSpeechEngine_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechEngineFlite) ClassNameInternalF() string {
	return n.QTextToSpeechEngine_PTR().ClassNameInternalF()
}

func NewQTextToSpeechEngineFliteFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechEngineFlite) {
	n = new(QTextToSpeechEngineFlite)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechEngineFlite")
	return
}

func (ptr *QTextToSpeechEngineFlite) DestroyQTextToSpeechEngineFlite() {
}

type QTextToSpeechEngineIos struct {
	QTextToSpeechEngine
}

type QTextToSpeechEngineIos_ITF interface {
	QTextToSpeechEngine_ITF
	QTextToSpeechEngineIos_PTR() *QTextToSpeechEngineIos
}

func (ptr *QTextToSpeechEngineIos) QTextToSpeechEngineIos_PTR() *QTextToSpeechEngineIos {
	return ptr
}

func (ptr *QTextToSpeechEngineIos) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechEngineIos) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextToSpeechEngine_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechEngineIos(ptr QTextToSpeechEngineIos_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngineIos_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechEngineIos) InitFromInternal(ptr uintptr, name string) {
	n.QTextToSpeechEngine_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechEngineIos) ClassNameInternalF() string {
	return n.QTextToSpeechEngine_PTR().ClassNameInternalF()
}

func NewQTextToSpeechEngineIosFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechEngineIos) {
	n = new(QTextToSpeechEngineIos)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechEngineIos")
	return
}

func (ptr *QTextToSpeechEngineIos) DestroyQTextToSpeechEngineIos() {
}

type QTextToSpeechEngineOsx struct {
	QTextToSpeechEngine
}

type QTextToSpeechEngineOsx_ITF interface {
	QTextToSpeechEngine_ITF
	QTextToSpeechEngineOsx_PTR() *QTextToSpeechEngineOsx
}

func (ptr *QTextToSpeechEngineOsx) QTextToSpeechEngineOsx_PTR() *QTextToSpeechEngineOsx {
	return ptr
}

func (ptr *QTextToSpeechEngineOsx) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechEngineOsx) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextToSpeechEngine_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechEngineOsx(ptr QTextToSpeechEngineOsx_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngineOsx_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechEngineOsx) InitFromInternal(ptr uintptr, name string) {
	n.QTextToSpeechEngine_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechEngineOsx) ClassNameInternalF() string {
	return n.QTextToSpeechEngine_PTR().ClassNameInternalF()
}

func NewQTextToSpeechEngineOsxFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechEngineOsx) {
	n = new(QTextToSpeechEngineOsx)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechEngineOsx")
	return
}

func (ptr *QTextToSpeechEngineOsx) DestroyQTextToSpeechEngineOsx() {
}

type QTextToSpeechEngineSapi struct {
	QTextToSpeechEngine
}

type QTextToSpeechEngineSapi_ITF interface {
	QTextToSpeechEngine_ITF
	QTextToSpeechEngineSapi_PTR() *QTextToSpeechEngineSapi
}

func (ptr *QTextToSpeechEngineSapi) QTextToSpeechEngineSapi_PTR() *QTextToSpeechEngineSapi {
	return ptr
}

func (ptr *QTextToSpeechEngineSapi) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechEngineSapi) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextToSpeechEngine_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechEngineSapi(ptr QTextToSpeechEngineSapi_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngineSapi_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechEngineSapi) InitFromInternal(ptr uintptr, name string) {
	n.QTextToSpeechEngine_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechEngineSapi) ClassNameInternalF() string {
	return n.QTextToSpeechEngine_PTR().ClassNameInternalF()
}

func NewQTextToSpeechEngineSapiFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechEngineSapi) {
	n = new(QTextToSpeechEngineSapi)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechEngineSapi")
	return
}

func (ptr *QTextToSpeechEngineSapi) DestroyQTextToSpeechEngineSapi() {
}

type QTextToSpeechEngineSpeechd struct {
	QTextToSpeechEngine
}

type QTextToSpeechEngineSpeechd_ITF interface {
	QTextToSpeechEngine_ITF
	QTextToSpeechEngineSpeechd_PTR() *QTextToSpeechEngineSpeechd
}

func (ptr *QTextToSpeechEngineSpeechd) QTextToSpeechEngineSpeechd_PTR() *QTextToSpeechEngineSpeechd {
	return ptr
}

func (ptr *QTextToSpeechEngineSpeechd) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechEngineSpeechd) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextToSpeechEngine_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechEngineSpeechd(ptr QTextToSpeechEngineSpeechd_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngineSpeechd_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechEngineSpeechd) InitFromInternal(ptr uintptr, name string) {
	n.QTextToSpeechEngine_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechEngineSpeechd) ClassNameInternalF() string {
	return n.QTextToSpeechEngine_PTR().ClassNameInternalF()
}

func NewQTextToSpeechEngineSpeechdFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechEngineSpeechd) {
	n = new(QTextToSpeechEngineSpeechd)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechEngineSpeechd")
	return
}

func (ptr *QTextToSpeechEngineSpeechd) DestroyQTextToSpeechEngineSpeechd() {
}

type QTextToSpeechEngineWinRT struct {
	QTextToSpeechEngine
}

type QTextToSpeechEngineWinRT_ITF interface {
	QTextToSpeechEngine_ITF
	QTextToSpeechEngineWinRT_PTR() *QTextToSpeechEngineWinRT
}

func (ptr *QTextToSpeechEngineWinRT) QTextToSpeechEngineWinRT_PTR() *QTextToSpeechEngineWinRT {
	return ptr
}

func (ptr *QTextToSpeechEngineWinRT) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngine_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechEngineWinRT) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QTextToSpeechEngine_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechEngineWinRT(ptr QTextToSpeechEngineWinRT_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechEngineWinRT_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechEngineWinRT) InitFromInternal(ptr uintptr, name string) {
	n.QTextToSpeechEngine_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechEngineWinRT) ClassNameInternalF() string {
	return n.QTextToSpeechEngine_PTR().ClassNameInternalF()
}

func NewQTextToSpeechEngineWinRTFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechEngineWinRT) {
	n = new(QTextToSpeechEngineWinRT)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechEngineWinRT")
	return
}

func (ptr *QTextToSpeechEngineWinRT) DestroyQTextToSpeechEngineWinRT() {
}

type QTextToSpeechPlugin struct {
	internal.Internal
}

type QTextToSpeechPlugin_ITF interface {
	QTextToSpeechPlugin_PTR() *QTextToSpeechPlugin
}

func (ptr *QTextToSpeechPlugin) QTextToSpeechPlugin_PTR() *QTextToSpeechPlugin {
	return ptr
}

func (ptr *QTextToSpeechPlugin) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QTextToSpeechPlugin) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQTextToSpeechPlugin(ptr QTextToSpeechPlugin_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechPlugin_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechPlugin) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQTextToSpeechPluginFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechPlugin) {
	n = new(QTextToSpeechPlugin)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechPlugin")
	return
}

func (ptr *QTextToSpeechPlugin) DestroyQTextToSpeechPlugin() {
}

func (ptr *QTextToSpeechPlugin) ConnectCreateTextToSpeechEngine(f func(parameters map[string]*core.QVariant, parent *core.QObject, errorString string) *QTextToSpeechEngine) {

	internal.CallLocalAndRegisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "ConnectCreateTextToSpeechEngine", "___REMOTE_CALLBACK___"}, f)
}

func (ptr *QTextToSpeechPlugin) DisconnectCreateTextToSpeechEngine() {

	internal.CallLocalAndDeregisterRemoteFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "DisconnectCreateTextToSpeechEngine"})
}

func (ptr *QTextToSpeechPlugin) CreateTextToSpeechEngine(parameters map[string]*core.QVariant, parent core.QObject_ITF, errorString string) *QTextToSpeechEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateTextToSpeechEngine", parameters, parent, errorString}).(*QTextToSpeechEngine)
}

func (ptr *QTextToSpeechPlugin) CreateTextToSpeechEngineDefault(parameters map[string]*core.QVariant, parent core.QObject_ITF, errorString string) *QTextToSpeechEngine {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "CreateTextToSpeechEngineDefault", parameters, parent, errorString}).(*QTextToSpeechEngine)
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_atList(v string, i int) *core.QVariant {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createTextToSpeechEngine_parameters_atList", v, i}).(*core.QVariant)
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_setList(key string, i core.QVariant_ITF) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createTextToSpeechEngine_parameters_setList", key, i})
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createTextToSpeechEngine_parameters_newList"}).(unsafe.Pointer)
}

func (ptr *QTextToSpeechPlugin) __createTextToSpeechEngine_parameters_keyList() []string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "__createTextToSpeechEngine_parameters_keyList"}).([]string)
}

func (ptr *QTextToSpeechPlugin) ____createTextToSpeechEngine_parameters_keyList_atList(i int) string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createTextToSpeechEngine_parameters_keyList_atList", i}).(string)
}

func (ptr *QTextToSpeechPlugin) ____createTextToSpeechEngine_parameters_keyList_setList(i string) {

	internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createTextToSpeechEngine_parameters_keyList_setList", i})
}

func (ptr *QTextToSpeechPlugin) ____createTextToSpeechEngine_parameters_keyList_newList() unsafe.Pointer {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "____createTextToSpeechEngine_parameters_keyList_newList"}).(unsafe.Pointer)
}

type QTextToSpeechPluginAndroid struct {
	core.QObject
	QTextToSpeechPlugin
}

type QTextToSpeechPluginAndroid_ITF interface {
	core.QObject_ITF
	QTextToSpeechPlugin_ITF
	QTextToSpeechPluginAndroid_PTR() *QTextToSpeechPluginAndroid
}

func (ptr *QTextToSpeechPluginAndroid) QTextToSpeechPluginAndroid_PTR() *QTextToSpeechPluginAndroid {
	return ptr
}

func (ptr *QTextToSpeechPluginAndroid) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechPluginAndroid) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QTextToSpeechPlugin_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechPluginAndroid(ptr QTextToSpeechPluginAndroid_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechPluginAndroid_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechPluginAndroid) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QTextToSpeechPlugin_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechPluginAndroid) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQTextToSpeechPluginAndroidFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechPluginAndroid) {
	n = new(QTextToSpeechPluginAndroid)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechPluginAndroid")
	return
}

func (ptr *QTextToSpeechPluginAndroid) DestroyQTextToSpeechPluginAndroid() {
}

type QTextToSpeechPluginFlite struct {
	core.QObject
	QTextToSpeechPlugin
}

type QTextToSpeechPluginFlite_ITF interface {
	core.QObject_ITF
	QTextToSpeechPlugin_ITF
	QTextToSpeechPluginFlite_PTR() *QTextToSpeechPluginFlite
}

func (ptr *QTextToSpeechPluginFlite) QTextToSpeechPluginFlite_PTR() *QTextToSpeechPluginFlite {
	return ptr
}

func (ptr *QTextToSpeechPluginFlite) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechPluginFlite) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QTextToSpeechPlugin_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechPluginFlite(ptr QTextToSpeechPluginFlite_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechPluginFlite_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechPluginFlite) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QTextToSpeechPlugin_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechPluginFlite) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQTextToSpeechPluginFliteFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechPluginFlite) {
	n = new(QTextToSpeechPluginFlite)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechPluginFlite")
	return
}

func (ptr *QTextToSpeechPluginFlite) DestroyQTextToSpeechPluginFlite() {
}

type QTextToSpeechPluginIos struct {
	core.QObject
	QTextToSpeechPlugin
}

type QTextToSpeechPluginIos_ITF interface {
	core.QObject_ITF
	QTextToSpeechPlugin_ITF
	QTextToSpeechPluginIos_PTR() *QTextToSpeechPluginIos
}

func (ptr *QTextToSpeechPluginIos) QTextToSpeechPluginIos_PTR() *QTextToSpeechPluginIos {
	return ptr
}

func (ptr *QTextToSpeechPluginIos) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechPluginIos) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QTextToSpeechPlugin_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechPluginIos(ptr QTextToSpeechPluginIos_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechPluginIos_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechPluginIos) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QTextToSpeechPlugin_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechPluginIos) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQTextToSpeechPluginIosFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechPluginIos) {
	n = new(QTextToSpeechPluginIos)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechPluginIos")
	return
}

func (ptr *QTextToSpeechPluginIos) DestroyQTextToSpeechPluginIos() {
}

type QTextToSpeechPluginOsx struct {
	core.QObject
	QTextToSpeechPlugin
}

type QTextToSpeechPluginOsx_ITF interface {
	core.QObject_ITF
	QTextToSpeechPlugin_ITF
	QTextToSpeechPluginOsx_PTR() *QTextToSpeechPluginOsx
}

func (ptr *QTextToSpeechPluginOsx) QTextToSpeechPluginOsx_PTR() *QTextToSpeechPluginOsx {
	return ptr
}

func (ptr *QTextToSpeechPluginOsx) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechPluginOsx) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QTextToSpeechPlugin_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechPluginOsx(ptr QTextToSpeechPluginOsx_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechPluginOsx_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechPluginOsx) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QTextToSpeechPlugin_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechPluginOsx) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQTextToSpeechPluginOsxFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechPluginOsx) {
	n = new(QTextToSpeechPluginOsx)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechPluginOsx")
	return
}

func (ptr *QTextToSpeechPluginOsx) DestroyQTextToSpeechPluginOsx() {
}

type QTextToSpeechPluginSapi struct {
	core.QObject
	QTextToSpeechPlugin
}

type QTextToSpeechPluginSapi_ITF interface {
	core.QObject_ITF
	QTextToSpeechPlugin_ITF
	QTextToSpeechPluginSapi_PTR() *QTextToSpeechPluginSapi
}

func (ptr *QTextToSpeechPluginSapi) QTextToSpeechPluginSapi_PTR() *QTextToSpeechPluginSapi {
	return ptr
}

func (ptr *QTextToSpeechPluginSapi) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechPluginSapi) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QTextToSpeechPlugin_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechPluginSapi(ptr QTextToSpeechPluginSapi_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechPluginSapi_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechPluginSapi) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QTextToSpeechPlugin_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechPluginSapi) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQTextToSpeechPluginSapiFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechPluginSapi) {
	n = new(QTextToSpeechPluginSapi)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechPluginSapi")
	return
}

func (ptr *QTextToSpeechPluginSapi) DestroyQTextToSpeechPluginSapi() {
}

type QTextToSpeechPluginSpeechd struct {
	core.QObject
	QTextToSpeechPlugin
}

type QTextToSpeechPluginSpeechd_ITF interface {
	core.QObject_ITF
	QTextToSpeechPlugin_ITF
	QTextToSpeechPluginSpeechd_PTR() *QTextToSpeechPluginSpeechd
}

func (ptr *QTextToSpeechPluginSpeechd) QTextToSpeechPluginSpeechd_PTR() *QTextToSpeechPluginSpeechd {
	return ptr
}

func (ptr *QTextToSpeechPluginSpeechd) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechPluginSpeechd) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QTextToSpeechPlugin_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechPluginSpeechd(ptr QTextToSpeechPluginSpeechd_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechPluginSpeechd_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechPluginSpeechd) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QTextToSpeechPlugin_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechPluginSpeechd) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQTextToSpeechPluginSpeechdFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechPluginSpeechd) {
	n = new(QTextToSpeechPluginSpeechd)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechPluginSpeechd")
	return
}

func (ptr *QTextToSpeechPluginSpeechd) DestroyQTextToSpeechPluginSpeechd() {
}

type QTextToSpeechPluginWinRT struct {
	core.QObject
	QTextToSpeechPlugin
}

type QTextToSpeechPluginWinRT_ITF interface {
	core.QObject_ITF
	QTextToSpeechPlugin_ITF
	QTextToSpeechPluginWinRT_PTR() *QTextToSpeechPluginWinRT
}

func (ptr *QTextToSpeechPluginWinRT) QTextToSpeechPluginWinRT_PTR() *QTextToSpeechPluginWinRT {
	return ptr
}

func (ptr *QTextToSpeechPluginWinRT) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QTextToSpeechPluginWinRT) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
		ptr.QTextToSpeechPlugin_PTR().SetPointer(p)
	}
}

func PointerFromQTextToSpeechPluginWinRT(ptr QTextToSpeechPluginWinRT_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechPluginWinRT_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechPluginWinRT) InitFromInternal(ptr uintptr, name string) {
	n.QObject_PTR().InitFromInternal(uintptr(ptr), name)
	n.QTextToSpeechPlugin_PTR().InitFromInternal(uintptr(ptr), name)

}

func (n *QTextToSpeechPluginWinRT) ClassNameInternalF() string {
	return n.QObject_PTR().ClassNameInternalF()
}

func NewQTextToSpeechPluginWinRTFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechPluginWinRT) {
	n = new(QTextToSpeechPluginWinRT)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechPluginWinRT")
	return
}

func (ptr *QTextToSpeechPluginWinRT) DestroyQTextToSpeechPluginWinRT() {
}

type QTextToSpeechProcessorFlite struct {
	internal.Internal
}

type QTextToSpeechProcessorFlite_ITF interface {
	QTextToSpeechProcessorFlite_PTR() *QTextToSpeechProcessorFlite
}

func (ptr *QTextToSpeechProcessorFlite) QTextToSpeechProcessorFlite_PTR() *QTextToSpeechProcessorFlite {
	return ptr
}

func (ptr *QTextToSpeechProcessorFlite) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QTextToSpeechProcessorFlite) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQTextToSpeechProcessorFlite(ptr QTextToSpeechProcessorFlite_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextToSpeechProcessorFlite_PTR().Pointer()
	}
	return nil
}

func (n *QTextToSpeechProcessorFlite) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQTextToSpeechProcessorFliteFromPointer(ptr unsafe.Pointer) (n *QTextToSpeechProcessorFlite) {
	n = new(QTextToSpeechProcessorFlite)
	n.InitFromInternal(uintptr(ptr), "speech.QTextToSpeechProcessorFlite")
	return
}

func (ptr *QTextToSpeechProcessorFlite) DestroyQTextToSpeechProcessorFlite() {
}

type QVoice struct {
	internal.Internal
}

type QVoice_ITF interface {
	QVoice_PTR() *QVoice
}

func (ptr *QVoice) QVoice_PTR() *QVoice {
	return ptr
}

func (ptr *QVoice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return unsafe.Pointer(ptr.Internal.Pointer())
	}
	return nil
}

func (ptr *QVoice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.Internal.SetPointer(uintptr(p))
	}
}

func PointerFromQVoice(ptr QVoice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVoice_PTR().Pointer()
	}
	return nil
}

func (n *QVoice) ClassNameInternalF() string {
	return n.Internal.ClassNameInternalF()
}

func NewQVoiceFromPointer(ptr unsafe.Pointer) (n *QVoice) {
	n = new(QVoice)
	n.InitFromInternal(uintptr(ptr), "speech.QVoice")
	return
}

func (ptr *QVoice) DestroyQVoice() {
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

func (ptr *QVoice) Age() QVoice__Age {

	return QVoice__Age(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Age"}).(float64))
}

func QVoice_AgeName(age QVoice__Age) string {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.QVoice_AgeName", "", age}).(string)
}

func (ptr *QVoice) AgeName(age QVoice__Age) string {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.QVoice_AgeName", "", age}).(string)
}

func (ptr *QVoice) Gender() QVoice__Gender {

	return QVoice__Gender(internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Gender"}).(float64))
}

func QVoice_GenderName(gender QVoice__Gender) string {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.QVoice_GenderName", "", gender}).(string)
}

func (ptr *QVoice) GenderName(gender QVoice__Gender) string {

	return internal.CallLocalFunction([]interface{}{"", "", "speech.QVoice_GenderName", "", gender}).(string)
}

func (ptr *QVoice) Name() string {

	return internal.CallLocalFunction([]interface{}{"", uintptr(ptr.Pointer()), ptr.ClassNameInternalF(), "Name"}).(string)
}

func init() {
	internal.ConstructorTable["speech.QTextToSpeech"] = NewQTextToSpeechFromPointer
	internal.ConstructorTable["speech.QTextToSpeechEngine"] = NewQTextToSpeechEngineFromPointer
	internal.ConstructorTable["speech.QTextToSpeechPlugin"] = NewQTextToSpeechPluginFromPointer
	internal.ConstructorTable["speech.QVoice"] = NewQVoiceFromPointer
}
