from apscheduler.schedulers.background import BackgroundScheduler

def my_task_function():
    # 这里写你的任务逻辑
    print("定时任务执行了！")

# 实例化调度器
scheduler = BackgroundScheduler()
scheduler.add_job(my_task_function, 'interval', seconds=5)

