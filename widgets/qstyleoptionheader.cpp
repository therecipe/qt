#include "qstyleoptionheader.h"
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
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

