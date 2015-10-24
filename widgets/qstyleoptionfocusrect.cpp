#include "qstyleoptionfocusrect.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QStyleOptionFocusRect>
#include "_cgo_export.h"

class MyQStyleOptionFocusRect: public QStyleOptionFocusRect {
public:
};

QtObjectPtr QStyleOptionFocusRect_NewQStyleOptionFocusRect(){
	return new QStyleOptionFocusRect();
}

QtObjectPtr QStyleOptionFocusRect_NewQStyleOptionFocusRect2(QtObjectPtr other){
	return new QStyleOptionFocusRect(*static_cast<QStyleOptionFocusRect*>(other));
}

