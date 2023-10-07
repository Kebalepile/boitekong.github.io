import sys
import time
import logging
from typing import List
from datetime import datetime
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.remote.webelement import WebElement
from selenium.webdriver.support import expected_conditions as EC
# from selenium.common.exceptions import StaleElementReferenceException
from selenium.webdriver.common.action_chains import ActionChains as AC
from spiders.types.types import Links, BlogPost

govPageLinks: dict = Links()


class Spider:
    Name = "gov-page"

    def __init__(self):

        self.AllowedDomains = [
            "https://www.govpage.co.za/",
            "https://www.govpage.co.za/latest-govpage-updates"
        ]
        # webdriver options
        opt = webdriver.ChromeOptions()
        # opt.add_argument("--headless") # enable headless for production
        self.driver = webdriver.Chrome(options=opt)
        # Set the window size to 768x1024 (tablet size)
        self.driver.set_window_size(768, 1024)

    def Launch(self):

        log.info(f"{self.Name} spider has Lunched ")

        self.driver.get(self.AllowedDomains[0])

        menu: WebElement = self.driver.find_element(
            By.CSS_SELECTOR, "*[aria-label='Menu']")

        menu.click()

        menuOptions: List[WebElement] = self.driver.find_elements(
            By.CSS_SELECTOR, "ul li.wsite-menu-item-wrap a.wsite-menu-item")

        url: (str | None)

        for o in menuOptions:

            if "govpage" in o.text.lower():
                url = o.get_attribute("href")
                break

        if url is not None:

            selector: str = ".blog-title-link"

            log.info(f"Loading {self.Name} vacancy page")

            self.driver.get(url)

            self.Emma(10)

            wait: WebDriverWait = WebDriverWait(self.driver, 10)
            elems: List[WebElement] = wait.until(
                EC.presence_of_all_elements_located(
                    (By.CSS_SELECTOR, selector))
            )

            AC(self.driver).scroll_to_element(elems[0])

            elems: List[WebElement] = self.driver.find_elements(
                By.CSS_SELECTOR, selector)

            vacanciesLink: (str | None)

            for e in elems:

                text: str = e.text.lower()
                # replace date str below with self.Date().lower()
                if "06 october 2023" in text:
                    govPageLinks["Title"] = text
                    vacanciesLink = e.get_attribute("href")
                    break

            if vacanciesLink is not None:
                self.departments(vacanciesLink)
            else:
                log.info(
                    f"{self.Name}, Sorry, No Government Job Posts for today")
                self.driver.close()

    def departments(self, url: str):

        log.info(f"{self.Name}, searching for latest government vacancies.")

        selector: str = "[id^='blog-post-'] a"

        self.driver.get(url)
        self.Emma(15)

        elems: List[WebElement] = self.driver.find_elements(
            By.CSS_SELECTOR, selector)

        if len(elems) > 0:

            for e in elems:

                text: str = e.text
                href: str = e.get_attribute("href")

                isTitle: bool = self.Date().lower() in text
                isPrivateSectorOpportunities: bool = "PRIVATE SECTOR OPPORTUNITIES".lower() in text

                if not isTitle and not isPrivateSectorOpportunities:
                    govPageLinks["Departments"][text] = href

            log.info(govPageLinks)
            for k in govPageLinks["Departments"]:
                blogpost = self.postContent(govPageLinks["Departments"][k])
                govPageLinks["BlogPosts"].append(blogpost)
            log.info(govPageLinks)

    def postContent(self, url: str):
        self.driver.get(url)
        self.Emma(15)

        WebDriverWait(self.driver, 10).until(
            EC.presence_of_all_elements_located(
                (By.CSS_SELECTOR, ".blog-post"))
        )

        selector: str = ".blog-title-link.blog-link"

        elems: List[WebElement] = self.driver.find_elements(
            By.CSS_SELECTOR, selector)
        if len(elems) > 0:
            e = elems[0]
            text = e.text
            href = e.get_attribute("href")

            date = self.driver.find_element(
                By.CSS_SELECTOR, ".blog-date > .date-text").text

            blogPost = BlogPost()
            blogPost["Title"] = text
            blogPost["Href"] = href
            blogPost["PostedDate"] = date

            elems = self.driver.find_elements(
                By.CSS_SELECTOR, ".blog-content > .paragraph")
            if len(elems):
                for e in elems:
                    text = e.text
                    content = blogPost["Content"]
                    content.append(text)
                    blogPost["Content"] = content
            else:
                src = self.driver.execute_script("""
                        const src = Array.from(document.getElementsByTagName('iframe')).filter(f =>{
                        
                            if (f.src.includes("drive.google")){
                                return f
                            }
                                        
                        }).map(f => f.src);
                        return src[0]; 
                    """)
                blogPost["Iframe"] = src
            return blogPost
        return "no blog post found"

    def Date(self) -> str:
        # current date
        date = datetime.now()
        return date.strftime("%d %B %Y")

    def Emma(self, seconds: float):

        time.sleep(seconds)


# Create a custom formatter for log messages
log_formatter = logging.Formatter(
    "%(asctime)s [%(levelname)s]: %(message)s", datefmt="%d %B %Y %H:%M:%S")
# Create a logger
log = logging.getLogger()
log.setLevel(logging.INFO)  # Set the logging level to INFO
# Create a console handler and set the formatter
console_handler = logging.StreamHandler(sys.stdout)
console_handler.setFormatter(log_formatter)

# Add the console handler to the logger
log.addHandler(console_handler)
# Now, you can use the logger to print messages with timestamps
# log.info("This is an informational message.")
# log.warning("This is a warning message.")
