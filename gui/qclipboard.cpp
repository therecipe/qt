#include "qclipboard.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QImage>
#include <QMimeData>
#include <QPixmap>
#include <QObject>
#include <QString>
#include <QClipboard>
#include "_cgo_export.h"

class MyQClipboard: public QClipboard {
public:
void Signal_Changed(QClipboard::Mode mode){callbackQClipboardChanged(this->objectName().toUtf8().data(), mode);};
void Signal_DataChanged(){callbackQClipboardDataChanged(this->objectName().toUtf8().data());};
void Signal_FindBufferChanged(){callbackQClipboardFindBufferChanged(this->objectName().toUtf8().data());};
void Signal_SelectionChanged(){callbackQClipboardSelectionChanged(this->objectName().toUtf8().data());};
};

void QClipboard_Clear(void* ptr, int mode){
	static_cast<QClipboard*>(ptr)->clear(static_cast<QClipboard::Mode>(mode));
}

void* QClipboard_MimeData(void* ptr, int mode){
	return const_cast<QMimeData*>(static_cast<QClipboard*>(ptr)->mimeData(static_cast<QClipboard::Mode>(mode)));
}

void QClipboard_SetMimeData(void* ptr, void* src, int mode){
	static_cast<QClipboard*>(ptr)->setMimeData(static_cast<QMimeData*>(src), static_cast<QClipboard::Mode>(mode));
}

void QClipboard_ConnectChanged(void* ptr){
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)(QClipboard::Mode)>(&QClipboard::changed), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)(QClipboard::Mode)>(&MyQClipboard::Signal_Changed));;
}

void QClipboard_DisconnectChanged(void* ptr){
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)(QClipboard::Mode)>(&QClipboard::changed), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)(QClipboard::Mode)>(&MyQClipboard::Signal_Changed));;
}

void QClipboard_ConnectDataChanged(void* ptr){
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::dataChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_DataChanged));;
}

void QClipboard_DisconnectDataChanged(void* ptr){
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::dataChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_DataChanged));;
}

void QClipboard_ConnectFindBufferChanged(void* ptr){
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::findBufferChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_FindBufferChanged));;
}

void QClipboard_DisconnectFindBufferChanged(void* ptr){
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::findBufferChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_FindBufferChanged));;
}

int QClipboard_OwnsClipboard(void* ptr){
	return static_cast<QClipboard*>(ptr)->ownsClipboard();
}

int QClipboard_OwnsFindBuffer(void* ptr){
	return static_cast<QClipboard*>(ptr)->ownsFindBuffer();
}

int QClipboard_OwnsSelection(void* ptr){
	return static_cast<QClipboard*>(ptr)->ownsSelection();
}

void QClipboard_ConnectSelectionChanged(void* ptr){
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::selectionChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_SelectionChanged));;
}

void QClipboard_DisconnectSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::selectionChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_SelectionChanged));;
}

void QClipboard_SetImage(void* ptr, void* image, int mode){
	static_cast<QClipboard*>(ptr)->setImage(*static_cast<QImage*>(image), static_cast<QClipboard::Mode>(mode));
}

void QClipboard_SetPixmap(void* ptr, void* pixmap, int mode){
	static_cast<QClipboard*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QClipboard::Mode>(mode));
}

void QClipboard_SetText(void* ptr, char* text, int mode){
	static_cast<QClipboard*>(ptr)->setText(QString(text), static_cast<QClipboard::Mode>(mode));
}

int QClipboard_SupportsFindBuffer(void* ptr){
	return static_cast<QClipboard*>(ptr)->supportsFindBuffer();
}

int QClipboard_SupportsSelection(void* ptr){
	return static_cast<QClipboard*>(ptr)->supportsSelection();
}

char* QClipboard_Text(void* ptr, int mode){
	return static_cast<QClipboard*>(ptr)->text(static_cast<QClipboard::Mode>(mode)).toUtf8().data();
}

