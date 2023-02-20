# REST-full-API-with-Go-
Preview : 
      <img width="1440" alt="Screenshot 2023-02-16 at 1 37 09 PM" src="https://user-images.githubusercontent.com/87073574/219305373-ab19a6bc-c798-4686-b954-7efbf45d6ba9.png">


# Retrieve all albums 
 
 curl http://localhost:8080/albums
      
# Retrieve single album using ID

      curl http://localhost:8080/albums/2


# Post data/album to server
      curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'


#reference 
-https://go.dev/doc/tutorial/web-service-gin
