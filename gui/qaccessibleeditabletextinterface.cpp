#include "qaccessibleeditabletextinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QAccessibleEditableTextInterface>
#include "_cgo_export.h"

class MyQAccessibleEditableTextInterface: public QAccessibleEditableTextInterface {
public:
};

void QAccessibleEditableTextInterface_DeleteText(QtObjectPtr ptr, int startOffset, int endOffset){
	static_cast<QAccessibleEditableTextInterface*>(ptr)->deleteText(startOffset, endOffset);
}

void QAccessibleEditableTextInterface_InsertText(QtObjectPtr ptr, int offset, char* text){
	static_cast<QAccessibleEditableTextInterface*>(ptr)->insertText(offset, QString(text));
}

void QAccessibleEditableTextInterface_ReplaceText(QtObjectPtr ptr, int startOffset, int endOffset, char* text){
	static_cast<QAccessibleEditableTextInterface*>(ptr)->replaceText(startOffset, endOffset, QString(text));
}

void QAccessibleEditableTextInterface_DestroyQAccessibleEditableTextInterface(QtObjectPtr ptr){
	static_cast<QAccessibleEditableTextInterface*>(ptr)->~QAccessibleEditableTextInterface();
}

