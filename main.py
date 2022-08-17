import requests
import base64

from random import randrange
from pexels_api import API
from PIL import Image, ImageDraw, ImageFont


def getWord(path):
    f = open(path)
    return f.read().split("\n")

def getCaption():
    verbLines = getWord("sexVerbs.txt")
    partLines = getWord("bodyParts.txt")
    objectLines = getWord("householdObjects.txt")

    verb = verbLines[randrange(0, len(verbLines))]
    part = partLines[randrange(0, len(partLines))]
    object = objectLines[randrange(0, len(objectLines))]
    caption = verb + " his " + part + " with a " + object
    return caption

def getPhotoURL():
    PEXELS_API_KEY = "abc123"
    api = API(PEXELS_API_KEY)
    api.search('men', 80, 1)
    photos = api.get_entries()
    sourceID = photos[randrange(1, 80)].id
    sourceURL = "https://api.pexels.com/v1/photos/" + str(sourceID)
    return sourceURL

def getPhoto():
    r = base64.b64encode(requests.get(getPhotoURL(), stream=True).content)
    with open("dude.jpg", "wb") as f:
        f.write(r)

def getFont():
    fonts = ["Herr_Von_Muellerhoff/HerrVonMuellerhoff-Regular.ttf", "Homemade_Apple/HomemadeApple-Regular.ttf", "Inspiration/Inspiration-Regular.ttf", "Pacifico/Pacifico-regular.ttf"]
    font = fonts[randrange(0, len(fonts) - 1)]
    return font 

def captionPhoto():
    im = Image.open("dude.jpg")
    draw = ImageDraw.Draw(im)
    font = ImageFont.truetype(getFont(), 20)
    draw.text((0, 0), getCaption(), (255, 255, 255), font)
    im.save("out.jpg")

def main():
    getPhoto()
    captionPhoto()

main()
