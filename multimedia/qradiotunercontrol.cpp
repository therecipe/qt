#include "qradiotunercontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRadioTuner>
#include <QObject>
#include <QRadioTunerControl>
#include "_cgo_export.h"

class MyQRadioTunerControl: public QRadioTunerControl {
public:
void Signal_AntennaConnectedChanged(bool connectionStatus){callbackQRadioTunerControlAntennaConnectedChanged(this->objectName().toUtf8().data(), connectionStatus);};
void Signal_BandChanged(QRadioTuner::Band band){callbackQRadioTunerControlBandChanged(this->objectName().toUtf8().data(), band);};
void Signal_FrequencyChanged(int frequency){callbackQRadioTunerControlFrequencyChanged(this->objectName().toUtf8().data(), frequency);};
void Signal_MutedChanged(bool muted){callbackQRadioTunerControlMutedChanged(this->objectName().toUtf8().data(), muted);};
void Signal_SearchingChanged(bool searching){callbackQRadioTunerControlSearchingChanged(this->objectName().toUtf8().data(), searching);};
void Signal_SignalStrengthChanged(int strength){callbackQRadioTunerControlSignalStrengthChanged(this->objectName().toUtf8().data(), strength);};
void Signal_StateChanged(QRadioTuner::State state){callbackQRadioTunerControlStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_StationFound(int frequency, QString stationId){callbackQRadioTunerControlStationFound(this->objectName().toUtf8().data(), frequency, stationId.toUtf8().data());};
void Signal_StereoStatusChanged(bool stereo){callbackQRadioTunerControlStereoStatusChanged(this->objectName().toUtf8().data(), stereo);};
void Signal_VolumeChanged(int volume){callbackQRadioTunerControlVolumeChanged(this->objectName().toUtf8().data(), volume);};
};

void QRadioTunerControl_ConnectAntennaConnectedChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::antennaConnectedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_AntennaConnectedChanged));;
}

void QRadioTunerControl_DisconnectAntennaConnectedChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::antennaConnectedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_AntennaConnectedChanged));;
}

int QRadioTunerControl_Band(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->band();
}

void QRadioTunerControl_ConnectBandChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::Band)>(&QRadioTunerControl::bandChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::Band)>(&MyQRadioTunerControl::Signal_BandChanged));;
}

void QRadioTunerControl_DisconnectBandChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::Band)>(&QRadioTunerControl::bandChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::Band)>(&MyQRadioTunerControl::Signal_BandChanged));;
}

void QRadioTunerControl_CancelSearch(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->cancelSearch();
}

int QRadioTunerControl_Error(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->error();
}

char* QRadioTunerControl_ErrorString(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->errorString().toUtf8().data();
}

int QRadioTunerControl_Frequency(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->frequency();
}

void QRadioTunerControl_ConnectFrequencyChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::frequencyChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_FrequencyChanged));;
}

void QRadioTunerControl_DisconnectFrequencyChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::frequencyChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_FrequencyChanged));;
}

int QRadioTunerControl_FrequencyStep(void* ptr, int band){
	return static_cast<QRadioTunerControl*>(ptr)->frequencyStep(static_cast<QRadioTuner::Band>(band));
}

int QRadioTunerControl_IsAntennaConnected(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->isAntennaConnected();
}

int QRadioTunerControl_IsBandSupported(void* ptr, int band){
	return static_cast<QRadioTunerControl*>(ptr)->isBandSupported(static_cast<QRadioTuner::Band>(band));
}

int QRadioTunerControl_IsMuted(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->isMuted();
}

int QRadioTunerControl_IsSearching(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->isSearching();
}

int QRadioTunerControl_IsStereo(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->isStereo();
}

void QRadioTunerControl_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::mutedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_MutedChanged));;
}

void QRadioTunerControl_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::mutedChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_MutedChanged));;
}

void QRadioTunerControl_SearchAllStations(void* ptr, int searchMode){
	static_cast<QRadioTunerControl*>(ptr)->searchAllStations(static_cast<QRadioTuner::SearchMode>(searchMode));
}

void QRadioTunerControl_SearchBackward(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->searchBackward();
}

void QRadioTunerControl_SearchForward(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->searchForward();
}

void QRadioTunerControl_ConnectSearchingChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::searchingChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_SearchingChanged));;
}

void QRadioTunerControl_DisconnectSearchingChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::searchingChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_SearchingChanged));;
}

void QRadioTunerControl_SetBand(void* ptr, int band){
	static_cast<QRadioTunerControl*>(ptr)->setBand(static_cast<QRadioTuner::Band>(band));
}

void QRadioTunerControl_SetFrequency(void* ptr, int frequency){
	static_cast<QRadioTunerControl*>(ptr)->setFrequency(frequency);
}

void QRadioTunerControl_SetMuted(void* ptr, int muted){
	static_cast<QRadioTunerControl*>(ptr)->setMuted(muted != 0);
}

void QRadioTunerControl_SetStereoMode(void* ptr, int mode){
	static_cast<QRadioTunerControl*>(ptr)->setStereoMode(static_cast<QRadioTuner::StereoMode>(mode));
}

void QRadioTunerControl_SetVolume(void* ptr, int volume){
	static_cast<QRadioTunerControl*>(ptr)->setVolume(volume);
}

int QRadioTunerControl_SignalStrength(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->signalStrength();
}

void QRadioTunerControl_ConnectSignalStrengthChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::signalStrengthChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_SignalStrengthChanged));;
}

void QRadioTunerControl_DisconnectSignalStrengthChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::signalStrengthChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_SignalStrengthChanged));;
}

void QRadioTunerControl_Start(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->start();
}

void QRadioTunerControl_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::State)>(&QRadioTunerControl::stateChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::State)>(&MyQRadioTunerControl::Signal_StateChanged));;
}

void QRadioTunerControl_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(QRadioTuner::State)>(&QRadioTunerControl::stateChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(QRadioTuner::State)>(&MyQRadioTunerControl::Signal_StateChanged));;
}

void QRadioTunerControl_ConnectStationFound(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int, QString)>(&QRadioTunerControl::stationFound), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int, QString)>(&MyQRadioTunerControl::Signal_StationFound));;
}

void QRadioTunerControl_DisconnectStationFound(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int, QString)>(&QRadioTunerControl::stationFound), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int, QString)>(&MyQRadioTunerControl::Signal_StationFound));;
}

int QRadioTunerControl_StereoMode(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->stereoMode();
}

void QRadioTunerControl_ConnectStereoStatusChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::stereoStatusChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_StereoStatusChanged));;
}

void QRadioTunerControl_DisconnectStereoStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(bool)>(&QRadioTunerControl::stereoStatusChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(bool)>(&MyQRadioTunerControl::Signal_StereoStatusChanged));;
}

void QRadioTunerControl_Stop(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->stop();
}

int QRadioTunerControl_Volume(void* ptr){
	return static_cast<QRadioTunerControl*>(ptr)->volume();
}

void QRadioTunerControl_ConnectVolumeChanged(void* ptr){
	QObject::connect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::volumeChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_VolumeChanged));;
}

void QRadioTunerControl_DisconnectVolumeChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTunerControl*>(ptr), static_cast<void (QRadioTunerControl::*)(int)>(&QRadioTunerControl::volumeChanged), static_cast<MyQRadioTunerControl*>(ptr), static_cast<void (MyQRadioTunerControl::*)(int)>(&MyQRadioTunerControl::Signal_VolumeChanged));;
}

void QRadioTunerControl_DestroyQRadioTunerControl(void* ptr){
	static_cast<QRadioTunerControl*>(ptr)->~QRadioTunerControl();
}

