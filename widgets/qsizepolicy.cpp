#include "qsizepolicy.h"
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QSizePolicy>
#include "_cgo_export.h"

class MyQSizePolicy: public QSizePolicy {
public:
};

QtObjectPtr QSizePolicy_NewQSizePolicy(){
	return new QSizePolicy();
}

QtObjectPtr QSizePolicy_NewQSizePolicy2(int horizontal, int vertical, int ty){
	return new QSizePolicy(static_cast<QSizePolicy::Policy>(horizontal), static_cast<QSizePolicy::Policy>(vertical), static_cast<QSizePolicy::ControlType>(ty));
}

int QSizePolicy_ControlType(QtObjectPtr ptr){
	return static_cast<QSizePolicy*>(ptr)->controlType();
}

int QSizePolicy_ExpandingDirections(QtObjectPtr ptr){
	return static_cast<QSizePolicy*>(ptr)->expandingDirections();
}

int QSizePolicy_HasHeightForWidth(QtObjectPtr ptr){
	return static_cast<QSizePolicy*>(ptr)->hasHeightForWidth();
}

int QSizePolicy_HasWidthForHeight(QtObjectPtr ptr){
	return static_cast<QSizePolicy*>(ptr)->hasWidthForHeight();
}

int QSizePolicy_HorizontalPolicy(QtObjectPtr ptr){
	return static_cast<QSizePolicy*>(ptr)->horizontalPolicy();
}

int QSizePolicy_HorizontalStretch(QtObjectPtr ptr){
	return static_cast<QSizePolicy*>(ptr)->horizontalStretch();
}

int QSizePolicy_RetainSizeWhenHidden(QtObjectPtr ptr){
	return static_cast<QSizePolicy*>(ptr)->retainSizeWhenHidden();
}

void QSizePolicy_SetControlType(QtObjectPtr ptr, int ty){
	static_cast<QSizePolicy*>(ptr)->setControlType(static_cast<QSizePolicy::ControlType>(ty));
}

void QSizePolicy_SetHeightForWidth(QtObjectPtr ptr, int dependent){
	static_cast<QSizePolicy*>(ptr)->setHeightForWidth(dependent != 0);
}

void QSizePolicy_SetHorizontalPolicy(QtObjectPtr ptr, int policy){
	static_cast<QSizePolicy*>(ptr)->setHorizontalPolicy(static_cast<QSizePolicy::Policy>(policy));
}

void QSizePolicy_SetHorizontalStretch(QtObjectPtr ptr, int stretchFactor){
	static_cast<QSizePolicy*>(ptr)->setHorizontalStretch(stretchFactor);
}

void QSizePolicy_SetRetainSizeWhenHidden(QtObjectPtr ptr, int retainSize){
	static_cast<QSizePolicy*>(ptr)->setRetainSizeWhenHidden(retainSize != 0);
}

void QSizePolicy_SetVerticalPolicy(QtObjectPtr ptr, int policy){
	static_cast<QSizePolicy*>(ptr)->setVerticalPolicy(static_cast<QSizePolicy::Policy>(policy));
}

void QSizePolicy_SetVerticalStretch(QtObjectPtr ptr, int stretchFactor){
	static_cast<QSizePolicy*>(ptr)->setVerticalStretch(stretchFactor);
}

void QSizePolicy_SetWidthForHeight(QtObjectPtr ptr, int dependent){
	static_cast<QSizePolicy*>(ptr)->setWidthForHeight(dependent != 0);
}

void QSizePolicy_Transpose(QtObjectPtr ptr){
	static_cast<QSizePolicy*>(ptr)->transpose();
}

int QSizePolicy_VerticalPolicy(QtObjectPtr ptr){
	return static_cast<QSizePolicy*>(ptr)->verticalPolicy();
}

int QSizePolicy_VerticalStretch(QtObjectPtr ptr){
	return static_cast<QSizePolicy*>(ptr)->verticalStretch();
}

