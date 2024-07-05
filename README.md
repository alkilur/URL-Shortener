# URL-Shortener REST API

Simple url shortener REST API service, implements URL save, delete and redirect by alias.
<br>

### Endpoints📞

| URL               | HTTP      | Action                        |
| ----------------  | --------- | ----------------------------- |
| '/{alias}'        | GET       | redirect by alias             |
| '/url/'           | POST (basic auth) | add new URL from request body |
| '/url/{alias}'    | DELETE (basic auth) | delete URL by alias         |

#
### Database table🔖

| Column     | Type         | Description                         |
| ---------- | ------------ | ----------------------------------- |
| id         | INT          | unique primary key                  |
| date       | VARCHAR(8)   | task completion date                |
| title      | TEXT         | task name                           |
| comment    | TEXT         | additional task comment             |
| repeat     | VARCHAR(128) | task repetition rule                |

#
### Run app🚀
<br>

1. **Build docker image:**
```bash
docker build -t url-shortener .
```
2. **Run docker container:**
```bash
docker run -d -p 8082:8082 url-shortener
```
3. **Use API by Endpoints.**
```bash
🚀
```

#
**original idea [@JustSkiv](https://github.com/JustSkiv)**
