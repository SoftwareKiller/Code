import requests
from bs4 import BeautifulSoup

# 假设这是你要爬取的网页的起始URL
start_url = 'http://m.uuzw.cc/l_132-132357/45775760.html'

# 用于存储所有页面内容的列表
all_pages_content = []

# 爬取页面的函数
def crawl_page(url):
    # 发送HTTP请求
    response = requests.get(url)
    # 确保请求成功
    if response.status_code == 200:
        # 使用BeautifulSoup解析HTML
        soup = BeautifulSoup(response.text, 'html.parser')
        # 假设下一页的标签是<a>标签，并且class为'next'
        next_page_tag = soup.find('a', class_='next')
        # 获取下一页的URL
        next_page_url = next_page_tag['href'] if next_page_tag else None
        # 获取当前页面的内容（这里只是示例，具体实现根据实际页面结构来）
        page_content = soup.find('div', class_='content')
        # 将当前页面的内容添加到列表中
        all_pages_content.append(page_content.text if page_content else 'No content found')
        return next_page_url
    else:
        print(f"Failed to retrieve {url}")
        return None

# 主循环，用于爬取所有页面
current_url = start_url
page_number = 1
while current_url:
    print(f"Crawling page {page_number}")
    current_url = crawl_page(current_url)
    page_number += 1

# 将所有页面的内容保存到本地文件
with open('crawled_content.txt', 'w', encoding='utf-8') as file:
    for content in all_pages_content:
        file.write(content + '\n\n')

print("Content has been saved to crawled_content.txt")