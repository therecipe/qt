#include "qcolormap.h"
#include <QModelIndex>
#include <QColor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QColormap>
#include "_cgo_export.h"

class MyQColormap: public QColormap {
public:
};

QtObjectPtr QColormap_NewQColormap(QtObjectPtr colormap){
	return new QColormap(*static_cast<QColormap*>(colormap));
}

int QColormap_Depth(QtObjectPtr ptr){
	return static_cast<QColormap*>(ptr)->depth();
}

int QColormap_Mode(QtObjectPtr ptr){
	return static_cast<QColormap*>(ptr)->mode();
}

int QColormap_Size(QtObjectPtr ptr){
	return static_cast<QColormap*>(ptr)->size();
}

void QColormap_DestroyQColormap(QtObjectPtr ptr){
	static_cast<QColormap*>(ptr)->~QColormap();
}

