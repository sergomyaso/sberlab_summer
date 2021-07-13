FROM alpine
COPY terrServ ./terrServ
COPY terraform ./terraform
RUN chmod +x /terrServ
ENTRYPOINT ["/terrServ"]