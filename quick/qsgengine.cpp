#include "qsgengine.h"
#include <QOpenGLContext>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QImage>
#include <QSGEngine>
#include "_cgo_export.h"

class MyQSGEngine: public QSGEngine {
public:
};

QtObjectPtr QSGEngine_NewQSGEngine(QtObjectPtr parent){
	return new QSGEngine(static_cast<QObject*>(parent));
}

QtObjectPtr QSGEngine_CreateRenderer(QtObjectPtr ptr){
	return static_cast<QSGEngine*>(ptr)->createRenderer();
}

QtObjectPtr QSGEngine_CreateTextureFromImage(QtObjectPtr ptr, QtObjectPtr image, int options){
	return static_cast<QSGEngine*>(ptr)->createTextureFromImage(*static_cast<QImage*>(image), static_cast<QSGEngine::CreateTextureOption>(options));
}

void QSGEngine_Initialize(QtObjectPtr ptr, QtObjectPtr context){
	static_cast<QSGEngine*>(ptr)->initialize(static_cast<QOpenGLContext*>(context));
}

void QSGEngine_Invalidate(QtObjectPtr ptr){
	static_cast<QSGEngine*>(ptr)->invalidate();
}

void QSGEngine_DestroyQSGEngine(QtObjectPtr ptr){
	static_cast<QSGEngine*>(ptr)->~QSGEngine();
}

