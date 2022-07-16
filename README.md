## Purpose
MegaCrawler is a scrapper that is based on colly, and updates information in database periodically. 

## Feature
* Service powered auto restart, can also run in clt.
* Log to service.
* Host a webserver on service.
* Built in webclient on clt to run task and check task at runtime.

## Example
In this scrapper it is intended to built plugins and do an empty import on them. Using `init()` to register website.

Then use `megaCrawler.Start()` to launch the crawler.

When the crawler is listening, you can use these flag to check or change the service:

* `--start string` Launch the selected website now.
* `--get string` Get the status of the selected website.
* `--list` List all current registered websites.

Note: `megaCrawler.Start()` is a blocking call.