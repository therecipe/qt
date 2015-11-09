#include "qradiodatacontrol.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QRadioData>
#include <QString>
#include <QRadioDataControl>
#include "_cgo_export.h"

class MyQRadioDataControl: public QRadioDataControl {
public:
void Signal_AlternativeFrequenciesEnabledChanged(bool enabled){callbackQRadioDataControlAlternativeFrequenciesEnabledChanged(this->objectName().toUtf8().data(), enabled);};
void Signal_ProgramTypeChanged(QRadioData::ProgramType programType){callbackQRadioDataControlProgramTypeChanged(this->objectName().toUtf8().data(), programType);};
void Signal_ProgramTypeNameChanged(QString programTypeName){callbackQRadioDataControlProgramTypeNameChanged(this->objectName().toUtf8().data(), programTypeName.toUtf8().data());};
void Signal_RadioTextChanged(QString radioText){callbackQRadioDataControlRadioTextChanged(this->objectName().toUtf8().data(), radioText.toUtf8().data());};
void Signal_StationIdChanged(QString stationId){callbackQRadioDataControlStationIdChanged(this->objectName().toUtf8().data(), stationId.toUtf8().data());};
void Signal_StationNameChanged(QString stationName){callbackQRadioDataControlStationNameChanged(this->objectName().toUtf8().data(), stationName.toUtf8().data());};
};

void QRadioDataControl_ConnectAlternativeFrequenciesEnabledChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(bool)>(&QRadioDataControl::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(bool)>(&MyQRadioDataControl::Signal_AlternativeFrequenciesEnabledChanged));;
}

void QRadioDataControl_DisconnectAlternativeFrequenciesEnabledChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(bool)>(&QRadioDataControl::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(bool)>(&MyQRadioDataControl::Signal_AlternativeFrequenciesEnabledChanged));;
}

int QRadioDataControl_Error(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->error();
}

char* QRadioDataControl_ErrorString(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->errorString().toUtf8().data();
}

int QRadioDataControl_IsAlternativeFrequenciesEnabled(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->isAlternativeFrequenciesEnabled();
}

int QRadioDataControl_ProgramType(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->programType();
}

void QRadioDataControl_ConnectProgramTypeChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QRadioData::ProgramType)>(&QRadioDataControl::programTypeChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QRadioData::ProgramType)>(&MyQRadioDataControl::Signal_ProgramTypeChanged));;
}

void QRadioDataControl_DisconnectProgramTypeChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QRadioData::ProgramType)>(&QRadioDataControl::programTypeChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QRadioData::ProgramType)>(&MyQRadioDataControl::Signal_ProgramTypeChanged));;
}

char* QRadioDataControl_ProgramTypeName(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->programTypeName().toUtf8().data();
}

void QRadioDataControl_ConnectProgramTypeNameChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::programTypeNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_ProgramTypeNameChanged));;
}

void QRadioDataControl_DisconnectProgramTypeNameChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::programTypeNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_ProgramTypeNameChanged));;
}

char* QRadioDataControl_RadioText(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->radioText().toUtf8().data();
}

void QRadioDataControl_ConnectRadioTextChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::radioTextChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_RadioTextChanged));;
}

void QRadioDataControl_DisconnectRadioTextChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::radioTextChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_RadioTextChanged));;
}

void QRadioDataControl_SetAlternativeFrequenciesEnabled(void* ptr, int enabled){
	static_cast<QRadioDataControl*>(ptr)->setAlternativeFrequenciesEnabled(enabled != 0);
}

char* QRadioDataControl_StationId(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->stationId().toUtf8().data();
}

void QRadioDataControl_ConnectStationIdChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationIdChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationIdChanged));;
}

void QRadioDataControl_DisconnectStationIdChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationIdChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationIdChanged));;
}

char* QRadioDataControl_StationName(void* ptr){
	return static_cast<QRadioDataControl*>(ptr)->stationName().toUtf8().data();
}

void QRadioDataControl_ConnectStationNameChanged(void* ptr){
	QObject::connect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationNameChanged));;
}

void QRadioDataControl_DisconnectStationNameChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioDataControl*>(ptr), static_cast<void (QRadioDataControl::*)(QString)>(&QRadioDataControl::stationNameChanged), static_cast<MyQRadioDataControl*>(ptr), static_cast<void (MyQRadioDataControl::*)(QString)>(&MyQRadioDataControl::Signal_StationNameChanged));;
}

void QRadioDataControl_DestroyQRadioDataControl(void* ptr){
	static_cast<QRadioDataControl*>(ptr)->~QRadioDataControl();
}

