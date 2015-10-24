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

QtObjectPtr QTcpSocket_NewQTcpSocket(QtObjectPtr parent){
	return new QTcpSocket(static_cast<QObject*>(parent));
}

void QTcpSocket_DestroyQTcpSocket(QtObjectPtr ptr){
	static_cast<QTcpSocket*>(ptr)->~QTcpSocket();
}

