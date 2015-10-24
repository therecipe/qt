#include "qstyleoptionheader.h"
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QString>
#include <QVariant>
#include <QStyleOptionHeader>
#include "_cgo_export.h"

class MyQStyleOptionHeader: public QStyleOptionHeader {
public:
};

QtObjectPtr QStyleOptionHeader_NewQStyleOptionHeader(){
	return new QStyleOptionHeader();
}

QtObjectPtr QStyleOptionHeader_NewQStyleOptionHeader2(QtObjectPtr other){
	return new QStyleOptionHeader(*static_cast<QStyleOptionHeader*>(other));
}

