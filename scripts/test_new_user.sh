curl -X POST http://localhost:3000/newuser \
     -H "Content-Type: application/json" \
     -d '{
          "first_name":"Niels",
          "last_name":"Koster",
          "username":"niels",
          "password":"apekop",
          "email":"niels@w3b.net",
          "phone":"0651123841",
          "street_name":"peppengouw",
          "house_number":"7",
          "city":"Almere",
          "postal_code":"1351 NA"
        }'
