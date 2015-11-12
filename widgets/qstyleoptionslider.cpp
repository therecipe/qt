#include "qstyleoptionslider.h"
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionSlider>
#include "_cgo_export.h"

class MyQStyleOptionSlider: public QStyleOptionSlider {
public:
};

void* QStyleOptionSlider_NewQStyleOptionSlider(){
	return new QStyleOptionSlider();
}

void* QStyleOptionSlider_NewQStyleOptionSlider2(void* other){
	return new QStyleOptionSlider(*static_cast<QStyleOptionSlider*>(other));
}

