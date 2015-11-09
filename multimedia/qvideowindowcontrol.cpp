#include "qvideowindowcontrol.h"
#include <QObject>
#include <QRect>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVideoWindowControl>
#include "_cgo_export.h"

class MyQVideoWindowControl: public QVideoWindowControl {
public:
void Signal_BrightnessChanged(int brightness){callbackQVideoWindowControlBrightnessChanged(this->objectName().toUtf8().data(), brightness);};
void Signal_ContrastChanged(int contrast){callbackQVideoWindowControlContrastChanged(this->objectName().toUtf8().data(), contrast);};
void Signal_FullScreenChanged(bool fullScreen){callbackQVideoWindowControlFullScreenChanged(this->objectName().toUtf8().data(), fullScreen);};
void Signal_HueChanged(int hue){callbackQVideoWindowControlHueChanged(this->objectName().toUtf8().data(), hue);};
void Signal_NativeSizeChanged(){callbackQVideoWindowControlNativeSizeChanged(this->objectName().toUtf8().data());};
void Signal_SaturationChanged(int saturation){callbackQVideoWindowControlSaturationChanged(this->objectName().toUtf8().data(), saturation);};
};

int QVideoWindowControl_AspectRatioMode(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->aspectRatioMode();
}

int QVideoWindowControl_Brightness(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->brightness();
}

void QVideoWindowControl_ConnectBrightnessChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::brightnessChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_BrightnessChanged));;
}

void QVideoWindowControl_DisconnectBrightnessChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::brightnessChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_BrightnessChanged));;
}

int QVideoWindowControl_Contrast(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->contrast();
}

void QVideoWindowControl_ConnectContrastChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::contrastChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_ContrastChanged));;
}

void QVideoWindowControl_DisconnectContrastChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::contrastChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_ContrastChanged));;
}

void QVideoWindowControl_ConnectFullScreenChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(bool)>(&QVideoWindowControl::fullScreenChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(bool)>(&MyQVideoWindowControl::Signal_FullScreenChanged));;
}

void QVideoWindowControl_DisconnectFullScreenChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(bool)>(&QVideoWindowControl::fullScreenChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(bool)>(&MyQVideoWindowControl::Signal_FullScreenChanged));;
}

int QVideoWindowControl_Hue(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->hue();
}

void QVideoWindowControl_ConnectHueChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::hueChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_HueChanged));;
}

void QVideoWindowControl_DisconnectHueChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::hueChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_HueChanged));;
}

int QVideoWindowControl_IsFullScreen(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->isFullScreen();
}

void QVideoWindowControl_ConnectNativeSizeChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)()>(&QVideoWindowControl::nativeSizeChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)()>(&MyQVideoWindowControl::Signal_NativeSizeChanged));;
}

void QVideoWindowControl_DisconnectNativeSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)()>(&QVideoWindowControl::nativeSizeChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)()>(&MyQVideoWindowControl::Signal_NativeSizeChanged));;
}

void QVideoWindowControl_Repaint(void* ptr){
	static_cast<QVideoWindowControl*>(ptr)->repaint();
}

int QVideoWindowControl_Saturation(void* ptr){
	return static_cast<QVideoWindowControl*>(ptr)->saturation();
}

void QVideoWindowControl_ConnectSaturationChanged(void* ptr){
	QObject::connect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::saturationChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_SaturationChanged));;
}

void QVideoWindowControl_DisconnectSaturationChanged(void* ptr){
	QObject::disconnect(static_cast<QVideoWindowControl*>(ptr), static_cast<void (QVideoWindowControl::*)(int)>(&QVideoWindowControl::saturationChanged), static_cast<MyQVideoWindowControl*>(ptr), static_cast<void (MyQVideoWindowControl::*)(int)>(&MyQVideoWindowControl::Signal_SaturationChanged));;
}

void QVideoWindowControl_SetAspectRatioMode(void* ptr, int mode){
	static_cast<QVideoWindowControl*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QVideoWindowControl_SetBrightness(void* ptr, int brightness){
	static_cast<QVideoWindowControl*>(ptr)->setBrightness(brightness);
}

void QVideoWindowControl_SetContrast(void* ptr, int contrast){
	static_cast<QVideoWindowControl*>(ptr)->setContrast(contrast);
}

void QVideoWindowControl_SetDisplayRect(void* ptr, void* rect){
	static_cast<QVideoWindowControl*>(ptr)->setDisplayRect(*static_cast<QRect*>(rect));
}

void QVideoWindowControl_SetFullScreen(void* ptr, int fullScreen){
	static_cast<QVideoWindowControl*>(ptr)->setFullScreen(fullScreen != 0);
}

void QVideoWindowControl_SetHue(void* ptr, int hue){
	static_cast<QVideoWindowControl*>(ptr)->setHue(hue);
}

void QVideoWindowControl_SetSaturation(void* ptr, int saturation){
	static_cast<QVideoWindowControl*>(ptr)->setSaturation(saturation);
}

void QVideoWindowControl_DestroyQVideoWindowControl(void* ptr){
	static_cast<QVideoWindowControl*>(ptr)->~QVideoWindowControl();
}

