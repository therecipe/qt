#include "qiconengine.h"
#include <QModelIndex>
#include <QRect>
#include <QPainter>
#include <QVariant>
#include <QUrl>
#include <QSize>
#include <QDataStream>
#include <QIcon>
#include <QString>
#include <QPixmap>
#include <QIconEngine>
#include "_cgo_export.h"

class MyQIconEngine: public QIconEngine {
public:
};

void QIconEngine_AddFile(QtObjectPtr ptr, char* fileName, QtObjectPtr size, int mode, int state){
	static_cast<QIconEngine*>(ptr)->addFile(QString(fileName), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIconEngine_AddPixmap(QtObjectPtr ptr, QtObjectPtr pixmap, int mode, int state){
	static_cast<QIconEngine*>(ptr)->addPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

QtObjectPtr QIconEngine_Clone(QtObjectPtr ptr){
	return static_cast<QIconEngine*>(ptr)->clone();
}

char* QIconEngine_IconName(QtObjectPtr ptr){
	return static_cast<QIconEngine*>(ptr)->iconName().toUtf8().data();
}

char* QIconEngine_Key(QtObjectPtr ptr){
	return static_cast<QIconEngine*>(ptr)->key().toUtf8().data();
}

void QIconEngine_Paint(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr rect, int mode, int state){
	static_cast<QIconEngine*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

int QIconEngine_Read(QtObjectPtr ptr, QtObjectPtr in){
	return static_cast<QIconEngine*>(ptr)->read(*static_cast<QDataStream*>(in));
}

int QIconEngine_Write(QtObjectPtr ptr, QtObjectPtr out){
	return static_cast<QIconEngine*>(ptr)->write(*static_cast<QDataStream*>(out));
}

void QIconEngine_DestroyQIconEngine(QtObjectPtr ptr){
	static_cast<QIconEngine*>(ptr)->~QIconEngine();
}

