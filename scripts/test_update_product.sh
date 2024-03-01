curl -X POST http://localhost:3000/updateproduct \
     -H 'Content-Type: application/json' \
     -d '{
           "id": 5,
           "title": "Bijgewerkt Product 2",
           "images": "afbeelding.jpg",
           "descr": "Dit is een bijgewerkte beschrijving van het product.",
           "specs": "Bijgewerkte specificaties van het product.",
           "price": 2099,
           "weight": 550,
           "length": 10,
           "width": 10,
           "height": 20
         }'
