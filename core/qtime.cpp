#include "qtime.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTime>
#include "_cgo_export.h"

class MyQTime: public QTime {
public:
};

QtObjectPtr QTime_NewQTime(){
	return new QTime();
}

QtObjectPtr QTime_NewQTime3(int h, int m, int s, int ms){
	return new QTime(h, m, s, ms);
}

int QTime_Elapsed(QtObjectPtr ptr){
	return static_cast<QTime*>(ptr)->elapsed();
}

int QTime_Hour(QtObjectPtr ptr){
	return static_cast<QTime*>(ptr)->hour();
}

int QTime_IsNull(QtObjectPtr ptr){
	return static_cast<QTime*>(ptr)->isNull();
}

int QTime_QTime_IsValid2(int h, int m, int s, int ms){
	return QTime::isValid(h, m, s, ms);
}

int QTime_IsValid(QtObjectPtr ptr){
	return static_cast<QTime*>(ptr)->isValid();
}

int QTime_Minute(QtObjectPtr ptr){
	return static_cast<QTime*>(ptr)->minute();
}

int QTime_Msec(QtObjectPtr ptr){
	return static_cast<QTime*>(ptr)->msec();
}

int QTime_MsecsSinceStartOfDay(QtObjectPtr ptr){
	return static_cast<QTime*>(ptr)->msecsSinceStartOfDay();
}

int QTime_MsecsTo(QtObjectPtr ptr, QtObjectPtr t){
	return static_cast<QTime*>(ptr)->msecsTo(*static_cast<QTime*>(t));
}

int QTime_Restart(QtObjectPtr ptr){
	return static_cast<QTime*>(ptr)->restart();
}

int QTime_Second(QtObjectPtr ptr){
	return static_cast<QTime*>(ptr)->second();
}

int QTime_SecsTo(QtObjectPtr ptr, QtObjectPtr t){
	return static_cast<QTime*>(ptr)->secsTo(*static_cast<QTime*>(t));
}

int QTime_SetHMS(QtObjectPtr ptr, int h, int m, int s, int ms){
	return static_cast<QTime*>(ptr)->setHMS(h, m, s, ms);
}

void QTime_Start(QtObjectPtr ptr){
	static_cast<QTime*>(ptr)->start();
}

char* QTime_ToString2(QtObjectPtr ptr, int format){
	return static_cast<QTime*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8().data();
}

char* QTime_ToString(QtObjectPtr ptr, char* format){
	return static_cast<QTime*>(ptr)->toString(QString(format)).toUtf8().data();
}

