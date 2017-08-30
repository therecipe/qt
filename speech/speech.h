// +build !minimal

#pragma once

#ifndef GO_QTSPEECH_H
#define GO_QTSPEECH_H

#include <stdint.h>

#ifdef __cplusplus
int QTextToSpeech_QTextToSpeech_QRegisterMetaType();
extern "C" {
#endif

struct QtSpeech_PackedString { char* data; long long len; };
struct QtSpeech_PackedList { void* data; long long len; };
struct QtSpeech_PackedString QTextToSpeech_QTextToSpeech_AvailableEngines();
void* QTextToSpeech_NewQTextToSpeech(void* parent);
void* QTextToSpeech_NewQTextToSpeech2(struct QtSpeech_PackedString engine, void* parent);
void QTextToSpeech_ConnectLocaleChanged(void* ptr);
void QTextToSpeech_DisconnectLocaleChanged(void* ptr);
void QTextToSpeech_LocaleChanged(void* ptr, void* locale);
void QTextToSpeech_Pause(void* ptr);
void QTextToSpeech_PauseDefault(void* ptr);
void QTextToSpeech_ConnectPitchChanged(void* ptr);
void QTextToSpeech_DisconnectPitchChanged(void* ptr);
void QTextToSpeech_PitchChanged(void* ptr, double pitch);
void QTextToSpeech_ConnectRateChanged(void* ptr);
void QTextToSpeech_DisconnectRateChanged(void* ptr);
void QTextToSpeech_RateChanged(void* ptr, double rate);
void QTextToSpeech_Resume(void* ptr);
void QTextToSpeech_ResumeDefault(void* ptr);
void QTextToSpeech_Say(void* ptr, struct QtSpeech_PackedString text);
void QTextToSpeech_SayDefault(void* ptr, struct QtSpeech_PackedString text);
void QTextToSpeech_SetLocale(void* ptr, void* locale);
void QTextToSpeech_SetLocaleDefault(void* ptr, void* locale);
void QTextToSpeech_SetPitch(void* ptr, double pitch);
void QTextToSpeech_SetPitchDefault(void* ptr, double pitch);
void QTextToSpeech_SetRate(void* ptr, double rate);
void QTextToSpeech_SetRateDefault(void* ptr, double rate);
void QTextToSpeech_SetVoice(void* ptr, void* voice);
void QTextToSpeech_SetVoiceDefault(void* ptr, void* voice);
void QTextToSpeech_SetVolume(void* ptr, double volume);
void QTextToSpeech_SetVolumeDefault(void* ptr, double volume);
void QTextToSpeech_ConnectStateChanged(void* ptr);
void QTextToSpeech_DisconnectStateChanged(void* ptr);
void QTextToSpeech_StateChanged(void* ptr, long long state);
void QTextToSpeech_Stop(void* ptr);
void QTextToSpeech_StopDefault(void* ptr);
void QTextToSpeech_ConnectVolumeChanged(void* ptr);
void QTextToSpeech_DisconnectVolumeChanged(void* ptr);
void QTextToSpeech_VolumeChanged(void* ptr, int volume);
void* QTextToSpeech_Locale(void* ptr);
struct QtSpeech_PackedList QTextToSpeech_AvailableLocales(void* ptr);
struct QtSpeech_PackedList QTextToSpeech_AvailableVoices(void* ptr);
void* QTextToSpeech_Voice(void* ptr);
long long QTextToSpeech_State(void* ptr);
double QTextToSpeech_Pitch(void* ptr);
double QTextToSpeech_Rate(void* ptr);
double QTextToSpeech_Volume(void* ptr);
void* QTextToSpeech___availableLocales_atList(void* ptr, int i);
void QTextToSpeech___availableLocales_setList(void* ptr, void* i);
void* QTextToSpeech___availableLocales_newList(void* ptr);
void* QTextToSpeech___availableVoices_atList(void* ptr, int i);
void QTextToSpeech___availableVoices_setList(void* ptr, void* i);
void* QTextToSpeech___availableVoices_newList(void* ptr);
void* QTextToSpeech___dynamicPropertyNames_atList(void* ptr, int i);
void QTextToSpeech___dynamicPropertyNames_setList(void* ptr, void* i);
void* QTextToSpeech___dynamicPropertyNames_newList(void* ptr);
void* QTextToSpeech___findChildren_atList2(void* ptr, int i);
void QTextToSpeech___findChildren_setList2(void* ptr, void* i);
void* QTextToSpeech___findChildren_newList2(void* ptr);
void* QTextToSpeech___findChildren_atList3(void* ptr, int i);
void QTextToSpeech___findChildren_setList3(void* ptr, void* i);
void* QTextToSpeech___findChildren_newList3(void* ptr);
void* QTextToSpeech___findChildren_atList(void* ptr, int i);
void QTextToSpeech___findChildren_setList(void* ptr, void* i);
void* QTextToSpeech___findChildren_newList(void* ptr);
void* QTextToSpeech___children_atList(void* ptr, int i);
void QTextToSpeech___children_setList(void* ptr, void* i);
void* QTextToSpeech___children_newList(void* ptr);
char QTextToSpeech_EventDefault(void* ptr, void* e);
char QTextToSpeech_EventFilterDefault(void* ptr, void* watched, void* event);
void QTextToSpeech_ChildEventDefault(void* ptr, void* event);
void QTextToSpeech_ConnectNotifyDefault(void* ptr, void* sign);
void QTextToSpeech_CustomEventDefault(void* ptr, void* event);
void QTextToSpeech_DeleteLaterDefault(void* ptr);
void QTextToSpeech_DisconnectNotifyDefault(void* ptr, void* sign);
void QTextToSpeech_TimerEventDefault(void* ptr, void* event);
void* QTextToSpeech_MetaObjectDefault(void* ptr);
void QTextToSpeechPlugin_DestroyQTextToSpeechPlugin(void* ptr);
void QTextToSpeechPlugin_DestroyQTextToSpeechPluginDefault(void* ptr);
void* QTextToSpeechPlugin___createTextToSpeechEngine_parameters_atList(void* ptr, struct QtSpeech_PackedString i);
void QTextToSpeechPlugin___createTextToSpeechEngine_parameters_setList(void* ptr, struct QtSpeech_PackedString key, void* i);
void* QTextToSpeechPlugin___createTextToSpeechEngine_parameters_newList(void* ptr);
struct QtSpeech_PackedList QTextToSpeechPlugin___createTextToSpeechEngine_keyList(void* ptr);
struct QtSpeech_PackedString QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_atList(void* ptr, int i);
void QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_setList(void* ptr, struct QtSpeech_PackedString i);
void* QTextToSpeechPlugin_____createTextToSpeechEngine_keyList_newList(void* ptr);
struct QtSpeech_PackedString QVoice_QVoice_AgeName(long long age);
struct QtSpeech_PackedString QVoice_QVoice_GenderName(long long gender);
void* QVoice_NewQVoice();
void* QVoice_NewQVoice2(void* other);
void QVoice_DestroyQVoice(void* ptr);
long long QVoice_Age(void* ptr);
long long QVoice_Gender(void* ptr);
struct QtSpeech_PackedString QVoice_Name(void* ptr);

#ifdef __cplusplus
}
#endif

#endif