#include "qerrormessage.h"
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QErrorMessage>
#include "_cgo_export.h"

class MyQErrorMessage: public QErrorMessage {
public:
};

QtObjectPtr QErrorMessage_NewQErrorMessage(QtObjectPtr parent){
	return new QErrorMessage(static_cast<QWidget*>(parent));
}

QtObjectPtr QErrorMessage_QErrorMessage_QtHandler(){
	return QErrorMessage::qtHandler();
}

void QErrorMessage_ShowMessage(QtObjectPtr ptr, char* message){
	QMetaObject::invokeMethod(static_cast<QErrorMessage*>(ptr), "showMessage", Q_ARG(QString, QString(message)));
}

void QErrorMessage_ShowMessage2(QtObjectPtr ptr, char* message, char* ty){
	QMetaObject::invokeMethod(static_cast<QErrorMessage*>(ptr), "showMessage", Q_ARG(QString, QString(message)), Q_ARG(QString, QString(ty)));
}

void QErrorMessage_DestroyQErrorMessage(QtObjectPtr ptr){
	static_cast<QErrorMessage*>(ptr)->~QErrorMessage();
}

