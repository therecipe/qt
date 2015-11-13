#include "qndefmessage.h"
#include <QNdefRecord>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNdefMessage>
#include "_cgo_export.h"

class MyQNdefMessage: public QNdefMessage {
public:
};

void* QNdefMessage_NewQNdefMessage(){
	return new QNdefMessage();
}

void* QNdefMessage_NewQNdefMessage3(void* message){
	return new QNdefMessage(*static_cast<QNdefMessage*>(message));
}

void* QNdefMessage_NewQNdefMessage2(void* record){
	return new QNdefMessage(*static_cast<QNdefRecord*>(record));
}

void* QNdefMessage_ToByteArray(void* ptr){
	return new QByteArray(static_cast<QNdefMessage*>(ptr)->toByteArray());
}

#include "qndefnfcurirecord.h"
#include <QModelIndex>
#include <QNdefRecord>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QNdefNfcUriRecord>
#include "_cgo_export.h"

class MyQNdefNfcUriRecord: public QNdefNfcUriRecord {
public:
};

void* QNdefNfcUriRecord_NewQNdefNfcUriRecord(){
	return new QNdefNfcUriRecord();
}

void* QNdefNfcUriRecord_NewQNdefNfcUriRecord2(void* other){
	return new QNdefNfcUriRecord(*static_cast<QNdefRecord*>(other));
}

void QNdefNfcUriRecord_SetUri(void* ptr, void* uri){
	static_cast<QNdefNfcUriRecord*>(ptr)->setUri(*static_cast<QUrl*>(uri));
}

#include "qndefnfctextrecord.h"
#include <QUrl>
#include <QModelIndex>
#include <QNdefRecord>
#include <QString>
#include <QVariant>
#include <QNdefNfcTextRecord>
#include "_cgo_export.h"

class MyQNdefNfcTextRecord: public QNdefNfcTextRecord {
public:
};

void* QNdefNfcTextRecord_NewQNdefNfcTextRecord(){
	return new QNdefNfcTextRecord();
}

void* QNdefNfcTextRecord_NewQNdefNfcTextRecord2(void* other){
	return new QNdefNfcTextRecord(*static_cast<QNdefRecord*>(other));
}

int QNdefNfcTextRecord_Encoding(void* ptr){
	return static_cast<QNdefNfcTextRecord*>(ptr)->encoding();
}

char* QNdefNfcTextRecord_Locale(void* ptr){
	return static_cast<QNdefNfcTextRecord*>(ptr)->locale().toUtf8().data();
}

void QNdefNfcTextRecord_SetEncoding(void* ptr, int encoding){
	static_cast<QNdefNfcTextRecord*>(ptr)->setEncoding(static_cast<QNdefNfcTextRecord::Encoding>(encoding));
}

void QNdefNfcTextRecord_SetLocale(void* ptr, char* locale){
	static_cast<QNdefNfcTextRecord*>(ptr)->setLocale(QString(locale));
}

void QNdefNfcTextRecord_SetText(void* ptr, char* text){
	static_cast<QNdefNfcTextRecord*>(ptr)->setText(QString(text));
}

char* QNdefNfcTextRecord_Text(void* ptr){
	return static_cast<QNdefNfcTextRecord*>(ptr)->text().toUtf8().data();
}

#include "qnearfieldtarget.h"
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QNearFieldTarget>
#include "_cgo_export.h"

class MyQNearFieldTarget: public QNearFieldTarget {
public:
void Signal_Disconnected(){callbackQNearFieldTargetDisconnected(this->objectName().toUtf8().data());};
void Signal_NdefMessagesWritten(){callbackQNearFieldTargetNdefMessagesWritten(this->objectName().toUtf8().data());};
};

int QNearFieldTarget_AccessMethods(void* ptr){
	return static_cast<QNearFieldTarget*>(ptr)->accessMethods();
}

void QNearFieldTarget_ConnectDisconnected(void* ptr){
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::disconnected), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_Disconnected));;
}

