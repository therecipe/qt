#include "qsettings.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextCodec>
#include <QSet>
#include <QObject>
#include <QString>
#include <QSettings>
#include "_cgo_export.h"

class MyQSettings: public QSettings {
public:
};

void* QSettings_NewQSettings3(int format, int scope, char* organization, char* application, void* parent){
	return new QSettings(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString(organization), QString(application), static_cast<QObject*>(parent));
}

void* QSettings_NewQSettings5(void* parent){
	return new QSettings(static_cast<QObject*>(parent));
}

void* QSettings_NewQSettings2(int scope, char* organization, char* application, void* parent){
	return new QSettings(static_cast<QSettings::Scope>(scope), QString(organization), QString(application), static_cast<QObject*>(parent));
}

void* QSettings_NewQSettings4(char* fileName, int format, void* parent){
	return new QSettings(QString(fileName), static_cast<QSettings::Format>(format), static_cast<QObject*>(parent));
}

void* QSettings_NewQSettings(char* organization, char* application, void* parent){
	return new QSettings(QString(organization), QString(application), static_cast<QObject*>(parent));
}

char* QSettings_AllKeys(void* ptr){
	return static_cast<QSettings*>(ptr)->allKeys().join("|").toUtf8().data();
}

char* QSettings_ApplicationName(void* ptr){
	return static_cast<QSettings*>(ptr)->applicationName().toUtf8().data();
}

void QSettings_BeginGroup(void* ptr, char* prefix){
	static_cast<QSettings*>(ptr)->beginGroup(QString(prefix));
}

int QSettings_BeginReadArray(void* ptr, char* prefix){
	return static_cast<QSettings*>(ptr)->beginReadArray(QString(prefix));
}

void QSettings_BeginWriteArray(void* ptr, char* prefix, int size){
	static_cast<QSettings*>(ptr)->beginWriteArray(QString(prefix), size);
}

char* QSettings_ChildGroups(void* ptr){
	return static_cast<QSettings*>(ptr)->childGroups().join("|").toUtf8().data();
}

char* QSettings_ChildKeys(void* ptr){
	return static_cast<QSettings*>(ptr)->childKeys().join("|").toUtf8().data();
}

void QSettings_Clear(void* ptr){
	static_cast<QSettings*>(ptr)->clear();
}

int QSettings_Contains(void* ptr, char* key){
	return static_cast<QSettings*>(ptr)->contains(QString(key));
}

int QSettings_QSettings_DefaultFormat(){
	return QSettings::defaultFormat();
}

void QSettings_EndArray(void* ptr){
	static_cast<QSettings*>(ptr)->endArray();
}

void QSettings_EndGroup(void* ptr){
	static_cast<QSettings*>(ptr)->endGroup();
}

int QSettings_FallbacksEnabled(void* ptr){
	return static_cast<QSettings*>(ptr)->fallbacksEnabled();
}

char* QSettings_FileName(void* ptr){
	return static_cast<QSettings*>(ptr)->fileName().toUtf8().data();
}

int QSettings_Format(void* ptr){
	return static_cast<QSettings*>(ptr)->format();
}

char* QSettings_Group(void* ptr){
	return static_cast<QSettings*>(ptr)->group().toUtf8().data();
}

void* QSettings_IniCodec(void* ptr){
	return static_cast<QSettings*>(ptr)->iniCodec();
}

int QSettings_IsWritable(void* ptr){
	return static_cast<QSettings*>(ptr)->isWritable();
}

char* QSettings_OrganizationName(void* ptr){
	return static_cast<QSettings*>(ptr)->organizationName().toUtf8().data();
}

int QSettings_Scope(void* ptr){
	return static_cast<QSettings*>(ptr)->scope();
}

void QSettings_SetArrayIndex(void* ptr, int i){
	static_cast<QSettings*>(ptr)->setArrayIndex(i);
}

void QSettings_QSettings_SetDefaultFormat(int format){
	QSettings::setDefaultFormat(static_cast<QSettings::Format>(format));
}

void QSettings_SetFallbacksEnabled(void* ptr, int b){
	static_cast<QSettings*>(ptr)->setFallbacksEnabled(b != 0);
}

void QSettings_SetIniCodec(void* ptr, void* codec){
	static_cast<QSettings*>(ptr)->setIniCodec(static_cast<QTextCodec*>(codec));
}

void QSettings_SetIniCodec2(void* ptr, char* codecName){
	static_cast<QSettings*>(ptr)->setIniCodec(const_cast<const char*>(codecName));
}

void QSettings_QSettings_SetPath(int format, int scope, char* path){
	QSettings::setPath(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString(path));
}

void QSettings_SetValue(void* ptr, char* key, void* value){
	static_cast<QSettings*>(ptr)->setValue(QString(key), *static_cast<QVariant*>(value));
}

int QSettings_Status(void* ptr){
	return static_cast<QSettings*>(ptr)->status();
}

void QSettings_Sync(void* ptr){
	static_cast<QSettings*>(ptr)->sync();
}

void* QSettings_Value(void* ptr, char* key, void* defaultValue){
	return new QVariant(static_cast<QSettings*>(ptr)->value(QString(key), *static_cast<QVariant*>(defaultValue)));
}

void QSettings_DestroyQSettings(void* ptr){
	static_cast<QSettings*>(ptr)->~QSettings();
}

