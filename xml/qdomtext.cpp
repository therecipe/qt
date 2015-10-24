#include "qdomtext.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomText>
#include "_cgo_export.h"

class MyQDomText: public QDomText {
public:
};

QtObjectPtr QDomText_NewQDomText(){
	return new QDomText();
}

QtObjectPtr QDomText_NewQDomText2(QtObjectPtr x){
	return new QDomText(*static_cast<QDomText*>(x));
}

int QDomText_NodeType(QtObjectPtr ptr){
	return static_cast<QDomText*>(ptr)->nodeType();
}