void QNearFieldTarget_DisconnectDisconnected(void* ptr){
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::disconnected), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_Disconnected));;
}

int QNearFieldTarget_HasNdefMessage(void* ptr){
	return static_cast<QNearFieldTarget*>(ptr)->hasNdefMessage();
}

int QNearFieldTarget_IsProcessingCommand(void* ptr){
	return static_cast<QNearFieldTarget*>(ptr)->isProcessingCommand();
}

void QNearFieldTarget_ConnectNdefMessagesWritten(void* ptr){
	QObject::connect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::ndefMessagesWritten), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_NdefMessagesWritten));;
}

void QNearFieldTarget_DisconnectNdefMessagesWritten(void* ptr){
	QObject::disconnect(static_cast<QNearFieldTarget*>(ptr), static_cast<void (QNearFieldTarget::*)()>(&QNearFieldTarget::ndefMessagesWritten), static_cast<MyQNearFieldTarget*>(ptr), static_cast<void (MyQNearFieldTarget::*)()>(&MyQNearFieldTarget::Signal_NdefMessagesWritten));;
}

int QNearFieldTarget_Type(void* ptr){
	return static_cast<QNearFieldTarget*>(ptr)->type();
}

void* QNearFieldTarget_Uid(void* ptr){
	return new QByteArray(static_cast<QNearFieldTarget*>(ptr)->uid());
}

void QNearFieldTarget_DestroyQNearFieldTarget(void* ptr){
	static_cast<QNearFieldTarget*>(ptr)->~QNearFieldTarget();
}

#include "qnearfieldsharetarget.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNearFieldShareManager>
#include <QNdefMessage>
#include <QNearFieldShareTarget>
#include "_cgo_export.h"

class MyQNearFieldShareTarget: public QNearFieldShareTarget {
public:
void Signal_Error(QNearFieldShareManager::ShareError error){callbackQNearFieldShareTargetError(this->objectName().toUtf8().data(), error);};
void Signal_ShareFinished(){callbackQNearFieldShareTargetShareFinished(this->objectName().toUtf8().data());};
};

void QNearFieldShareTarget_Cancel(void* ptr){
	static_cast<QNearFieldShareTarget*>(ptr)->cancel();
}

void QNearFieldShareTarget_ConnectError(void* ptr){
	QObject::connect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareTarget::error), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareTarget::Signal_Error));;
}

void QNearFieldShareTarget_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareTarget::error), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareTarget::Signal_Error));;
}

int QNearFieldShareTarget_IsShareInProgress(void* ptr){
	return static_cast<QNearFieldShareTarget*>(ptr)->isShareInProgress();
}

int QNearFieldShareTarget_Share(void* ptr, void* message){
	return static_cast<QNearFieldShareTarget*>(ptr)->share(*static_cast<QNdefMessage*>(message));
}

int QNearFieldShareTarget_ShareError(void* ptr){
	return static_cast<QNearFieldShareTarget*>(ptr)->shareError();
}

void QNearFieldShareTarget_ConnectShareFinished(void* ptr){
	QObject::connect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)()>(&QNearFieldShareTarget::shareFinished), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)()>(&MyQNearFieldShareTarget::Signal_ShareFinished));;
}

void QNearFieldShareTarget_DisconnectShareFinished(void* ptr){
	QObject::disconnect(static_cast<QNearFieldShareTarget*>(ptr), static_cast<void (QNearFieldShareTarget::*)()>(&QNearFieldShareTarget::shareFinished), static_cast<MyQNearFieldShareTarget*>(ptr), static_cast<void (MyQNearFieldShareTarget::*)()>(&MyQNearFieldShareTarget::Signal_ShareFinished));;
}

int QNearFieldShareTarget_ShareModes(void* ptr){
	return static_cast<QNearFieldShareTarget*>(ptr)->shareModes();
}

void QNearFieldShareTarget_DestroyQNearFieldShareTarget(void* ptr){
	static_cast<QNearFieldShareTarget*>(ptr)->~QNearFieldShareTarget();
}

