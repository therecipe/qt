#include "qdomnamednodemap.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomNamedNodeMap>
#include "_cgo_export.h"

class MyQDomNamedNodeMap: public QDomNamedNodeMap {
public:
};

void* QDomNamedNodeMap_NewQDomNamedNodeMap(){
	return new QDomNamedNodeMap();
}

void* QDomNamedNodeMap_NewQDomNamedNodeMap2(void* n){
	return new QDomNamedNodeMap(*static_cast<QDomNamedNodeMap*>(n));
}

int QDomNamedNodeMap_Contains(void* ptr, char* name){
	return static_cast<QDomNamedNodeMap*>(ptr)->contains(QString(name));
}

int QDomNamedNodeMap_Count(void* ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->count();
}

int QDomNamedNodeMap_IsEmpty(void* ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->isEmpty();
}

int QDomNamedNodeMap_Length(void* ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->length();
}

int QDomNamedNodeMap_Size(void* ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->size();
}

void QDomNamedNodeMap_DestroyQDomNamedNodeMap(void* ptr){
	static_cast<QDomNamedNodeMap*>(ptr)->~QDomNamedNodeMap();
}

