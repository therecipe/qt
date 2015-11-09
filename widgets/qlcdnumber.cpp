#include "qlcdnumber.h"
#include <QModelIndex>
#include <QWidget>
#include <QObject>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QLCDNumber>
#include "_cgo_export.h"

class MyQLCDNumber: public QLCDNumber {
public:
void Signal_Overflow(){callbackQLCDNumberOverflow(this->objectName().toUtf8().data());};
};

int QLCDNumber_IntValue(void* ptr){
	return static_cast<QLCDNumber*>(ptr)->intValue();
}

int QLCDNumber_Mode(void* ptr){
	return static_cast<QLCDNumber*>(ptr)->mode();
}

int QLCDNumber_SegmentStyle(void* ptr){
	return static_cast<QLCDNumber*>(ptr)->segmentStyle();
}

void QLCDNumber_SetMode(void* ptr, int v){
	static_cast<QLCDNumber*>(ptr)->setMode(static_cast<QLCDNumber::Mode>(v));
}

void QLCDNumber_SetSegmentStyle(void* ptr, int v){
	static_cast<QLCDNumber*>(ptr)->setSegmentStyle(static_cast<QLCDNumber::SegmentStyle>(v));
}

void QLCDNumber_SetSmallDecimalPoint(void* ptr, int v){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setSmallDecimalPoint", Q_ARG(bool, v != 0));
}

int QLCDNumber_SmallDecimalPoint(void* ptr){
	return static_cast<QLCDNumber*>(ptr)->smallDecimalPoint();
}

void* QLCDNumber_NewQLCDNumber(void* parent){
	return new QLCDNumber(static_cast<QWidget*>(parent));
}

int QLCDNumber_CheckOverflow2(void* ptr, int num){
	return static_cast<QLCDNumber*>(ptr)->checkOverflow(num);
}

int QLCDNumber_DigitCount(void* ptr){
	return static_cast<QLCDNumber*>(ptr)->digitCount();
}

void QLCDNumber_Display(void* ptr, char* s){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "display", Q_ARG(QString, QString(s)));
}

void QLCDNumber_Display3(void* ptr, int num){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "display", Q_ARG(int, num));
}

void QLCDNumber_ConnectOverflow(void* ptr){
	QObject::connect(static_cast<QLCDNumber*>(ptr), static_cast<void (QLCDNumber::*)()>(&QLCDNumber::overflow), static_cast<MyQLCDNumber*>(ptr), static_cast<void (MyQLCDNumber::*)()>(&MyQLCDNumber::Signal_Overflow));;
}

void QLCDNumber_DisconnectOverflow(void* ptr){
	QObject::disconnect(static_cast<QLCDNumber*>(ptr), static_cast<void (QLCDNumber::*)()>(&QLCDNumber::overflow), static_cast<MyQLCDNumber*>(ptr), static_cast<void (MyQLCDNumber::*)()>(&MyQLCDNumber::Signal_Overflow));;
}

void QLCDNumber_SetBinMode(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setBinMode");
}

void QLCDNumber_SetDecMode(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setDecMode");
}

void QLCDNumber_SetDigitCount(void* ptr, int numDigits){
	static_cast<QLCDNumber*>(ptr)->setDigitCount(numDigits);
}

void QLCDNumber_SetHexMode(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setHexMode");
}

void QLCDNumber_SetOctMode(void* ptr){
	QMetaObject::invokeMethod(static_cast<QLCDNumber*>(ptr), "setOctMode");
}

void QLCDNumber_DestroyQLCDNumber(void* ptr){
	static_cast<QLCDNumber*>(ptr)->~QLCDNumber();
}

