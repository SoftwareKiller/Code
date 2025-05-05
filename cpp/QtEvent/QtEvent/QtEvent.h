#pragma once

#include <QtWidgets/QMainWindow>
#include "ui_QtEvent.h"

class QtEvent : public QMainWindow
{
	Q_OBJECT

public:
	QtEvent(QWidget* parent = nullptr);
	~QtEvent();

private:
	Ui::QtEventClass ui;
};
