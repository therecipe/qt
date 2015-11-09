#include "qprogressbar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QObject>
#include <QMetaObject>
#include <QProgressBar>
#include "_cgo_export.h"

class MyQProgressBar: public QProgressBar {
public:
void Signal_ValueChanged(int value){callbackQProgressBarValueChanged(this->objectName().toUtf8().data(), value);};
};

int QProgressBar_Alignment(void* ptr){
	return static_cast<QProgressBar*>(ptr)->alignment();
}

char* QProgressBar_Format(void* ptr){
	return static_cast<QProgressBar*>(ptr)->format().toUtf8().data();
}

int QProgressBar_InvertedAppearance(void* ptr){
	return static_cast<QProgressBar*>(ptr)->invertedAppearance();
}

int QProgressBar_IsTextVisible(void* ptr){
	return static_cast<QProgressBar*>(ptr)->isTextVisible();
}

int QProgressBar_Maximum(void* ptr){
	return static_cast<QProgressBar*>(ptr)->maximum();
}

int QProgressBar_Minimum(void* ptr){
	return static_cast<QProgressBar*>(ptr)->minimum();
}

int QProgressBar_Orientation(void* ptr){
	return static_cast<QProgressBar*>(ptr)->orientation();
}

void QProgressBar_ResetFormat(void* ptr){
	static_cast<QProgressBar*>(ptr)->resetFormat();
}

void QProgressBar_SetAlignment(void* ptr, int alignment){
	static_cast<QProgressBar*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QProgressBar_SetFormat(void* ptr, char* format){
	static_cast<QProgressBar*>(ptr)->setFormat(QString(format));
}

void QProgressBar_SetInvertedAppearance(void* ptr, int invert){
	static_cast<QProgressBar*>(ptr)->setInvertedAppearance(invert != 0);
}

void QProgressBar_SetMaximum(void* ptr, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setMaximum", Q_ARG(int, maximum));
}

void QProgressBar_SetMinimum(void* ptr, int minimum){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setMinimum", Q_ARG(int, minimum));
}

void QProgressBar_SetOrientation(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setOrientation", Q_ARG(Qt::Orientation, static_cast<Qt::Orientation>(v)));
}

void QProgressBar_SetTextDirection(void* ptr, int textDirection){
	static_cast<QProgressBar*>(ptr)->setTextDirection(static_cast<QProgressBar::Direction>(textDirection));
}

void QProgressBar_SetTextVisible(void* ptr, int visible){
	static_cast<QProgressBar*>(ptr)->setTextVisible(visible != 0);
}

void QProgressBar_SetValue(void* ptr, int value){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setValue", Q_ARG(int, value));
}

char* QProgressBar_Text(void* ptr){
	return static_cast<QProgressBar*>(ptr)->text().toUtf8().data();
}

int QProgressBar_TextDirection(void* ptr){
	return static_cast<QProgressBar*>(ptr)->textDirection();
}

int QProgressBar_Value(void* ptr){
	return static_cast<QProgressBar*>(ptr)->value();
}

void* QProgressBar_NewQProgressBar(void* parent){
	return new QProgressBar(static_cast<QWidget*>(parent));
}

void QProgressBar_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "reset");
}

void QProgressBar_SetRange(void* ptr, int minimum, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressBar*>(ptr), "setRange", Q_ARG(int, minimum), Q_ARG(int, maximum));
}

void QProgressBar_ConnectValueChanged(void* ptr){
	QObject::connect(static_cast<QProgressBar*>(ptr), static_cast<void (QProgressBar::*)(int)>(&QProgressBar::valueChanged), static_cast<MyQProgressBar*>(ptr), static_cast<void (MyQProgressBar::*)(int)>(&MyQProgressBar::Signal_ValueChanged));;
}

void QProgressBar_DisconnectValueChanged(void* ptr){
	QObject::disconnect(static_cast<QProgressBar*>(ptr), static_cast<void (QProgressBar::*)(int)>(&QProgressBar::valueChanged), static_cast<MyQProgressBar*>(ptr), static_cast<void (MyQProgressBar::*)(int)>(&MyQProgressBar::Signal_ValueChanged));;
}

void QProgressBar_DestroyQProgressBar(void* ptr){
	static_cast<QProgressBar*>(ptr)->~QProgressBar();
}

