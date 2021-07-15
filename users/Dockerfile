FROM adoptopenjdk/openjdk11:alpine AS build
COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src
RUN ./gradlew build --no-daemon 

FROM adoptopenjdk/openjdk11:alpine

EXPOSE 8080

RUN mkdir /app
ARG AZ_DATABASE_NAME
ARG AZ_SQL_SERVER_PASSWORD
ENV AZ_DATABASE_NAME=${AZ_DATABASE_NAME}
ENV AZ_SQL_SERVER_PASSWORD=${AZ_SQL_SERVER_PASSWORD}
COPY --from=build /home/gradle/src/build/libs/*.jar /app/spring-boot-application.jar

ENTRYPOINT ["java", "-jar","/app/spring-boot-application.jar"]