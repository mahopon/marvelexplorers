from dotenv import load_dotenv
from scrap.charScrapper import charSrapper

def main():
    load_dotenv()
    scrape_characters()

if "__name__" == "__main__":
    main()