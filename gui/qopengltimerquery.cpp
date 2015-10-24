#include "qopengltimerquery.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QOpenGLTimerQuery>
#include "_cgo_export.h"

class MyQOpenGLTimerQuery: public QOpenGLTimerQuery {
public:
};

QtObjectPtr QOpenGLTimerQuery_NewQOpenGLTimerQuery(QtObjectPtr parent){
	return new QOpenGLTimerQuery(static_cast<QObject*>(parent));
}

void QOpenGLTimerQuery_Begin(QtObjectPtr ptr){
	static_cast<QOpenGLTimerQuery*>(ptr)->begin();
}

int QOpenGLTimerQuery_Create(QtObjectPtr ptr){
	return static_cast<QOpenGLTimerQuery*>(ptr)->create();
}

void QOpenGLTimerQuery_Destroy(QtObjectPtr ptr){
	static_cast<QOpenGLTimerQuery*>(ptr)->destroy();
}

void QOpenGLTimerQuery_End(QtObjectPtr ptr){
	static_cast<QOpenGLTimerQuery*>(ptr)->end();
}

int QOpenGLTimerQuery_IsCreated(QtObjectPtr ptr){
	return static_cast<QOpenGLTimerQuery*>(ptr)->isCreated();
}

int QOpenGLTimerQuery_IsResultAvailable(QtObjectPtr ptr){
	return static_cast<QOpenGLTimerQuery*>(ptr)->isResultAvailable();
}

void QOpenGLTimerQuery_RecordTimestamp(QtObjectPtr ptr){
	static_cast<QOpenGLTimerQuery*>(ptr)->recordTimestamp();
}

void QOpenGLTimerQuery_DestroyQOpenGLTimerQuery(QtObjectPtr ptr){
	static_cast<QOpenGLTimerQuery*>(ptr)->~QOpenGLTimerQuery();
}

