# Chill as a Service
Chill as a Service (CaaS) aims to provide what the internet needs most: more chill. 
Backed by the chillest of algorithms, CaaS provides a robust API for all of your chill 
needs.

# API
## /chill
You've got an exam comming up. The boss needs that new software package
by Friday. Your significant other is really putting on the pressure
to try that new hookah bar that you'd really rather not go visit. What
ever the occasion, some times you just need to chill. CaaS has you 
covered. Just simply hit the the `/chill` endpoint and you will get 
back a JSON response with a chill phrase to help bring you serenity.

Example request:
```
http://chill.io/chill
```

Example Response:
```json
{ "chill": "super chill" }
```

## /chillify
Some one just handed you a large document. Just a single glance, and you could
already tell: this was seriously going to harsh your mellow. Who hasn't this 
happened to. The `/chillify` endpoint is their to remedy this exact situation.
Simply provide it with the unchill text, and `/chillify` will send you back
a much more chill version of the text.

Example Request:
```
http://chill.io/chillify?text=I%20hate%20talking%20with%20my%20kombucha%20supplier.
```

Example Response:
```json
{ "chill_text": "I love talking with my chill kombucha supplier." }
```

## /states_chill
Where ever you are, sometimes you just need a chill image. In this respect,
the `/states_chill` endpoint has your back. Just provide the state you would
like to see a chill image over, and `/states_chill` will send you back a chill
image url. The image will resonate with the chillest vibes of the provided
state.

Example Request:
```
http://chill.io/states_chill?state=MN
```

Example Response:
```json
{ "chill_image": "http://mycdn.chill.io/MN.jpg" }
```
