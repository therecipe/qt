#include "qsslellipticcurve.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QSslEllipticCurve>
#include "_cgo_export.h"

class MyQSslEllipticCurve: public QSslEllipticCurve {
public:
};

void* QSslEllipticCurve_NewQSslEllipticCurve(){
	return new QSslEllipticCurve();
}

int QSslEllipticCurve_IsValid(void* ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->isValid();
}

int QSslEllipticCurve_IsTlsNamedCurve(void* ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->isTlsNamedCurve();
}

char* QSslEllipticCurve_LongName(void* ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->longName().toUtf8().data();
}

char* QSslEllipticCurve_ShortName(void* ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->shortName().toUtf8().data();
}

