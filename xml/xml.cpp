#define protected public

#include "xml.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QDomAttr>
#include <QDomCDATASection>
#include <QDomCharacterData>
#include <QDomComment>
#include <QDomDocument>
#include <QDomDocumentFragment>
#include <QDomDocumentType>
#include <QDomElement>
#include <QDomEntity>
#include <QDomEntityReference>
#include <QDomImplementation>
#include <QDomNamedNodeMap>
#include <QDomNode>
#include <QDomNodeList>
#include <QDomNotation>
#include <QDomProcessingInstruction>
#include <QDomText>
#include <QIODevice>
#include <QLatin1String>
#include <QString>
#include <QTextStream>
#include <QXmlAttributes>
#include <QXmlContentHandler>
#include <QXmlDTDHandler>
#include <QXmlDeclHandler>
#include <QXmlDefaultHandler>
#include <QXmlEntityResolver>
#include <QXmlErrorHandler>
#include <QXmlInputSource>
#include <QXmlLexicalHandler>
#include <QXmlLocator>
#include <QXmlName>
#include <QXmlNamespaceSupport>
#include <QXmlParseException>
#include <QXmlReader>
#include <QXmlSimpleReader>

void* QDomAttr_NewQDomAttr(){
	return new QDomAttr();
}

void* QDomAttr_NewQDomAttr2(void* x){
	return new QDomAttr(*static_cast<QDomAttr*>(x));
}

char* QDomAttr_Name(void* ptr){
	return static_cast<QDomAttr*>(ptr)->name().toUtf8().data();
}

int QDomAttr_NodeType(void* ptr){
	return static_cast<QDomAttr*>(ptr)->nodeType();
}

void QDomAttr_SetValue(void* ptr, char* v){
	static_cast<QDomAttr*>(ptr)->setValue(QString(v));
}

int QDomAttr_Specified(void* ptr){
	return static_cast<QDomAttr*>(ptr)->specified();
}

char* QDomAttr_Value(void* ptr){
	return static_cast<QDomAttr*>(ptr)->value().toUtf8().data();
}

void* QDomCDATASection_NewQDomCDATASection(){
	return new QDomCDATASection();
}

void* QDomCDATASection_NewQDomCDATASection2(void* x){
	return new QDomCDATASection(*static_cast<QDomCDATASection*>(x));
}

int QDomCDATASection_NodeType(void* ptr){
	return static_cast<QDomCDATASection*>(ptr)->nodeType();
}

void* QDomCharacterData_NewQDomCharacterData(){
	return new QDomCharacterData();
}

void* QDomCharacterData_NewQDomCharacterData2(void* x){
	return new QDomCharacterData(*static_cast<QDomCharacterData*>(x));
}

void QDomCharacterData_AppendData(void* ptr, char* arg){
	static_cast<QDomCharacterData*>(ptr)->appendData(QString(arg));
}

char* QDomCharacterData_Data(void* ptr){
	return static_cast<QDomCharacterData*>(ptr)->data().toUtf8().data();
}

int QDomCharacterData_Length(void* ptr){
	return static_cast<QDomCharacterData*>(ptr)->length();
}

int QDomCharacterData_NodeType(void* ptr){
	return static_cast<QDomCharacterData*>(ptr)->nodeType();
}

void QDomCharacterData_SetData(void* ptr, char* v){
	static_cast<QDomCharacterData*>(ptr)->setData(QString(v));
}

void* QDomComment_NewQDomComment(){
	return new QDomComment();
}

void* QDomComment_NewQDomComment2(void* x){
	return new QDomComment(*static_cast<QDomComment*>(x));
}

int QDomComment_NodeType(void* ptr){
	return static_cast<QDomComment*>(ptr)->nodeType();
}

void* QDomDocument_NewQDomDocument(){
	return new QDomDocument();
}

void* QDomDocument_NewQDomDocument4(void* x){
	return new QDomDocument(*static_cast<QDomDocument*>(x));
}

void* QDomDocument_NewQDomDocument3(void* doctype){
	return new QDomDocument(*static_cast<QDomDocumentType*>(doctype));
}

void* QDomDocument_NewQDomDocument2(char* name){
	return new QDomDocument(QString(name));
}

int QDomDocument_NodeType(void* ptr){
	return static_cast<QDomDocument*>(ptr)->nodeType();
}

