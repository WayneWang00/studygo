目录结构：
main---|
       html-----|
                |__css  样式
                |__img  图片
                |__js   

       view-----|
                |__index.html    列表页
                |__context.html  详情页
                |__search.html   搜索页
                |__searchmsg.html搜索到的内容

       infos.go 爬取页面上的信息
 
       main.go  从数据库中拿到数据，并显示在页面上，搜索数据高亮显示

       news.go  数据库连接，redis做mysql缓存             
	
       页面跳转示意图