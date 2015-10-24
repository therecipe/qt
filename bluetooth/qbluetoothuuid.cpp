#include "qbluetoothuuid.h"
#include <QModelIndex>
#include <QUuid>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QBluetoothUuid>
#include "_cgo_export.h"

class MyQBluetoothUuid: public QBluetoothUuid {
public:
};

QtObjectPtr QBluetoothUuid_NewQBluetoothUuid(){
	return new QBluetoothUuid();
}

QtObjectPtr QBluetoothUuid_NewQBluetoothUuid4(int uuid){
	return new QBluetoothUuid(static_cast<QBluetoothUuid::CharacteristicType>(uuid));
}

QtObjectPtr QBluetoothUuid_NewQBluetoothUuid5(int uuid){
	return new QBluetoothUuid(static_cast<QBluetoothUuid::DescriptorType>(uuid));
}

QtObjectPtr QBluetoothUuid_NewQBluetoothUuid2(int uuid){
	return new QBluetoothUuid(static_cast<QBluetoothUuid::ProtocolUuid>(uuid));
}

QtObjectPtr QBluetoothUuid_NewQBluetoothUuid3(int uuid){
	return new QBluetoothUuid(static_cast<QBluetoothUuid::ServiceClassUuid>(uuid));
}

QtObjectPtr QBluetoothUuid_NewQBluetoothUuid10(QtObjectPtr uuid){
	return new QBluetoothUuid(*static_cast<QBluetoothUuid*>(uuid));
}

QtObjectPtr QBluetoothUuid_NewQBluetoothUuid9(char* uuid){
	return new QBluetoothUuid(QString(uuid));
}

QtObjectPtr QBluetoothUuid_NewQBluetoothUuid11(QtObjectPtr uuid){
	return new QBluetoothUuid(*static_cast<QUuid*>(uuid));
}

char* QBluetoothUuid_QBluetoothUuid_CharacteristicToString(int uuid){
	return QBluetoothUuid::characteristicToString(static_cast<QBluetoothUuid::CharacteristicType>(uuid)).toUtf8().data();
}

char* QBluetoothUuid_QBluetoothUuid_DescriptorToString(int uuid){
	return QBluetoothUuid::descriptorToString(static_cast<QBluetoothUuid::DescriptorType>(uuid)).toUtf8().data();
}

int QBluetoothUuid_MinimumSize(QtObjectPtr ptr){
	return static_cast<QBluetoothUuid*>(ptr)->minimumSize();
}

char* QBluetoothUuid_QBluetoothUuid_ProtocolToString(int uuid){
	return QBluetoothUuid::protocolToString(static_cast<QBluetoothUuid::ProtocolUuid>(uuid)).toUtf8().data();
}

char* QBluetoothUuid_QBluetoothUuid_ServiceClassToString(int uuid){
	return QBluetoothUuid::serviceClassToString(static_cast<QBluetoothUuid::ServiceClassUuid>(uuid)).toUtf8().data();
}

void QBluetoothUuid_DestroyQBluetoothUuid(QtObjectPtr ptr){
	static_cast<QBluetoothUuid*>(ptr)->~QBluetoothUuid();
}

