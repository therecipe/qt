#include "qprogressbar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QMetaObject>
#include <QObject>
#include <QProgressBar>
#include "_cgo_export.h"

class MyQProgressBar: public QProgressBar {
public:
void Signal_ValueChanged(int value){callbackQProgressBarValueChanged(this->objectName().toUtf8().data(), value);};
};

int QProgressBar_Alignment(QtObjectPtr ptr){
	return static_cast<QProgressBar*>(ptr)->alignment();
}

char* QProgressBar_Format(QtObjectPtr ptr){
	return static_cast<QProgressBar*>(ptr)->format().toUtf8().data();
}

int QProgressBar_InvertedAppearance(QtObjectPtr ptr){
	return static_cast<QProgressBar*>(ptr)->invertedAppearance();
}

int QProgressBar_IsTextVisible(QtObjectPtr ptr){
	return static_cast<QProgressBar*>(ptr)->isTextVisible();
}

int QProgressBar_Maximum(QtObjectPtr ptr){
	return static_cast<QProgressBar*>(ptr)->maximum();
}

int QProgressBar_Minimum(QtObjectPtr ptr){
	return static_cast<QProgressBar*>(ptr)->minimum();
}

int QProgressBar_Orientation(QtObjectPtr ptr){
	return static_cast<QProgressBar*>(ptr)->orientation();
}

void QProgressBar_ResetFormat(QtObjectPtr ptr){
	static_cast<QProgressBar*>(ptr)->resetFormat();
}

void QProgressBar_SetAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QProgressBar*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QProgressBar_SetFormat(QtObjectPtr ptr, char* format){
	static_cast<QProgressBar*>(ptr)->setFormat(QString(format));
}

void QProgressBar_SetInvertedAppearance(QtObjectPtr ptr, int invert){
	static_cast<QProgressBar*>(ptr)->setInvertedAppearance(invert != 0);
}

void QProgressBar_SetMaximum(QtObjectPtr ptr, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setMaximum", Q_ARG(int, maximum));
}

void QProgressBar_SetMinimum(QtObjectPtr ptr, int minimum){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setMinimum", Q_ARG(int, minimum));
}

void QProgressBar_SetOrientation(QtObjectPtr ptr, int v){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setOrientation", Q_ARG(Qt::Orientation, static_cast<Qt::Orientation>(v)));
}

void QProgressBar_SetTextDirection(QtObjectPtr ptr, int textDirection){
	static_cast<QProgressBar*>(ptr)->setTextDirection(static_cast<QProgressBar::Direction>(textDirection));
}

void QProgressBar_SetTextVisible(QtObjectPtr ptr, int visible){
	static_cast<QProgressBar*>(ptr)->setTextVisible(visible != 0);
}

void QProgressBar_SetValue(QtObjectPtr ptr, int value){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setValue", Q_ARG(int, value));
}

char* QProgressBar_Text(QtObjectPtr ptr){
	return static_cast<QProgressBar*>(ptr)->text().toUtf8().data();
}

int QProgressBar_TextDirection(QtObjectPtr ptr){
	return static_cast<QProgressBar*>(ptr)->textDirection();
}

int QProgressBar_Value(QtObjectPtr ptr){
	return static_cast<QProgressBar*>(ptr)->value();
}

QtObjectPtr QProgressBar_NewQProgressBar(QtObjectPtr parent){
	return new QProgressBar(static_cast<QWidget*>(parent));
}

void QProgressBar_Reset(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "reset");
}

void QProgressBar_SetRange(QtObjectPtr ptr, int minimum, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setRange", Q_ARG(int, minimum), Q_ARG(int, maximum));
}

void QProgressBar_ConnectValueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QProgressBar*>(ptr), static_cast<void (QProgressBar::*)(int)>(&QProgressBar::valueChanged), static_cast<MyQProgressBar*>(ptr), static_cast<void (MyQProgressBar::*)(int)>(&MyQProgressBar::Signal_ValueChanged));;
}

void QProgressBar_DisconnectValueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QProgressBar*>(ptr), static_cast<void (QProgressBar::*)(int)>(&QProgressBar::valueChanged), static_cast<MyQProgressBar*>(ptr), static_cast<void (MyQProgressBar::*)(int)>(&MyQProgressBar::Signal_ValueChanged));;
}

void QProgressBar_DestroyQProgressBar(QtObjectPtr ptr){
	static_cast<QProgressBar*>(ptr)->~QProgressBar();
}

