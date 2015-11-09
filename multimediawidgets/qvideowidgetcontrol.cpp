#include "qvideowidgetcontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVideoWidget>
#include <QObject>
#include <QVideoWidgetControl>
#include "_cgo_export.h"

class MyQVideoWidgetControl: public QVideoWidgetControl {
public:
void Signal_BrightnessChanged(int brightness){callbackQVideoWidgetControlBrightnessChanged(this->objectName().toUtf8().data(), brightness);};
void Signal_ContrastChanged(int contrast){callbackQVideoWidgetControlContrastChanged(this->objectName().toUtf8().data(), contrast);};
void Signal_FullScreenChanged(bool fullScreen){callbackQVideoWidgetControlFullScreenChanged(this->objectName().toUtf8().data(), fullScreen);};
void Signal_HueChanged(int hue){callbackQVideoWidgetControlHueChanged(this->objectName().toUtf8().data(), hue);};
void Signal_SaturationChanged(int saturation){callbackQVideoWidgetControlSaturationChanged(this->objectName().toUtf8().data(), saturation);};
};

int QVideoWidgetControl_AspectRatioMode(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->aspectRatioMode();
}

int QVideoWidgetControl_Brightness(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->brightness();
}

void QVideoWidgetControl_ConnectBrightnessChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::brightnessChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_BrightnessChanged));;
}

void QVideoWidgetControl_DisconnectBrightnessChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::brightnessChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_BrightnessChanged));;
}

int QVideoWidgetControl_Contrast(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->contrast();
}

void QVideoWidgetControl_ConnectContrastChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::contrastChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_ContrastChanged));;
}

void QVideoWidgetControl_DisconnectContrastChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::contrastChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_ContrastChanged));;
}

void QVideoWidgetControl_ConnectFullScreenChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(bool)>(&QVideoWidgetControl::fullScreenChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(bool)>(&MyQVideoWidgetControl::Signal_FullScreenChanged));;
}

void QVideoWidgetControl_DisconnectFullScreenChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(bool)>(&QVideoWidgetControl::fullScreenChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(bool)>(&MyQVideoWidgetControl::Signal_FullScreenChanged));;
}

int QVideoWidgetControl_Hue(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->hue();
}

void QVideoWidgetControl_ConnectHueChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::hueChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_HueChanged));;
}

void QVideoWidgetControl_DisconnectHueChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::hueChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_HueChanged));;
}

int QVideoWidgetControl_IsFullScreen(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->isFullScreen();
}

int QVideoWidgetControl_Saturation(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->saturation();
}

void QVideoWidgetControl_ConnectSaturationChanged(void* ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::saturationChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_SaturationChanged));;
}

void QVideoWidgetControl_DisconnectSaturationChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::saturationChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_SaturationChanged));;
}

void QVideoWidgetControl_SetAspectRatioMode(void* ptr, int mode){
	static_cast<QVideoWidgetControl*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QVideoWidgetControl_SetBrightness(void* ptr, int brightness){
	static_cast<QVideoWidgetControl*>(ptr)->setBrightness(brightness);
}

void QVideoWidgetControl_SetContrast(void* ptr, int contrast){
	static_cast<QVideoWidgetControl*>(ptr)->setContrast(contrast);
}

void QVideoWidgetControl_SetFullScreen(void* ptr, int fullScreen){
	static_cast<QVideoWidgetControl*>(ptr)->setFullScreen(fullScreen != 0);
}

void QVideoWidgetControl_SetHue(void* ptr, int hue){
	static_cast<QVideoWidgetControl*>(ptr)->setHue(hue);
}

void QVideoWidgetControl_SetSaturation(void* ptr, int saturation){
	static_cast<QVideoWidgetControl*>(ptr)->setSaturation(saturation);
}

void* QVideoWidgetControl_VideoWidget(void* ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->videoWidget();
}

void QVideoWidgetControl_DestroyQVideoWidgetControl(void* ptr){
	static_cast<QVideoWidgetControl*>(ptr)->~QVideoWidgetControl();
}

