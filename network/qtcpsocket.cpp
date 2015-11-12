#include "qtcpsocket.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QTcpSocket>
#include "_cgo_export.h"

class MyQTcpSocket: public QTcpSocket {
public:
};

void* QTcpSocket_NewQTcpSocket(void* parent){
	return new QTcpSocket(static_cast<QObject*>(parent));
}

void QTcpSocket_DestroyQTcpSocket(void* ptr){
	static_cast<QTcpSocket*>(ptr)->~QTcpSocket();
}

