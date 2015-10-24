#include "qtextblock.h"
#include <QTextBlockUserData>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextDocument>
#include <QTextBlock>
#include "_cgo_export.h"

class MyQTextBlock: public QTextBlock {
public:
};

int QTextBlock_IsValid(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->isValid();
}

QtObjectPtr QTextBlock_NewQTextBlock(QtObjectPtr other){
	return new QTextBlock(*static_cast<QTextBlock*>(other));
}

int QTextBlock_BlockFormatIndex(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->blockFormatIndex();
}

int QTextBlock_CharFormatIndex(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->charFormatIndex();
}

void QTextBlock_ClearLayout(QtObjectPtr ptr){
	static_cast<QTextBlock*>(ptr)->clearLayout();
}

int QTextBlock_Contains(QtObjectPtr ptr, int position){
	return static_cast<QTextBlock*>(ptr)->contains(position);
}

int QTextBlock_BlockNumber(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->blockNumber();
}

QtObjectPtr QTextBlock_Document(QtObjectPtr ptr){
	return const_cast<QTextDocument*>(static_cast<QTextBlock*>(ptr)->document());
}

int QTextBlock_FirstLineNumber(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->firstLineNumber();
}

int QTextBlock_IsVisible(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->isVisible();
}

QtObjectPtr QTextBlock_Layout(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->layout();
}

int QTextBlock_Length(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->length();
}

int QTextBlock_LineCount(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->lineCount();
}

int QTextBlock_Position(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->position();
}

int QTextBlock_Revision(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->revision();
}

void QTextBlock_SetLineCount(QtObjectPtr ptr, int count){
	static_cast<QTextBlock*>(ptr)->setLineCount(count);
}

void QTextBlock_SetRevision(QtObjectPtr ptr, int rev){
	static_cast<QTextBlock*>(ptr)->setRevision(rev);
}

void QTextBlock_SetUserData(QtObjectPtr ptr, QtObjectPtr data){
	static_cast<QTextBlock*>(ptr)->setUserData(static_cast<QTextBlockUserData*>(data));
}

void QTextBlock_SetUserState(QtObjectPtr ptr, int state){
	static_cast<QTextBlock*>(ptr)->setUserState(state);
}

void QTextBlock_SetVisible(QtObjectPtr ptr, int visible){
	static_cast<QTextBlock*>(ptr)->setVisible(visible != 0);
}

char* QTextBlock_Text(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->text().toUtf8().data();
}

int QTextBlock_TextDirection(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->textDirection();
}

QtObjectPtr QTextBlock_TextList(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->textList();
}

QtObjectPtr QTextBlock_UserData(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->userData();
}

int QTextBlock_UserState(QtObjectPtr ptr){
	return static_cast<QTextBlock*>(ptr)->userState();
}

