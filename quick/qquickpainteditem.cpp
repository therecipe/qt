#include "qquickpainteditem.h"
#include <QSize>
#include <QPainter>
#include <QVariant>
#include <QUrl>
#include <QObject>
#include <QRect>
#include <QString>
#include <QModelIndex>
#include <QColor>
#include <QQuickPaintedItem>
#include "_cgo_export.h"

class MyQQuickPaintedItem: public QQuickPaintedItem {
public:
void Signal_ContentsScaleChanged(){callbackQQuickPaintedItemContentsScaleChanged(this->objectName().toUtf8().data());};
void Signal_ContentsSizeChanged(){callbackQQuickPaintedItemContentsSizeChanged(this->objectName().toUtf8().data());};
void Signal_FillColorChanged(){callbackQQuickPaintedItemFillColorChanged(this->objectName().toUtf8().data());};
void Signal_RenderTargetChanged(){callbackQQuickPaintedItemRenderTargetChanged(this->objectName().toUtf8().data());};
};

double QQuickPaintedItem_ContentsScale(void* ptr){
	return static_cast<double>(static_cast<QQuickPaintedItem*>(ptr)->contentsScale());
}

void* QQuickPaintedItem_FillColor(void* ptr){
	return new QColor(static_cast<QQuickPaintedItem*>(ptr)->fillColor());
}

int QQuickPaintedItem_RenderTarget(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->renderTarget();
}

void QQuickPaintedItem_SetContentsScale(void* ptr, double v){
	static_cast<QQuickPaintedItem*>(ptr)->setContentsScale(static_cast<qreal>(v));
}

void QQuickPaintedItem_SetContentsSize(void* ptr, void* v){
	static_cast<QQuickPaintedItem*>(ptr)->setContentsSize(*static_cast<QSize*>(v));
}

void QQuickPaintedItem_SetFillColor(void* ptr, void* v){
	static_cast<QQuickPaintedItem*>(ptr)->setFillColor(*static_cast<QColor*>(v));
}

void QQuickPaintedItem_SetRenderTarget(void* ptr, int target){
	static_cast<QQuickPaintedItem*>(ptr)->setRenderTarget(static_cast<QQuickPaintedItem::RenderTarget>(target));
}

int QQuickPaintedItem_Antialiasing(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->antialiasing();
}

void QQuickPaintedItem_ConnectContentsScaleChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsScaleChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsScaleChanged));;
}

void QQuickPaintedItem_DisconnectContentsScaleChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsScaleChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsScaleChanged));;
}

void QQuickPaintedItem_ConnectContentsSizeChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsSizeChanged));;
}

void QQuickPaintedItem_DisconnectContentsSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsSizeChanged));;
}

void QQuickPaintedItem_ConnectFillColorChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::fillColorChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_FillColorChanged));;
}

void QQuickPaintedItem_DisconnectFillColorChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::fillColorChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_FillColorChanged));;
}

int QQuickPaintedItem_IsTextureProvider(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->isTextureProvider();
}

int QQuickPaintedItem_Mipmap(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->mipmap();
}

int QQuickPaintedItem_OpaquePainting(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->opaquePainting();
}

void QQuickPaintedItem_Paint(void* ptr, void* painter){
	static_cast<QQuickPaintedItem*>(ptr)->paint(static_cast<QPainter*>(painter));
}

int QQuickPaintedItem_PerformanceHints(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->performanceHints();
}

void QQuickPaintedItem_ConnectRenderTargetChanged(void* ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::renderTargetChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_RenderTargetChanged));;
}

void QQuickPaintedItem_DisconnectRenderTargetChanged(void* ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::renderTargetChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_RenderTargetChanged));;
}

void QQuickPaintedItem_ResetContentsSize(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->resetContentsSize();
}

void QQuickPaintedItem_SetAntialiasing(void* ptr, int enable){
	static_cast<QQuickPaintedItem*>(ptr)->setAntialiasing(enable != 0);
}

void QQuickPaintedItem_SetMipmap(void* ptr, int enable){
	static_cast<QQuickPaintedItem*>(ptr)->setMipmap(enable != 0);
}

void QQuickPaintedItem_SetOpaquePainting(void* ptr, int opaque){
	static_cast<QQuickPaintedItem*>(ptr)->setOpaquePainting(opaque != 0);
}

void QQuickPaintedItem_SetPerformanceHint(void* ptr, int hint, int enabled){
	static_cast<QQuickPaintedItem*>(ptr)->setPerformanceHint(static_cast<QQuickPaintedItem::PerformanceHint>(hint), enabled != 0);
}

void QQuickPaintedItem_SetPerformanceHints(void* ptr, int hints){
	static_cast<QQuickPaintedItem*>(ptr)->setPerformanceHints(static_cast<QQuickPaintedItem::PerformanceHint>(hints));
}

void* QQuickPaintedItem_TextureProvider(void* ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->textureProvider();
}

void QQuickPaintedItem_Update(void* ptr, void* rect){
	static_cast<QQuickPaintedItem*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QQuickPaintedItem_DestroyQQuickPaintedItem(void* ptr){
	static_cast<QQuickPaintedItem*>(ptr)->~QQuickPaintedItem();
}

