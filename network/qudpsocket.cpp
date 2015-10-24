#include "qudpsocket.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHostAddress>
#include <QNetworkInterface>
#include <QObject>
#include <QString>
#include <QUdpSocket>
#include "_cgo_export.h"

class MyQUdpSocket: public QUdpSocket {
public:
};

QtObjectPtr QUdpSocket_NewQUdpSocket(QtObjectPtr parent){
	return new QUdpSocket(static_cast<QObject*>(parent));
}

int QUdpSocket_HasPendingDatagrams(QtObjectPtr ptr){
	return static_cast<QUdpSocket*>(ptr)->hasPendingDatagrams();
}

int QUdpSocket_JoinMulticastGroup(QtObjectPtr ptr, QtObjectPtr groupAddress){
	return static_cast<QUdpSocket*>(ptr)->joinMulticastGroup(*static_cast<QHostAddress*>(groupAddress));
}

int QUdpSocket_JoinMulticastGroup2(QtObjectPtr ptr, QtObjectPtr groupAddress, QtObjectPtr iface){
	return static_cast<QUdpSocket*>(ptr)->joinMulticastGroup(*static_cast<QHostAddress*>(groupAddress), *static_cast<QNetworkInterface*>(iface));
}

int QUdpSocket_LeaveMulticastGroup(QtObjectPtr ptr, QtObjectPtr groupAddress){
	return static_cast<QUdpSocket*>(ptr)->leaveMulticastGroup(*static_cast<QHostAddress*>(groupAddress));
}

int QUdpSocket_LeaveMulticastGroup2(QtObjectPtr ptr, QtObjectPtr groupAddress, QtObjectPtr iface){
	return static_cast<QUdpSocket*>(ptr)->leaveMulticastGroup(*static_cast<QHostAddress*>(groupAddress), *static_cast<QNetworkInterface*>(iface));
}

void QUdpSocket_SetMulticastInterface(QtObjectPtr ptr, QtObjectPtr iface){
	static_cast<QUdpSocket*>(ptr)->setMulticastInterface(*static_cast<QNetworkInterface*>(iface));
}

void QUdpSocket_DestroyQUdpSocket(QtObjectPtr ptr){
	static_cast<QUdpSocket*>(ptr)->~QUdpSocket();
}

