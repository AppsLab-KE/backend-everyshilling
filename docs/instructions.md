# To test the mock api server try out these commands:

- From your terminal run:

``` docker compose up ``` - This starts the prism container

```prism mock openapi.yml``` -

- In a new terminal using the local host returned test out your endpoint. Example

```curl http://127.0.0.1:4010/currencies ```

Note : Docker compose might generate errors...still trying to figure out the correct path

