#include "qstyleoptionslider.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionSlider>
#include "_cgo_export.h"

class MyQStyleOptionSlider: public QStyleOptionSlider {
public:
};

QtObjectPtr QStyleOptionSlider_NewQStyleOptionSlider(){
	return new QStyleOptionSlider();
}

QtObjectPtr QStyleOptionSlider_NewQStyleOptionSlider2(QtObjectPtr other){
	return new QStyleOptionSlider(*static_cast<QStyleOptionSlider*>(other));
}

