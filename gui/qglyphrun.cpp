#include "qglyphrun.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QRectF>
#include <QRawFont>
#include <QGlyphRun>
#include "_cgo_export.h"

class MyQGlyphRun: public QGlyphRun {
public:
};

QtObjectPtr QGlyphRun_NewQGlyphRun(){
	return new QGlyphRun();
}

QtObjectPtr QGlyphRun_NewQGlyphRun2(QtObjectPtr other){
	return new QGlyphRun(*static_cast<QGlyphRun*>(other));
}

void QGlyphRun_Clear(QtObjectPtr ptr){
	static_cast<QGlyphRun*>(ptr)->clear();
}

int QGlyphRun_Flags(QtObjectPtr ptr){
	return static_cast<QGlyphRun*>(ptr)->flags();
}

int QGlyphRun_IsEmpty(QtObjectPtr ptr){
	return static_cast<QGlyphRun*>(ptr)->isEmpty();
}

int QGlyphRun_IsRightToLeft(QtObjectPtr ptr){
	return static_cast<QGlyphRun*>(ptr)->isRightToLeft();
}

int QGlyphRun_Overline(QtObjectPtr ptr){
	return static_cast<QGlyphRun*>(ptr)->overline();
}

void QGlyphRun_SetBoundingRect(QtObjectPtr ptr, QtObjectPtr boundingRect){
	static_cast<QGlyphRun*>(ptr)->setBoundingRect(*static_cast<QRectF*>(boundingRect));
}

void QGlyphRun_SetFlag(QtObjectPtr ptr, int flag, int enabled){
	static_cast<QGlyphRun*>(ptr)->setFlag(static_cast<QGlyphRun::GlyphRunFlag>(flag), enabled != 0);
}

void QGlyphRun_SetFlags(QtObjectPtr ptr, int flags){
	static_cast<QGlyphRun*>(ptr)->setFlags(static_cast<QGlyphRun::GlyphRunFlag>(flags));
}

void QGlyphRun_SetOverline(QtObjectPtr ptr, int overline){
	static_cast<QGlyphRun*>(ptr)->setOverline(overline != 0);
}

void QGlyphRun_SetRawFont(QtObjectPtr ptr, QtObjectPtr rawFont){
	static_cast<QGlyphRun*>(ptr)->setRawFont(*static_cast<QRawFont*>(rawFont));
}

void QGlyphRun_SetRightToLeft(QtObjectPtr ptr, int rightToLeft){
	static_cast<QGlyphRun*>(ptr)->setRightToLeft(rightToLeft != 0);
}

void QGlyphRun_SetStrikeOut(QtObjectPtr ptr, int strikeOut){
	static_cast<QGlyphRun*>(ptr)->setStrikeOut(strikeOut != 0);
}

void QGlyphRun_SetUnderline(QtObjectPtr ptr, int underline){
	static_cast<QGlyphRun*>(ptr)->setUnderline(underline != 0);
}

int QGlyphRun_StrikeOut(QtObjectPtr ptr){
	return static_cast<QGlyphRun*>(ptr)->strikeOut();
}

void QGlyphRun_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QGlyphRun*>(ptr)->swap(*static_cast<QGlyphRun*>(other));
}

int QGlyphRun_Underline(QtObjectPtr ptr){
	return static_cast<QGlyphRun*>(ptr)->underline();
}

void QGlyphRun_DestroyQGlyphRun(QtObjectPtr ptr){
	static_cast<QGlyphRun*>(ptr)->~QGlyphRun();
}

