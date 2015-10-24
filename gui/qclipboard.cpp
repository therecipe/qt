#include "qclipboard.h"
#include <QPixmap>
#include <QObject>
#include <QImage>
#include <QMimeData>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QClipboard>
#include "_cgo_export.h"

class MyQClipboard: public QClipboard {
public:
void Signal_Changed(QClipboard::Mode mode){callbackQClipboardChanged(this->objectName().toUtf8().data(), mode);};
void Signal_DataChanged(){callbackQClipboardDataChanged(this->objectName().toUtf8().data());};
void Signal_FindBufferChanged(){callbackQClipboardFindBufferChanged(this->objectName().toUtf8().data());};
void Signal_SelectionChanged(){callbackQClipboardSelectionChanged(this->objectName().toUtf8().data());};
};

void QClipboard_Clear(QtObjectPtr ptr, int mode){
	static_cast<QClipboard*>(ptr)->clear(static_cast<QClipboard::Mode>(mode));
}

QtObjectPtr QClipboard_MimeData(QtObjectPtr ptr, int mode){
	return const_cast<QMimeData*>(static_cast<QClipboard*>(ptr)->mimeData(static_cast<QClipboard::Mode>(mode)));
}

void QClipboard_SetMimeData(QtObjectPtr ptr, QtObjectPtr src, int mode){
	static_cast<QClipboard*>(ptr)->setMimeData(static_cast<QMimeData*>(src), static_cast<QClipboard::Mode>(mode));
}

void QClipboard_ConnectChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)(QClipboard::Mode)>(&QClipboard::changed), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)(QClipboard::Mode)>(&MyQClipboard::Signal_Changed));;
}

void QClipboard_DisconnectChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)(QClipboard::Mode)>(&QClipboard::changed), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)(QClipboard::Mode)>(&MyQClipboard::Signal_Changed));;
}

void QClipboard_ConnectDataChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::dataChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_DataChanged));;
}

void QClipboard_DisconnectDataChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::dataChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_DataChanged));;
}

void QClipboard_ConnectFindBufferChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::findBufferChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_FindBufferChanged));;
}

void QClipboard_DisconnectFindBufferChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::findBufferChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_FindBufferChanged));;
}

int QClipboard_OwnsClipboard(QtObjectPtr ptr){
	return static_cast<QClipboard*>(ptr)->ownsClipboard();
}

int QClipboard_OwnsFindBuffer(QtObjectPtr ptr){
	return static_cast<QClipboard*>(ptr)->ownsFindBuffer();
}

int QClipboard_OwnsSelection(QtObjectPtr ptr){
	return static_cast<QClipboard*>(ptr)->ownsSelection();
}

void QClipboard_ConnectSelectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::selectionChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_SelectionChanged));;
}

void QClipboard_DisconnectSelectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::selectionChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_SelectionChanged));;
}

void QClipboard_SetImage(QtObjectPtr ptr, QtObjectPtr image, int mode){
	static_cast<QClipboard*>(ptr)->setImage(*static_cast<QImage*>(image), static_cast<QClipboard::Mode>(mode));
}

void QClipboard_SetPixmap(QtObjectPtr ptr, QtObjectPtr pixmap, int mode){
	static_cast<QClipboard*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QClipboard::Mode>(mode));
}

void QClipboard_SetText(QtObjectPtr ptr, char* text, int mode){
	static_cast<QClipboard*>(ptr)->setText(QString(text), static_cast<QClipboard::Mode>(mode));
}

int QClipboard_SupportsFindBuffer(QtObjectPtr ptr){
	return static_cast<QClipboard*>(ptr)->supportsFindBuffer();
}

int QClipboard_SupportsSelection(QtObjectPtr ptr){
	return static_cast<QClipboard*>(ptr)->supportsSelection();
}

char* QClipboard_Text(QtObjectPtr ptr, int mode){
	return static_cast<QClipboard*>(ptr)->text(static_cast<QClipboard::Mode>(mode)).toUtf8().data();
}

