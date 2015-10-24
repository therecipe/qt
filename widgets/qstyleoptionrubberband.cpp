#include "qstyleoptionrubberband.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionRubberBand>
#include "_cgo_export.h"

class MyQStyleOptionRubberBand: public QStyleOptionRubberBand {
public:
};

QtObjectPtr QStyleOptionRubberBand_NewQStyleOptionRubberBand(){
	return new QStyleOptionRubberBand();
}

QtObjectPtr QStyleOptionRubberBand_NewQStyleOptionRubberBand2(QtObjectPtr other){
	return new QStyleOptionRubberBand(*static_cast<QStyleOptionRubberBand*>(other));
}

