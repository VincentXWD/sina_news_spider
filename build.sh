go build *.go
mv fakeHeader sina_news_spider
scp sina_news_spider ubuntu@123.207.173.126:~/sina_news_spider
