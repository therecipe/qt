#include "qslider.h"
#include <QModelIndex>
#include <QWidget>
#include <QEvent>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSlider>
#include "_cgo_export.h"

class MyQSlider: public QSlider {
public:
};

void QSlider_SetTickInterval(void* ptr, int ti){
	static_cast<QSlider*>(ptr)->setTickInterval(ti);
}

void QSlider_SetTickPosition(void* ptr, int position){
	static_cast<QSlider*>(ptr)->setTickPosition(static_cast<QSlider::TickPosition>(position));
}

int QSlider_TickInterval(void* ptr){
	return static_cast<QSlider*>(ptr)->tickInterval();
}

int QSlider_TickPosition(void* ptr){
	return static_cast<QSlider*>(ptr)->tickPosition();
}

void* QSlider_NewQSlider(void* parent){
	return new QSlider(static_cast<QWidget*>(parent));
}

void* QSlider_NewQSlider2(int orientation, void* parent){
	return new QSlider(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

int QSlider_Event(void* ptr, void* event){
	return static_cast<QSlider*>(ptr)->event(static_cast<QEvent*>(event));
}

void QSlider_DestroyQSlider(void* ptr){
	static_cast<QSlider*>(ptr)->~QSlider();
}

