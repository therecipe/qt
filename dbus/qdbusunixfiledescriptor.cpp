#include "qdbusunixfiledescriptor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusUnixFileDescriptor>
#include "_cgo_export.h"

class MyQDBusUnixFileDescriptor: public QDBusUnixFileDescriptor {
public:
};

QtObjectPtr QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor(){
	return new QDBusUnixFileDescriptor();
}

QtObjectPtr QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(QtObjectPtr other){
	return new QDBusUnixFileDescriptor(*static_cast<QDBusUnixFileDescriptor*>(other));
}

QtObjectPtr QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(int fileDescriptor){
	return new QDBusUnixFileDescriptor(fileDescriptor);
}

int QDBusUnixFileDescriptor_FileDescriptor(QtObjectPtr ptr){
	return static_cast<QDBusUnixFileDescriptor*>(ptr)->fileDescriptor();
}

int QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported(){
	return QDBusUnixFileDescriptor::isSupported();
}

int QDBusUnixFileDescriptor_IsValid(QtObjectPtr ptr){
	return static_cast<QDBusUnixFileDescriptor*>(ptr)->isValid();
}

void QDBusUnixFileDescriptor_SetFileDescriptor(QtObjectPtr ptr, int fileDescriptor){
	static_cast<QDBusUnixFileDescriptor*>(ptr)->setFileDescriptor(fileDescriptor);
}

void QDBusUnixFileDescriptor_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QDBusUnixFileDescriptor*>(ptr)->swap(*static_cast<QDBusUnixFileDescriptor*>(other));
}

void QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(QtObjectPtr ptr){
	static_cast<QDBusUnixFileDescriptor*>(ptr)->~QDBusUnixFileDescriptor();
}

