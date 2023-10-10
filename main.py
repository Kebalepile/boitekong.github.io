from spiders.public.govpage.govpageSpider import Spider
from spiders.private.govpage.govpageSpider import Spider as PrivateSpider


def main():

    govpage_spider = Spider()
    govpage_spider.Launch()

    govpagePrivateSector = PrivateSpider()
    govpagePrivateSector.Launch()
main()