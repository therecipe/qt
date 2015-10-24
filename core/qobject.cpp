#include "qobject.h"
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QThread>
#include <QObject>
#include "_cgo_export.h"

class MyQObject: public QObject {
public:
void Signal_Destroyed(QObject * obj){callbackQObjectDestroyed(this->objectName().toUtf8().data(), obj);};
void Signal_ObjectNameChanged(const QString & objectName){callbackQObjectObjectNameChanged(this->objectName().toUtf8().data(), objectName.toUtf8().data());};
};

void QObject_InstallEventFilter(QtObjectPtr ptr, QtObjectPtr filterObj){
	static_cast<QObject*>(ptr)->installEventFilter(static_cast<QObject*>(filterObj));
}

char* QObject_ObjectName(QtObjectPtr ptr){
	return static_cast<QObject*>(ptr)->objectName().toUtf8().data();
}

void QObject_SetObjectName(QtObjectPtr ptr, char* name){
	static_cast<QObject*>(ptr)->setObjectName(QString(name));
}

QtObjectPtr QObject_NewQObject(QtObjectPtr parent){
	return new QObject(static_cast<QObject*>(parent));
}

int QObject_BlockSignals(QtObjectPtr ptr, int block){
	return static_cast<QObject*>(ptr)->blockSignals(block != 0);
}

void QObject_DeleteLater(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QObject*>(ptr), "deleteLater");
}

void QObject_ConnectDestroyed(QtObjectPtr ptr){
	QObject::connect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed));;
}

void QObject_DisconnectDestroyed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed));;
}

void QObject_DumpObjectInfo(QtObjectPtr ptr){
	static_cast<QObject*>(ptr)->dumpObjectInfo();
}

void QObject_DumpObjectTree(QtObjectPtr ptr){
	static_cast<QObject*>(ptr)->dumpObjectTree();
}

int QObject_Event(QtObjectPtr ptr, QtObjectPtr e){
	return static_cast<QObject*>(ptr)->event(static_cast<QEvent*>(e));
}

int QObject_EventFilter(QtObjectPtr ptr, QtObjectPtr watched, QtObjectPtr event){
	return static_cast<QObject*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QObject_Inherits(QtObjectPtr ptr, char* className){
	return static_cast<QObject*>(ptr)->inherits(const_cast<const char*>(className));
}

int QObject_IsWidgetType(QtObjectPtr ptr){
	return static_cast<QObject*>(ptr)->isWidgetType();
}

int QObject_IsWindowType(QtObjectPtr ptr){
	return static_cast<QObject*>(ptr)->isWindowType();
}

void QObject_KillTimer(QtObjectPtr ptr, int id){
	static_cast<QObject*>(ptr)->killTimer(id);
}

QtObjectPtr QObject_MetaObject(QtObjectPtr ptr){
	return const_cast<QMetaObject*>(static_cast<QObject*>(ptr)->metaObject());
}

void QObject_MoveToThread(QtObjectPtr ptr, QtObjectPtr targetThread){
	static_cast<QObject*>(ptr)->moveToThread(static_cast<QThread*>(targetThread));
}

void QObject_ConnectObjectNameChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged));;
}

void QObject_DisconnectObjectNameChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged));;
}

QtObjectPtr QObject_Parent(QtObjectPtr ptr){
	return static_cast<QObject*>(ptr)->parent();
}

char* QObject_Property(QtObjectPtr ptr, char* name){
	return static_cast<QObject*>(ptr)->property(const_cast<const char*>(name)).toString().toUtf8().data();
}

void QObject_RemoveEventFilter(QtObjectPtr ptr, QtObjectPtr obj){
	static_cast<QObject*>(ptr)->removeEventFilter(static_cast<QObject*>(obj));
}

void QObject_SetParent(QtObjectPtr ptr, QtObjectPtr parent){
	static_cast<QObject*>(ptr)->setParent(static_cast<QObject*>(parent));
}

int QObject_SetProperty(QtObjectPtr ptr, char* name, char* value){
	return static_cast<QObject*>(ptr)->setProperty(const_cast<const char*>(name), QVariant(value));
}

int QObject_StartTimer(QtObjectPtr ptr, int interval, int timerType){
	return static_cast<QObject*>(ptr)->startTimer(interval, static_cast<Qt::TimerType>(timerType));
}

int QObject_SignalsBlocked(QtObjectPtr ptr){
	return static_cast<QObject*>(ptr)->signalsBlocked();
}

QtObjectPtr QObject_Thread(QtObjectPtr ptr){
	return static_cast<QObject*>(ptr)->thread();
}

char* QObject_QObject_Tr(char* sourceText, char* disambiguation, int n){
	return QObject::tr(const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8().data();
}

void QObject_DestroyQObject(QtObjectPtr ptr){
	static_cast<QObject*>(ptr)->~QObject();
}

