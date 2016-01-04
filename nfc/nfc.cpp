#define protected public

#include "nfc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QNdefFilter>
#include <QNdefMessage>
#include <QNdefNfcSmartPosterRecord>
#include <QNdefNfcTextRecord>
#include <QNdefNfcUriRecord>
#include <QNdefRecord>
#include <QNearFieldManager>
#include <QNearFieldShareManager>
#include <QNearFieldShareTarget>
#include <QNearFieldTarget>
#include <QObject>
#include <QQmlNdefRecord>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>

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

void* QNdefNfcSmartPosterRecord_Uri(void* ptr){
	return new QUrl(static_cast<QNdefNfcSmartPosterRecord*>(ptr)->uri());
}

void QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(void* ptr){
	static_cast<QNdefNfcSmartPosterRecord*>(ptr)->~QNdefNfcSmartPosterRecord();
}

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

void* QNdefNfcUriRecord_NewQNdefNfcUriRecord(){
	return new QNdefNfcUriRecord();
}

void* QNdefNfcUriRecord_NewQNdefNfcUriRecord2(void* other){
	return new QNdefNfcUriRecord(*static_cast<QNdefRecord*>(other));
}

void QNdefNfcUriRecord_SetUri(void* ptr, void* uri){
	static_cast<QNdefNfcUriRecord*>(ptr)->setUri(*static_cast<QUrl*>(uri));
}

void* QNdefNfcUriRecord_Uri(void* ptr){
	return new QUrl(static_cast<QNdefNfcUriRecord*>(ptr)->uri());
}

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

class MyQNearFieldManager: public QNearFieldManager {
public:
	void Signal_TargetDetected(QNearFieldTarget * target) { callbackQNearFieldManagerTargetDetected(this, this->objectName().toUtf8().data(), target); };
	void Signal_TargetLost(QNearFieldTarget * target) { callbackQNearFieldManagerTargetLost(this, this->objectName().toUtf8().data(), target); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldManagerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNearFieldManagerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNearFieldManagerCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void QNearFieldManager_TargetDetected(void* ptr, void* target){
	static_cast<QNearFieldManager*>(ptr)->targetDetected(static_cast<QNearFieldTarget*>(target));
}

void QNearFieldManager_ConnectTargetLost(void* ptr){
	QObject::connect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));;
}

void QNearFieldManager_DisconnectTargetLost(void* ptr){
	QObject::disconnect(static_cast<QNearFieldManager*>(ptr), static_cast<void (QNearFieldManager::*)(QNearFieldTarget *)>(&QNearFieldManager::targetLost), static_cast<MyQNearFieldManager*>(ptr), static_cast<void (MyQNearFieldManager::*)(QNearFieldTarget *)>(&MyQNearFieldManager::Signal_TargetLost));;
}

void QNearFieldManager_TargetLost(void* ptr, void* target){
	static_cast<QNearFieldManager*>(ptr)->targetLost(static_cast<QNearFieldTarget*>(target));
}

int QNearFieldManager_UnregisterNdefMessageHandler(void* ptr, int handlerId){
	return static_cast<QNearFieldManager*>(ptr)->unregisterNdefMessageHandler(handlerId);
}

void QNearFieldManager_DestroyQNearFieldManager(void* ptr){
	static_cast<QNearFieldManager*>(ptr)->~QNearFieldManager();
}

