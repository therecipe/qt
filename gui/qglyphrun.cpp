#include "qglyphrun.h"
#include <QRawFont>
#include <QRectF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QGlyphRun>
#include "_cgo_export.h"

class MyQGlyphRun: public QGlyphRun {
public:
};

void* QGlyphRun_NewQGlyphRun(){
	return new QGlyphRun();
}

void* QGlyphRun_NewQGlyphRun2(void* other){
	return new QGlyphRun(*static_cast<QGlyphRun*>(other));
}

void QGlyphRun_Clear(void* ptr){
	static_cast<QGlyphRun*>(ptr)->clear();
}

int QGlyphRun_Flags(void* ptr){
	return static_cast<QGlyphRun*>(ptr)->flags();
}

int QGlyphRun_IsEmpty(void* ptr){
	return static_cast<QGlyphRun*>(ptr)->isEmpty();
}

int QGlyphRun_IsRightToLeft(void* ptr){
	return static_cast<QGlyphRun*>(ptr)->isRightToLeft();
}

int QGlyphRun_Overline(void* ptr){
	return static_cast<QGlyphRun*>(ptr)->overline();
}

void QGlyphRun_SetBoundingRect(void* ptr, void* boundingRect){
	static_cast<QGlyphRun*>(ptr)->setBoundingRect(*static_cast<QRectF*>(boundingRect));
}

void QGlyphRun_SetFlag(void* ptr, int flag, int enabled){
	static_cast<QGlyphRun*>(ptr)->setFlag(static_cast<QGlyphRun::GlyphRunFlag>(flag), enabled != 0);
}

void QGlyphRun_SetFlags(void* ptr, int flags){
	static_cast<QGlyphRun*>(ptr)->setFlags(static_cast<QGlyphRun::GlyphRunFlag>(flags));
}

void QGlyphRun_SetOverline(void* ptr, int overline){
	static_cast<QGlyphRun*>(ptr)->setOverline(overline != 0);
}

void QGlyphRun_SetRawFont(void* ptr, void* rawFont){
	static_cast<QGlyphRun*>(ptr)->setRawFont(*static_cast<QRawFont*>(rawFont));
}

void QGlyphRun_SetRightToLeft(void* ptr, int rightToLeft){
	static_cast<QGlyphRun*>(ptr)->setRightToLeft(rightToLeft != 0);
}

void QGlyphRun_SetStrikeOut(void* ptr, int strikeOut){
	static_cast<QGlyphRun*>(ptr)->setStrikeOut(strikeOut != 0);
}

void QGlyphRun_SetUnderline(void* ptr, int underline){
	static_cast<QGlyphRun*>(ptr)->setUnderline(underline != 0);
}

int QGlyphRun_StrikeOut(void* ptr){
	return static_cast<QGlyphRun*>(ptr)->strikeOut();
}

void QGlyphRun_Swap(void* ptr, void* other){
	static_cast<QGlyphRun*>(ptr)->swap(*static_cast<QGlyphRun*>(other));
}

int QGlyphRun_Underline(void* ptr){
	return static_cast<QGlyphRun*>(ptr)->underline();
}

void QGlyphRun_DestroyQGlyphRun(void* ptr){
	static_cast<QGlyphRun*>(ptr)->~QGlyphRun();
}

