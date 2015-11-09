#include "qobject.h"
#include <QMetaObject>
#include <QEvent>
#include <QThread>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include "_cgo_export.h"

class MyQObject: public QObject {
public:
void Signal_Destroyed(QObject * obj){callbackQObjectDestroyed(this->objectName().toUtf8().data(), obj);};
void Signal_ObjectNameChanged(const QString & objectName){callbackQObjectObjectNameChanged(this->objectName().toUtf8().data(), objectName.toUtf8().data());};
};

void QObject_InstallEventFilter(void* ptr, void* filterObj){
	static_cast<QObject*>(ptr)->installEventFilter(static_cast<QObject*>(filterObj));
}

char* QObject_ObjectName(void* ptr){
	return static_cast<QObject*>(ptr)->objectName().toUtf8().data();
}

void QObject_SetObjectName(void* ptr, char* name){
	static_cast<QObject*>(ptr)->setObjectName(QString(name));
}

void* QObject_NewQObject(void* parent){
	return new QObject(static_cast<QObject*>(parent));
}

int QObject_BlockSignals(void* ptr, int block){
	return static_cast<QObject*>(ptr)->blockSignals(block != 0);
}

void QObject_DeleteLater(void* ptr){
	QMetaObject::invokeMethod(static_cast<QObject*>(ptr), "deleteLater");
}

void QObject_ConnectDestroyed(void* ptr){
	QObject::connect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed));;
}

void QObject_DisconnectDestroyed(void* ptr){
	QObject::disconnect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed));;
}

void QObject_DumpObjectInfo(void* ptr){
	static_cast<QObject*>(ptr)->dumpObjectInfo();
}

void QObject_DumpObjectTree(void* ptr){
	static_cast<QObject*>(ptr)->dumpObjectTree();
}

int QObject_Event(void* ptr, void* e){
	return static_cast<QObject*>(ptr)->event(static_cast<QEvent*>(e));
}

int QObject_EventFilter(void* ptr, void* watched, void* event){
	return static_cast<QObject*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QObject_FindChild(void* ptr, char* name, int options){
	return static_cast<QObject*>(ptr)->findChild<QObject*>(QString(name), static_cast<Qt::FindChildOption>(options));
}

int QObject_Inherits(void* ptr, char* className){
	return static_cast<QObject*>(ptr)->inherits(const_cast<const char*>(className));
}

int QObject_IsWidgetType(void* ptr){
	return static_cast<QObject*>(ptr)->isWidgetType();
}

int QObject_IsWindowType(void* ptr){
	return static_cast<QObject*>(ptr)->isWindowType();
}

void QObject_KillTimer(void* ptr, int id){
	static_cast<QObject*>(ptr)->killTimer(id);
}

void* QObject_MetaObject(void* ptr){
	return const_cast<QMetaObject*>(static_cast<QObject*>(ptr)->metaObject());
}

void QObject_MoveToThread(void* ptr, void* targetThread){
	static_cast<QObject*>(ptr)->moveToThread(static_cast<QThread*>(targetThread));
}

void QObject_ConnectObjectNameChanged(void* ptr){
	QObject::connect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged));;
}

void QObject_DisconnectObjectNameChanged(void* ptr){
	QObject::disconnect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged));;
}

void* QObject_Parent(void* ptr){
	return static_cast<QObject*>(ptr)->parent();
}

void* QObject_Property(void* ptr, char* name){
	return new QVariant(static_cast<QObject*>(ptr)->property(const_cast<const char*>(name)));
}

void QObject_RemoveEventFilter(void* ptr, void* obj){
	static_cast<QObject*>(ptr)->removeEventFilter(static_cast<QObject*>(obj));
}

void QObject_SetParent(void* ptr, void* parent){
	static_cast<QObject*>(ptr)->setParent(static_cast<QObject*>(parent));
}

int QObject_SetProperty(void* ptr, char* name, void* value){
	return static_cast<QObject*>(ptr)->setProperty(const_cast<const char*>(name), *static_cast<QVariant*>(value));
}

int QObject_StartTimer(void* ptr, int interval, int timerType){
	return static_cast<QObject*>(ptr)->startTimer(interval, static_cast<Qt::TimerType>(timerType));
}

int QObject_SignalsBlocked(void* ptr){
	return static_cast<QObject*>(ptr)->signalsBlocked();
}

void* QObject_Thread(void* ptr){
	return static_cast<QObject*>(ptr)->thread();
}

char* QObject_QObject_Tr(char* sourceText, char* disambiguation, int n){
	return QObject::tr(const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8().data();
}

void QObject_DestroyQObject(void* ptr){
	static_cast<QObject*>(ptr)->~QObject();
}

