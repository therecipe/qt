#include "qvariant.h"
#include <QJsonObject>
#include <QChar>
#include <QPointF>
#include <QDateTime>
#include <QPersistentModelIndex>
#include <QString>
#include <QByteArray>
#include <QJsonValue>
#include <QDate>
#include <QPoint>
#include <QSizeF>
#include <QModelIndex>
#include <QRectF>
#include <QJsonArray>
#include <QSize>
#include <QRect>
#include <QLine>
#include <QEasingCurve>
#include <QRegExp>
#include <QLocale>
#include <QJsonDocument>
#include <QTime>
#include <QVariant>
#include <QRegularExpression>
#include <QLineF>
#include <QUuid>
#include <QBitArray>
#include <QDataStream>
#include <QLatin1String>
#include <QUrl>
#include <QVariant>
#include "_cgo_export.h"

class MyQVariant: public QVariant {
public:
};

void* QVariant_NewQVariant20(void* c){
	return new QVariant(*static_cast<QChar*>(c));
}

void* QVariant_NewQVariant18(void* val){
	return new QVariant(*static_cast<QLatin1String*>(val));
}

void* QVariant_NewQVariant11(int val){
	return new QVariant(val != 0);
}

void* QVariant_NewQVariant16(void* val){
	return new QVariant(*static_cast<QBitArray*>(val));
}

void* QVariant_NewQVariant15(void* val){
	return new QVariant(*static_cast<QByteArray*>(val));
}

void* QVariant_NewQVariant21(void* val){
	return new QVariant(*static_cast<QDate*>(val));
}

void* QVariant_NewQVariant23(void* val){
	return new QVariant(*static_cast<QDateTime*>(val));
}

void* QVariant_NewQVariant39(void* val){
	return new QVariant(*static_cast<QEasingCurve*>(val));
}

void* QVariant_NewQVariant45(void* val){
	return new QVariant(*static_cast<QJsonArray*>(val));
}

void* QVariant_NewQVariant46(void* val){
	return new QVariant(*static_cast<QJsonDocument*>(val));
}

void* QVariant_NewQVariant44(void* val){
	return new QVariant(*static_cast<QJsonObject*>(val));
}

void* QVariant_NewQVariant43(void* val){
	return new QVariant(*static_cast<QJsonValue*>(val));
}

void* QVariant_NewQVariant31(void* val){
	return new QVariant(*static_cast<QLine*>(val));
}

void* QVariant_NewQVariant32(void* val){
	return new QVariant(*static_cast<QLineF*>(val));
}

void* QVariant_NewQVariant35(void* l){
	return new QVariant(*static_cast<QLocale*>(l));
}

void* QVariant_NewQVariant41(void* val){
	return new QVariant(*static_cast<QModelIndex*>(val));
}

void* QVariant_NewQVariant42(void* val){
	return new QVariant(*static_cast<QPersistentModelIndex*>(val));
}

void* QVariant_NewQVariant29(void* val){
	return new QVariant(*static_cast<QPoint*>(val));
}

void* QVariant_NewQVariant30(void* val){
	return new QVariant(*static_cast<QPointF*>(val));
}

void* QVariant_NewQVariant33(void* val){
	return new QVariant(*static_cast<QRect*>(val));
}

void* QVariant_NewQVariant34(void* val){
	return new QVariant(*static_cast<QRectF*>(val));
}

void* QVariant_NewQVariant36(void* regExp){
	return new QVariant(*static_cast<QRegExp*>(regExp));
}

void* QVariant_NewQVariant37(void* re){
	return new QVariant(*static_cast<QRegularExpression*>(re));
}

void* QVariant_NewQVariant27(void* val){
	return new QVariant(*static_cast<QSize*>(val));
}

void* QVariant_NewQVariant28(void* val){
	return new QVariant(*static_cast<QSizeF*>(val));
}

void* QVariant_NewQVariant17(char* val){
	return new QVariant(QString(val));
}

