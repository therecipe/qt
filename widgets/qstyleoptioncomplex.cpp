#include "qstyleoptioncomplex.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QStyleOptionComplex>
#include "_cgo_export.h"

class MyQStyleOptionComplex: public QStyleOptionComplex {
public:
};

QtObjectPtr QStyleOptionComplex_NewQStyleOptionComplex2(QtObjectPtr other){
	return new QStyleOptionComplex(*static_cast<QStyleOptionComplex*>(other));
}

QtObjectPtr QStyleOptionComplex_NewQStyleOptionComplex(int version, int ty){
	return new QStyleOptionComplex(version, ty);
}

