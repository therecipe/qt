#include "qstyleoptionframe.h"
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOptionFrame>
#include "_cgo_export.h"

class MyQStyleOptionFrame: public QStyleOptionFrame {
public:
};

void* QStyleOptionFrame_NewQStyleOptionFrame(){
	return new QStyleOptionFrame();
}

void* QStyleOptionFrame_NewQStyleOptionFrame2(void* other){
	return new QStyleOptionFrame(*static_cast<QStyleOptionFrame*>(other));
}

