#include "qwebchannelabstracttransport.h"
#include <QUrl>
#include <QModelIndex>
#include <QJsonObject>
#include <QWebChannel>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QWebChannelAbstractTransport>
#include "_cgo_export.h"

class MyQWebChannelAbstractTransport: public QWebChannelAbstractTransport {
public:
};

void QWebChannelAbstractTransport_SendMessage(QtObjectPtr ptr, QtObjectPtr message){
	QMetaObject::invokeMethod(static_cast<QWebChannelAbstractTransport*>(ptr), "sendMessage", Q_ARG(QJsonObject, *static_cast<QJsonObject*>(message)));
}

void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(QtObjectPtr ptr){
	static_cast<QWebChannelAbstractTransport*>(ptr)->~QWebChannelAbstractTransport();
}

