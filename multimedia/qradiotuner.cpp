#include "qradiotuner.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QRadioTuner>
#include "_cgo_export.h"

class MyQRadioTuner: public QRadioTuner {
public:
void Signal_AntennaConnectedChanged(bool connectionStatus){callbackQRadioTunerAntennaConnectedChanged(this->objectName().toUtf8().data(), connectionStatus);};
void Signal_BandChanged(QRadioTuner::Band band){callbackQRadioTunerBandChanged(this->objectName().toUtf8().data(), band);};
void Signal_FrequencyChanged(int frequency){callbackQRadioTunerFrequencyChanged(this->objectName().toUtf8().data(), frequency);};
void Signal_MutedChanged(bool muted){callbackQRadioTunerMutedChanged(this->objectName().toUtf8().data(), muted);};
void Signal_SearchingChanged(bool searching){callbackQRadioTunerSearchingChanged(this->objectName().toUtf8().data(), searching);};
void Signal_SignalStrengthChanged(int strength){callbackQRadioTunerSignalStrengthChanged(this->objectName().toUtf8().data(), strength);};
void Signal_StateChanged(QRadioTuner::State state){callbackQRadioTunerStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_StationFound(int frequency, QString stationId){callbackQRadioTunerStationFound(this->objectName().toUtf8().data(), frequency, stationId.toUtf8().data());};
void Signal_StereoStatusChanged(bool stereo){callbackQRadioTunerStereoStatusChanged(this->objectName().toUtf8().data(), stereo);};
void Signal_VolumeChanged(int volume){callbackQRadioTunerVolumeChanged(this->objectName().toUtf8().data(), volume);};
};

int QRadioTuner_Band(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->band();
}

int QRadioTuner_Frequency(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->frequency();
}

int QRadioTuner_IsAntennaConnected(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->isAntennaConnected();
}

int QRadioTuner_IsMuted(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->isMuted();
}

int QRadioTuner_IsSearching(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->isSearching();
}

int QRadioTuner_IsStereo(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->isStereo();
}

void* QRadioTuner_RadioData(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->radioData();
}

void QRadioTuner_SetMuted(void* ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

void QRadioTuner_SetStereoMode(void* ptr, int mode){
	static_cast<QRadioTuner*>(ptr)->setStereoMode(static_cast<QRadioTuner::StereoMode>(mode));
}

void QRadioTuner_SetVolume(void* ptr, int volume){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setVolume", Q_ARG(int, volume));
}

int QRadioTuner_SignalStrength(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->signalStrength();
}

int QRadioTuner_StereoMode(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->stereoMode();
}

int QRadioTuner_Volume(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->volume();
}

void* QRadioTuner_NewQRadioTuner(void* parent){
	return new QRadioTuner(static_cast<QObject*>(parent));
}

void QRadioTuner_ConnectAntennaConnectedChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::antennaConnectedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_AntennaConnectedChanged));;
}

void QRadioTuner_DisconnectAntennaConnectedChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::antennaConnectedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_AntennaConnectedChanged));;
}

void QRadioTuner_ConnectBandChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::Band)>(&QRadioTuner::bandChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::Band)>(&MyQRadioTuner::Signal_BandChanged));;
}

void QRadioTuner_DisconnectBandChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::Band)>(&QRadioTuner::bandChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::Band)>(&MyQRadioTuner::Signal_BandChanged));;
}

void QRadioTuner_CancelSearch(void* ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "cancelSearch");
}

int QRadioTuner_Error(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->error();
}

char* QRadioTuner_ErrorString(void* ptr){
	return static_cast<QRadioTuner*>(ptr)->errorString().toUtf8().data();
}

void QRadioTuner_ConnectFrequencyChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::frequencyChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_FrequencyChanged));;
}

void QRadioTuner_DisconnectFrequencyChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::frequencyChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_FrequencyChanged));;
}

int QRadioTuner_FrequencyStep(void* ptr, int band){
	return static_cast<QRadioTuner*>(ptr)->frequencyStep(static_cast<QRadioTuner::Band>(band));
}

int QRadioTuner_IsBandSupported(void* ptr, int band){
	return static_cast<QRadioTuner*>(ptr)->isBandSupported(static_cast<QRadioTuner::Band>(band));
}

void QRadioTuner_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::mutedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_MutedChanged));;
}

void QRadioTuner_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::mutedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_MutedChanged));;
}

void QRadioTuner_SearchAllStations(void* ptr, int searchMode){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchAllStations", Q_ARG(QRadioTuner::SearchMode, static_cast<QRadioTuner::SearchMode>(searchMode)));
}

void QRadioTuner_SearchBackward(void* ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchBackward");
}

void QRadioTuner_SearchForward(void* ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchForward");
}

void QRadioTuner_ConnectSearchingChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::searchingChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_SearchingChanged));;
}

void QRadioTuner_DisconnectSearchingChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::searchingChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_SearchingChanged));;
}

void QRadioTuner_SetBand(void* ptr, int band){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setBand", Q_ARG(QRadioTuner::Band, static_cast<QRadioTuner::Band>(band)));
}

void QRadioTuner_SetFrequency(void* ptr, int frequency){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setFrequency", Q_ARG(int, frequency));
}

void QRadioTuner_ConnectSignalStrengthChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::signalStrengthChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_SignalStrengthChanged));;
}

void QRadioTuner_DisconnectSignalStrengthChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::signalStrengthChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_SignalStrengthChanged));;
}

void QRadioTuner_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "start");
}

void QRadioTuner_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::State)>(&QRadioTuner::stateChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::State)>(&MyQRadioTuner::Signal_StateChanged));;
}

void QRadioTuner_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::State)>(&QRadioTuner::stateChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::State)>(&MyQRadioTuner::Signal_StateChanged));;
}

void QRadioTuner_ConnectStationFound(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int, QString)>(&QRadioTuner::stationFound), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int, QString)>(&MyQRadioTuner::Signal_StationFound));;
}

void QRadioTuner_DisconnectStationFound(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int, QString)>(&QRadioTuner::stationFound), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int, QString)>(&MyQRadioTuner::Signal_StationFound));;
}

void QRadioTuner_ConnectStereoStatusChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::stereoStatusChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_StereoStatusChanged));;
}

void QRadioTuner_DisconnectStereoStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::stereoStatusChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_StereoStatusChanged));;
}

void QRadioTuner_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "stop");
}

void QRadioTuner_ConnectVolumeChanged(void* ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::volumeChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_VolumeChanged));;
}

void QRadioTuner_DisconnectVolumeChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::volumeChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_VolumeChanged));;
}

void QRadioTuner_DestroyQRadioTuner(void* ptr){
	static_cast<QRadioTuner*>(ptr)->~QRadioTuner();
}

