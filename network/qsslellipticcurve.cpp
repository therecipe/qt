#include "qsslellipticcurve.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSslEllipticCurve>
#include "_cgo_export.h"

class MyQSslEllipticCurve: public QSslEllipticCurve {
public:
};

QtObjectPtr QSslEllipticCurve_NewQSslEllipticCurve(){
	return new QSslEllipticCurve();
}

int QSslEllipticCurve_IsValid(QtObjectPtr ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->isValid();
}

int QSslEllipticCurve_IsTlsNamedCurve(QtObjectPtr ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->isTlsNamedCurve();
}

char* QSslEllipticCurve_LongName(QtObjectPtr ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->longName().toUtf8().data();
}

char* QSslEllipticCurve_ShortName(QtObjectPtr ptr){
	return static_cast<QSslEllipticCurve*>(ptr)->shortName().toUtf8().data();
}

