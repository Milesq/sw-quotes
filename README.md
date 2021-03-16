# Star Wars quotes

This repo is an engine to search movie scene by quote

## Intended Use

I'll create discord bot which will consume this API and publish Star Wars scene

## Descritpion

In movies.config.yml you can define named scenes, which can be achieved later by their names

In movies directory, you should place movies & srt files belonging to them

# Queries

## Simple query

You can send simple query as follows `"begin phrase"-"end phrase"`.

## Explicit declare movie ID

Sometimes you want to be sure this engine will respond with gif from specific movie. You can use movie id. The query will be in this form `#movieID"begin phrase"-"end phrase"`

## Offset

After phrase you can declare offset (in milliseconds)

## Phrase number

When engine find multiple matching phrases in subs, engine will return an error with found scenes and their ids. Then you can pass scene id between `[` and `]`

## Full example

```
#1"You turned her against me"(-2)[3]-"I will do what I must"(4)[5])
```
