#include "qstyleoptionheader.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionHeader>
#include "_cgo_export.h"

class MyQStyleOptionHeader: public QStyleOptionHeader {
public:
};

void* QStyleOptionHeader_NewQStyleOptionHeader(){
	return new QStyleOptionHeader();
}

void* QStyleOptionHeader_NewQStyleOptionHeader2(void* other){
	return new QStyleOptionHeader(*static_cast<QStyleOptionHeader*>(other));
}

