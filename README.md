# mini-olx

> A simple Golang API connected to a MySql database to **list**, **create** and **delete** adverts


## How to setup:
- Clone project
- Create **.env** file following **.env.example**
- Run `docker-compose up -d` to set up the MySql database and frontend
- Run `go run main.go`


## API
**List adverts**
```
curl --location --request GET 'http://localhost:4040/api/ads'
```

**Create an advert**
```
curl --location --request POST 'http://localhost:4040/api/ads' \
--form 'title="Sofá cinzento 5"' \
--form 'description="Lorem Ipsum dasda"' \
--form 'price="350.5"' \
--form 'ad_image=@"/Users/isabelsantos/Pictures/91072971_3683624258330948_398133950691672064_n.jpeg"'
```


**Delete an advert**
```
curl --location --request DELETE 'http://localhost:4040/api/ads/{advert_id}'
```
