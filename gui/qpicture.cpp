#include "qpicture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPainter>
#include <QRect>
#include <QIODevice>
#include <QPicture>
#include "_cgo_export.h"

class MyQPicture: public QPicture {
public:
};

int QPicture_IsNull(void* ptr){
	return static_cast<QPicture*>(ptr)->isNull();
}

int QPicture_Load2(void* ptr, void* dev, char* format){
	return static_cast<QPicture*>(ptr)->load(static_cast<QIODevice*>(dev), const_cast<const char*>(format));
}

int QPicture_Load(void* ptr, char* fileName, char* format){
	return static_cast<QPicture*>(ptr)->load(QString(fileName), const_cast<const char*>(format));
}

int QPicture_Play(void* ptr, void* painter){
	return static_cast<QPicture*>(ptr)->play(static_cast<QPainter*>(painter));
}

int QPicture_Save2(void* ptr, void* dev, char* format){
	return static_cast<QPicture*>(ptr)->save(static_cast<QIODevice*>(dev), const_cast<const char*>(format));
}

int QPicture_Save(void* ptr, char* fileName, char* format){
	return static_cast<QPicture*>(ptr)->save(QString(fileName), const_cast<const char*>(format));
}

void QPicture_SetBoundingRect(void* ptr, void* r){
	static_cast<QPicture*>(ptr)->setBoundingRect(*static_cast<QRect*>(r));
}

void QPicture_Swap(void* ptr, void* other){
	static_cast<QPicture*>(ptr)->swap(*static_cast<QPicture*>(other));
}

void QPicture_DestroyQPicture(void* ptr){
	static_cast<QPicture*>(ptr)->~QPicture();
}

