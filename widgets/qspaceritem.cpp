#include "qspaceritem.h"
#include <QSize>
#include <QRect>
#include <QSizePolicy>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSpacerItem>
#include "_cgo_export.h"

class MyQSpacerItem: public QSpacerItem {
public:
};

void* QSpacerItem_NewQSpacerItem(int w, int h, int hPolicy, int vPolicy){
	return new QSpacerItem(w, h, static_cast<QSizePolicy::Policy>(hPolicy), static_cast<QSizePolicy::Policy>(vPolicy));
}

void QSpacerItem_ChangeSize(void* ptr, int w, int h, int hPolicy, int vPolicy){
	static_cast<QSpacerItem*>(ptr)->changeSize(w, h, static_cast<QSizePolicy::Policy>(hPolicy), static_cast<QSizePolicy::Policy>(vPolicy));
}

int QSpacerItem_ExpandingDirections(void* ptr){
	return static_cast<QSpacerItem*>(ptr)->expandingDirections();
}

int QSpacerItem_IsEmpty(void* ptr){
	return static_cast<QSpacerItem*>(ptr)->isEmpty();
}

void QSpacerItem_SetGeometry(void* ptr, void* r){
	static_cast<QSpacerItem*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void* QSpacerItem_SpacerItem(void* ptr){
	return static_cast<QSpacerItem*>(ptr)->spacerItem();
}

void QSpacerItem_DestroyQSpacerItem(void* ptr){
	static_cast<QSpacerItem*>(ptr)->~QSpacerItem();
}

