#include "qerrormessage.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QErrorMessage>
#include "_cgo_export.h"

class MyQErrorMessage: public QErrorMessage {
public:
};

void* QErrorMessage_NewQErrorMessage(void* parent){
	return new QErrorMessage(static_cast<QWidget*>(parent));
}

void* QErrorMessage_QErrorMessage_QtHandler(){
	return QErrorMessage::qtHandler();
}

void QErrorMessage_ShowMessage(void* ptr, char* message){
	QMetaObject::invokeMethod(static_cast<QErrorMessage*>(ptr), "showMessage", Q_ARG(QString, QString(message)));
}

void QErrorMessage_ShowMessage2(void* ptr, char* message, char* ty){
	QMetaObject::invokeMethod(static_cast<QErrorMessage*>(ptr), "showMessage", Q_ARG(QString, QString(message)), Q_ARG(QString, QString(ty)));
}

void QErrorMessage_DestroyQErrorMessage(void* ptr){
	static_cast<QErrorMessage*>(ptr)->~QErrorMessage();
}

