#include "qradiotuner.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
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

int QRadioTuner_Band(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->band();
}

int QRadioTuner_Frequency(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->frequency();
}

int QRadioTuner_IsAntennaConnected(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->isAntennaConnected();
}

int QRadioTuner_IsMuted(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->isMuted();
}

int QRadioTuner_IsSearching(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->isSearching();
}

int QRadioTuner_IsStereo(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->isStereo();
}

QtObjectPtr QRadioTuner_RadioData(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->radioData();
}

void QRadioTuner_SetMuted(QtObjectPtr ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

void QRadioTuner_SetStereoMode(QtObjectPtr ptr, int mode){
	static_cast<QRadioTuner*>(ptr)->setStereoMode(static_cast<QRadioTuner::StereoMode>(mode));
}

void QRadioTuner_SetVolume(QtObjectPtr ptr, int volume){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setVolume", Q_ARG(int, volume));
}

int QRadioTuner_SignalStrength(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->signalStrength();
}

int QRadioTuner_StereoMode(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->stereoMode();
}

int QRadioTuner_Volume(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->volume();
}

QtObjectPtr QRadioTuner_NewQRadioTuner(QtObjectPtr parent){
	return new QRadioTuner(static_cast<QObject*>(parent));
}

void QRadioTuner_ConnectAntennaConnectedChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::antennaConnectedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_AntennaConnectedChanged));;
}

void QRadioTuner_DisconnectAntennaConnectedChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::antennaConnectedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_AntennaConnectedChanged));;
}

void QRadioTuner_ConnectBandChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::Band)>(&QRadioTuner::bandChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::Band)>(&MyQRadioTuner::Signal_BandChanged));;
}

void QRadioTuner_DisconnectBandChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::Band)>(&QRadioTuner::bandChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::Band)>(&MyQRadioTuner::Signal_BandChanged));;
}

void QRadioTuner_CancelSearch(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "cancelSearch");
}

int QRadioTuner_Error(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->error();
}

char* QRadioTuner_ErrorString(QtObjectPtr ptr){
	return static_cast<QRadioTuner*>(ptr)->errorString().toUtf8().data();
}

void QRadioTuner_ConnectFrequencyChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::frequencyChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_FrequencyChanged));;
}

void QRadioTuner_DisconnectFrequencyChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::frequencyChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_FrequencyChanged));;
}

int QRadioTuner_FrequencyStep(QtObjectPtr ptr, int band){
	return static_cast<QRadioTuner*>(ptr)->frequencyStep(static_cast<QRadioTuner::Band>(band));
}

int QRadioTuner_IsBandSupported(QtObjectPtr ptr, int band){
	return static_cast<QRadioTuner*>(ptr)->isBandSupported(static_cast<QRadioTuner::Band>(band));
}

void QRadioTuner_ConnectMutedChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::mutedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_MutedChanged));;
}

void QRadioTuner_DisconnectMutedChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::mutedChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_MutedChanged));;
}

void QRadioTuner_SearchAllStations(QtObjectPtr ptr, int searchMode){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchAllStations", Q_ARG(QRadioTuner::SearchMode, static_cast<QRadioTuner::SearchMode>(searchMode)));
}

void QRadioTuner_SearchBackward(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchBackward");
}

void QRadioTuner_SearchForward(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "searchForward");
}

void QRadioTuner_ConnectSearchingChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::searchingChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_SearchingChanged));;
}

void QRadioTuner_DisconnectSearchingChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::searchingChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_SearchingChanged));;
}

void QRadioTuner_SetBand(QtObjectPtr ptr, int band){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setBand", Q_ARG(QRadioTuner::Band, static_cast<QRadioTuner::Band>(band)));
}

void QRadioTuner_SetFrequency(QtObjectPtr ptr, int frequency){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "setFrequency", Q_ARG(int, frequency));
}

void QRadioTuner_ConnectSignalStrengthChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::signalStrengthChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_SignalStrengthChanged));;
}

void QRadioTuner_DisconnectSignalStrengthChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::signalStrengthChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_SignalStrengthChanged));;
}

void QRadioTuner_Start(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "start");
}

void QRadioTuner_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::State)>(&QRadioTuner::stateChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::State)>(&MyQRadioTuner::Signal_StateChanged));;
}

void QRadioTuner_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(QRadioTuner::State)>(&QRadioTuner::stateChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(QRadioTuner::State)>(&MyQRadioTuner::Signal_StateChanged));;
}

void QRadioTuner_ConnectStationFound(QtObjectPtr ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int, QString)>(&QRadioTuner::stationFound), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int, QString)>(&MyQRadioTuner::Signal_StationFound));;
}

void QRadioTuner_DisconnectStationFound(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int, QString)>(&QRadioTuner::stationFound), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int, QString)>(&MyQRadioTuner::Signal_StationFound));;
}

void QRadioTuner_ConnectStereoStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::stereoStatusChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_StereoStatusChanged));;
}

void QRadioTuner_DisconnectStereoStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(bool)>(&QRadioTuner::stereoStatusChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(bool)>(&MyQRadioTuner::Signal_StereoStatusChanged));;
}

void QRadioTuner_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QRadioTuner*>(ptr), "stop");
}

void QRadioTuner_ConnectVolumeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::volumeChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_VolumeChanged));;
}

void QRadioTuner_DisconnectVolumeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QRadioTuner*>(ptr), static_cast<void (QRadioTuner::*)(int)>(&QRadioTuner::volumeChanged), static_cast<MyQRadioTuner*>(ptr), static_cast<void (MyQRadioTuner::*)(int)>(&MyQRadioTuner::Signal_VolumeChanged));;
}

void QRadioTuner_DestroyQRadioTuner(QtObjectPtr ptr){
	static_cast<QRadioTuner*>(ptr)->~QRadioTuner();
}

