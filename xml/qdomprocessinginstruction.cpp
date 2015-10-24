#include "qdomprocessinginstruction.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomProcessingInstruction>
#include "_cgo_export.h"

class MyQDomProcessingInstruction: public QDomProcessingInstruction {
public:
};

QtObjectPtr QDomProcessingInstruction_NewQDomProcessingInstruction(){
	return new QDomProcessingInstruction();
}

QtObjectPtr QDomProcessingInstruction_NewQDomProcessingInstruction2(QtObjectPtr x){
	return new QDomProcessingInstruction(*static_cast<QDomProcessingInstruction*>(x));
}

char* QDomProcessingInstruction_Data(QtObjectPtr ptr){
	return static_cast<QDomProcessingInstruction*>(ptr)->data().toUtf8().data();
}

int QDomProcessingInstruction_NodeType(QtObjectPtr ptr){
	return static_cast<QDomProcessingInstruction*>(ptr)->nodeType();
}

void QDomProcessingInstruction_SetData(QtObjectPtr ptr, char* d){
	static_cast<QDomProcessingInstruction*>(ptr)->setData(QString(d));
}

char* QDomProcessingInstruction_Target(QtObjectPtr ptr){
	return static_cast<QDomProcessingInstruction*>(ptr)->target().toUtf8().data();
}

