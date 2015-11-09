#include "qradiodata.h"
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaObject>
#include <QObject>
#include <QRadioData>
#include "_cgo_export.h"

class MyQRadioData: public QRadioData {
public:
void Signal_AlternativeFrequenciesEnabledChanged(bool enabled){callbackQRadioDataAlternativeFrequenciesEnabledChanged(this->objectName().toUtf8().data(), enabled);};
void Signal_ProgramTypeChanged(QRadioData::ProgramType programType){callbackQRadioDataProgramTypeChanged(this->objectName().toUtf8().data(), programType);};
void Signal_ProgramTypeNameChanged(QString programTypeName){callbackQRadioDataProgramTypeNameChanged(this->objectName().toUtf8().data(), programTypeName.toUtf8().data());};
void Signal_RadioTextChanged(QString radioText){callbackQRadioDataRadioTextChanged(this->objectName().toUtf8().data(), radioText.toUtf8().data());};
void Signal_StationIdChanged(QString stationId){callbackQRadioDataStationIdChanged(this->objectName().toUtf8().data(), stationId.toUtf8().data());};
void Signal_StationNameChanged(QString stationName){callbackQRadioDataStationNameChanged(this->objectName().toUtf8().data(), stationName.toUtf8().data());};
};

int QRadioData_IsAlternativeFrequenciesEnabled(void* ptr){
	return static_cast<QRadioData*>(ptr)->isAlternativeFrequenciesEnabled();
}

int QRadioData_ProgramType(void* ptr){
	return static_cast<QRadioData*>(ptr)->programType();
}

char* QRadioData_ProgramTypeName(void* ptr){
	return static_cast<QRadioData*>(ptr)->programTypeName().toUtf8().data();
}

char* QRadioData_RadioText(void* ptr){
	return static_cast<QRadioData*>(ptr)->radioText().toUtf8().data();
}

void QRadioData_SetAlternativeFrequenciesEnabled(void* ptr, int enabled){
	QMetaObject::invokeMethod(static_cast<QRadioData*>(ptr), "setAlternativeFrequenciesEnabled", Q_ARG(bool, enabled != 0));
}

char* QRadioData_StationId(void* ptr){
	return static_cast<QRadioData*>(ptr)->stationId().toUtf8().data();
}

char* QRadioData_StationName(void* ptr){
	return static_cast<QRadioData*>(ptr)->stationName().toUtf8().data();
}

void* QRadioData_NewQRadioData(void* mediaObject, void* parent){
	return new QRadioData(static_cast<QMediaObject*>(mediaObject), static_cast<QObject*>(parent));
}

void QRadioData_ConnectAlternativeFrequenciesEnabledChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(bool)>(&QRadioData::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(bool)>(&MyQRadioData::Signal_AlternativeFrequenciesEnabledChanged));;
}

void QRadioData_DisconnectAlternativeFrequenciesEnabledChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(bool)>(&QRadioData::alternativeFrequenciesEnabledChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(bool)>(&MyQRadioData::Signal_AlternativeFrequenciesEnabledChanged));;
}

int QRadioData_Error(void* ptr){
	return static_cast<QRadioData*>(ptr)->error();
}

char* QRadioData_ErrorString(void* ptr){
	return static_cast<QRadioData*>(ptr)->errorString().toUtf8().data();
}

void* QRadioData_MediaObject(void* ptr){
	return static_cast<QRadioData*>(ptr)->mediaObject();
}

void QRadioData_ConnectProgramTypeChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QRadioData::ProgramType)>(&QRadioData::programTypeChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QRadioData::ProgramType)>(&MyQRadioData::Signal_ProgramTypeChanged));;
}

void QRadioData_DisconnectProgramTypeChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QRadioData::ProgramType)>(&QRadioData::programTypeChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QRadioData::ProgramType)>(&MyQRadioData::Signal_ProgramTypeChanged));;
}

void QRadioData_ConnectProgramTypeNameChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::programTypeNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_ProgramTypeNameChanged));;
}

void QRadioData_DisconnectProgramTypeNameChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::programTypeNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_ProgramTypeNameChanged));;
}

void QRadioData_ConnectRadioTextChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::radioTextChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_RadioTextChanged));;
}

void QRadioData_DisconnectRadioTextChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::radioTextChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_RadioTextChanged));;
}

void QRadioData_ConnectStationIdChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationIdChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationIdChanged));;
}

void QRadioData_DisconnectStationIdChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationIdChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationIdChanged));;
}

void QRadioData_ConnectStationNameChanged(void* ptr){
	QObject::connect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationNameChanged));;
}

void QRadioData_DisconnectStationNameChanged(void* ptr){
	QObject::disconnect(static_cast<QRadioData*>(ptr), static_cast<void (QRadioData::*)(QString)>(&QRadioData::stationNameChanged), static_cast<MyQRadioData*>(ptr), static_cast<void (MyQRadioData::*)(QString)>(&MyQRadioData::Signal_StationNameChanged));;
}

void QRadioData_DestroyQRadioData(void* ptr){
	static_cast<QRadioData*>(ptr)->~QRadioData();
}

