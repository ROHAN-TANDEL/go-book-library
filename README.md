# go-book-library

	/*
		get:
		curl -X GET 'http://localhost:8080/get-books'

		get:
		curl -X GET 'http://localhost:8080/get-book/28'

		put:
		curl -X PUT "http://localhost:8080/replace-book/19" -d '{"title": "RAHi 2", "language": "Arabic", "summary": "Sparrow and the rain", "isbn": "122222WRRHAN", "publisher": "Trains and Tracks Co", "publication_date":"2000-10-10"}' -H "Content-Type: application/json"

		post:
		curl -X POST "http://localhost:8080/add-book" -d '{"title": "The Power Of Subconcious Mind", "language": "Arabic", "summary": "Sparrow and the rain", "isbn": "889a891244403E", "publisher": "Trains and Tracks Co", "publication_date":"2000-10-10"}' -H "Content-Type: application/json"

		patch:
		curl -X PATCH "http://localhost:8080/upgrade-book/28" -d '{"title": "Power of Habits", "language": "Arabic", "summary": "Sparrow and the rain", "isbn": "122222WWWARR", "publisher": "Trains and Tracks Co", "publication_date":"2000-10-10"}' -H "Content-Type: application/json"

		delete:
		curl -X DELETE "http://localhost:8080/remove-book/28"
	*/