#include "qscreen.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QScreen>
#include "_cgo_export.h"

class MyQScreen: public QScreen {
public:
void Signal_OrientationChanged(Qt::ScreenOrientation orientation){callbackQScreenOrientationChanged(this->objectName().toUtf8().data(), orientation);};
void Signal_PrimaryOrientationChanged(Qt::ScreenOrientation orientation){callbackQScreenPrimaryOrientationChanged(this->objectName().toUtf8().data(), orientation);};
};

int QScreen_Depth(void* ptr){
	return static_cast<QScreen*>(ptr)->depth();
}

double QScreen_DevicePixelRatio(void* ptr){
	return static_cast<double>(static_cast<QScreen*>(ptr)->devicePixelRatio());
}

double QScreen_LogicalDotsPerInch(void* ptr){
	return static_cast<double>(static_cast<QScreen*>(ptr)->logicalDotsPerInch());
}

double QScreen_LogicalDotsPerInchX(void* ptr){
	return static_cast<double>(static_cast<QScreen*>(ptr)->logicalDotsPerInchX());
}

double QScreen_LogicalDotsPerInchY(void* ptr){
	return static_cast<double>(static_cast<QScreen*>(ptr)->logicalDotsPerInchY());
}

char* QScreen_Name(void* ptr){
	return static_cast<QScreen*>(ptr)->name().toUtf8().data();
}

int QScreen_NativeOrientation(void* ptr){
	return static_cast<QScreen*>(ptr)->nativeOrientation();
}

int QScreen_Orientation(void* ptr){
	return static_cast<QScreen*>(ptr)->orientation();
}

double QScreen_PhysicalDotsPerInch(void* ptr){
	return static_cast<double>(static_cast<QScreen*>(ptr)->physicalDotsPerInch());
}

double QScreen_PhysicalDotsPerInchX(void* ptr){
	return static_cast<double>(static_cast<QScreen*>(ptr)->physicalDotsPerInchX());
}

double QScreen_PhysicalDotsPerInchY(void* ptr){
	return static_cast<double>(static_cast<QScreen*>(ptr)->physicalDotsPerInchY());
}

int QScreen_PrimaryOrientation(void* ptr){
	return static_cast<QScreen*>(ptr)->primaryOrientation();
}

double QScreen_RefreshRate(void* ptr){
	return static_cast<double>(static_cast<QScreen*>(ptr)->refreshRate());
}

int QScreen_AngleBetween(void* ptr, int a, int b){
	return static_cast<QScreen*>(ptr)->angleBetween(static_cast<Qt::ScreenOrientation>(a), static_cast<Qt::ScreenOrientation>(b));
}

int QScreen_IsLandscape(void* ptr, int o){
	return static_cast<QScreen*>(ptr)->isLandscape(static_cast<Qt::ScreenOrientation>(o));
}

int QScreen_IsPortrait(void* ptr, int o){
	return static_cast<QScreen*>(ptr)->isPortrait(static_cast<Qt::ScreenOrientation>(o));
}

void QScreen_ConnectOrientationChanged(void* ptr){
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::orientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_OrientationChanged));;
}

void QScreen_DisconnectOrientationChanged(void* ptr){
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::orientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_OrientationChanged));;
}

int QScreen_OrientationUpdateMask(void* ptr){
	return static_cast<QScreen*>(ptr)->orientationUpdateMask();
}

void QScreen_ConnectPrimaryOrientationChanged(void* ptr){
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::primaryOrientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_PrimaryOrientationChanged));;
}

void QScreen_DisconnectPrimaryOrientationChanged(void* ptr){
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::primaryOrientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_PrimaryOrientationChanged));;
}

void QScreen_SetOrientationUpdateMask(void* ptr, int mask){
	static_cast<QScreen*>(ptr)->setOrientationUpdateMask(static_cast<Qt::ScreenOrientation>(mask));
}

void QScreen_DestroyQScreen(void* ptr){
	static_cast<QScreen*>(ptr)->~QScreen();
}

