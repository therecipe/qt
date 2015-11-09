#include "qdial.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QDial>
#include "_cgo_export.h"

class MyQDial: public QDial {
public:
};

int QDial_NotchSize(void* ptr){
	return static_cast<QDial*>(ptr)->notchSize();
}

double QDial_NotchTarget(void* ptr){
	return static_cast<double>(static_cast<QDial*>(ptr)->notchTarget());
}

int QDial_NotchesVisible(void* ptr){
	return static_cast<QDial*>(ptr)->notchesVisible();
}

void QDial_SetNotchesVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QDial*>(ptr), "setNotchesVisible", Q_ARG(bool, visible != 0));
}

void QDial_SetWrapping(void* ptr, int on){
	QMetaObject::invokeMethod(static_cast<QDial*>(ptr), "setWrapping", Q_ARG(bool, on != 0));
}

int QDial_Wrapping(void* ptr){
	return static_cast<QDial*>(ptr)->wrapping();
}

void* QDial_NewQDial(void* parent){
	return new QDial(static_cast<QWidget*>(parent));
}

void QDial_DestroyQDial(void* ptr){
	static_cast<QDial*>(ptr)->~QDial();
}

