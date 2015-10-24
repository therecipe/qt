#include "qdial.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QMetaObject>
#include <QDial>
#include "_cgo_export.h"

class MyQDial: public QDial {
public:
};

int QDial_NotchSize(QtObjectPtr ptr){
	return static_cast<QDial*>(ptr)->notchSize();
}

int QDial_NotchesVisible(QtObjectPtr ptr){
	return static_cast<QDial*>(ptr)->notchesVisible();
}

void QDial_SetNotchesVisible(QtObjectPtr ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QDial*>(ptr), "setNotchesVisible", Q_ARG(bool, visible != 0));
}

void QDial_SetWrapping(QtObjectPtr ptr, int on){
	QMetaObject::invokeMethod(static_cast<QDial*>(ptr), "setWrapping", Q_ARG(bool, on != 0));
}

int QDial_Wrapping(QtObjectPtr ptr){
	return static_cast<QDial*>(ptr)->wrapping();
}

QtObjectPtr QDial_NewQDial(QtObjectPtr parent){
	return new QDial(static_cast<QWidget*>(parent));
}

void QDial_DestroyQDial(QtObjectPtr ptr){
	static_cast<QDial*>(ptr)->~QDial();
}

