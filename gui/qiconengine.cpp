#include "qiconengine.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIcon>
#include <QSize>
#include <QRect>
#include <QPixmap>
#include <QString>
#include <QPainter>
#include <QDataStream>
#include <QIconEngine>
#include "_cgo_export.h"

class MyQIconEngine: public QIconEngine {
public:
};

void QIconEngine_AddFile(void* ptr, char* fileName, void* size, int mode, int state){
	static_cast<QIconEngine*>(ptr)->addFile(QString(fileName), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIconEngine_AddPixmap(void* ptr, void* pixmap, int mode, int state){
	static_cast<QIconEngine*>(ptr)->addPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void* QIconEngine_Clone(void* ptr){
	return static_cast<QIconEngine*>(ptr)->clone();
}

char* QIconEngine_IconName(void* ptr){
	return static_cast<QIconEngine*>(ptr)->iconName().toUtf8().data();
}

char* QIconEngine_Key(void* ptr){
	return static_cast<QIconEngine*>(ptr)->key().toUtf8().data();
}

void QIconEngine_Paint(void* ptr, void* painter, void* rect, int mode, int state){
	static_cast<QIconEngine*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

int QIconEngine_Read(void* ptr, void* in){
	return static_cast<QIconEngine*>(ptr)->read(*static_cast<QDataStream*>(in));
}

int QIconEngine_Write(void* ptr, void* out){
	return static_cast<QIconEngine*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QIconEngine_DestroyQIconEngine(void* ptr){
	static_cast<QIconEngine*>(ptr)->~QIconEngine();
}

