#pragma once
#pragma execution_character_set("utf-8")

#include <QtWidgets/QMainWindow>
#include "ui_QtMainWindow.h"

class QtMainWindow : public QMainWindow
{
	Q_OBJECT

public:
	QtMainWindow(QWidget* parent = nullptr);
	~QtMainWindow();

private:
	Ui::QtMainWindowClass ui;
};
