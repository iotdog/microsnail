
curl -X POST --header 'Content-Type: application/json' --header 'Accept: application/json' \
-d '{
   "phone": "12341234000",
   "password": "123456",
   "verif": "111111"
 }' 'http://localhost:60001/api/register'
