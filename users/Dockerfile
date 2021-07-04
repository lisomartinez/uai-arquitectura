FROM gradle:7-jdk11-openj9 AS build
WORKDIR /opt/app
COPY ./ /opt/app/
RUN gradle build --stacktrace

FROM openjdk:11

LABEL maintaner="lisandromartinez@gmail.com"

COPY --from=build /opt/app/users/build/libs/user-service.0.0.1-SNAPSHOT.jar  /opt/app/users.jar
EXPOSE 8080
ENTRYPOINT ["java","-jar","/opt/app/users.jar"]
