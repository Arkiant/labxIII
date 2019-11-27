# labxIII
3th edition LabX travelgate

## Descripción
Crear un chatbot con DialogFlow que tenga interpretación de lenguaje natural, crear comandos de voz con los cuales poder terminar con una reserva a hotelX

## Primeros pasos

- Tener instalado Golang 1.13 https://golang.org/dl/
- DialogFlow
- Webhook:
    - URL
    - Basic AUTH
    - Headers
    - https://cloud.google.com/dialogflow/docs/fulfillment-overview

## Arquitectura
<img src="./arquitectura.png">

## Propuesta

Crear un servicio en lenguaje golang (el webhook) que conecta con hotelx y poder realizar una conversación desde OK Google.

Crearemos 2 Intents:

- Search
- Book

### Search
Intentará descifrar fechas y lugar y obtener la opción mas barata que tenga dispo, para ello tendremos que generar una petición a hotel-search con los plugins cheapest, una petición a hotel-list y una a quote para comprobar que esté disponible.

#### Parámetros necesarios

- fecha de inicio
- fecha final
- personas
- lugar

### Book
Mediante lenguaje natural le diremos a Google nuestros datos necesarios para poder realizar una reserva


