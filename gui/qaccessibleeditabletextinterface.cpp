#include "qaccessibleeditabletextinterface.h"
#include <QModelIndex>
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAccessibleEditableTextInterface>
#include "_cgo_export.h"

class MyQAccessibleEditableTextInterface: public QAccessibleEditableTextInterface {
public:
};

void QAccessibleEditableTextInterface_DeleteText(void* ptr, int startOffset, int endOffset){
	static_cast<QAccessibleEditableTextInterface*>(ptr)->deleteText(startOffset, endOffset);
}

void QAccessibleEditableTextInterface_InsertText(void* ptr, int offset, char* text){
	static_cast<QAccessibleEditableTextInterface*>(ptr)->insertText(offset, QString(text));
}

void QAccessibleEditableTextInterface_ReplaceText(void* ptr, int startOffset, int endOffset, char* text){
	static_cast<QAccessibleEditableTextInterface*>(ptr)->replaceText(startOffset, endOffset, QString(text));
}

void QAccessibleEditableTextInterface_DestroyQAccessibleEditableTextInterface(void* ptr){
	static_cast<QAccessibleEditableTextInterface*>(ptr)->~QAccessibleEditableTextInterface();
}

