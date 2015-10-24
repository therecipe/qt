#include "qdomcdatasection.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomCDATASection>
#include "_cgo_export.h"

class MyQDomCDATASection: public QDomCDATASection {
public:
};

QtObjectPtr QDomCDATASection_NewQDomCDATASection(){
	return new QDomCDATASection();
}

QtObjectPtr QDomCDATASection_NewQDomCDATASection2(QtObjectPtr x){
	return new QDomCDATASection(*static_cast<QDomCDATASection*>(x));
}

int QDomCDATASection_NodeType(QtObjectPtr ptr){
	return static_cast<QDomCDATASection*>(ptr)->nodeType();
}

