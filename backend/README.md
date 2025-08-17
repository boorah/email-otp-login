# email-otp-login

A simple backend service in golang for implementing email otp based login

## Building the image

```
docker build -t email-otp-login .
```

## Running container

```
docker run --rm  -p 8080:8080 --name test-container --env-file ./.env.container email-otp-login
```

Make sure to create a file named `.env.container` which has all the environment variables
