#include "qspaceritem.h"
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QSizePolicy>
#include <QRect>
#include <QString>
#include <QVariant>
#include <QSpacerItem>
#include "_cgo_export.h"

class MyQSpacerItem: public QSpacerItem {
public:
};

QtObjectPtr QSpacerItem_NewQSpacerItem(int w, int h, int hPolicy, int vPolicy){
	return new QSpacerItem(w, h, static_cast<QSizePolicy::Policy>(hPolicy), static_cast<QSizePolicy::Policy>(vPolicy));
}

void QSpacerItem_ChangeSize(QtObjectPtr ptr, int w, int h, int hPolicy, int vPolicy){
	static_cast<QSpacerItem*>(ptr)->changeSize(w, h, static_cast<QSizePolicy::Policy>(hPolicy), static_cast<QSizePolicy::Policy>(vPolicy));
}

int QSpacerItem_ExpandingDirections(QtObjectPtr ptr){
	return static_cast<QSpacerItem*>(ptr)->expandingDirections();
}

int QSpacerItem_IsEmpty(QtObjectPtr ptr){
	return static_cast<QSpacerItem*>(ptr)->isEmpty();
}

void QSpacerItem_SetGeometry(QtObjectPtr ptr, QtObjectPtr r){
	static_cast<QSpacerItem*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

QtObjectPtr QSpacerItem_SpacerItem(QtObjectPtr ptr){
	return static_cast<QSpacerItem*>(ptr)->spacerItem();
}

void QSpacerItem_DestroyQSpacerItem(QtObjectPtr ptr){
	static_cast<QSpacerItem*>(ptr)->~QSpacerItem();
}

