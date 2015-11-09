#include "qiodevice.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QObject>
#include <QIODevice>
#include "_cgo_export.h"

class MyQIODevice: public QIODevice {
public:
void Signal_AboutToClose(){callbackQIODeviceAboutToClose(this->objectName().toUtf8().data());};
void Signal_ReadChannelFinished(){callbackQIODeviceReadChannelFinished(this->objectName().toUtf8().data());};
void Signal_ReadyRead(){callbackQIODeviceReadyRead(this->objectName().toUtf8().data());};
};

int QIODevice_GetChar(void* ptr, char* c){
	return static_cast<QIODevice*>(ptr)->getChar(c);
}

int QIODevice_PutChar(void* ptr, char* c){
	return static_cast<QIODevice*>(ptr)->putChar(*c);
}

void QIODevice_ConnectAboutToClose(void* ptr){
	QObject::connect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::aboutToClose), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_AboutToClose));;
}

void QIODevice_DisconnectAboutToClose(void* ptr){
	QObject::disconnect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::aboutToClose), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_AboutToClose));;
}

int QIODevice_AtEnd(void* ptr){
	return static_cast<QIODevice*>(ptr)->atEnd();
}

int QIODevice_CanReadLine(void* ptr){
	return static_cast<QIODevice*>(ptr)->canReadLine();
}

void QIODevice_Close(void* ptr){
	static_cast<QIODevice*>(ptr)->close();
}

char* QIODevice_ErrorString(void* ptr){
	return static_cast<QIODevice*>(ptr)->errorString().toUtf8().data();
}

int QIODevice_IsOpen(void* ptr){
	return static_cast<QIODevice*>(ptr)->isOpen();
}

int QIODevice_IsReadable(void* ptr){
	return static_cast<QIODevice*>(ptr)->isReadable();
}

int QIODevice_IsSequential(void* ptr){
	return static_cast<QIODevice*>(ptr)->isSequential();
}

int QIODevice_IsTextModeEnabled(void* ptr){
	return static_cast<QIODevice*>(ptr)->isTextModeEnabled();
}

int QIODevice_IsWritable(void* ptr){
	return static_cast<QIODevice*>(ptr)->isWritable();
}

int QIODevice_Open(void* ptr, int mode){
	return static_cast<QIODevice*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

int QIODevice_OpenMode(void* ptr){
	return static_cast<QIODevice*>(ptr)->openMode();
}

void* QIODevice_ReadAll(void* ptr){
	return new QByteArray(static_cast<QIODevice*>(ptr)->readAll());
}

void QIODevice_ConnectReadChannelFinished(void* ptr){
	QObject::connect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::readChannelFinished), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_ReadChannelFinished));;
}

void QIODevice_DisconnectReadChannelFinished(void* ptr){
	QObject::disconnect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::readChannelFinished), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_ReadChannelFinished));;
}

void QIODevice_ConnectReadyRead(void* ptr){
	QObject::connect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::readyRead), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_ReadyRead));;
}

void QIODevice_DisconnectReadyRead(void* ptr){
	QObject::disconnect(static_cast<QIODevice*>(ptr), static_cast<void (QIODevice::*)()>(&QIODevice::readyRead), static_cast<MyQIODevice*>(ptr), static_cast<void (MyQIODevice::*)()>(&MyQIODevice::Signal_ReadyRead));;
}

int QIODevice_Reset(void* ptr){
	return static_cast<QIODevice*>(ptr)->reset();
}

void QIODevice_SetTextModeEnabled(void* ptr, int enabled){
	static_cast<QIODevice*>(ptr)->setTextModeEnabled(enabled != 0);
}

void QIODevice_UngetChar(void* ptr, char* c){
	static_cast<QIODevice*>(ptr)->ungetChar(*c);
}

int QIODevice_WaitForBytesWritten(void* ptr, int msecs){
	return static_cast<QIODevice*>(ptr)->waitForBytesWritten(msecs);
}

int QIODevice_WaitForReadyRead(void* ptr, int msecs){
	return static_cast<QIODevice*>(ptr)->waitForReadyRead(msecs);
}

void QIODevice_DestroyQIODevice(void* ptr){
	static_cast<QIODevice*>(ptr)->~QIODevice();
}

