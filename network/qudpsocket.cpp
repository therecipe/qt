#include "qudpsocket.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkInterface>
#include <QObject>
#include <QHostAddress>
#include <QString>
#include <QUdpSocket>
#include "_cgo_export.h"

class MyQUdpSocket: public QUdpSocket {
public:
};

void* QUdpSocket_NewQUdpSocket(void* parent){
	return new QUdpSocket(static_cast<QObject*>(parent));
}

int QUdpSocket_HasPendingDatagrams(void* ptr){
	return static_cast<QUdpSocket*>(ptr)->hasPendingDatagrams();
}

int QUdpSocket_JoinMulticastGroup(void* ptr, void* groupAddress){
	return static_cast<QUdpSocket*>(ptr)->joinMulticastGroup(*static_cast<QHostAddress*>(groupAddress));
}

int QUdpSocket_JoinMulticastGroup2(void* ptr, void* groupAddress, void* iface){
	return static_cast<QUdpSocket*>(ptr)->joinMulticastGroup(*static_cast<QHostAddress*>(groupAddress), *static_cast<QNetworkInterface*>(iface));
}

int QUdpSocket_LeaveMulticastGroup(void* ptr, void* groupAddress){
	return static_cast<QUdpSocket*>(ptr)->leaveMulticastGroup(*static_cast<QHostAddress*>(groupAddress));
}

int QUdpSocket_LeaveMulticastGroup2(void* ptr, void* groupAddress, void* iface){
	return static_cast<QUdpSocket*>(ptr)->leaveMulticastGroup(*static_cast<QHostAddress*>(groupAddress), *static_cast<QNetworkInterface*>(iface));
}

void QUdpSocket_SetMulticastInterface(void* ptr, void* iface){
	static_cast<QUdpSocket*>(ptr)->setMulticastInterface(*static_cast<QNetworkInterface*>(iface));
}

void QUdpSocket_DestroyQUdpSocket(void* ptr){
	static_cast<QUdpSocket*>(ptr)->~QUdpSocket();
}

