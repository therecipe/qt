#include "qvideowidgetcontrol.h"
#include <QVideoWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
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

#include "qcameraviewfinder.h"
#include <QCamera>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QCameraViewfinder>
#include "_cgo_export.h"

class MyQCameraViewfinder: public QCameraViewfinder {
public:
};

void* QCameraViewfinder_NewQCameraViewfinder(void* parent){
	return new QCameraViewfinder(static_cast<QWidget*>(parent));
}

void* QCameraViewfinder_MediaObject(void* ptr){
	return static_cast<QCameraViewfinder*>(ptr)->mediaObject();
}

void QCameraViewfinder_DestroyQCameraViewfinder(void* ptr){
	static_cast<QCameraViewfinder*>(ptr)->~QCameraViewfinder();
}

#include "qvideowidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QWidget>
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

#include "qgraphicsvideoitem.h"
#include <QUrl>
#include <QStyleOption>
#include <QSizeF>
#include <QWidget>
#include <QPointF>
#include <QString>
#include <QModelIndex>
#include <QGraphicsItem>
#include <QStyle>
#include <QPainter>
#include <QVariant>
#include <QPoint>
#include <QSize>
#include <QStyleOptionGraphicsItem>
#include <QGraphicsVideoItem>
#include "_cgo_export.h"

class MyQGraphicsVideoItem: public QGraphicsVideoItem {
public:
};

void* QGraphicsVideoItem_NewQGraphicsVideoItem(void* parent){
	return new QGraphicsVideoItem(static_cast<QGraphicsItem*>(parent));
}

int QGraphicsVideoItem_AspectRatioMode(void* ptr){
	return static_cast<QGraphicsVideoItem*>(ptr)->aspectRatioMode();
}

void* QGraphicsVideoItem_MediaObject(void* ptr){
	return static_cast<QGraphicsVideoItem*>(ptr)->mediaObject();
}

void QGraphicsVideoItem_Paint(void* ptr, void* painter, void* option, void* widget){
	static_cast<QGraphicsVideoItem*>(ptr)->paint(static_cast<QPainter*>(painter), static_cast<QStyleOptionGraphicsItem*>(option), static_cast<QWidget*>(widget));
}

void QGraphicsVideoItem_SetAspectRatioMode(void* ptr, int mode){
	static_cast<QGraphicsVideoItem*>(ptr)->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

void QGraphicsVideoItem_SetOffset(void* ptr, void* offset){
	static_cast<QGraphicsVideoItem*>(ptr)->setOffset(*static_cast<QPointF*>(offset));
}

void QGraphicsVideoItem_SetSize(void* ptr, void* size){
	static_cast<QGraphicsVideoItem*>(ptr)->setSize(*static_cast<QSizeF*>(size));
}

void QGraphicsVideoItem_DestroyQGraphicsVideoItem(void* ptr){
	static_cast<QGraphicsVideoItem*>(ptr)->~QGraphicsVideoItem();
}