#include "qqmlndefrecord.h"
#include <QNdefRecord>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlNdefRecord>
#include "_cgo_export.h"

class MyQQmlNdefRecord: public QQmlNdefRecord {
public:
void Signal_RecordChanged(){callbackQQmlNdefRecordRecordChanged(this->objectName().toUtf8().data());};
void Signal_TypeChanged(){callbackQQmlNdefRecordTypeChanged(this->objectName().toUtf8().data());};
void Signal_TypeNameFormatChanged(){callbackQQmlNdefRecordTypeNameFormatChanged(this->objectName().toUtf8().data());};
};

int QQmlNdefRecord_TypeNameFormat(void* ptr){
	return static_cast<QQmlNdefRecord*>(ptr)->typeNameFormat();
}

void* QQmlNdefRecord_NewQQmlNdefRecord(void* parent){
	return new QQmlNdefRecord(static_cast<QObject*>(parent));
}

void* QQmlNdefRecord_NewQQmlNdefRecord2(void* record, void* parent){
	return new QQmlNdefRecord(*static_cast<QNdefRecord*>(record), static_cast<QObject*>(parent));
}

void QQmlNdefRecord_ConnectRecordChanged(void* ptr){
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::recordChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_RecordChanged));;
}

void QQmlNdefRecord_DisconnectRecordChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::recordChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_RecordChanged));;
}

void QQmlNdefRecord_SetRecord(void* ptr, void* record){
	static_cast<QQmlNdefRecord*>(ptr)->setRecord(*static_cast<QNdefRecord*>(record));
}

void QQmlNdefRecord_SetType(void* ptr, char* newtype){
	static_cast<QQmlNdefRecord*>(ptr)->setType(QString(newtype));
}

void QQmlNdefRecord_SetTypeNameFormat(void* ptr, int newTypeNameFormat){
	static_cast<QQmlNdefRecord*>(ptr)->setTypeNameFormat(static_cast<QQmlNdefRecord::TypeNameFormat>(newTypeNameFormat));
}

char* QQmlNdefRecord_Type(void* ptr){
	return static_cast<QQmlNdefRecord*>(ptr)->type().toUtf8().data();
}

void QQmlNdefRecord_ConnectTypeChanged(void* ptr){
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeChanged));;
}

void QQmlNdefRecord_DisconnectTypeChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeChanged));;
}

void QQmlNdefRecord_ConnectTypeNameFormatChanged(void* ptr){
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));;
}

void QQmlNdefRecord_DisconnectTypeNameFormatChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));;
}

#include "qnearfieldsharemanager.h"
#include <QModelIndex>
#include <QObject>
#include <QNearFieldShareTarget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QNearFieldShareManager>
#include "_cgo_export.h"

class MyQNearFieldShareManager: public QNearFieldShareManager {
public:
void Signal_Error(QNearFieldShareManager::ShareError error){callbackQNearFieldShareManagerError(this->objectName().toUtf8().data(), error);};
void Signal_ShareModesChanged(QNearFieldShareManager::ShareModes modes){callbackQNearFieldShareManagerShareModesChanged(this->objectName().toUtf8().data(), modes);};
void Signal_TargetDetected(QNearFieldShareTarget * shareTarget){callbackQNearFieldShareManagerTargetDetected(this->objectName().toUtf8().data(), shareTarget);};
};

void* QNearFieldShareManager_NewQNearFieldShareManager(void* parent){
	return new QNearFieldShareManager(static_cast<QObject*>(parent));
}

void QNearFieldShareManager_ConnectError(void* ptr){
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareManager::error), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareManager::Signal_Error));;
}

void QNearFieldShareManager_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&QNearFieldShareManager::error), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareError)>(&MyQNearFieldShareManager::Signal_Error));;
}

void QNearFieldShareManager_SetShareModes(void* ptr, int mode){
	static_cast<QNearFieldShareManager*>(ptr)->setShareModes(static_cast<QNearFieldShareManager::ShareMode>(mode));
}

