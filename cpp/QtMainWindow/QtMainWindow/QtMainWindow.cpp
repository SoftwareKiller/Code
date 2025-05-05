#include "QtMainWindow.h"
#include <QMenuBar>
#include <QToolBar>
#include <QPushButton>
#include <QLabel>
#include <QDockWidget>
#include <QTextEdit>

QtMainWindow::QtMainWindow(QWidget* parent)
	: QMainWindow(parent)
{
	// 重置窗口大小
	resize(600, 400);

	// 菜单栏创建，最多只能有一个
	QMenuBar* bar = menuBar();
	setMenuBar(bar);

	// 创建菜单
	QMenu* fileMenu = bar->addMenu("文件");
	QMenu* editMenu = bar->addMenu("编辑");

	// 创建菜单项
	QAction* newAction = fileMenu->addAction("新建");
	// 添加分隔符
	fileMenu->addSeparator();

	QAction* openAction = fileMenu->addAction("打开");

	// 工具栏，可以有多个
	QToolBar* toolBar = new QToolBar(this);
	/*addToolBar(toolBar);*/
	addToolBar(Qt::LeftToolBarArea, toolBar);

	// 只允许 左右停靠
	toolBar->setAllowedAreas(Qt::LeftToolBarArea | Qt::RightToolBarArea);

	// 设置浮动
	toolBar->setFloatable(false);

	// 设置移动(优先级最高)
	toolBar->setMovable(false);

	// 工具栏中添加内容
	toolBar->addAction(newAction);
	toolBar->addSeparator();
	toolBar->addAction(openAction);
	toolBar->addSeparator();

	// 工具栏中添加控件
	QPushButton* btn = new QPushButton("BTN", this);
	toolBar->addWidget(btn);

	// 状态栏 最多有一个
	QStatusBar* stBar = statusBar();
	setStatusBar(stBar);
	// 放标签控件
	QLabel* label = new QLabel("提示信息", this);
	stBar->addWidget(label);

	QLabel* labelR = new QLabel("右侧提示信息", this);
	stBar->addPermanentWidget(labelR);

	// 铆接部件 (浮动窗口) 可以有多个
	QDockWidget* dockWidget = new QDockWidget("浮动", this);
	addDockWidget(Qt::BottomDockWidgetArea, dockWidget);

	// 设置区域只允许上下
	dockWidget->setAllowedAreas(Qt::TopDockWidgetArea | Qt::BottomDockWidgetArea);

	// 设置核心
	QTextEdit* edit = new QTextEdit(this);
	setCentralWidget(edit);
}

QtMainWindow::~QtMainWindow()
{
}
