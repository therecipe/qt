#include "qdomnode.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextStream>
#include <QString>
#include <QDomNode>
#include "_cgo_export.h"

class MyQDomNode: public QDomNode {
public:
};

QtObjectPtr QDomNode_NewQDomNode(){
	return new QDomNode();
}

QtObjectPtr QDomNode_NewQDomNode2(QtObjectPtr n){
	return new QDomNode(*static_cast<QDomNode*>(n));
}

void QDomNode_Clear(QtObjectPtr ptr){
	static_cast<QDomNode*>(ptr)->clear();
}

int QDomNode_ColumnNumber(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->columnNumber();
}

int QDomNode_HasAttributes(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->hasAttributes();
}

int QDomNode_HasChildNodes(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->hasChildNodes();
}

int QDomNode_IsAttr(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isAttr();
}

int QDomNode_IsCDATASection(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isCDATASection();
}

int QDomNode_IsCharacterData(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isCharacterData();
}

int QDomNode_IsComment(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isComment();
}

int QDomNode_IsDocument(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isDocument();
}

int QDomNode_IsDocumentFragment(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isDocumentFragment();
}

int QDomNode_IsDocumentType(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isDocumentType();
}

int QDomNode_IsElement(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isElement();
}

int QDomNode_IsEntity(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isEntity();
}

int QDomNode_IsEntityReference(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isEntityReference();
}

int QDomNode_IsNotation(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isNotation();
}

int QDomNode_IsNull(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isNull();
}

int QDomNode_IsProcessingInstruction(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isProcessingInstruction();
}

int QDomNode_IsSupported(QtObjectPtr ptr, char* feature, char* version){
	return static_cast<QDomNode*>(ptr)->isSupported(QString(feature), QString(version));
}

int QDomNode_IsText(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->isText();
}

int QDomNode_LineNumber(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->lineNumber();
}

char* QDomNode_LocalName(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->localName().toUtf8().data();
}

char* QDomNode_NamespaceURI(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->namespaceURI().toUtf8().data();
}

char* QDomNode_NodeName(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->nodeName().toUtf8().data();
}

int QDomNode_NodeType(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->nodeType();
}

char* QDomNode_NodeValue(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->nodeValue().toUtf8().data();
}

void QDomNode_Normalize(QtObjectPtr ptr){
	static_cast<QDomNode*>(ptr)->normalize();
}

char* QDomNode_Prefix(QtObjectPtr ptr){
	return static_cast<QDomNode*>(ptr)->prefix().toUtf8().data();
}

void QDomNode_Save(QtObjectPtr ptr, QtObjectPtr stream, int indent, int encodingPolicy){
	static_cast<QDomNode*>(ptr)->save(*static_cast<QTextStream*>(stream), indent, static_cast<QDomNode::EncodingPolicy>(encodingPolicy));
}

void QDomNode_SetNodeValue(QtObjectPtr ptr, char* v){
	static_cast<QDomNode*>(ptr)->setNodeValue(QString(v));
}

void QDomNode_SetPrefix(QtObjectPtr ptr, char* pre){
	static_cast<QDomNode*>(ptr)->setPrefix(QString(pre));
}

void QDomNode_DestroyQDomNode(QtObjectPtr ptr){
	static_cast<QDomNode*>(ptr)->~QDomNode();
}

