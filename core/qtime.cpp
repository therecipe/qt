#include "qtime.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QTime>
#include "_cgo_export.h"

class MyQTime: public QTime {
public:
};

void* QTime_NewQTime(){
	return new QTime();
}

void* QTime_NewQTime3(int h, int m, int s, int ms){
	return new QTime(h, m, s, ms);
}

int QTime_Elapsed(void* ptr){
	return static_cast<QTime*>(ptr)->elapsed();
}

int QTime_Hour(void* ptr){
	return static_cast<QTime*>(ptr)->hour();
}

int QTime_IsNull(void* ptr){
	return static_cast<QTime*>(ptr)->isNull();
}

int QTime_QTime_IsValid2(int h, int m, int s, int ms){
	return QTime::isValid(h, m, s, ms);
}

int QTime_IsValid(void* ptr){
	return static_cast<QTime*>(ptr)->isValid();
}

int QTime_Minute(void* ptr){
	return static_cast<QTime*>(ptr)->minute();
}

int QTime_Msec(void* ptr){
	return static_cast<QTime*>(ptr)->msec();
}

int QTime_MsecsSinceStartOfDay(void* ptr){
	return static_cast<QTime*>(ptr)->msecsSinceStartOfDay();
}

int QTime_MsecsTo(void* ptr, void* t){
	return static_cast<QTime*>(ptr)->msecsTo(*static_cast<QTime*>(t));
}

int QTime_Restart(void* ptr){
	return static_cast<QTime*>(ptr)->restart();
}

int QTime_Second(void* ptr){
	return static_cast<QTime*>(ptr)->second();
}

int QTime_SecsTo(void* ptr, void* t){
	return static_cast<QTime*>(ptr)->secsTo(*static_cast<QTime*>(t));
}

int QTime_SetHMS(void* ptr, int h, int m, int s, int ms){
	return static_cast<QTime*>(ptr)->setHMS(h, m, s, ms);
}

void QTime_Start(void* ptr){
	static_cast<QTime*>(ptr)->start();
}

char* QTime_ToString2(void* ptr, int format){
	return static_cast<QTime*>(ptr)->toString(static_cast<Qt::DateFormat>(format)).toUtf8().data();
}

char* QTime_ToString(void* ptr, char* format){
	return static_cast<QTime*>(ptr)->toString(QString(format)).toUtf8().data();
}

