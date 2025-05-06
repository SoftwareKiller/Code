#pragma once
#pragma execution_character_set("utf-8")

#include <QtWidgets/QMainWindow>
#include <QWidget>
#include "ui_QtPainterEvent.h"

class QtPainterEvent : public QMainWindow
{
	Q_OBJECT

public:
	QtPainterEvent(QWidget* parent = nullptr);
	~QtPainterEvent();

	// 绘图事件
	void paintEvent(QPaintEvent*);

private:
	Ui::QtPainterEventClass ui;
};
