# For running the mockserver using prism
FROM stoplight/prism:4

ADD everyshilling.yaml /usr/src/prism/
EXPOSE 4010

#the path test-api.yaml to be changed to everyshilling.yaml once backend devs are done
CMD ["mock","-h","'0.0.0.0'","/usr/src/prism/everyshilling.yaml"]

