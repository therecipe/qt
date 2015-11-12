#include "qcolormap.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QColor>
#include <QColormap>
#include "_cgo_export.h"

class MyQColormap: public QColormap {
public:
};

void* QColormap_NewQColormap(void* colormap){
	return new QColormap(*static_cast<QColormap*>(colormap));
}

int QColormap_Depth(void* ptr){
	return static_cast<QColormap*>(ptr)->depth();
}

int QColormap_Mode(void* ptr){
	return static_cast<QColormap*>(ptr)->mode();
}

int QColormap_Size(void* ptr){
	return static_cast<QColormap*>(ptr)->size();
}

void QColormap_DestroyQColormap(void* ptr){
	static_cast<QColormap*>(ptr)->~QColormap();
}

