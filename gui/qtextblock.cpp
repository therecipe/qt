#include "qtextblock.h"
#include <QUrl>
#include <QModelIndex>
#include <QTextBlockUserData>
#include <QTextDocument>
#include <QString>
#include <QVariant>
#include <QTextBlock>
#include "_cgo_export.h"

class MyQTextBlock: public QTextBlock {
public:
};

int QTextBlock_IsValid(void* ptr){
	return static_cast<QTextBlock*>(ptr)->isValid();
}

void* QTextBlock_NewQTextBlock(void* other){
	return new QTextBlock(*static_cast<QTextBlock*>(other));
}

int QTextBlock_BlockFormatIndex(void* ptr){
	return static_cast<QTextBlock*>(ptr)->blockFormatIndex();
}

int QTextBlock_CharFormatIndex(void* ptr){
	return static_cast<QTextBlock*>(ptr)->charFormatIndex();
}

void QTextBlock_ClearLayout(void* ptr){
	static_cast<QTextBlock*>(ptr)->clearLayout();
}

int QTextBlock_Contains(void* ptr, int position){
	return static_cast<QTextBlock*>(ptr)->contains(position);
}

int QTextBlock_BlockNumber(void* ptr){
	return static_cast<QTextBlock*>(ptr)->blockNumber();
}

void* QTextBlock_Document(void* ptr){
	return const_cast<QTextDocument*>(static_cast<QTextBlock*>(ptr)->document());
}

int QTextBlock_FirstLineNumber(void* ptr){
	return static_cast<QTextBlock*>(ptr)->firstLineNumber();
}

int QTextBlock_IsVisible(void* ptr){
	return static_cast<QTextBlock*>(ptr)->isVisible();
}

void* QTextBlock_Layout(void* ptr){
	return static_cast<QTextBlock*>(ptr)->layout();
}

int QTextBlock_Length(void* ptr){
	return static_cast<QTextBlock*>(ptr)->length();
}

int QTextBlock_LineCount(void* ptr){
	return static_cast<QTextBlock*>(ptr)->lineCount();
}

int QTextBlock_Position(void* ptr){
	return static_cast<QTextBlock*>(ptr)->position();
}

int QTextBlock_Revision(void* ptr){
	return static_cast<QTextBlock*>(ptr)->revision();
}

void QTextBlock_SetLineCount(void* ptr, int count){
	static_cast<QTextBlock*>(ptr)->setLineCount(count);
}

void QTextBlock_SetRevision(void* ptr, int rev){
	static_cast<QTextBlock*>(ptr)->setRevision(rev);
}

void QTextBlock_SetUserData(void* ptr, void* data){
	static_cast<QTextBlock*>(ptr)->setUserData(static_cast<QTextBlockUserData*>(data));
}

void QTextBlock_SetUserState(void* ptr, int state){
	static_cast<QTextBlock*>(ptr)->setUserState(state);
}

void QTextBlock_SetVisible(void* ptr, int visible){
	static_cast<QTextBlock*>(ptr)->setVisible(visible != 0);
}

char* QTextBlock_Text(void* ptr){
	return static_cast<QTextBlock*>(ptr)->text().toUtf8().data();
}

int QTextBlock_TextDirection(void* ptr){
	return static_cast<QTextBlock*>(ptr)->textDirection();
}

void* QTextBlock_TextList(void* ptr){
	return static_cast<QTextBlock*>(ptr)->textList();
}

void* QTextBlock_UserData(void* ptr){
	return static_cast<QTextBlock*>(ptr)->userData();
}

int QTextBlock_UserState(void* ptr){
	return static_cast<QTextBlock*>(ptr)->userState();
}

