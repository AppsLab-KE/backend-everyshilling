FROM stoplight/prism:4
ADD openapi.yaml /usr/src/prism/
EXPOSE 4010
CMD ["mock","-h","'0.0.0.0'","/usr/src/prism/openapi.yaml"]