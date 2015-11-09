#include "qstyleoptioncomplex.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionComplex>
#include "_cgo_export.h"

class MyQStyleOptionComplex: public QStyleOptionComplex {
public:
};

void* QStyleOptionComplex_NewQStyleOptionComplex2(void* other){
	return new QStyleOptionComplex(*static_cast<QStyleOptionComplex*>(other));
}

void* QStyleOptionComplex_NewQStyleOptionComplex(int version, int ty){
	return new QStyleOptionComplex(version, ty);
}

