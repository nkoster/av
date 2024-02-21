curl -X POST http://localhost:3000/newproduct \
     -H 'Content-Type: application/json' \
     -d '{
           "title": "Nieuw Product",
           "url_title": "nieuw-product",
           "images": "afbeelding1.jpg,afbeelding2.png",
           "descr": "Dit is een beschrijving van het nieuwe product.",
           "specs": "Specificaties van het product.",
           "price": 1999,
           "weight": 500,
           "length": 0,
           "width": 0,
           "height": 0
         }'