void QNearFieldManager_TimerEvent(void* ptr, void* event){
	static_cast<MyQNearFieldManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldManager_TimerEventDefault(void* ptr, void* event){
	static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldManager_ChildEvent(void* ptr, void* event){
	static_cast<MyQNearFieldManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldManager_ChildEventDefault(void* ptr, void* event){
	static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldManager_CustomEvent(void* ptr, void* event){
	static_cast<MyQNearFieldManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNearFieldManager_CustomEventDefault(void* ptr, void* event){
	static_cast<QNearFieldManager*>(ptr)->QNearFieldManager::customEvent(static_cast<QEvent*>(event));
}

class MyQNearFieldShareManager: public QNearFieldShareManager {
public:
	void Signal_Error(QNearFieldShareManager::ShareError error) { callbackQNearFieldShareManagerError(this, this->objectName().toUtf8().data(), error); };
	void Signal_ShareModesChanged(QNearFieldShareManager::ShareModes modes) { callbackQNearFieldShareManagerShareModesChanged(this, this->objectName().toUtf8().data(), modes); };
	void Signal_TargetDetected(QNearFieldShareTarget * shareTarget) { callbackQNearFieldShareManagerTargetDetected(this, this->objectName().toUtf8().data(), shareTarget); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldShareManagerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNearFieldShareManagerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNearFieldShareManagerCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void QNearFieldShareManager_Error(void* ptr, int error){
	static_cast<QNearFieldShareManager*>(ptr)->error(static_cast<QNearFieldShareManager::ShareError>(error));
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

void QNearFieldShareManager_ShareModesChanged(void* ptr, int modes){
	static_cast<QNearFieldShareManager*>(ptr)->shareModesChanged(static_cast<QNearFieldShareManager::ShareMode>(modes));
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

void QNearFieldShareManager_TargetDetected(void* ptr, void* shareTarget){
	static_cast<QNearFieldShareManager*>(ptr)->targetDetected(static_cast<QNearFieldShareTarget*>(shareTarget));
}

void QNearFieldShareManager_DestroyQNearFieldShareManager(void* ptr){
	static_cast<QNearFieldShareManager*>(ptr)->~QNearFieldShareManager();
}

void QNearFieldShareManager_TimerEvent(void* ptr, void* event){
	static_cast<MyQNearFieldShareManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldShareManager_TimerEventDefault(void* ptr, void* event){
	static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldShareManager_ChildEvent(void* ptr, void* event){
	static_cast<MyQNearFieldShareManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldShareManager_ChildEventDefault(void* ptr, void* event){
	static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldShareManager_CustomEvent(void* ptr, void* event){
	static_cast<MyQNearFieldShareManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNearFieldShareManager_CustomEventDefault(void* ptr, void* event){
	static_cast<QNearFieldShareManager*>(ptr)->QNearFieldShareManager::customEvent(static_cast<QEvent*>(event));
}

class MyQNearFieldShareTarget: public QNearFieldShareTarget {
public:
	void Signal_Error(QNearFieldShareManager::ShareError error) { callbackQNearFieldShareTargetError(this, this->objectName().toUtf8().data(), error); };
	void Signal_ShareFinished() { callbackQNearFieldShareTargetShareFinished(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldShareTargetTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNearFieldShareTargetChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNearFieldShareTargetCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void QNearFieldShareTarget_Error(void* ptr, int error){
	static_cast<QNearFieldShareTarget*>(ptr)->error(static_cast<QNearFieldShareManager::ShareError>(error));
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

void QNearFieldShareTarget_ShareFinished(void* ptr){
	static_cast<QNearFieldShareTarget*>(ptr)->shareFinished();
}

int QNearFieldShareTarget_ShareModes(void* ptr){
	return static_cast<QNearFieldShareTarget*>(ptr)->shareModes();
}

void QNearFieldShareTarget_DestroyQNearFieldShareTarget(void* ptr){
	static_cast<QNearFieldShareTarget*>(ptr)->~QNearFieldShareTarget();
}

void QNearFieldShareTarget_TimerEvent(void* ptr, void* event){
	static_cast<MyQNearFieldShareTarget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldShareTarget_TimerEventDefault(void* ptr, void* event){
	static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldShareTarget_ChildEvent(void* ptr, void* event){
	static_cast<MyQNearFieldShareTarget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldShareTarget_ChildEventDefault(void* ptr, void* event){
	static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldShareTarget_CustomEvent(void* ptr, void* event){
	static_cast<MyQNearFieldShareTarget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNearFieldShareTarget_CustomEventDefault(void* ptr, void* event){
	static_cast<QNearFieldShareTarget*>(ptr)->QNearFieldShareTarget::customEvent(static_cast<QEvent*>(event));
}

class MyQNearFieldTarget: public QNearFieldTarget {
public:
	void Signal_Disconnected() { callbackQNearFieldTargetDisconnected(this, this->objectName().toUtf8().data()); };
	void Signal_NdefMessagesWritten() { callbackQNearFieldTargetNdefMessagesWritten(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQNearFieldTargetTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNearFieldTargetChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNearFieldTargetCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void QNearFieldTarget_Disconnected(void* ptr){
	static_cast<QNearFieldTarget*>(ptr)->disconnected();
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

void QNearFieldTarget_NdefMessagesWritten(void* ptr){
	static_cast<QNearFieldTarget*>(ptr)->ndefMessagesWritten();
}

int QNearFieldTarget_Type(void* ptr){
	return static_cast<QNearFieldTarget*>(ptr)->type();
}

void* QNearFieldTarget_Uid(void* ptr){
	return new QByteArray(static_cast<QNearFieldTarget*>(ptr)->uid());
}

void* QNearFieldTarget_Url(void* ptr){
	return new QUrl(static_cast<QNearFieldTarget*>(ptr)->url());
}

void QNearFieldTarget_DestroyQNearFieldTarget(void* ptr){
	static_cast<QNearFieldTarget*>(ptr)->~QNearFieldTarget();
}

void QNearFieldTarget_TimerEvent(void* ptr, void* event){
	static_cast<MyQNearFieldTarget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldTarget_TimerEventDefault(void* ptr, void* event){
	static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNearFieldTarget_ChildEvent(void* ptr, void* event){
	static_cast<MyQNearFieldTarget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldTarget_ChildEventDefault(void* ptr, void* event){
	static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::childEvent(static_cast<QChildEvent*>(event));
}

void QNearFieldTarget_CustomEvent(void* ptr, void* event){
	static_cast<MyQNearFieldTarget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNearFieldTarget_CustomEventDefault(void* ptr, void* event){
	static_cast<QNearFieldTarget*>(ptr)->QNearFieldTarget::customEvent(static_cast<QEvent*>(event));
}

class MyQQmlNdefRecord: public QQmlNdefRecord {
public:
	void Signal_RecordChanged() { callbackQQmlNdefRecordRecordChanged(this, this->objectName().toUtf8().data()); };
	void Signal_TypeChanged() { callbackQQmlNdefRecordTypeChanged(this, this->objectName().toUtf8().data()); };
	void Signal_TypeNameFormatChanged() { callbackQQmlNdefRecordTypeNameFormatChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQQmlNdefRecordTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQQmlNdefRecordChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQQmlNdefRecordCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void QQmlNdefRecord_RecordChanged(void* ptr){
	static_cast<QQmlNdefRecord*>(ptr)->recordChanged();
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

void QQmlNdefRecord_TypeChanged(void* ptr){
	static_cast<QQmlNdefRecord*>(ptr)->typeChanged();
}

void QQmlNdefRecord_ConnectTypeNameFormatChanged(void* ptr){
	QObject::connect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));;
}

void QQmlNdefRecord_DisconnectTypeNameFormatChanged(void* ptr){
	QObject::disconnect(static_cast<QQmlNdefRecord*>(ptr), static_cast<void (QQmlNdefRecord::*)()>(&QQmlNdefRecord::typeNameFormatChanged), static_cast<MyQQmlNdefRecord*>(ptr), static_cast<void (MyQQmlNdefRecord::*)()>(&MyQQmlNdefRecord::Signal_TypeNameFormatChanged));;
}

void QQmlNdefRecord_TypeNameFormatChanged(void* ptr){
	static_cast<QQmlNdefRecord*>(ptr)->typeNameFormatChanged();
}

void QQmlNdefRecord_TimerEvent(void* ptr, void* event){
	static_cast<MyQQmlNdefRecord*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlNdefRecord_TimerEventDefault(void* ptr, void* event){
	static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::timerEvent(static_cast<QTimerEvent*>(event));
}

void QQmlNdefRecord_ChildEvent(void* ptr, void* event){
	static_cast<MyQQmlNdefRecord*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QQmlNdefRecord_ChildEventDefault(void* ptr, void* event){
	static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::childEvent(static_cast<QChildEvent*>(event));
}

void QQmlNdefRecord_CustomEvent(void* ptr, void* event){
	static_cast<MyQQmlNdefRecord*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QQmlNdefRecord_CustomEventDefault(void* ptr, void* event){
	static_cast<QQmlNdefRecord*>(ptr)->QQmlNdefRecord::customEvent(static_cast<QEvent*>(event));
}

