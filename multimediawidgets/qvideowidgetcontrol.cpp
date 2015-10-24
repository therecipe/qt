#include "qvideowidgetcontrol.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QVideoWidget>
#include <QString>
#include <QVariant>
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

int QVideoWidgetControl_AspectRatioMode(QtObjectPtr ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->aspectRatioMode();
}

int QVideoWidgetControl_Brightness(QtObjectPtr ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->brightness();
}

void QVideoWidgetControl_ConnectBrightnessChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::brightnessChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_BrightnessChanged));;
}

void QVideoWidgetControl_DisconnectBrightnessChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::brightnessChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_BrightnessChanged));;
}

int QVideoWidgetControl_Contrast(QtObjectPtr ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->contrast();
}

void QVideoWidgetControl_ConnectContrastChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::contrastChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_ContrastChanged));;
}

void QVideoWidgetControl_DisconnectContrastChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::contrastChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_ContrastChanged));;
}

void QVideoWidgetControl_ConnectFullScreenChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(bool)>(&QVideoWidgetControl::fullScreenChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(bool)>(&MyQVideoWidgetControl::Signal_FullScreenChanged));;
}

void QVideoWidgetControl_DisconnectFullScreenChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(bool)>(&QVideoWidgetControl::fullScreenChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(bool)>(&MyQVideoWidgetControl::Signal_FullScreenChanged));;
}

int QVideoWidgetControl_Hue(QtObjectPtr ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->hue();
}

void QVideoWidgetControl_ConnectHueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::hueChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_HueChanged));;
}

void QVideoWidgetControl_DisconnectHueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::hueChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_HueChanged));;
}

int QVideoWidgetControl_IsFullScreen(QtObjectPtr ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->isFullScreen();
}

int QVideoWidgetControl_Saturation(QtObjectPtr ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->saturation();
}

void QVideoWidgetControl_ConnectSaturationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::saturationChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_SaturationChanged));;
}

void QVideoWidgetControl_DisconnectSaturationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QVideoWidgetControl*>(ptr), static_cast<void (QVideoWidgetControl::*)(int)>(&QVideoWidgetControl::saturationChanged), static_cast<MyQVideoWidgetControl*>(ptr), static_cast<void (MyQVideoWidgetControl::*)(int)>(&MyQVideoWidgetControl::Signal_SaturationChanged));;
}

void QVideoWidgetControl_SetAspectRatioMode(QtObjectPtr ptr, int mode){
	static_cast<QVideoWidgetControl*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QVideoWidgetControl_SetBrightness(QtObjectPtr ptr, int brightness){
	static_cast<QVideoWidgetControl*>(ptr)->setBrightness(brightness);
}

void QVideoWidgetControl_SetContrast(QtObjectPtr ptr, int contrast){
	static_cast<QVideoWidgetControl*>(ptr)->setContrast(contrast);
}

void QVideoWidgetControl_SetFullScreen(QtObjectPtr ptr, int fullScreen){
	static_cast<QVideoWidgetControl*>(ptr)->setFullScreen(fullScreen != 0);
}

void QVideoWidgetControl_SetHue(QtObjectPtr ptr, int hue){
	static_cast<QVideoWidgetControl*>(ptr)->setHue(hue);
}

void QVideoWidgetControl_SetSaturation(QtObjectPtr ptr, int saturation){
	static_cast<QVideoWidgetControl*>(ptr)->setSaturation(saturation);
}

QtObjectPtr QVideoWidgetControl_VideoWidget(QtObjectPtr ptr){
	return static_cast<QVideoWidgetControl*>(ptr)->videoWidget();
}

void QVideoWidgetControl_DestroyQVideoWidgetControl(QtObjectPtr ptr){
	static_cast<QVideoWidgetControl*>(ptr)->~QVideoWidgetControl();
}

