# Golang NLU Chat Bot
This repo is to support a POC using various Golang ML/NLU/NLP libraries to build a chat bot for order pizzas. The goal
is being able to determine the intent of a message sent to the bot, then pull the correct entities from that message.
Being a POC, this code attempts to keep things simple by only performing two tasks:
- Determine the intent of a message, in this case "are you ordering a pizza" or "are you inquiring about store hours"
- Determine the entities of a pizza order, in this case what size and toppings are you ordering

In the spirit of keeping it simple, this repo does not try to tackle conversation handling, just a simple "1 question, 1
response" chat bot. However, it should lay decent groundwork for expanding the bot.

## Intents
Intents will be determined using Bayesian classification with the following library: https://github.com/jbrukh/bayesian

## Named Entities
Named Entity Recognition will be performed using the following library: https://github.com/jdkato/prose

## REST Framework
The bot will be wrapped in a REST framework, specifically the Echo framework: https://echo.labstack.com/