#include "qdomnamednodemap.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QDomNamedNodeMap>
#include "_cgo_export.h"

class MyQDomNamedNodeMap: public QDomNamedNodeMap {
public:
};

QtObjectPtr QDomNamedNodeMap_NewQDomNamedNodeMap(){
	return new QDomNamedNodeMap();
}

QtObjectPtr QDomNamedNodeMap_NewQDomNamedNodeMap2(QtObjectPtr n){
	return new QDomNamedNodeMap(*static_cast<QDomNamedNodeMap*>(n));
}

int QDomNamedNodeMap_Contains(QtObjectPtr ptr, char* name){
	return static_cast<QDomNamedNodeMap*>(ptr)->contains(QString(name));
}

int QDomNamedNodeMap_Count(QtObjectPtr ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->count();
}

int QDomNamedNodeMap_IsEmpty(QtObjectPtr ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->isEmpty();
}

int QDomNamedNodeMap_Length(QtObjectPtr ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->length();
}

int QDomNamedNodeMap_Size(QtObjectPtr ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->size();
}

void QDomNamedNodeMap_DestroyQDomNamedNodeMap(QtObjectPtr ptr){
	static_cast<QDomNamedNodeMap*>(ptr)->~QDomNamedNodeMap();
}

