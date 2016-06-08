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
```json
{ "text": "I hate talking with my kombucha supplier." }
```

Example Resposne:
```json
{ "chill text": "I love talking with my chill kombucha supplier." }
```
