#include "qvideowidget.h"
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QVideoWidget>
#include "_cgo_export.h"

class MyQVideoWidget: public QVideoWidget {
public:
void Signal_BrightnessChanged(int brightness){callbackQVideoWidgetBrightnessChanged(this->objectName().toUtf8().data(), brightness);};
void Signal_ContrastChanged(int contrast){callbackQVideoWidgetContrastChanged(this->objectName().toUtf8().data(), contrast);};
void Signal_FullScreenChanged(bool fullScreen){callbackQVideoWidgetFullScreenChanged(this->objectName().toUtf8().data(), fullScreen);};
void Signal_HueChanged(int hue){callbackQVideoWidgetHueChanged(this->objectName().toUtf8().data(), hue);};
void Signal_SaturationChanged(int saturation){callbackQVideoWidgetSaturationChanged(this->objectName().toUtf8().data(), saturation);};
};

int QVideoWidget_AspectRatioMode(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->aspectRatioMode();
}

int QVideoWidget_Brightness(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->brightness();
}

int QVideoWidget_Contrast(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->contrast();
}

int QVideoWidget_Hue(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->hue();
}

void* QVideoWidget_MediaObject(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->mediaObject();
}

int QVideoWidget_Saturation(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->saturation();
}

void QVideoWidget_SetAspectRatioMode(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setAspectRatioMode", Q_ARG(Qt::AspectRatioMode, static_cast<Qt::AspectRatioMode>(mode)));
}

void QVideoWidget_SetBrightness(void* ptr, int brightness){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setBrightness", Q_ARG(int, brightness));
}

void QVideoWidget_SetContrast(void* ptr, int contrast){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setContrast", Q_ARG(int, contrast));
}

void QVideoWidget_SetFullScreen(void* ptr, int fullScreen){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setFullScreen", Q_ARG(bool, fullScreen != 0));
}

void QVideoWidget_SetHue(void* ptr, int hue){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setHue", Q_ARG(int, hue));
}

void QVideoWidget_SetSaturation(void* ptr, int saturation){
	QMetaObject::invokeMethod(static_cast<QVideoWidget*>(ptr), "setSaturation", Q_ARG(int, saturation));
}

void* QVideoWidget_NewQVideoWidget(void* parent){
	return new QVideoWidget(static_cast<QWidget*>(parent));
}

void QVideoWidget_ConnectBrightnessChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::brightnessChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_BrightnessChanged));;
}

void QVideoWidget_DisconnectBrightnessChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::brightnessChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_BrightnessChanged));;
}

void QVideoWidget_ConnectContrastChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::contrastChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_ContrastChanged));;
}

void QVideoWidget_DisconnectContrastChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::contrastChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_ContrastChanged));;
}

void QVideoWidget_ConnectFullScreenChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(bool)>(&QVideoWidget::fullScreenChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(bool)>(&MyQVideoWidget::Signal_FullScreenChanged));;
}

void QVideoWidget_DisconnectFullScreenChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(bool)>(&QVideoWidget::fullScreenChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(bool)>(&MyQVideoWidget::Signal_FullScreenChanged));;
}

void QVideoWidget_ConnectHueChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::hueChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_HueChanged));;
}

void QVideoWidget_DisconnectHueChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::hueChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_HueChanged));;
}

int QVideoWidget_IsFullScreen(void* ptr){
	return static_cast<QVideoWidget*>(ptr)->isFullScreen();
}

void QVideoWidget_ConnectSaturationChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::saturationChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_SaturationChanged));;
}

void QVideoWidget_DisconnectSaturationChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidget*>(ptr), static_cast<void (QVideoWidget::*)(int)>(&QVideoWidget::saturationChanged), static_cast<MyQVideoWidget*>(ptr), static_cast<void (MyQVideoWidget::*)(int)>(&MyQVideoWidget::Signal_SaturationChanged));;
}

void QVideoWidget_DestroyQVideoWidget(void* ptr){
	static_cast<QVideoWidget*>(ptr)->~QVideoWidget();
}

