from django.http import HttpResponse

from TestModel.models import Test

# def testdb(request):
#     test1 = Test(name='runoob')
#     test1.save()
#     return HttpResponse("<p>数据添加成功</p>")

def testdb(request):
    response = ""
    response1 = ""
    # 全表查询
    listTest = Test.objects.all()
    # 条件过滤
    response2 = Test.objects.filter(id=1)
    # 获取单个对象
    response3 = Test.objects.get(id=1)
    # 分页
    # OFFSET 0 LIMIT 2
    Test.objects.order_by('name')[0:2]
