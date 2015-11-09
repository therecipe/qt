#include "qsplashscreen.h"
#include <QMetaObject>
#include <QWidget>
#include <QVariant>
#include <QUrl>
#include <QObject>
#include <QColor>
#include <QString>
#include <QModelIndex>
#include <QPixmap>
#include <QSplashScreen>
#include "_cgo_export.h"

class MyQSplashScreen: public QSplashScreen {
public:
void Signal_MessageChanged(const QString & message){callbackQSplashScreenMessageChanged(this->objectName().toUtf8().data(), message.toUtf8().data());};
};

void* QSplashScreen_NewQSplashScreen2(void* parent, void* pixmap, int f){
	return new QSplashScreen(static_cast<QWidget*>(parent), *static_cast<QPixmap*>(pixmap), static_cast<Qt::WindowType>(f));
}

void* QSplashScreen_NewQSplashScreen(void* pixmap, int f){
	return new QSplashScreen(*static_cast<QPixmap*>(pixmap), static_cast<Qt::WindowType>(f));
}

void QSplashScreen_ClearMessage(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSplashScreen*>(ptr), "clearMessage");
}

void QSplashScreen_Finish(void* ptr, void* mainWin){
	static_cast<QSplashScreen*>(ptr)->finish(static_cast<QWidget*>(mainWin));
}

char* QSplashScreen_Message(void* ptr){
	return static_cast<QSplashScreen*>(ptr)->message().toUtf8().data();
}

void QSplashScreen_ConnectMessageChanged(void* ptr){
	QObject::connect(static_cast<QSplashScreen*>(ptr), static_cast<void (QSplashScreen::*)(const QString &)>(&QSplashScreen::messageChanged), static_cast<MyQSplashScreen*>(ptr), static_cast<void (MyQSplashScreen::*)(const QString &)>(&MyQSplashScreen::Signal_MessageChanged));;
}

void QSplashScreen_DisconnectMessageChanged(void* ptr){
	QObject::disconnect(static_cast<QSplashScreen*>(ptr), static_cast<void (QSplashScreen::*)(const QString &)>(&QSplashScreen::messageChanged), static_cast<MyQSplashScreen*>(ptr), static_cast<void (MyQSplashScreen::*)(const QString &)>(&MyQSplashScreen::Signal_MessageChanged));;
}

void QSplashScreen_Repaint(void* ptr){
	static_cast<QSplashScreen*>(ptr)->repaint();
}

void QSplashScreen_SetPixmap(void* ptr, void* pixmap){
	static_cast<QSplashScreen*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap));
}

void QSplashScreen_ShowMessage(void* ptr, char* message, int alignment, void* color){
	QMetaObject::invokeMethod(static_cast<QSplashScreen*>(ptr), "showMessage", Q_ARG(QString, QString(message)), Q_ARG(int, alignment), Q_ARG(QColor, *static_cast<QColor*>(color)));
}

void QSplashScreen_DestroyQSplashScreen(void* ptr){
	static_cast<QSplashScreen*>(ptr)->~QSplashScreen();
}

