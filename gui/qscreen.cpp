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

int QScreen_Depth(QtObjectPtr ptr){
	return static_cast<QScreen*>(ptr)->depth();
}

char* QScreen_Name(QtObjectPtr ptr){
	return static_cast<QScreen*>(ptr)->name().toUtf8().data();
}

int QScreen_NativeOrientation(QtObjectPtr ptr){
	return static_cast<QScreen*>(ptr)->nativeOrientation();
}

int QScreen_Orientation(QtObjectPtr ptr){
	return static_cast<QScreen*>(ptr)->orientation();
}

int QScreen_PrimaryOrientation(QtObjectPtr ptr){
	return static_cast<QScreen*>(ptr)->primaryOrientation();
}

int QScreen_AngleBetween(QtObjectPtr ptr, int a, int b){
	return static_cast<QScreen*>(ptr)->angleBetween(static_cast<Qt::ScreenOrientation>(a), static_cast<Qt::ScreenOrientation>(b));
}

int QScreen_IsLandscape(QtObjectPtr ptr, int o){
	return static_cast<QScreen*>(ptr)->isLandscape(static_cast<Qt::ScreenOrientation>(o));
}

int QScreen_IsPortrait(QtObjectPtr ptr, int o){
	return static_cast<QScreen*>(ptr)->isPortrait(static_cast<Qt::ScreenOrientation>(o));
}

void QScreen_ConnectOrientationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::orientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_OrientationChanged));;
}

void QScreen_DisconnectOrientationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::orientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_OrientationChanged));;
}

int QScreen_OrientationUpdateMask(QtObjectPtr ptr){
	return static_cast<QScreen*>(ptr)->orientationUpdateMask();
}

void QScreen_ConnectPrimaryOrientationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::primaryOrientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_PrimaryOrientationChanged));;
}

void QScreen_DisconnectPrimaryOrientationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::primaryOrientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_PrimaryOrientationChanged));;
}

void QScreen_SetOrientationUpdateMask(QtObjectPtr ptr, int mask){
	static_cast<QScreen*>(ptr)->setOrientationUpdateMask(static_cast<Qt::ScreenOrientation>(mask));
}

void QScreen_DestroyQScreen(QtObjectPtr ptr){
	static_cast<QScreen*>(ptr)->~QScreen();
}

