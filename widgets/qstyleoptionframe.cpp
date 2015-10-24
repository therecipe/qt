#include "qstyleoptionframe.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionFrame>
#include "_cgo_export.h"

class MyQStyleOptionFrame: public QStyleOptionFrame {
public:
};

QtObjectPtr QStyleOptionFrame_NewQStyleOptionFrame(){
	return new QStyleOptionFrame();
}

QtObjectPtr QStyleOptionFrame_NewQStyleOptionFrame2(QtObjectPtr other){
	return new QStyleOptionFrame(*static_cast<QStyleOptionFrame*>(other));
}