int QNearFieldShareManager_ShareError(void* ptr){
	return static_cast<QNearFieldShareManager*>(ptr)->shareError();
}

int QNearFieldShareManager_ShareModes(void* ptr){
	return static_cast<QNearFieldShareManager*>(ptr)->shareModes();
}

void QNearFieldShareManager_ConnectShareModesChanged(void* ptr){
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&QNearFieldShareManager::shareModesChanged), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&MyQNearFieldShareManager::Signal_ShareModesChanged));;
}

void QNearFieldShareManager_DisconnectShareModesChanged(void* ptr){
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&QNearFieldShareManager::shareModesChanged), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareManager::ShareModes)>(&MyQNearFieldShareManager::Signal_ShareModesChanged));;
}

int QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes(){
	return QNearFieldShareManager::supportedShareModes();
}

void QNearFieldShareManager_ConnectTargetDetected(void* ptr){
	QObject::connect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareTarget *)>(&QNearFieldShareManager::targetDetected), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareTarget *)>(&MyQNearFieldShareManager::Signal_TargetDetected));;
}

void QNearFieldShareManager_DisconnectTargetDetected(void* ptr){
	QObject::disconnect(static_cast<QNearFieldShareManager*>(ptr), static_cast<void (QNearFieldShareManager::*)(QNearFieldShareTarget *)>(&QNearFieldShareManager::targetDetected), static_cast<MyQNearFieldShareManager*>(ptr), static_cast<void (MyQNearFieldShareManager::*)(QNearFieldShareTarget *)>(&MyQNearFieldShareManager::Signal_TargetDetected));;
}

void QNearFieldShareManager_DestroyQNearFieldShareManager(void* ptr){
	static_cast<QNearFieldShareManager*>(ptr)->~QNearFieldShareManager();
}

#include "qndefnfcsmartposterrecord.h"
#include <QNdefRecord>
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNdefNfcUriRecord>
#include <QNdefNfcTextRecord>
#include <QNdefNfcSmartPosterRecord>
#include "_cgo_export.h"

class MyQNdefNfcSmartPosterRecord: public QNdefNfcSmartPosterRecord {
public:
};

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord(){
	return new QNdefNfcSmartPosterRecord();
}

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(void* other){
	return new QNdefNfcSmartPosterRecord(*static_cast<QNdefNfcSmartPosterRecord*>(other));
}

void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(void* other){
	return new QNdefNfcSmartPosterRecord(*static_cast<QNdefRecord*>(other));
}

int QNdefNfcSmartPosterRecord_Action(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->action();
}

void QNdefNfcSmartPosterRecord_AddIcon2(void* ptr, void* ty, void* data){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addIcon(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(data));
}

int QNdefNfcSmartPosterRecord_AddTitle(void* ptr, void* text){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addTitle(*static_cast<QNdefNfcTextRecord*>(text));
}

int QNdefNfcSmartPosterRecord_AddTitle2(void* ptr, char* text, char* locale, int encoding){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->addTitle(QString(text), QString(locale), static_cast<QNdefNfcTextRecord::Encoding>(encoding));
}

int QNdefNfcSmartPosterRecord_HasAction(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasAction();
}

int QNdefNfcSmartPosterRecord_HasIcon(void* ptr, void* mimetype){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasIcon(*static_cast<QByteArray*>(mimetype));
}

int QNdefNfcSmartPosterRecord_HasSize(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasSize();
}

int QNdefNfcSmartPosterRecord_HasTitle(void* ptr, char* locale){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasTitle(QString(locale));
}

int QNdefNfcSmartPosterRecord_HasTypeInfo(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->hasTypeInfo();
}

void* QNdefNfcSmartPosterRecord_Icon(void* ptr, void* mimetype){
	return new QByteArray(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->icon(*static_cast<QByteArray*>(mimetype)));
}

