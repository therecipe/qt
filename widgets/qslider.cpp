#include "qslider.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QEvent>
#include <QWidget>
#include <QSlider>
#include "_cgo_export.h"

class MyQSlider: public QSlider {
public:
};

void QSlider_SetTickInterval(QtObjectPtr ptr, int ti){
	static_cast<QSlider*>(ptr)->setTickInterval(ti);
}

void QSlider_SetTickPosition(QtObjectPtr ptr, int position){
	static_cast<QSlider*>(ptr)->setTickPosition(static_cast<QSlider::TickPosition>(position));
}

int QSlider_TickInterval(QtObjectPtr ptr){
	return static_cast<QSlider*>(ptr)->tickInterval();
}

int QSlider_TickPosition(QtObjectPtr ptr){
	return static_cast<QSlider*>(ptr)->tickPosition();
}

QtObjectPtr QSlider_NewQSlider(QtObjectPtr parent){
	return new QSlider(static_cast<QWidget*>(parent));
}

QtObjectPtr QSlider_NewQSlider2(int orientation, QtObjectPtr parent){
	return new QSlider(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

int QSlider_Event(QtObjectPtr ptr, QtObjectPtr event){
	return static_cast<QSlider*>(ptr)->event(static_cast<QEvent*>(event));
}

void QSlider_DestroyQSlider(QtObjectPtr ptr){
	static_cast<QSlider*>(ptr)->~QSlider();
}

