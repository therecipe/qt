#include "qstatusbar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QMetaObject>
#include <QObject>
#include <QStatusBar>
#include "_cgo_export.h"

class MyQStatusBar: public QStatusBar {
public:
void Signal_MessageChanged(const QString & message){callbackQStatusBarMessageChanged(this->objectName().toUtf8().data(), message.toUtf8().data());};
};

int QStatusBar_IsSizeGripEnabled(QtObjectPtr ptr){
	return static_cast<QStatusBar*>(ptr)->isSizeGripEnabled();
}

void QStatusBar_SetSizeGripEnabled(QtObjectPtr ptr, int v){
	static_cast<QStatusBar*>(ptr)->setSizeGripEnabled(v != 0);
}

QtObjectPtr QStatusBar_NewQStatusBar(QtObjectPtr parent){
	return new QStatusBar(static_cast<QWidget*>(parent));
}

void QStatusBar_AddPermanentWidget(QtObjectPtr ptr, QtObjectPtr widget, int stretch){
	static_cast<QStatusBar*>(ptr)->addPermanentWidget(static_cast<QWidget*>(widget), stretch);
}

void QStatusBar_AddWidget(QtObjectPtr ptr, QtObjectPtr widget, int stretch){
	static_cast<QStatusBar*>(ptr)->addWidget(static_cast<QWidget*>(widget), stretch);
}

void QStatusBar_ClearMessage(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QStatusBar*>(ptr), "clearMessage");
}

char* QStatusBar_CurrentMessage(QtObjectPtr ptr){
	return static_cast<QStatusBar*>(ptr)->currentMessage().toUtf8().data();
}

int QStatusBar_InsertPermanentWidget(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch){
	return static_cast<QStatusBar*>(ptr)->insertPermanentWidget(index, static_cast<QWidget*>(widget), stretch);
}

int QStatusBar_InsertWidget(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch){
	return static_cast<QStatusBar*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget), stretch);
}

void QStatusBar_ConnectMessageChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QStatusBar*>(ptr), static_cast<void (QStatusBar::*)(const QString &)>(&QStatusBar::messageChanged), static_cast<MyQStatusBar*>(ptr), static_cast<void (MyQStatusBar::*)(const QString &)>(&MyQStatusBar::Signal_MessageChanged));;
}

void QStatusBar_DisconnectMessageChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStatusBar*>(ptr), static_cast<void (QStatusBar::*)(const QString &)>(&QStatusBar::messageChanged), static_cast<MyQStatusBar*>(ptr), static_cast<void (MyQStatusBar::*)(const QString &)>(&MyQStatusBar::Signal_MessageChanged));;
}

void QStatusBar_RemoveWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QStatusBar*>(ptr)->removeWidget(static_cast<QWidget*>(widget));
}

void QStatusBar_ShowMessage(QtObjectPtr ptr, char* message, int timeout){
	QMetaObject::invokeMethod(static_cast<QStatusBar*>(ptr), "showMessage", Q_ARG(QString, QString(message)), Q_ARG(int, timeout));
}

void QStatusBar_DestroyQStatusBar(QtObjectPtr ptr){
	static_cast<QStatusBar*>(ptr)->~QStatusBar();
}

