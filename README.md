# Golang NLU Chat Bot
This repo is to support a POC using various Golang ML/NLU/NLP libraries to build a chat bot for order pizzas. The goal
is being able to determine the intent of a message sent to the bot, then pull the correct entities from that message.

## Intents
Intents will be determined using Bayesian classification with the following library: https://github.com/jbrukh/bayesian

## Named Entities
Named Entity Recognition will be performed using the following library: https://github.com/jdkato/prose

## REST Framework
The bot will be wrapped in a REST framework, specifically the Echo framework: https://echo.labstack.com/