# GolangBackendLvl1CourseWork
Bitme URL Shortener is an easy-to-use app which provides new and improved experience in shortening your URLs.
link: https://bitmegb.herokuapp.com/home

It consists of 3 main components:
-Frontend UI, written in HTML that provides a user-friendly experience of inputting your URL (https://bitmegb.herokuapp.com/su/*Short URL*) and getting a shortened version as well as various statistics gathered from people using the shortened link that can be gathered by redirecting with you admin URL https://bitmegb.herokuapp.com/getData/*Admin URL*.

-Backend part, written in GO that generates shortened URLs, writes and gets data from a data base and sends all required information to the Frontend part for the user. The router part is written using Open API and go chi router(https://github.com/go-chi/chi).

-Database that stores long URLs and their short versions as well as data calculated by the backend part of the app. Right now the data base is implemented using goleveldb(https://github.com/syndtr/goleveldb).

-Hexagonal architecture is implemented for easy component swap

-Postgres implementation is planned fr future releases.