int QNdefNfcSmartPosterRecord_IconCount(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->iconCount();
}

int QNdefNfcSmartPosterRecord_RemoveIcon2(void* ptr, void* ty){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeIcon(*static_cast<QByteArray*>(ty));
}

int QNdefNfcSmartPosterRecord_RemoveTitle(void* ptr, void* text){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeTitle(*static_cast<QNdefNfcTextRecord*>(text));
}

int QNdefNfcSmartPosterRecord_RemoveTitle2(void* ptr, char* locale){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->removeTitle(QString(locale));
}

void QNdefNfcSmartPosterRecord_SetAction(void* ptr, int act){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setAction(static_cast<QNdefNfcSmartPosterRecord::Action>(act));
}

void QNdefNfcSmartPosterRecord_SetTypeInfo(void* ptr, void* ty){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setTypeInfo(*static_cast<QByteArray*>(ty));
}

void QNdefNfcSmartPosterRecord_SetUri(void* ptr, void* url){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setUri(*static_cast<QNdefNfcUriRecord*>(url));
}

void QNdefNfcSmartPosterRecord_SetUri2(void* ptr, void* url){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->setUri(*static_cast<QUrl*>(url));
}

char* QNdefNfcSmartPosterRecord_Title(void* ptr, char* locale){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->title(QString(locale)).toUtf8().data();
}

int QNdefNfcSmartPosterRecord_TitleCount(void* ptr){
	return static_cast<QNdefNfcSmartPosterRecord*>(ptr)->titleCount();
}

void* QNdefNfcSmartPosterRecord_TypeInfo(void* ptr){
	return new QByteArray(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->typeInfo());
}

void QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(void* ptr){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->~QNdefNfcSmartPosterRecord();
}

#include "qnearfieldmanager.h"
#include <QVariant>
#include <QUrl>
#include <QObject>
#include <QString>
#include <QModelIndex>
#include <QNdefFilter>
#include <QNdefRecord>
#include <QByteArray>
#include <QNearFieldTarget>
#include <QNearFieldManager>
#include "_cgo_export.h"

class MyQNearFieldManager: public QNearFieldManager {
public:
void Signal_TargetDetected(QNearFieldTarget * target){callbackQNearFieldManagerTargetDetected(this->objectName().toUtf8().data(), target);};
void Signal_TargetLost(QNearFieldTarget * target){callbackQNearFieldManagerTargetLost(this->objectName().toUtf8().data(), target);};
};

int QNearFieldManager_RegisterNdefMessageHandler(void* ptr, void* object, char* method){
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(static_cast<QObject*>(object), const_cast<const char*>(method));
}

int QNearFieldManager_StartTargetDetection(void* ptr){
	return static_cast<QNearFieldManager*>(ptr)->startTargetDetection();
}

void* QNearFieldManager_NewQNearFieldManager(void* parent){
	return new QNearFieldManager(static_cast<QObject*>(parent));
}

int QNearFieldManager_IsAvailable(void* ptr){
	return static_cast<QNearFieldManager*>(ptr)->isAvailable();
}

int QNearFieldManager_RegisterNdefMessageHandler2(void* ptr, int typeNameFormat, void* ty, void* object, char* method){
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat), *static_cast<QByteArray*>(ty), static_cast<QObject*>(object), const_cast<const char*>(method));
}

int QNearFieldManager_RegisterNdefMessageHandler3(void* ptr, void* filter, void* object, char* method){
	return static_cast<QNearFieldManager*>(ptr)->registerNdefMessageHandler(*static_cast<QNdefFilter*>(filter), static_cast<QObject*>(object), const_cast<const char*>(method));
}

void QNearFieldManager_SetTargetAccessModes(void* ptr, int accessModes){
	static_cast<QNearFieldManager*>(ptr)->setTargetAccessModes(static_cast<QNearFieldManager::TargetAccessMode>(accessModes));
}

void QNearFieldManager_StopTargetDetection(void* ptr){
	static_cast<QNearFieldManager*>(ptr)->stopTargetDetection();
}