int QDomDocument_SetContent7(void* ptr, void* dev, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QIODevice*>(dev), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent3(void* ptr, void* dev, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QIODevice*>(dev), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent8(void* ptr, void* source, void* reader, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QXmlInputSource*>(source), static_cast<QXmlReader*>(reader), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent4(void* ptr, void* source, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QXmlInputSource*>(source), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent6(void* ptr, void* buffer, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(*static_cast<QByteArray*>(buffer), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent(void* ptr, void* data, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(*static_cast<QByteArray*>(data), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent5(void* ptr, char* text, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(QString(text), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent2(void* ptr, char* text, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(QString(text), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

void* QDomDocument_ToByteArray(void* ptr, int indent){
	return new QByteArray(static_cast<QDomDocument*>(ptr)->toByteArray(indent));
}

char* QDomDocument_ToString(void* ptr, int indent){
	return static_cast<QDomDocument*>(ptr)->toString(indent).toUtf8().data();
}

void QDomDocument_DestroyQDomDocument(void* ptr){
	static_cast<QDomDocument*>(ptr)->~QDomDocument();
}

void* QDomDocumentFragment_NewQDomDocumentFragment(){
	return new QDomDocumentFragment();
}

void* QDomDocumentFragment_NewQDomDocumentFragment2(void* x){
	return new QDomDocumentFragment(*static_cast<QDomDocumentFragment*>(x));
}

int QDomDocumentFragment_NodeType(void* ptr){
	return static_cast<QDomDocumentFragment*>(ptr)->nodeType();
}

void* QDomDocumentType_NewQDomDocumentType(){
	return new QDomDocumentType();
}

void* QDomDocumentType_NewQDomDocumentType2(void* n){
	return new QDomDocumentType(*static_cast<QDomDocumentType*>(n));
}

char* QDomDocumentType_InternalSubset(void* ptr){
	return static_cast<QDomDocumentType*>(ptr)->internalSubset().toUtf8().data();
}

char* QDomDocumentType_Name(void* ptr){
	return static_cast<QDomDocumentType*>(ptr)->name().toUtf8().data();
}

int QDomDocumentType_NodeType(void* ptr){
	return static_cast<QDomDocumentType*>(ptr)->nodeType();
}

char* QDomDocumentType_PublicId(void* ptr){
	return static_cast<QDomDocumentType*>(ptr)->publicId().toUtf8().data();
}

char* QDomDocumentType_SystemId(void* ptr){
	return static_cast<QDomDocumentType*>(ptr)->systemId().toUtf8().data();
}

void* QDomElement_NewQDomElement(){
	return new QDomElement();
}

void* QDomElement_NewQDomElement2(void* x){
	return new QDomElement(*static_cast<QDomElement*>(x));
}

char* QDomElement_Attribute(void* ptr, char* name, char* defValue){
	return static_cast<QDomElement*>(ptr)->attribute(QString(name), QString(defValue)).toUtf8().data();
}

char* QDomElement_AttributeNS(void* ptr, char* nsURI, char* localName, char* defValue){
	return static_cast<QDomElement*>(ptr)->attributeNS(QString(nsURI), QString(localName), QString(defValue)).toUtf8().data();
}

int QDomElement_HasAttribute(void* ptr, char* name){
	return static_cast<QDomElement*>(ptr)->hasAttribute(QString(name));
}

int QDomElement_HasAttributeNS(void* ptr, char* nsURI, char* localName){
	return static_cast<QDomElement*>(ptr)->hasAttributeNS(QString(nsURI), QString(localName));
}

int QDomElement_NodeType(void* ptr){
	return static_cast<QDomElement*>(ptr)->nodeType();
}

void QDomElement_RemoveAttribute(void* ptr, char* name){
	static_cast<QDomElement*>(ptr)->removeAttribute(QString(name));
}

void QDomElement_RemoveAttributeNS(void* ptr, char* nsURI, char* localName){
	static_cast<QDomElement*>(ptr)->removeAttributeNS(QString(nsURI), QString(localName));
}

void QDomElement_SetAttribute(void* ptr, char* name, char* value){
	static_cast<QDomElement*>(ptr)->setAttribute(QString(name), QString(value));
}

void QDomElement_SetAttribute2(void* ptr, char* name, int value){
	static_cast<QDomElement*>(ptr)->setAttribute(QString(name), value);
}

void QDomElement_SetAttributeNS(void* ptr, char* nsURI, char* qName, char* value){
	static_cast<QDomElement*>(ptr)->setAttributeNS(QString(nsURI), QString(qName), QString(value));
}

void QDomElement_SetAttributeNS2(void* ptr, char* nsURI, char* qName, int value){
	static_cast<QDomElement*>(ptr)->setAttributeNS(QString(nsURI), QString(qName), value);
}

void QDomElement_SetTagName(void* ptr, char* name){
	static_cast<QDomElement*>(ptr)->setTagName(QString(name));
}

char* QDomElement_TagName(void* ptr){
	return static_cast<QDomElement*>(ptr)->tagName().toUtf8().data();
}

char* QDomElement_Text(void* ptr){
	return static_cast<QDomElement*>(ptr)->text().toUtf8().data();
}

void* QDomEntity_NewQDomEntity(){
	return new QDomEntity();
}

void* QDomEntity_NewQDomEntity2(void* x){
	return new QDomEntity(*static_cast<QDomEntity*>(x));
}

int QDomEntity_NodeType(void* ptr){
	return static_cast<QDomEntity*>(ptr)->nodeType();
}

char* QDomEntity_NotationName(void* ptr){
	return static_cast<QDomEntity*>(ptr)->notationName().toUtf8().data();
}

char* QDomEntity_PublicId(void* ptr){
	return static_cast<QDomEntity*>(ptr)->publicId().toUtf8().data();
}

char* QDomEntity_SystemId(void* ptr){
	return static_cast<QDomEntity*>(ptr)->systemId().toUtf8().data();
}

void* QDomEntityReference_NewQDomEntityReference(){
	return new QDomEntityReference();
}

void* QDomEntityReference_NewQDomEntityReference2(void* x){
	return new QDomEntityReference(*static_cast<QDomEntityReference*>(x));
}

int QDomEntityReference_NodeType(void* ptr){
	return static_cast<QDomEntityReference*>(ptr)->nodeType();
}

void* QDomImplementation_NewQDomImplementation(){
	return new QDomImplementation();
}

void* QDomImplementation_NewQDomImplementation2(void* x){
	return new QDomImplementation(*static_cast<QDomImplementation*>(x));
}

int QDomImplementation_HasFeature(void* ptr, char* feature, char* version){
	return static_cast<QDomImplementation*>(ptr)->hasFeature(QString(feature), QString(version));
}

int QDomImplementation_QDomImplementation_InvalidDataPolicy(){
	return QDomImplementation::invalidDataPolicy();
}

int QDomImplementation_IsNull(void* ptr){
	return static_cast<QDomImplementation*>(ptr)->isNull();
}

void QDomImplementation_QDomImplementation_SetInvalidDataPolicy(int policy){
	QDomImplementation::setInvalidDataPolicy(static_cast<QDomImplementation::InvalidDataPolicy>(policy));
}

void QDomImplementation_DestroyQDomImplementation(void* ptr){
	static_cast<QDomImplementation*>(ptr)->~QDomImplementation();
}

void* QDomNamedNodeMap_NewQDomNamedNodeMap(){
	return new QDomNamedNodeMap();
}

void* QDomNamedNodeMap_NewQDomNamedNodeMap2(void* n){
	return new QDomNamedNodeMap(*static_cast<QDomNamedNodeMap*>(n));
}

int QDomNamedNodeMap_Contains(void* ptr, char* name){
	return static_cast<QDomNamedNodeMap*>(ptr)->contains(QString(name));
}

int QDomNamedNodeMap_Count(void* ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->count();
}

int QDomNamedNodeMap_IsEmpty(void* ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->isEmpty();
}

int QDomNamedNodeMap_Length(void* ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->length();
}

int QDomNamedNodeMap_Size(void* ptr){
	return static_cast<QDomNamedNodeMap*>(ptr)->size();
}

void QDomNamedNodeMap_DestroyQDomNamedNodeMap(void* ptr){
	static_cast<QDomNamedNodeMap*>(ptr)->~QDomNamedNodeMap();
}

void* QDomNode_NewQDomNode(){
	return new QDomNode();
}

void* QDomNode_NewQDomNode2(void* n){
	return new QDomNode(*static_cast<QDomNode*>(n));
}

void QDomNode_Clear(void* ptr){
	static_cast<QDomNode*>(ptr)->clear();
}

int QDomNode_ColumnNumber(void* ptr){
	return static_cast<QDomNode*>(ptr)->columnNumber();
}

int QDomNode_HasAttributes(void* ptr){
	return static_cast<QDomNode*>(ptr)->hasAttributes();
}

int QDomNode_HasChildNodes(void* ptr){
	return static_cast<QDomNode*>(ptr)->hasChildNodes();
}

int QDomNode_IsAttr(void* ptr){
	return static_cast<QDomNode*>(ptr)->isAttr();
}

int QDomNode_IsCDATASection(void* ptr){
	return static_cast<QDomNode*>(ptr)->isCDATASection();
}

int QDomNode_IsCharacterData(void* ptr){
	return static_cast<QDomNode*>(ptr)->isCharacterData();
}

int QDomNode_IsComment(void* ptr){
	return static_cast<QDomNode*>(ptr)->isComment();
}

int QDomNode_IsDocument(void* ptr){
	return static_cast<QDomNode*>(ptr)->isDocument();
}

int QDomNode_IsDocumentFragment(void* ptr){
	return static_cast<QDomNode*>(ptr)->isDocumentFragment();
}

int QDomNode_IsDocumentType(void* ptr){
	return static_cast<QDomNode*>(ptr)->isDocumentType();
}

int QDomNode_IsElement(void* ptr){
	return static_cast<QDomNode*>(ptr)->isElement();
}

int QDomNode_IsEntity(void* ptr){
	return static_cast<QDomNode*>(ptr)->isEntity();
}

int QDomNode_IsEntityReference(void* ptr){
	return static_cast<QDomNode*>(ptr)->isEntityReference();
}

int QDomNode_IsNotation(void* ptr){
	return static_cast<QDomNode*>(ptr)->isNotation();
}

int QDomNode_IsNull(void* ptr){
	return static_cast<QDomNode*>(ptr)->isNull();
}

int QDomNode_IsProcessingInstruction(void* ptr){
	return static_cast<QDomNode*>(ptr)->isProcessingInstruction();
}

int QDomNode_IsSupported(void* ptr, char* feature, char* version){
	return static_cast<QDomNode*>(ptr)->isSupported(QString(feature), QString(version));
}

int QDomNode_IsText(void* ptr){
	return static_cast<QDomNode*>(ptr)->isText();
}

int QDomNode_LineNumber(void* ptr){
	return static_cast<QDomNode*>(ptr)->lineNumber();
}

char* QDomNode_LocalName(void* ptr){
	return static_cast<QDomNode*>(ptr)->localName().toUtf8().data();
}

char* QDomNode_NamespaceURI(void* ptr){
	return static_cast<QDomNode*>(ptr)->namespaceURI().toUtf8().data();
}

char* QDomNode_NodeName(void* ptr){
	return static_cast<QDomNode*>(ptr)->nodeName().toUtf8().data();
}

int QDomNode_NodeType(void* ptr){
	return static_cast<QDomNode*>(ptr)->nodeType();
}

char* QDomNode_NodeValue(void* ptr){
	return static_cast<QDomNode*>(ptr)->nodeValue().toUtf8().data();
}

void QDomNode_Normalize(void* ptr){
	static_cast<QDomNode*>(ptr)->normalize();
}

char* QDomNode_Prefix(void* ptr){
	return static_cast<QDomNode*>(ptr)->prefix().toUtf8().data();
}

void QDomNode_Save(void* ptr, void* stream, int indent, int encodingPolicy){
	static_cast<QDomNode*>(ptr)->save(*static_cast<QTextStream*>(stream), indent, static_cast<QDomNode::EncodingPolicy>(encodingPolicy));
}

void QDomNode_SetNodeValue(void* ptr, char* v){
	static_cast<QDomNode*>(ptr)->setNodeValue(QString(v));
}

void QDomNode_SetPrefix(void* ptr, char* pre){
	static_cast<QDomNode*>(ptr)->setPrefix(QString(pre));
}

void QDomNode_DestroyQDomNode(void* ptr){
	static_cast<QDomNode*>(ptr)->~QDomNode();
}

void* QDomNodeList_NewQDomNodeList(){
	return new QDomNodeList();
}

void* QDomNodeList_NewQDomNodeList2(void* n){
	return new QDomNodeList(*static_cast<QDomNodeList*>(n));
}

int QDomNodeList_Count(void* ptr){
	return static_cast<QDomNodeList*>(ptr)->count();
}

int QDomNodeList_IsEmpty(void* ptr){
	return static_cast<QDomNodeList*>(ptr)->isEmpty();
}

int QDomNodeList_Length(void* ptr){
	return static_cast<QDomNodeList*>(ptr)->length();
}

int QDomNodeList_Size(void* ptr){
	return static_cast<QDomNodeList*>(ptr)->size();
}

void QDomNodeList_DestroyQDomNodeList(void* ptr){
	static_cast<QDomNodeList*>(ptr)->~QDomNodeList();
}

void* QDomNotation_NewQDomNotation(){
	return new QDomNotation();
}

void* QDomNotation_NewQDomNotation2(void* x){
	return new QDomNotation(*static_cast<QDomNotation*>(x));
}

int QDomNotation_NodeType(void* ptr){
	return static_cast<QDomNotation*>(ptr)->nodeType();
}

char* QDomNotation_PublicId(void* ptr){
	return static_cast<QDomNotation*>(ptr)->publicId().toUtf8().data();
}

char* QDomNotation_SystemId(void* ptr){
	return static_cast<QDomNotation*>(ptr)->systemId().toUtf8().data();
}

void* QDomProcessingInstruction_NewQDomProcessingInstruction(){
	return new QDomProcessingInstruction();
}

void* QDomProcessingInstruction_NewQDomProcessingInstruction2(void* x){
	return new QDomProcessingInstruction(*static_cast<QDomProcessingInstruction*>(x));
}

char* QDomProcessingInstruction_Data(void* ptr){
	return static_cast<QDomProcessingInstruction*>(ptr)->data().toUtf8().data();
}

int QDomProcessingInstruction_NodeType(void* ptr){
	return static_cast<QDomProcessingInstruction*>(ptr)->nodeType();
}

void QDomProcessingInstruction_SetData(void* ptr, char* d){
	static_cast<QDomProcessingInstruction*>(ptr)->setData(QString(d));
}

char* QDomProcessingInstruction_Target(void* ptr){
	return static_cast<QDomProcessingInstruction*>(ptr)->target().toUtf8().data();
}

void* QDomText_NewQDomText(){
	return new QDomText();
}

void* QDomText_NewQDomText2(void* x){
	return new QDomText(*static_cast<QDomText*>(x));
}

int QDomText_NodeType(void* ptr){
	return static_cast<QDomText*>(ptr)->nodeType();
}

class MyQXmlAttributes: public QXmlAttributes {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQXmlAttributes() : QXmlAttributes() {};
};

void* QXmlAttributes_NewQXmlAttributes(){
	return new MyQXmlAttributes();
}

void QXmlAttributes_DestroyQXmlAttributes(void* ptr){
	static_cast<QXmlAttributes*>(ptr)->~QXmlAttributes();
}

void QXmlAttributes_Append(void* ptr, char* qName, char* uri, char* localPart, char* value){
	static_cast<QXmlAttributes*>(ptr)->append(QString(qName), QString(uri), QString(localPart), QString(value));
}

void QXmlAttributes_Clear(void* ptr){
	static_cast<QXmlAttributes*>(ptr)->clear();
}

int QXmlAttributes_Count(void* ptr){
	return static_cast<QXmlAttributes*>(ptr)->count();
}

int QXmlAttributes_Index2(void* ptr, void* qName){
	return static_cast<QXmlAttributes*>(ptr)->index(*static_cast<QLatin1String*>(qName));
}

int QXmlAttributes_Index(void* ptr, char* qName){
	return static_cast<QXmlAttributes*>(ptr)->index(QString(qName));
}

int QXmlAttributes_Index3(void* ptr, char* uri, char* localPart){
	return static_cast<QXmlAttributes*>(ptr)->index(QString(uri), QString(localPart));
}

int QXmlAttributes_Length(void* ptr){
	return static_cast<QXmlAttributes*>(ptr)->length();
}

char* QXmlAttributes_LocalName(void* ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->localName(index).toUtf8().data();
}

char* QXmlAttributes_QName(void* ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->qName(index).toUtf8().data();
}

char* QXmlAttributes_Type2(void* ptr, char* qName){
	return static_cast<QXmlAttributes*>(ptr)->type(QString(qName)).toUtf8().data();
}

char* QXmlAttributes_Type3(void* ptr, char* uri, char* localName){
	return static_cast<QXmlAttributes*>(ptr)->type(QString(uri), QString(localName)).toUtf8().data();
}

char* QXmlAttributes_Type(void* ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->type(index).toUtf8().data();
}

char* QXmlAttributes_Uri(void* ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->uri(index).toUtf8().data();
}

char* QXmlAttributes_Value3(void* ptr, void* qName){
	return static_cast<QXmlAttributes*>(ptr)->value(*static_cast<QLatin1String*>(qName)).toUtf8().data();
}

char* QXmlAttributes_Value2(void* ptr, char* qName){
	return static_cast<QXmlAttributes*>(ptr)->value(QString(qName)).toUtf8().data();
}

char* QXmlAttributes_Value4(void* ptr, char* uri, char* localName){
	return static_cast<QXmlAttributes*>(ptr)->value(QString(uri), QString(localName)).toUtf8().data();
}

char* QXmlAttributes_Value(void* ptr, int index){
	return static_cast<QXmlAttributes*>(ptr)->value(index).toUtf8().data();
}

char* QXmlAttributes_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlAttributes*>(static_cast<QXmlAttributes*>(ptr))) {
		return static_cast<MyQXmlAttributes*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlAttributes_BASE").toUtf8().data();
}

void QXmlAttributes_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlAttributes*>(static_cast<QXmlAttributes*>(ptr))) {
		static_cast<MyQXmlAttributes*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQXmlContentHandler: public QXmlContentHandler {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

int QXmlContentHandler_Characters(void* ptr, char* ch){
	return static_cast<QXmlContentHandler*>(ptr)->characters(QString(ch));
}

int QXmlContentHandler_EndDocument(void* ptr){
	return static_cast<QXmlContentHandler*>(ptr)->endDocument();
}

int QXmlContentHandler_EndElement(void* ptr, char* namespaceURI, char* localName, char* qName){
	return static_cast<QXmlContentHandler*>(ptr)->endElement(QString(namespaceURI), QString(localName), QString(qName));
}

int QXmlContentHandler_EndPrefixMapping(void* ptr, char* prefix){
	return static_cast<QXmlContentHandler*>(ptr)->endPrefixMapping(QString(prefix));
}

char* QXmlContentHandler_ErrorString(void* ptr){
	return static_cast<QXmlContentHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlContentHandler_IgnorableWhitespace(void* ptr, char* ch){
	return static_cast<QXmlContentHandler*>(ptr)->ignorableWhitespace(QString(ch));
}

int QXmlContentHandler_ProcessingInstruction(void* ptr, char* target, char* data){
	return static_cast<QXmlContentHandler*>(ptr)->processingInstruction(QString(target), QString(data));
}

void QXmlContentHandler_SetDocumentLocator(void* ptr, void* locator){
	static_cast<QXmlContentHandler*>(ptr)->setDocumentLocator(static_cast<QXmlLocator*>(locator));
}

int QXmlContentHandler_SkippedEntity(void* ptr, char* name){
	return static_cast<QXmlContentHandler*>(ptr)->skippedEntity(QString(name));
}

int QXmlContentHandler_StartDocument(void* ptr){
	return static_cast<QXmlContentHandler*>(ptr)->startDocument();
}

int QXmlContentHandler_StartElement(void* ptr, char* namespaceURI, char* localName, char* qName, void* atts){
	return static_cast<QXmlContentHandler*>(ptr)->startElement(QString(namespaceURI), QString(localName), QString(qName), *static_cast<QXmlAttributes*>(atts));
}

int QXmlContentHandler_StartPrefixMapping(void* ptr, char* prefix, char* uri){
	return static_cast<QXmlContentHandler*>(ptr)->startPrefixMapping(QString(prefix), QString(uri));
}

void QXmlContentHandler_DestroyQXmlContentHandler(void* ptr){
	static_cast<QXmlContentHandler*>(ptr)->~QXmlContentHandler();
}

char* QXmlContentHandler_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlContentHandler*>(static_cast<QXmlContentHandler*>(ptr))) {
		return static_cast<MyQXmlContentHandler*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlContentHandler_BASE").toUtf8().data();
}

void QXmlContentHandler_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlContentHandler*>(static_cast<QXmlContentHandler*>(ptr))) {
		static_cast<MyQXmlContentHandler*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQXmlDTDHandler: public QXmlDTDHandler {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

char* QXmlDTDHandler_ErrorString(void* ptr){
	return static_cast<QXmlDTDHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlDTDHandler_NotationDecl(void* ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlDTDHandler*>(ptr)->notationDecl(QString(name), QString(publicId), QString(systemId));
}

int QXmlDTDHandler_UnparsedEntityDecl(void* ptr, char* name, char* publicId, char* systemId, char* notationName){
	return static_cast<QXmlDTDHandler*>(ptr)->unparsedEntityDecl(QString(name), QString(publicId), QString(systemId), QString(notationName));
}

void QXmlDTDHandler_DestroyQXmlDTDHandler(void* ptr){
	static_cast<QXmlDTDHandler*>(ptr)->~QXmlDTDHandler();
}

char* QXmlDTDHandler_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlDTDHandler*>(static_cast<QXmlDTDHandler*>(ptr))) {
		return static_cast<MyQXmlDTDHandler*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlDTDHandler_BASE").toUtf8().data();
}

void QXmlDTDHandler_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlDTDHandler*>(static_cast<QXmlDTDHandler*>(ptr))) {
		static_cast<MyQXmlDTDHandler*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQXmlDeclHandler: public QXmlDeclHandler {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

int QXmlDeclHandler_AttributeDecl(void* ptr, char* eName, char* aName, char* ty, char* valueDefault, char* value){
	return static_cast<QXmlDeclHandler*>(ptr)->attributeDecl(QString(eName), QString(aName), QString(ty), QString(valueDefault), QString(value));
}

char* QXmlDeclHandler_ErrorString(void* ptr){
	return static_cast<QXmlDeclHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlDeclHandler_ExternalEntityDecl(void* ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlDeclHandler*>(ptr)->externalEntityDecl(QString(name), QString(publicId), QString(systemId));
}

int QXmlDeclHandler_InternalEntityDecl(void* ptr, char* name, char* value){
	return static_cast<QXmlDeclHandler*>(ptr)->internalEntityDecl(QString(name), QString(value));
}

void QXmlDeclHandler_DestroyQXmlDeclHandler(void* ptr){
	static_cast<QXmlDeclHandler*>(ptr)->~QXmlDeclHandler();
}

char* QXmlDeclHandler_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlDeclHandler*>(static_cast<QXmlDeclHandler*>(ptr))) {
		return static_cast<MyQXmlDeclHandler*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlDeclHandler_BASE").toUtf8().data();
}

void QXmlDeclHandler_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlDeclHandler*>(static_cast<QXmlDeclHandler*>(ptr))) {
		static_cast<MyQXmlDeclHandler*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQXmlDefaultHandler: public QXmlDefaultHandler {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQXmlDefaultHandler() : QXmlDefaultHandler() {};
	void setDocumentLocator(QXmlLocator * locator) { callbackQXmlDefaultHandlerSetDocumentLocator(this, this->objectNameAbs().toUtf8().data(), locator); };
};

void* QXmlDefaultHandler_NewQXmlDefaultHandler(){
	return new MyQXmlDefaultHandler();
}

void QXmlDefaultHandler_DestroyQXmlDefaultHandler(void* ptr){
	static_cast<QXmlDefaultHandler*>(ptr)->~QXmlDefaultHandler();
}

int QXmlDefaultHandler_AttributeDecl(void* ptr, char* eName, char* aName, char* ty, char* valueDefault, char* value){
	return static_cast<QXmlDefaultHandler*>(ptr)->attributeDecl(QString(eName), QString(aName), QString(ty), QString(valueDefault), QString(value));
}

int QXmlDefaultHandler_Characters(void* ptr, char* ch){
	return static_cast<QXmlDefaultHandler*>(ptr)->characters(QString(ch));
}

int QXmlDefaultHandler_Comment(void* ptr, char* ch){
	return static_cast<QXmlDefaultHandler*>(ptr)->comment(QString(ch));
}

int QXmlDefaultHandler_EndCDATA(void* ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->endCDATA();
}

int QXmlDefaultHandler_EndDTD(void* ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->endDTD();
}

int QXmlDefaultHandler_EndDocument(void* ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->endDocument();
}

int QXmlDefaultHandler_EndElement(void* ptr, char* namespaceURI, char* localName, char* qName){
	return static_cast<QXmlDefaultHandler*>(ptr)->endElement(QString(namespaceURI), QString(localName), QString(qName));
}

int QXmlDefaultHandler_EndEntity(void* ptr, char* name){
	return static_cast<QXmlDefaultHandler*>(ptr)->endEntity(QString(name));
}

int QXmlDefaultHandler_EndPrefixMapping(void* ptr, char* prefix){
	return static_cast<QXmlDefaultHandler*>(ptr)->endPrefixMapping(QString(prefix));
}

int QXmlDefaultHandler_Error(void* ptr, void* exception){
	return static_cast<QXmlDefaultHandler*>(ptr)->error(*static_cast<QXmlParseException*>(exception));
}

char* QXmlDefaultHandler_ErrorString(void* ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlDefaultHandler_ExternalEntityDecl(void* ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlDefaultHandler*>(ptr)->externalEntityDecl(QString(name), QString(publicId), QString(systemId));
}

int QXmlDefaultHandler_FatalError(void* ptr, void* exception){
	return static_cast<QXmlDefaultHandler*>(ptr)->fatalError(*static_cast<QXmlParseException*>(exception));
}

int QXmlDefaultHandler_IgnorableWhitespace(void* ptr, char* ch){
	return static_cast<QXmlDefaultHandler*>(ptr)->ignorableWhitespace(QString(ch));
}

int QXmlDefaultHandler_InternalEntityDecl(void* ptr, char* name, char* value){
	return static_cast<QXmlDefaultHandler*>(ptr)->internalEntityDecl(QString(name), QString(value));
}

int QXmlDefaultHandler_NotationDecl(void* ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlDefaultHandler*>(ptr)->notationDecl(QString(name), QString(publicId), QString(systemId));
}

int QXmlDefaultHandler_ProcessingInstruction(void* ptr, char* target, char* data){
	return static_cast<QXmlDefaultHandler*>(ptr)->processingInstruction(QString(target), QString(data));
}

void QXmlDefaultHandler_SetDocumentLocator(void* ptr, void* locator){
	static_cast<MyQXmlDefaultHandler*>(ptr)->setDocumentLocator(static_cast<QXmlLocator*>(locator));
}

void QXmlDefaultHandler_SetDocumentLocatorDefault(void* ptr, void* locator){
	static_cast<QXmlDefaultHandler*>(ptr)->QXmlDefaultHandler::setDocumentLocator(static_cast<QXmlLocator*>(locator));
}

int QXmlDefaultHandler_SkippedEntity(void* ptr, char* name){
	return static_cast<QXmlDefaultHandler*>(ptr)->skippedEntity(QString(name));
}

int QXmlDefaultHandler_StartCDATA(void* ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->startCDATA();
}

int QXmlDefaultHandler_StartDTD(void* ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlDefaultHandler*>(ptr)->startDTD(QString(name), QString(publicId), QString(systemId));
}

int QXmlDefaultHandler_StartDocument(void* ptr){
	return static_cast<QXmlDefaultHandler*>(ptr)->startDocument();
}

int QXmlDefaultHandler_StartElement(void* ptr, char* namespaceURI, char* localName, char* qName, void* atts){
	return static_cast<QXmlDefaultHandler*>(ptr)->startElement(QString(namespaceURI), QString(localName), QString(qName), *static_cast<QXmlAttributes*>(atts));
}

int QXmlDefaultHandler_StartEntity(void* ptr, char* name){
	return static_cast<QXmlDefaultHandler*>(ptr)->startEntity(QString(name));
}

int QXmlDefaultHandler_StartPrefixMapping(void* ptr, char* prefix, char* uri){
	return static_cast<QXmlDefaultHandler*>(ptr)->startPrefixMapping(QString(prefix), QString(uri));
}

int QXmlDefaultHandler_UnparsedEntityDecl(void* ptr, char* name, char* publicId, char* systemId, char* notationName){
	return static_cast<QXmlDefaultHandler*>(ptr)->unparsedEntityDecl(QString(name), QString(publicId), QString(systemId), QString(notationName));
}

int QXmlDefaultHandler_Warning(void* ptr, void* exception){
	return static_cast<QXmlDefaultHandler*>(ptr)->warning(*static_cast<QXmlParseException*>(exception));
}

char* QXmlDefaultHandler_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlDefaultHandler*>(static_cast<QXmlDefaultHandler*>(ptr))) {
		return static_cast<MyQXmlDefaultHandler*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlDefaultHandler_BASE").toUtf8().data();
}

void QXmlDefaultHandler_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlDefaultHandler*>(static_cast<QXmlDefaultHandler*>(ptr))) {
		static_cast<MyQXmlDefaultHandler*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQXmlEntityResolver: public QXmlEntityResolver {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

char* QXmlEntityResolver_ErrorString(void* ptr){
	return static_cast<QXmlEntityResolver*>(ptr)->errorString().toUtf8().data();
}

void QXmlEntityResolver_DestroyQXmlEntityResolver(void* ptr){
	static_cast<QXmlEntityResolver*>(ptr)->~QXmlEntityResolver();
}

char* QXmlEntityResolver_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlEntityResolver*>(static_cast<QXmlEntityResolver*>(ptr))) {
		return static_cast<MyQXmlEntityResolver*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlEntityResolver_BASE").toUtf8().data();
}

void QXmlEntityResolver_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlEntityResolver*>(static_cast<QXmlEntityResolver*>(ptr))) {
		static_cast<MyQXmlEntityResolver*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQXmlErrorHandler: public QXmlErrorHandler {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

int QXmlErrorHandler_Error(void* ptr, void* exception){
	return static_cast<QXmlErrorHandler*>(ptr)->error(*static_cast<QXmlParseException*>(exception));
}

char* QXmlErrorHandler_ErrorString(void* ptr){
	return static_cast<QXmlErrorHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlErrorHandler_FatalError(void* ptr, void* exception){
	return static_cast<QXmlErrorHandler*>(ptr)->fatalError(*static_cast<QXmlParseException*>(exception));
}

int QXmlErrorHandler_Warning(void* ptr, void* exception){
	return static_cast<QXmlErrorHandler*>(ptr)->warning(*static_cast<QXmlParseException*>(exception));
}

void QXmlErrorHandler_DestroyQXmlErrorHandler(void* ptr){
	static_cast<QXmlErrorHandler*>(ptr)->~QXmlErrorHandler();
}

char* QXmlErrorHandler_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlErrorHandler*>(static_cast<QXmlErrorHandler*>(ptr))) {
		return static_cast<MyQXmlErrorHandler*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlErrorHandler_BASE").toUtf8().data();
}

void QXmlErrorHandler_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlErrorHandler*>(static_cast<QXmlErrorHandler*>(ptr))) {
		static_cast<MyQXmlErrorHandler*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQXmlInputSource: public QXmlInputSource {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQXmlInputSource() : QXmlInputSource() {};
	MyQXmlInputSource(QIODevice *dev) : QXmlInputSource(dev) {};
	void fetchData() { callbackQXmlInputSourceFetchData(this, this->objectNameAbs().toUtf8().data()); };
	void reset() { callbackQXmlInputSourceReset(this, this->objectNameAbs().toUtf8().data()); };
	void setData(const QString & dat) { callbackQXmlInputSourceSetData(this, this->objectNameAbs().toUtf8().data(), dat.toUtf8().data()); };
};

void* QXmlInputSource_NewQXmlInputSource(){
	return new MyQXmlInputSource();
}

void* QXmlInputSource_NewQXmlInputSource2(void* dev){
	return new MyQXmlInputSource(static_cast<QIODevice*>(dev));
}

char* QXmlInputSource_Data(void* ptr){
	return static_cast<QXmlInputSource*>(ptr)->data().toUtf8().data();
}

void QXmlInputSource_FetchData(void* ptr){
	static_cast<MyQXmlInputSource*>(ptr)->fetchData();
}

void QXmlInputSource_FetchDataDefault(void* ptr){
	static_cast<QXmlInputSource*>(ptr)->QXmlInputSource::fetchData();
}

char* QXmlInputSource_FromRawData(void* ptr, void* data, int beginning){
	return static_cast<QXmlInputSource*>(ptr)->fromRawData(*static_cast<QByteArray*>(data), beginning != 0).toUtf8().data();
}

void QXmlInputSource_Reset(void* ptr){
	static_cast<MyQXmlInputSource*>(ptr)->reset();
}

void QXmlInputSource_ResetDefault(void* ptr){
	static_cast<QXmlInputSource*>(ptr)->QXmlInputSource::reset();
}

void QXmlInputSource_SetData(void* ptr, char* dat){
	static_cast<MyQXmlInputSource*>(ptr)->setData(QString(dat));
}

void QXmlInputSource_SetDataDefault(void* ptr, char* dat){
	static_cast<QXmlInputSource*>(ptr)->QXmlInputSource::setData(QString(dat));
}

void QXmlInputSource_DestroyQXmlInputSource(void* ptr){
	static_cast<QXmlInputSource*>(ptr)->~QXmlInputSource();
}

char* QXmlInputSource_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlInputSource*>(static_cast<QXmlInputSource*>(ptr))) {
		return static_cast<MyQXmlInputSource*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlInputSource_BASE").toUtf8().data();
}

void QXmlInputSource_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlInputSource*>(static_cast<QXmlInputSource*>(ptr))) {
		static_cast<MyQXmlInputSource*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQXmlLexicalHandler: public QXmlLexicalHandler {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

int QXmlLexicalHandler_Comment(void* ptr, char* ch){
	return static_cast<QXmlLexicalHandler*>(ptr)->comment(QString(ch));
}

int QXmlLexicalHandler_EndCDATA(void* ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->endCDATA();
}

int QXmlLexicalHandler_EndDTD(void* ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->endDTD();
}

int QXmlLexicalHandler_EndEntity(void* ptr, char* name){
	return static_cast<QXmlLexicalHandler*>(ptr)->endEntity(QString(name));
}

char* QXmlLexicalHandler_ErrorString(void* ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlLexicalHandler_StartCDATA(void* ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->startCDATA();
}

int QXmlLexicalHandler_StartDTD(void* ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlLexicalHandler*>(ptr)->startDTD(QString(name), QString(publicId), QString(systemId));
}

int QXmlLexicalHandler_StartEntity(void* ptr, char* name){
	return static_cast<QXmlLexicalHandler*>(ptr)->startEntity(QString(name));
}

void QXmlLexicalHandler_DestroyQXmlLexicalHandler(void* ptr){
	static_cast<QXmlLexicalHandler*>(ptr)->~QXmlLexicalHandler();
}

char* QXmlLexicalHandler_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlLexicalHandler*>(static_cast<QXmlLexicalHandler*>(ptr))) {
		return static_cast<MyQXmlLexicalHandler*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlLexicalHandler_BASE").toUtf8().data();
}

void QXmlLexicalHandler_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlLexicalHandler*>(static_cast<QXmlLexicalHandler*>(ptr))) {
		static_cast<MyQXmlLexicalHandler*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQXmlLocator: public QXmlLocator {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

int QXmlLocator_ColumnNumber(void* ptr){
	return static_cast<QXmlLocator*>(ptr)->columnNumber();
}

int QXmlLocator_LineNumber(void* ptr){
	return static_cast<QXmlLocator*>(ptr)->lineNumber();
}

void QXmlLocator_DestroyQXmlLocator(void* ptr){
	static_cast<QXmlLocator*>(ptr)->~QXmlLocator();
}

char* QXmlLocator_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlLocator*>(static_cast<QXmlLocator*>(ptr))) {
		return static_cast<MyQXmlLocator*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlLocator_BASE").toUtf8().data();
}

void QXmlLocator_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlLocator*>(static_cast<QXmlLocator*>(ptr))) {
		static_cast<MyQXmlLocator*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QXmlNamespaceSupport_NewQXmlNamespaceSupport(){
	return new QXmlNamespaceSupport();
}

void QXmlNamespaceSupport_PopContext(void* ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->popContext();
}

char* QXmlNamespaceSupport_Prefix(void* ptr, char* uri){
	return static_cast<QXmlNamespaceSupport*>(ptr)->prefix(QString(uri)).toUtf8().data();
}

char* QXmlNamespaceSupport_Prefixes(void* ptr){
	return static_cast<QXmlNamespaceSupport*>(ptr)->prefixes().join(",,,").toUtf8().data();
}

char* QXmlNamespaceSupport_Prefixes2(void* ptr, char* uri){
	return static_cast<QXmlNamespaceSupport*>(ptr)->prefixes(QString(uri)).join(",,,").toUtf8().data();
}

void QXmlNamespaceSupport_PushContext(void* ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->pushContext();
}

void QXmlNamespaceSupport_Reset(void* ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->reset();
}

void QXmlNamespaceSupport_SetPrefix(void* ptr, char* pre, char* uri){
	static_cast<QXmlNamespaceSupport*>(ptr)->setPrefix(QString(pre), QString(uri));
}

char* QXmlNamespaceSupport_Uri(void* ptr, char* prefix){
	return static_cast<QXmlNamespaceSupport*>(ptr)->uri(QString(prefix)).toUtf8().data();
}

void QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(void* ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->~QXmlNamespaceSupport();
}

void* QXmlParseException_NewQXmlParseException(char* name, int c, int l, char* p, char* s){
	return new QXmlParseException(QString(name), c, l, QString(p), QString(s));
}

void* QXmlParseException_NewQXmlParseException2(void* other){
	return new QXmlParseException(*static_cast<QXmlParseException*>(other));
}

int QXmlParseException_ColumnNumber(void* ptr){
	return static_cast<QXmlParseException*>(ptr)->columnNumber();
}

int QXmlParseException_LineNumber(void* ptr){
	return static_cast<QXmlParseException*>(ptr)->lineNumber();
}

char* QXmlParseException_Message(void* ptr){
	return static_cast<QXmlParseException*>(ptr)->message().toUtf8().data();
}

char* QXmlParseException_PublicId(void* ptr){
	return static_cast<QXmlParseException*>(ptr)->publicId().toUtf8().data();
}

char* QXmlParseException_SystemId(void* ptr){
	return static_cast<QXmlParseException*>(ptr)->systemId().toUtf8().data();
}

void QXmlParseException_DestroyQXmlParseException(void* ptr){
	static_cast<QXmlParseException*>(ptr)->~QXmlParseException();
}

class MyQXmlReader: public QXmlReader {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void* QXmlReader_DTDHandler(void* ptr){
	return static_cast<QXmlReader*>(ptr)->DTDHandler();
}

void* QXmlReader_ContentHandler(void* ptr){
	return static_cast<QXmlReader*>(ptr)->contentHandler();
}

void* QXmlReader_DeclHandler(void* ptr){
	return static_cast<QXmlReader*>(ptr)->declHandler();
}

void* QXmlReader_EntityResolver(void* ptr){
	return static_cast<QXmlReader*>(ptr)->entityResolver();
}

void* QXmlReader_ErrorHandler(void* ptr){
	return static_cast<QXmlReader*>(ptr)->errorHandler();
}

int QXmlReader_Feature(void* ptr, char* name, int ok){
	return static_cast<QXmlReader*>(ptr)->feature(QString(name), NULL);
}

int QXmlReader_HasFeature(void* ptr, char* name){
	return static_cast<QXmlReader*>(ptr)->hasFeature(QString(name));
}

int QXmlReader_HasProperty(void* ptr, char* name){
	return static_cast<QXmlReader*>(ptr)->hasProperty(QString(name));
}

void* QXmlReader_LexicalHandler(void* ptr){
	return static_cast<QXmlReader*>(ptr)->lexicalHandler();
}

int QXmlReader_Parse(void* ptr, void* input){
	return static_cast<QXmlReader*>(ptr)->parse(static_cast<QXmlInputSource*>(input));
}

void* QXmlReader_Property(void* ptr, char* name, int ok){
	return static_cast<QXmlReader*>(ptr)->property(QString(name), NULL);
}

void QXmlReader_SetContentHandler(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setContentHandler(static_cast<QXmlContentHandler*>(handler));
}

void QXmlReader_SetDTDHandler(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setDTDHandler(static_cast<QXmlDTDHandler*>(handler));
}

void QXmlReader_SetDeclHandler(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setDeclHandler(static_cast<QXmlDeclHandler*>(handler));
}

void QXmlReader_SetEntityResolver(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setEntityResolver(static_cast<QXmlEntityResolver*>(handler));
}

void QXmlReader_SetErrorHandler(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setErrorHandler(static_cast<QXmlErrorHandler*>(handler));
}

void QXmlReader_SetFeature(void* ptr, char* name, int value){
	static_cast<QXmlReader*>(ptr)->setFeature(QString(name), value != 0);
}

void QXmlReader_SetLexicalHandler(void* ptr, void* handler){
	static_cast<QXmlReader*>(ptr)->setLexicalHandler(static_cast<QXmlLexicalHandler*>(handler));
}

void QXmlReader_DestroyQXmlReader(void* ptr){
	static_cast<QXmlReader*>(ptr)->~QXmlReader();
}

char* QXmlReader_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlReader*>(static_cast<QXmlReader*>(ptr))) {
		return static_cast<MyQXmlReader*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlReader_BASE").toUtf8().data();
}

void QXmlReader_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlReader*>(static_cast<QXmlReader*>(ptr))) {
		static_cast<MyQXmlReader*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQXmlSimpleReader: public QXmlSimpleReader {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQXmlSimpleReader() : QXmlSimpleReader() {};
	void setContentHandler(QXmlContentHandler * handler) { callbackQXmlSimpleReaderSetContentHandler(this, this->objectNameAbs().toUtf8().data(), handler); };
	void setDTDHandler(QXmlDTDHandler * handler) { callbackQXmlSimpleReaderSetDTDHandler(this, this->objectNameAbs().toUtf8().data(), handler); };
	void setDeclHandler(QXmlDeclHandler * handler) { callbackQXmlSimpleReaderSetDeclHandler(this, this->objectNameAbs().toUtf8().data(), handler); };
	void setEntityResolver(QXmlEntityResolver * handler) { callbackQXmlSimpleReaderSetEntityResolver(this, this->objectNameAbs().toUtf8().data(), handler); };
	void setErrorHandler(QXmlErrorHandler * handler) { callbackQXmlSimpleReaderSetErrorHandler(this, this->objectNameAbs().toUtf8().data(), handler); };
	void setFeature(const QString & name, bool enable) { callbackQXmlSimpleReaderSetFeature(this, this->objectNameAbs().toUtf8().data(), name.toUtf8().data(), enable); };
	void setLexicalHandler(QXmlLexicalHandler * handler) { callbackQXmlSimpleReaderSetLexicalHandler(this, this->objectNameAbs().toUtf8().data(), handler); };
};

void* QXmlSimpleReader_DTDHandler(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->DTDHandler();
}

void* QXmlSimpleReader_NewQXmlSimpleReader(){
	return new MyQXmlSimpleReader();
}

void* QXmlSimpleReader_ContentHandler(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->contentHandler();
}

void* QXmlSimpleReader_DeclHandler(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->declHandler();
}

void* QXmlSimpleReader_EntityResolver(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->entityResolver();
}

void* QXmlSimpleReader_ErrorHandler(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->errorHandler();
}

int QXmlSimpleReader_Feature(void* ptr, char* name, int ok){
	return static_cast<QXmlSimpleReader*>(ptr)->feature(QString(name), NULL);
}

int QXmlSimpleReader_HasFeature(void* ptr, char* name){
	return static_cast<QXmlSimpleReader*>(ptr)->hasFeature(QString(name));
}

int QXmlSimpleReader_HasProperty(void* ptr, char* name){
	return static_cast<QXmlSimpleReader*>(ptr)->hasProperty(QString(name));
}

void* QXmlSimpleReader_LexicalHandler(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->lexicalHandler();
}

int QXmlSimpleReader_Parse(void* ptr, void* input){
	return static_cast<QXmlSimpleReader*>(ptr)->parse(*static_cast<QXmlInputSource*>(input));
}

int QXmlSimpleReader_ParseContinue(void* ptr){
	return static_cast<QXmlSimpleReader*>(ptr)->parseContinue();
}

void* QXmlSimpleReader_Property(void* ptr, char* name, int ok){
	return static_cast<QXmlSimpleReader*>(ptr)->property(QString(name), NULL);
}

void QXmlSimpleReader_SetContentHandler(void* ptr, void* handler){
	static_cast<MyQXmlSimpleReader*>(ptr)->setContentHandler(static_cast<QXmlContentHandler*>(handler));
}

void QXmlSimpleReader_SetContentHandlerDefault(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setContentHandler(static_cast<QXmlContentHandler*>(handler));
}

void QXmlSimpleReader_SetDTDHandler(void* ptr, void* handler){
	static_cast<MyQXmlSimpleReader*>(ptr)->setDTDHandler(static_cast<QXmlDTDHandler*>(handler));
}

void QXmlSimpleReader_SetDTDHandlerDefault(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setDTDHandler(static_cast<QXmlDTDHandler*>(handler));
}

void QXmlSimpleReader_SetDeclHandler(void* ptr, void* handler){
	static_cast<MyQXmlSimpleReader*>(ptr)->setDeclHandler(static_cast<QXmlDeclHandler*>(handler));
}

void QXmlSimpleReader_SetDeclHandlerDefault(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setDeclHandler(static_cast<QXmlDeclHandler*>(handler));
}

void QXmlSimpleReader_SetEntityResolver(void* ptr, void* handler){
	static_cast<MyQXmlSimpleReader*>(ptr)->setEntityResolver(static_cast<QXmlEntityResolver*>(handler));
}

void QXmlSimpleReader_SetEntityResolverDefault(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setEntityResolver(static_cast<QXmlEntityResolver*>(handler));
}

void QXmlSimpleReader_SetErrorHandler(void* ptr, void* handler){
	static_cast<MyQXmlSimpleReader*>(ptr)->setErrorHandler(static_cast<QXmlErrorHandler*>(handler));
}

void QXmlSimpleReader_SetErrorHandlerDefault(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setErrorHandler(static_cast<QXmlErrorHandler*>(handler));
}

void QXmlSimpleReader_SetFeature(void* ptr, char* name, int enable){
	static_cast<MyQXmlSimpleReader*>(ptr)->setFeature(QString(name), enable != 0);
}

void QXmlSimpleReader_SetFeatureDefault(void* ptr, char* name, int enable){
	static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setFeature(QString(name), enable != 0);
}

void QXmlSimpleReader_SetLexicalHandler(void* ptr, void* handler){
	static_cast<MyQXmlSimpleReader*>(ptr)->setLexicalHandler(static_cast<QXmlLexicalHandler*>(handler));
}

void QXmlSimpleReader_SetLexicalHandlerDefault(void* ptr, void* handler){
	static_cast<QXmlSimpleReader*>(ptr)->QXmlSimpleReader::setLexicalHandler(static_cast<QXmlLexicalHandler*>(handler));
}

void QXmlSimpleReader_DestroyQXmlSimpleReader(void* ptr){
	static_cast<QXmlSimpleReader*>(ptr)->~QXmlSimpleReader();
}

char* QXmlSimpleReader_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQXmlSimpleReader*>(static_cast<QXmlSimpleReader*>(ptr))) {
		return static_cast<MyQXmlSimpleReader*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QXmlSimpleReader_BASE").toUtf8().data();
}

void QXmlSimpleReader_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQXmlSimpleReader*>(static_cast<QXmlSimpleReader*>(ptr))) {
		static_cast<MyQXmlSimpleReader*>(ptr)->setObjectNameAbs(QString(name));
	}
}

