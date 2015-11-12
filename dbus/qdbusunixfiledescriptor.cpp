#include "qdbusunixfiledescriptor.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusUnixFileDescriptor>
#include "_cgo_export.h"

class MyQDBusUnixFileDescriptor: public QDBusUnixFileDescriptor {
public:
};

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor(){
	return new QDBusUnixFileDescriptor();
}

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(void* other){
	return new QDBusUnixFileDescriptor(*static_cast<QDBusUnixFileDescriptor*>(other));
}

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(int fileDescriptor){
	return new QDBusUnixFileDescriptor(fileDescriptor);
}

int QDBusUnixFileDescriptor_FileDescriptor(void* ptr){
	return static_cast<QDBusUnixFileDescriptor*>(ptr)->fileDescriptor();
}

int QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported(){
	return QDBusUnixFileDescriptor::isSupported();
}

int QDBusUnixFileDescriptor_IsValid(void* ptr){
	return static_cast<QDBusUnixFileDescriptor*>(ptr)->isValid();
}

void QDBusUnixFileDescriptor_SetFileDescriptor(void* ptr, int fileDescriptor){
	static_cast<QDBusUnixFileDescriptor*>(ptr)->setFileDescriptor(fileDescriptor);
}

void QDBusUnixFileDescriptor_Swap(void* ptr, void* other){
	static_cast<QDBusUnixFileDescriptor*>(ptr)->swap(*static_cast<QDBusUnixFileDescriptor*>(other));
}

void QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(void* ptr){
	static_cast<QDBusUnixFileDescriptor*>(ptr)->~QDBusUnixFileDescriptor();
}

