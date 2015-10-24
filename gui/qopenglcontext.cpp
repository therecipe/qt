#include "qopenglcontext.h"
#include <QSurface>
#include <QObject>
#include <QByteArray>
#include <QOpenGLVersionProfile>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScreen>
#include <QSurfaceFormat>
#include <QOpenGLContext>
#include "_cgo_export.h"

class MyQOpenGLContext: public QOpenGLContext {
public:
void Signal_AboutToBeDestroyed(){callbackQOpenGLContextAboutToBeDestroyed(this->objectName().toUtf8().data());};
};

QtObjectPtr QOpenGLContext_NewQOpenGLContext(QtObjectPtr parent){
	return new QOpenGLContext(static_cast<QObject*>(parent));
}

void QOpenGLContext_ConnectAboutToBeDestroyed(QtObjectPtr ptr){
	QObject::connect(static_cast<QOpenGLContext*>(ptr), static_cast<void (QOpenGLContext::*)()>(&QOpenGLContext::aboutToBeDestroyed), static_cast<MyQOpenGLContext*>(ptr), static_cast<void (MyQOpenGLContext::*)()>(&MyQOpenGLContext::Signal_AboutToBeDestroyed));;
}

void QOpenGLContext_DisconnectAboutToBeDestroyed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QOpenGLContext*>(ptr), static_cast<void (QOpenGLContext::*)()>(&QOpenGLContext::aboutToBeDestroyed), static_cast<MyQOpenGLContext*>(ptr), static_cast<void (MyQOpenGLContext::*)()>(&MyQOpenGLContext::Signal_AboutToBeDestroyed));;
}

int QOpenGLContext_QOpenGLContext_AreSharing(QtObjectPtr first, QtObjectPtr second){
	return QOpenGLContext::areSharing(static_cast<QOpenGLContext*>(first), static_cast<QOpenGLContext*>(second));
}

int QOpenGLContext_Create(QtObjectPtr ptr){
	return static_cast<QOpenGLContext*>(ptr)->create();
}

QtObjectPtr QOpenGLContext_QOpenGLContext_CurrentContext(){
	return QOpenGLContext::currentContext();
}

void QOpenGLContext_DoneCurrent(QtObjectPtr ptr){
	static_cast<QOpenGLContext*>(ptr)->doneCurrent();
}

QtObjectPtr QOpenGLContext_Functions(QtObjectPtr ptr){
	return static_cast<QOpenGLContext*>(ptr)->functions();
}

QtObjectPtr QOpenGLContext_QOpenGLContext_GlobalShareContext(){
	return QOpenGLContext::globalShareContext();
}

int QOpenGLContext_HasExtension(QtObjectPtr ptr, QtObjectPtr extension){
	return static_cast<QOpenGLContext*>(ptr)->hasExtension(*static_cast<QByteArray*>(extension));
}

int QOpenGLContext_IsOpenGLES(QtObjectPtr ptr){
	return static_cast<QOpenGLContext*>(ptr)->isOpenGLES();
}

int QOpenGLContext_IsValid(QtObjectPtr ptr){
	return static_cast<QOpenGLContext*>(ptr)->isValid();
}

int QOpenGLContext_MakeCurrent(QtObjectPtr ptr, QtObjectPtr surface){
	return static_cast<QOpenGLContext*>(ptr)->makeCurrent(static_cast<QSurface*>(surface));
}

char* QOpenGLContext_NativeHandle(QtObjectPtr ptr){
	return static_cast<QOpenGLContext*>(ptr)->nativeHandle().toString().toUtf8().data();
}

void QOpenGLContext_QOpenGLContext_OpenGLModuleHandle(){
	QOpenGLContext::openGLModuleHandle();
}

int QOpenGLContext_QOpenGLContext_OpenGLModuleType(){
	return QOpenGLContext::openGLModuleType();
}

void QOpenGLContext_SetFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QOpenGLContext*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QOpenGLContext_SetNativeHandle(QtObjectPtr ptr, char* handle){
	static_cast<QOpenGLContext*>(ptr)->setNativeHandle(QVariant(handle));
}

QtObjectPtr QOpenGLContext_Screen(QtObjectPtr ptr){
	return static_cast<QOpenGLContext*>(ptr)->screen();
}

void QOpenGLContext_SetScreen(QtObjectPtr ptr, QtObjectPtr screen){
	static_cast<QOpenGLContext*>(ptr)->setScreen(static_cast<QScreen*>(screen));
}

void QOpenGLContext_SetShareContext(QtObjectPtr ptr, QtObjectPtr shareContext){
	static_cast<QOpenGLContext*>(ptr)->setShareContext(static_cast<QOpenGLContext*>(shareContext));
}

QtObjectPtr QOpenGLContext_ShareContext(QtObjectPtr ptr){
	return static_cast<QOpenGLContext*>(ptr)->shareContext();
}

QtObjectPtr QOpenGLContext_ShareGroup(QtObjectPtr ptr){
	return static_cast<QOpenGLContext*>(ptr)->shareGroup();
}

int QOpenGLContext_QOpenGLContext_SupportsThreadedOpenGL(){
	return QOpenGLContext::supportsThreadedOpenGL();
}

QtObjectPtr QOpenGLContext_Surface(QtObjectPtr ptr){
	return static_cast<QOpenGLContext*>(ptr)->surface();
}

void QOpenGLContext_SwapBuffers(QtObjectPtr ptr, QtObjectPtr surface){
	static_cast<QOpenGLContext*>(ptr)->swapBuffers(static_cast<QSurface*>(surface));
}

QtObjectPtr QOpenGLContext_VersionFunctions(QtObjectPtr ptr, QtObjectPtr versionProfile){
	return static_cast<QOpenGLContext*>(ptr)->versionFunctions(*static_cast<QOpenGLVersionProfile*>(versionProfile));
}

void QOpenGLContext_DestroyQOpenGLContext(QtObjectPtr ptr){
	static_cast<QOpenGLContext*>(ptr)->~QOpenGLContext();
}

