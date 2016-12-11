# sina_news_spider
Here's a web spider aims to crawl the news from sina news in english.
Because of golang's feature, you just need to copy the binary file to your linux server and run it.
When you run the sina_news_spider. Here're two main directories named 'result' and 'catalog' the spider will create.

./result/ saves the news divided by the special coverage and subject. Each special coverage also has a .usns file saves the urls of subjects under this special coverage.
./result/[special coverage]/[subject]/ saves the news with the news' id and .sns suffix.

./catalog/ saves all paths of files that this spider created.
./catalog/[special coverage]/ saves the files' paths relate this special coverage.

\*.snp saves the news' paths.
./catalog/\*.usns saves the subjects' paths under the special coverage.
./result/\*.usns saves the urls.
./result/[special coverage]/[subject]/\*.sns saves the news with 4 lines are subject, title, time, article.

Enjoy it and forgive my poor English :-)
