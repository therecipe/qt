#include "qsizepolicy.h"
#include <QModelIndex>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSizePolicy>
#include "_cgo_export.h"

class MyQSizePolicy: public QSizePolicy {
public:
};

void* QSizePolicy_NewQSizePolicy(){
	return new QSizePolicy();
}

void* QSizePolicy_NewQSizePolicy2(int horizontal, int vertical, int ty){
	return new QSizePolicy(static_cast<QSizePolicy::Policy>(horizontal), static_cast<QSizePolicy::Policy>(vertical), static_cast<QSizePolicy::ControlType>(ty));
}

int QSizePolicy_ControlType(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->controlType();
}

int QSizePolicy_ExpandingDirections(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->expandingDirections();
}

int QSizePolicy_HasHeightForWidth(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->hasHeightForWidth();
}

int QSizePolicy_HasWidthForHeight(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->hasWidthForHeight();
}

int QSizePolicy_HorizontalPolicy(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->horizontalPolicy();
}

int QSizePolicy_HorizontalStretch(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->horizontalStretch();
}

int QSizePolicy_RetainSizeWhenHidden(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->retainSizeWhenHidden();
}

void QSizePolicy_SetControlType(void* ptr, int ty){
	static_cast<QSizePolicy*>(ptr)->setControlType(static_cast<QSizePolicy::ControlType>(ty));
}

void QSizePolicy_SetHeightForWidth(void* ptr, int dependent){
	static_cast<QSizePolicy*>(ptr)->setHeightForWidth(dependent != 0);
}

void QSizePolicy_SetHorizontalPolicy(void* ptr, int policy){
	static_cast<QSizePolicy*>(ptr)->setHorizontalPolicy(static_cast<QSizePolicy::Policy>(policy));
}

void QSizePolicy_SetHorizontalStretch(void* ptr, int stretchFactor){
	static_cast<QSizePolicy*>(ptr)->setHorizontalStretch(stretchFactor);
}

void QSizePolicy_SetRetainSizeWhenHidden(void* ptr, int retainSize){
	static_cast<QSizePolicy*>(ptr)->setRetainSizeWhenHidden(retainSize != 0);
}

void QSizePolicy_SetVerticalPolicy(void* ptr, int policy){
	static_cast<QSizePolicy*>(ptr)->setVerticalPolicy(static_cast<QSizePolicy::Policy>(policy));
}

void QSizePolicy_SetVerticalStretch(void* ptr, int stretchFactor){
	static_cast<QSizePolicy*>(ptr)->setVerticalStretch(stretchFactor);
}

void QSizePolicy_SetWidthForHeight(void* ptr, int dependent){
	static_cast<QSizePolicy*>(ptr)->setWidthForHeight(dependent != 0);
}

void QSizePolicy_Transpose(void* ptr){
	static_cast<QSizePolicy*>(ptr)->transpose();
}

int QSizePolicy_VerticalPolicy(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->verticalPolicy();
}

int QSizePolicy_VerticalStretch(void* ptr){
	return static_cast<QSizePolicy*>(ptr)->verticalStretch();
}

