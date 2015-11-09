#include "qstatusbar.h"
#include <QWidget>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStatusBar>
#include "_cgo_export.h"

class MyQStatusBar: public QStatusBar {
public:
void Signal_MessageChanged(const QString & message){callbackQStatusBarMessageChanged(this->objectName().toUtf8().data(), message.toUtf8().data());};
};

int QStatusBar_IsSizeGripEnabled(void* ptr){
	return static_cast<QStatusBar*>(ptr)->isSizeGripEnabled();
}

void QStatusBar_SetSizeGripEnabled(void* ptr, int v){
	static_cast<QStatusBar*>(ptr)->setSizeGripEnabled(v != 0);
}

void* QStatusBar_NewQStatusBar(void* parent){
	return new QStatusBar(static_cast<QWidget*>(parent));
}

void QStatusBar_AddPermanentWidget(void* ptr, void* widget, int stretch){
	static_cast<QStatusBar*>(ptr)->addPermanentWidget(static_cast<QWidget*>(widget), stretch);
}

void QStatusBar_AddWidget(void* ptr, void* widget, int stretch){
	static_cast<QStatusBar*>(ptr)->addWidget(static_cast<QWidget*>(widget), stretch);
}

void QStatusBar_ClearMessage(void* ptr){
	QMetaObject::invokeMethod(static_cast<QStatusBar*>(ptr), "clearMessage");
}

char* QStatusBar_CurrentMessage(void* ptr){
	return static_cast<QStatusBar*>(ptr)->currentMessage().toUtf8().data();
}

int QStatusBar_InsertPermanentWidget(void* ptr, int index, void* widget, int stretch){
	return static_cast<QStatusBar*>(ptr)->insertPermanentWidget(index, static_cast<QWidget*>(widget), stretch);
}

int QStatusBar_InsertWidget(void* ptr, int index, void* widget, int stretch){
	return static_cast<QStatusBar*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget), stretch);
}

void QStatusBar_ConnectMessageChanged(void* ptr){
	QObject::connect(static_cast<QStatusBar*>(ptr), static_cast<void (QStatusBar::*)(const QString &)>(&QStatusBar::messageChanged), static_cast<MyQStatusBar*>(ptr), static_cast<void (MyQStatusBar::*)(const QString &)>(&MyQStatusBar::Signal_MessageChanged));;
}

void QStatusBar_DisconnectMessageChanged(void* ptr){
	QObject::disconnect(static_cast<QStatusBar*>(ptr), static_cast<void (QStatusBar::*)(const QString &)>(&QStatusBar::messageChanged), static_cast<MyQStatusBar*>(ptr), static_cast<void (MyQStatusBar::*)(const QString &)>(&MyQStatusBar::Signal_MessageChanged));;
}

void QStatusBar_RemoveWidget(void* ptr, void* widget){
	static_cast<QStatusBar*>(ptr)->removeWidget(static_cast<QWidget*>(widget));
}

void QStatusBar_ShowMessage(void* ptr, char* message, int timeout){
	QMetaObject::invokeMethod(static_cast<QStatusBar*>(ptr), "showMessage", Q_ARG(QString, QString(message)), Q_ARG(int, timeout));
}

void QStatusBar_DestroyQStatusBar(void* ptr){
	static_cast<QStatusBar*>(ptr)->~QStatusBar();
}

