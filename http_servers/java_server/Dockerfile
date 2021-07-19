FROM openjdk:15-jdk-alpine
COPY ./out/artifacts/java_server_jar/java_server.jar ./app.jar
EXPOSE 8080
ENTRYPOINT ["java","-jar","/app.jar"]