#include "qopengltimemonitor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QOpenGLTimeMonitor>
#include "_cgo_export.h"

class MyQOpenGLTimeMonitor: public QOpenGLTimeMonitor {
public:
};

QtObjectPtr QOpenGLTimeMonitor_NewQOpenGLTimeMonitor(QtObjectPtr parent){
	return new QOpenGLTimeMonitor(static_cast<QObject*>(parent));
}

int QOpenGLTimeMonitor_Create(QtObjectPtr ptr){
	return static_cast<QOpenGLTimeMonitor*>(ptr)->create();
}

void QOpenGLTimeMonitor_Destroy(QtObjectPtr ptr){
	static_cast<QOpenGLTimeMonitor*>(ptr)->destroy();
}

int QOpenGLTimeMonitor_IsCreated(QtObjectPtr ptr){
	return static_cast<QOpenGLTimeMonitor*>(ptr)->isCreated();
}

int QOpenGLTimeMonitor_IsResultAvailable(QtObjectPtr ptr){
	return static_cast<QOpenGLTimeMonitor*>(ptr)->isResultAvailable();
}

int QOpenGLTimeMonitor_RecordSample(QtObjectPtr ptr){
	return static_cast<QOpenGLTimeMonitor*>(ptr)->recordSample();
}

void QOpenGLTimeMonitor_Reset(QtObjectPtr ptr){
	static_cast<QOpenGLTimeMonitor*>(ptr)->reset();
}

int QOpenGLTimeMonitor_SampleCount(QtObjectPtr ptr){
	return static_cast<QOpenGLTimeMonitor*>(ptr)->sampleCount();
}

void QOpenGLTimeMonitor_SetSampleCount(QtObjectPtr ptr, int sampleCount){
	static_cast<QOpenGLTimeMonitor*>(ptr)->setSampleCount(sampleCount);
}

void QOpenGLTimeMonitor_DestroyQOpenGLTimeMonitor(QtObjectPtr ptr){
	static_cast<QOpenGLTimeMonitor*>(ptr)->~QOpenGLTimeMonitor();
}

