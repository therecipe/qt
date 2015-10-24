#include "qdomdocumentfragment.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomDocument>
#include <QDomDocumentFragment>
#include "_cgo_export.h"

class MyQDomDocumentFragment: public QDomDocumentFragment {
public:
};

QtObjectPtr QDomDocumentFragment_NewQDomDocumentFragment(){
	return new QDomDocumentFragment();
}

QtObjectPtr QDomDocumentFragment_NewQDomDocumentFragment2(QtObjectPtr x){
	return new QDomDocumentFragment(*static_cast<QDomDocumentFragment*>(x));
}

int QDomDocumentFragment_NodeType(QtObjectPtr ptr){
	return static_cast<QDomDocumentFragment*>(ptr)->nodeType();
}

