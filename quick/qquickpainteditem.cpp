#include "qquickpainteditem.h"
#include <QString>
#include <QUrl>
#include <QObject>
#include <QRect>
#include <QColor>
#include <QVariant>
#include <QModelIndex>
#include <QPainter>
#include <QSize>
#include <QQuickPaintedItem>
#include "_cgo_export.h"

class MyQQuickPaintedItem: public QQuickPaintedItem {
public:
void Signal_ContentsScaleChanged(){callbackQQuickPaintedItemContentsScaleChanged(this->objectName().toUtf8().data());};
void Signal_ContentsSizeChanged(){callbackQQuickPaintedItemContentsSizeChanged(this->objectName().toUtf8().data());};
void Signal_FillColorChanged(){callbackQQuickPaintedItemFillColorChanged(this->objectName().toUtf8().data());};
void Signal_RenderTargetChanged(){callbackQQuickPaintedItemRenderTargetChanged(this->objectName().toUtf8().data());};
};

int QQuickPaintedItem_RenderTarget(QtObjectPtr ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->renderTarget();
}

void QQuickPaintedItem_SetContentsSize(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QQuickPaintedItem*>(ptr)->setContentsSize(*static_cast<QSize*>(v));
}

void QQuickPaintedItem_SetFillColor(QtObjectPtr ptr, QtObjectPtr v){
	static_cast<QQuickPaintedItem*>(ptr)->setFillColor(*static_cast<QColor*>(v));
}

void QQuickPaintedItem_SetRenderTarget(QtObjectPtr ptr, int target){
	static_cast<QQuickPaintedItem*>(ptr)->setRenderTarget(static_cast<QQuickPaintedItem::RenderTarget>(target));
}

int QQuickPaintedItem_Antialiasing(QtObjectPtr ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->antialiasing();
}

void QQuickPaintedItem_ConnectContentsScaleChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsScaleChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsScaleChanged));;
}

void QQuickPaintedItem_DisconnectContentsScaleChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsScaleChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsScaleChanged));;
}

void QQuickPaintedItem_ConnectContentsSizeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsSizeChanged));;
}

void QQuickPaintedItem_DisconnectContentsSizeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::contentsSizeChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_ContentsSizeChanged));;
}

void QQuickPaintedItem_ConnectFillColorChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::fillColorChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_FillColorChanged));;
}

void QQuickPaintedItem_DisconnectFillColorChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::fillColorChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_FillColorChanged));;
}

int QQuickPaintedItem_IsTextureProvider(QtObjectPtr ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->isTextureProvider();
}

int QQuickPaintedItem_Mipmap(QtObjectPtr ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->mipmap();
}

int QQuickPaintedItem_OpaquePainting(QtObjectPtr ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->opaquePainting();
}

void QQuickPaintedItem_Paint(QtObjectPtr ptr, QtObjectPtr painter){
	static_cast<QQuickPaintedItem*>(ptr)->paint(static_cast<QPainter*>(painter));
}

int QQuickPaintedItem_PerformanceHints(QtObjectPtr ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->performanceHints();
}

void QQuickPaintedItem_ConnectRenderTargetChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::renderTargetChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_RenderTargetChanged));;
}

void QQuickPaintedItem_DisconnectRenderTargetChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QQuickPaintedItem*>(ptr), static_cast<void (QQuickPaintedItem::*)()>(&QQuickPaintedItem::renderTargetChanged), static_cast<MyQQuickPaintedItem*>(ptr), static_cast<void (MyQQuickPaintedItem::*)()>(&MyQQuickPaintedItem::Signal_RenderTargetChanged));;
}

void QQuickPaintedItem_ResetContentsSize(QtObjectPtr ptr){
	static_cast<QQuickPaintedItem*>(ptr)->resetContentsSize();
}

void QQuickPaintedItem_SetAntialiasing(QtObjectPtr ptr, int enable){
	static_cast<QQuickPaintedItem*>(ptr)->setAntialiasing(enable != 0);
}

void QQuickPaintedItem_SetMipmap(QtObjectPtr ptr, int enable){
	static_cast<QQuickPaintedItem*>(ptr)->setMipmap(enable != 0);
}

void QQuickPaintedItem_SetOpaquePainting(QtObjectPtr ptr, int opaque){
	static_cast<QQuickPaintedItem*>(ptr)->setOpaquePainting(opaque != 0);
}

void QQuickPaintedItem_SetPerformanceHint(QtObjectPtr ptr, int hint, int enabled){
	static_cast<QQuickPaintedItem*>(ptr)->setPerformanceHint(static_cast<QQuickPaintedItem::PerformanceHint>(hint), enabled != 0);
}

void QQuickPaintedItem_SetPerformanceHints(QtObjectPtr ptr, int hints){
	static_cast<QQuickPaintedItem*>(ptr)->setPerformanceHints(static_cast<QQuickPaintedItem::PerformanceHint>(hints));
}

QtObjectPtr QQuickPaintedItem_TextureProvider(QtObjectPtr ptr){
	return static_cast<QQuickPaintedItem*>(ptr)->textureProvider();
}

void QQuickPaintedItem_Update(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QQuickPaintedItem*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QQuickPaintedItem_DestroyQQuickPaintedItem(QtObjectPtr ptr){
	static_cast<QQuickPaintedItem*>(ptr)->~QQuickPaintedItem();
}

