#include "qsgengine.h"
#include <QModelIndex>
#include <QImage>
#include <QObject>
#include <QOpenGLContext>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSGEngine>
#include "_cgo_export.h"

class MyQSGEngine: public QSGEngine {
public:
};

void* QSGEngine_NewQSGEngine(void* parent){
	return new QSGEngine(static_cast<QObject*>(parent));
}

void* QSGEngine_CreateRenderer(void* ptr){
	return static_cast<QSGEngine*>(ptr)->createRenderer();
}

void* QSGEngine_CreateTextureFromImage(void* ptr, void* image, int options){
	return static_cast<QSGEngine*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QSGEngine::CreateTextureOption>(options));
}

void QSGEngine_Initialize(void* ptr, void* context){
	static_cast<QSGEngine*>(ptr)->initialize(static_cast<QOpenGLContext*>(context));
}

void QSGEngine_Invalidate(void* ptr){
	static_cast<QSGEngine*>(ptr)->invalidate();
}

void QSGEngine_DestroyQSGEngine(void* ptr){
	static_cast<QSGEngine*>(ptr)->~QSGEngine();
}