int QNearFieldManager_TargetAccessModes(void* ptr){
	return static_cast<QNearFieldManager*>(ptr)->targetAccessModes();
}

void QNearFieldManager_ConnectTargetDetected(void* ptr){
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetDetected), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetDetected));;
}

void QNearFieldManager_DisconnectTargetDetected(void* ptr){
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetDetected), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetDetected));;
}

void QNearFieldManager_ConnectTargetLost(void* ptr){
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));;
}

void QNearFieldManager_DisconnectTargetLost(void* ptr){
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));;
}

int QNearFieldManager_UnregisterNdefMessageHandler(void* ptr, int handlerId){
	return static_cast<QNearFieldManager*>(ptr)->unregisterNdefMessageHandler(handlerId);
}

void QNearFieldManager_DestroyQNearFieldManager(void* ptr){
	static_cast<QNearFieldManager*>(ptr)->~QNearFieldManager();
}

#include "qndeffilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNdefFilter>
#include "_cgo_export.h"

class MyQNdefFilter: public QNdefFilter {
public:
};

void* QNdefFilter_NewQNdefFilter(){
	return new QNdefFilter();
}

void* QNdefFilter_NewQNdefFilter2(void* other){
	return new QNdefFilter(*static_cast<QNdefFilter*>(other));
}

void QNdefFilter_Clear(void* ptr){
	static_cast<QNdefFilter*>(ptr)->clear();
}

int QNdefFilter_OrderMatch(void* ptr){
	return static_cast<QNdefFilter*>(ptr)->orderMatch();
}

int QNdefFilter_RecordCount(void* ptr){
	return static_cast<QNdefFilter*>(ptr)->recordCount();
}

void QNdefFilter_SetOrderMatch(void* ptr, int on){
	static_cast<QNdefFilter*>(ptr)->setOrderMatch(on != 0);
}

void QNdefFilter_DestroyQNdefFilter(void* ptr){
	static_cast<QNdefFilter*>(ptr)->~QNdefFilter();
}

#include "qndefrecord.h"
#include <QByteArray>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNdefRecord>
#include "_cgo_export.h"

class MyQNdefRecord: public QNdefRecord {
public:
};

void* QNdefRecord_NewQNdefRecord(){
	return new QNdefRecord();
}

void* QNdefRecord_NewQNdefRecord2(void* other){
	return new QNdefRecord(*static_cast<QNdefRecord*>(other));
}

void* QNdefRecord_Id(void* ptr){
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->id());
}

int QNdefRecord_IsEmpty(void* ptr){
	return static_cast<QNdefRecord*>(ptr)->isEmpty();
}

void* QNdefRecord_Payload(void* ptr){
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->payload());
}

void QNdefRecord_SetId(void* ptr, void* id){
	static_cast<QNdefRecord*>(ptr)->setId(*static_cast<QByteArray*>(id));
}

void QNdefRecord_SetPayload(void* ptr, void* payload){
	static_cast<QNdefRecord*>(ptr)->setPayload(*static_cast<QByteArray*>(payload));
}

void QNdefRecord_SetType(void* ptr, void* ty){
	static_cast<QNdefRecord*>(ptr)->setType(*static_cast<QByteArray*>(ty));
}

void QNdefRecord_SetTypeNameFormat(void* ptr, int typeNameFormat){
	static_cast<QNdefRecord*>(ptr)->setTypeNameFormat(static_cast<QNdefRecord::TypeNameFormat>(typeNameFormat));
}

void* QNdefRecord_Type(void* ptr){
	return new QByteArray(static_cast<QNdefRecord*>(ptr)->type());
}

int QNdefRecord_TypeNameFormat(void* ptr){
	return static_cast<QNdefRecord*>(ptr)->typeNameFormat();
}

void QNdefRecord_DestroyQNdefRecord(void* ptr){
	static_cast<QNdefRecord*>(ptr)->~QNdefRecord();
}

