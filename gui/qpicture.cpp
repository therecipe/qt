#include "qpicture.h"
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QRect>
#include <QPainter>
#include <QString>
#include <QVariant>
#include <QPicture>
#include "_cgo_export.h"

class MyQPicture: public QPicture {
public:
};

int QPicture_IsNull(QtObjectPtr ptr){
	return static_cast<QPicture*>(ptr)->isNull();
}

int QPicture_Load2(QtObjectPtr ptr, QtObjectPtr dev, char* format){
	return static_cast<QPicture*>(ptr)->load(static_cast<QIODevice*>(dev), const_cast<const char*>(format));
}

int QPicture_Load(QtObjectPtr ptr, char* fileName, char* format){
	return static_cast<QPicture*>(ptr)->load(QString(fileName), const_cast<const char*>(format));
}

int QPicture_Play(QtObjectPtr ptr, QtObjectPtr painter){
	return static_cast<QPicture*>(ptr)->play(static_cast<QPainter*>(painter));
}

int QPicture_Save2(QtObjectPtr ptr, QtObjectPtr dev, char* format){
	return static_cast<QPicture*>(ptr)->save(static_cast<QIODevice*>(dev), const_cast<const char*>(format));
}

int QPicture_Save(QtObjectPtr ptr, char* fileName, char* format){
	return static_cast<QPicture*>(ptr)->save(QString(fileName), const_cast<const char*>(format));
}

void QPicture_SetBoundingRect(QtObjectPtr ptr, QtObjectPtr r){
	static_cast<QPicture*>(ptr)->setBoundingRect(*static_cast<QRect*>(r));
}

void QPicture_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QPicture*>(ptr)->swap(*static_cast<QPicture*>(other));
}

void QPicture_DestroyQPicture(QtObjectPtr ptr){
	static_cast<QPicture*>(ptr)->~QPicture();
}

