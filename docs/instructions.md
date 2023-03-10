# Docker Service for Spotlight Prism

### To test the mock api server Run

 N/B: Docker compose might generate errors...if the everyshilling.yml file is not placed under the root folder.

- Under the **root directory** on your terminal run:


```
docker-compose up
```

 


```
prism mock everyshilling.yml
```

- In either a new terminal or your browser test out your endpoint using curl command.
Replace currencies with your endpoint
```
curl http://127.0.0.1:4010/currencies
```


## Side note : test-api.yaml files to be deleted once all the endpoints are included in the everyshilling,yml