void* QVariant_NewQVariant19(char* val){
	return new QVariant(QString(val).split("|", QString::SkipEmptyParts));
}

void* QVariant_NewQVariant22(void* val){
	return new QVariant(*static_cast<QTime*>(val));
}

void* QVariant_NewQVariant38(void* val){
	return new QVariant(*static_cast<QUrl*>(val));
}

void* QVariant_NewQVariant40(void* val){
	return new QVariant(*static_cast<QUuid*>(val));
}

void* QVariant_NewQVariant5(void* p){
	return new QVariant(*static_cast<QVariant*>(p));
}

void* QVariant_NewQVariant14(char* val){
	return new QVariant(const_cast<const char*>(val));
}

void* QVariant_NewQVariant7(int val){
	return new QVariant(val);
}

void* QVariant_ToByteArray(void* ptr){
	return new QByteArray(static_cast<QVariant*>(ptr)->toByteArray());
}

void* QVariant_ToDateTime(void* ptr){
	return new QDateTime(static_cast<QVariant*>(ptr)->toDateTime());
}

void* QVariant_ToEasingCurve(void* ptr){
	return new QEasingCurve(static_cast<QVariant*>(ptr)->toEasingCurve());
}

void* QVariant_ToRegExp(void* ptr){
	return new QRegExp(static_cast<QVariant*>(ptr)->toRegExp());
}

void* QVariant_ToRegularExpression(void* ptr){
	return new QRegularExpression(static_cast<QVariant*>(ptr)->toRegularExpression());
}

char* QVariant_ToStringList(void* ptr){
	return static_cast<QVariant*>(ptr)->toStringList().join("|").toUtf8().data();
}

void QVariant_DestroyQVariant(void* ptr){
	static_cast<QVariant*>(ptr)->~QVariant();
}

void* QVariant_NewQVariant(){
	return new QVariant();
}

void* QVariant_NewQVariant6(void* s){
	return new QVariant(*static_cast<QDataStream*>(s));
}

void* QVariant_NewQVariant47(void* other){
	return new QVariant(*static_cast<QVariant*>(other));
}

void QVariant_Clear(void* ptr){
	static_cast<QVariant*>(ptr)->clear();
}

int QVariant_Convert(void* ptr, int targetTypeId){
	return static_cast<QVariant*>(ptr)->convert(targetTypeId);
}

int QVariant_IsNull(void* ptr){
	return static_cast<QVariant*>(ptr)->isNull();
}

int QVariant_IsValid(void* ptr){
	return static_cast<QVariant*>(ptr)->isValid();
}

void QVariant_Swap(void* ptr, void* other){
	static_cast<QVariant*>(ptr)->swap(*static_cast<QVariant*>(other));
}

int QVariant_ToBool(void* ptr){
	return static_cast<QVariant*>(ptr)->toBool();
}

int QVariant_ToInt(void* ptr, int ok){
	return static_cast<QVariant*>(ptr)->toInt(NULL);
}

void* QVariant_ToJsonArray(void* ptr){
	return new QJsonArray(static_cast<QVariant*>(ptr)->toJsonArray());
}

void* QVariant_ToJsonDocument(void* ptr){
	return new QJsonDocument(static_cast<QVariant*>(ptr)->toJsonDocument());
}

void* QVariant_ToJsonObject(void* ptr){
	return new QJsonObject(static_cast<QVariant*>(ptr)->toJsonObject());
}

void* QVariant_ToModelIndex(void* ptr){
	return static_cast<QVariant*>(ptr)->toModelIndex().internalPointer();
}

double QVariant_ToReal(void* ptr, int ok){
	return static_cast<double>(static_cast<QVariant*>(ptr)->toReal(NULL));
}

char* QVariant_ToString(void* ptr){
	return static_cast<QVariant*>(ptr)->toString().toUtf8().data();
}

int QVariant_UserType(void* ptr){
	return static_cast<QVariant*>(ptr)->userType();
}

