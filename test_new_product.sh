curl -X POST http://localhost:3000/newproduct \
     -H 'Content-Type: application/json' \
     -d '{
           "title": "Nieuw Product",
           "url_title": "nieuw-product",
           "image": "afbeelding.jpg",
           "descr": "Dit is een beschrijving van het nieuwe product.",
           "specs": "Specificaties van het product.",
           "price": 1999,
           "weight": 500
         }'
