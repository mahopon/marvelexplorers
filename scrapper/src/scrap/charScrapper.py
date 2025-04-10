import os
import hashlib
import time
import requests
import dotenv
from classes.Character import Character

def scrape_characters():
    PUBLIC_KEY = os.getenv("PUBLIC_KEY")
    PRIVATE_KEY = os.getenv("PRIVATE_KEY")
    TIMESTAMP = str(time.time())
    HASH = hashlib.md5((TIMESTAMP + PRIVATE_KEY + PUBLIC_KEY).encode("utf-8")).hexdigest()
    offset = 0
    URL = f"https://gateway.marvel.com/v1/public/characters?apikey={PUBLIC_KEY}&hash={HASH}&ts={TIMESTAMP}&offset={str(offset)}"
    r = requests.get(URL)
    response = r.json()
    chars = []
    for res in response["data"]["results"]:
        newChar = Character(res["id"], res["name"], res["description"], res["modified"], res["thumbnail"]["path"], res["thumbnail"]["extension"], res["resourceURI"])
        chars.append(newChar)
    print(chars)

if __name__ == "__main__":
    dotenv.load_dotenv()
    scrape_characters()