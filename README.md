
# Go-Kit Pattern

This is just standar pattern create with go kit command

## Authors

- [@giriaditya](https://www.github.com/giriaditya)

## Deployment

To deploy this project run

```bash
  go build
```

## API Reference

#### Create books

```http
  POST /create
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `title`   | `string` | **Required**. Id of item to fetch |
| `author`  | `string` | **Required**. Id of item to fetch |

#### Get item

```http
  POST /update
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of item to fetch    |
| `title`   | `string` | **Required**. Id of item to fetch |
| `author`  | `string` | **Required**. Id of item to fetch |

