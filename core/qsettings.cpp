#include "qsettings.h"
#include <QUrl>
#include <QModelIndex>
#include <QSet>
#include <QObject>
#include <QTextCodec>
#include <QString>
#include <QVariant>
#include <QSettings>
#include "_cgo_export.h"

class MyQSettings: public QSettings {
public:
};

QtObjectPtr QSettings_NewQSettings3(int format, int scope, char* organization, char* application, QtObjectPtr parent){
	return new QSettings(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString(organization), QString(application), static_cast<QObject*>(parent));
}

QtObjectPtr QSettings_NewQSettings5(QtObjectPtr parent){
	return new QSettings(static_cast<QObject*>(parent));
}

QtObjectPtr QSettings_NewQSettings2(int scope, char* organization, char* application, QtObjectPtr parent){
	return new QSettings(static_cast<QSettings::Scope>(scope), QString(organization), QString(application), static_cast<QObject*>(parent));
}

QtObjectPtr QSettings_NewQSettings4(char* fileName, int format, QtObjectPtr parent){
	return new QSettings(QString(fileName), static_cast<QSettings::Format>(format), static_cast<QObject*>(parent));
}

QtObjectPtr QSettings_NewQSettings(char* organization, char* application, QtObjectPtr parent){
	return new QSettings(QString(organization), QString(application), static_cast<QObject*>(parent));
}

char* QSettings_AllKeys(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->allKeys().join("|").toUtf8().data();
}

char* QSettings_ApplicationName(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->applicationName().toUtf8().data();
}

void QSettings_BeginGroup(QtObjectPtr ptr, char* prefix){
	static_cast<QSettings*>(ptr)->beginGroup(QString(prefix));
}

int QSettings_BeginReadArray(QtObjectPtr ptr, char* prefix){
	return static_cast<QSettings*>(ptr)->beginReadArray(QString(prefix));
}

void QSettings_BeginWriteArray(QtObjectPtr ptr, char* prefix, int size){
	static_cast<QSettings*>(ptr)->beginWriteArray(QString(prefix), size);
}

char* QSettings_ChildGroups(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->childGroups().join("|").toUtf8().data();
}

char* QSettings_ChildKeys(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->childKeys().join("|").toUtf8().data();
}

void QSettings_Clear(QtObjectPtr ptr){
	static_cast<QSettings*>(ptr)->clear();
}

int QSettings_Contains(QtObjectPtr ptr, char* key){
	return static_cast<QSettings*>(ptr)->contains(QString(key));
}

int QSettings_QSettings_DefaultFormat(){
	return QSettings::defaultFormat();
}

void QSettings_EndArray(QtObjectPtr ptr){
	static_cast<QSettings*>(ptr)->endArray();
}

void QSettings_EndGroup(QtObjectPtr ptr){
	static_cast<QSettings*>(ptr)->endGroup();
}

int QSettings_FallbacksEnabled(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->fallbacksEnabled();
}

char* QSettings_FileName(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->fileName().toUtf8().data();
}

int QSettings_Format(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->format();
}

char* QSettings_Group(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->group().toUtf8().data();
}

QtObjectPtr QSettings_IniCodec(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->iniCodec();
}

int QSettings_IsWritable(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->isWritable();
}

char* QSettings_OrganizationName(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->organizationName().toUtf8().data();
}

int QSettings_Scope(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->scope();
}

void QSettings_SetArrayIndex(QtObjectPtr ptr, int i){
	static_cast<QSettings*>(ptr)->setArrayIndex(i);
}

void QSettings_QSettings_SetDefaultFormat(int format){
	QSettings::setDefaultFormat(static_cast<QSettings::Format>(format));
}

void QSettings_SetFallbacksEnabled(QtObjectPtr ptr, int b){
	static_cast<QSettings*>(ptr)->setFallbacksEnabled(b != 0);
}

void QSettings_SetIniCodec(QtObjectPtr ptr, QtObjectPtr codec){
	static_cast<QSettings*>(ptr)->setIniCodec(static_cast<QTextCodec*>(codec));
}

void QSettings_SetIniCodec2(QtObjectPtr ptr, char* codecName){
	static_cast<QSettings*>(ptr)->setIniCodec(const_cast<const char*>(codecName));
}

void QSettings_QSettings_SetPath(int format, int scope, char* path){
	QSettings::setPath(static_cast<QSettings::Format>(format), static_cast<QSettings::Scope>(scope), QString(path));
}

void QSettings_SetValue(QtObjectPtr ptr, char* key, char* value){
	static_cast<QSettings*>(ptr)->setValue(QString(key), QVariant(value));
}

int QSettings_Status(QtObjectPtr ptr){
	return static_cast<QSettings*>(ptr)->status();
}

void QSettings_Sync(QtObjectPtr ptr){
	static_cast<QSettings*>(ptr)->sync();
}

char* QSettings_Value(QtObjectPtr ptr, char* key, char* defaultValue){
	return static_cast<QSettings*>(ptr)->value(QString(key), QVariant(defaultValue)).toString().toUtf8().data();
}

void QSettings_DestroyQSettings(QtObjectPtr ptr){
	static_cast<QSettings*>(ptr)->~QSettings();
}

