from pymongo import MongoClient
from bson import json_util
import urllib.request

client = MongoClient()
db = client['bonpreu']
prods = db['prods']

empty = bytearray()
empty.extend(b'[]')
baseUrl = 'https://www.compraonline.bonpreuesclat.cat/products/page-{0}/products.json'
i = 0
while True:
    url = baseUrl.format(i)
    data = urllib.request.urlopen(url).read()
    if data == empty:
        break
    data = json_util.loads(data)
    prods.insert_many(data)
    i += 1
